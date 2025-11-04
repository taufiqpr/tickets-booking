package repository

import (
	"context"
	"fmt"
	pb "ticket-booking/proto/schedule"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ScheduleRepository interface {
	Create(ctx context.Context, req *pb.CreateScheduleRequest) (*pb.Schedule, error)
	GetByID(ctx context.Context, id int64) (*pb.Schedule, error)
	List(ctx context.Context, origin, destination, departureDate string, page, limit int32) ([]*pb.Schedule, int32, error)
}

type scheduleRepository struct {
	db *pgxpool.Pool
}

func NewScheduleRepository(db *pgxpool.Pool) ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (r *scheduleRepository) Create(ctx context.Context, req *pb.CreateScheduleRequest) (*pb.Schedule, error) {
	var id int64
	var trainName string

	// Get train name
	err := r.db.QueryRow(ctx, `SELECT name FROM trains WHERE id = $1`, req.TrainId).Scan(&trainName)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(ctx, `
		INSERT INTO schedules (train_id, origin, destination, departure_time, arrival_time, price, available_seats, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, (SELECT capacity FROM trains WHERE id = $1), NOW()) RETURNING id`,
		req.TrainId, req.Origin, req.Destination, req.DepartureTime, req.ArrivalTime, req.Price).Scan(&id)
	if err != nil {
		return nil, err
	}

	// Get available seats from train capacity
	var availableSeats int32
	err = r.db.QueryRow(ctx, `SELECT capacity FROM trains WHERE id = $1`, req.TrainId).Scan(&availableSeats)
	if err != nil {
		availableSeats = 0
	}

	return &pb.Schedule{
		Id:             id,
		TrainId:        req.TrainId,
		TrainName:      trainName,
		Origin:         req.Origin,
		Destination:    req.Destination,
		DepartureTime:  req.DepartureTime,
		ArrivalTime:    req.ArrivalTime,
		Price:          req.Price,
		AvailableSeats: availableSeats,
	}, nil
}

func (r *scheduleRepository) GetByID(ctx context.Context, id int64) (*pb.Schedule, error) {
	var schedule pb.Schedule
	var trainName string

	err := r.db.QueryRow(ctx, `
		SELECT s.id, s.train_id, t.name, s.origin, s.destination, 
		       s.departure_time, s.arrival_time, s.price, s.available_seats
		FROM schedules s
		LEFT JOIN trains t ON s.train_id = t.id
		WHERE s.id = $1 AND s.deleted_at IS NULL`, id).Scan(
		&schedule.Id, &schedule.TrainId, &trainName, &schedule.Origin, &schedule.Destination,
		&schedule.DepartureTime, &schedule.ArrivalTime, &schedule.Price, &schedule.AvailableSeats)
	if err != nil {
		return nil, err
	}

	schedule.TrainName = trainName
	return &schedule, nil
}

func (r *scheduleRepository) List(ctx context.Context, origin, destination, departureDate string, page, limit int32) ([]*pb.Schedule, int32, error) {
	offset := (page - 1) * limit

	query := `
		SELECT s.id, s.train_id, t.name, s.origin, s.destination, 
		       s.departure_time, s.arrival_time, s.price, s.available_seats
		FROM schedules s
		LEFT JOIN trains t ON s.train_id = t.id
		WHERE s.deleted_at IS NULL`

	args := []interface{}{}
	argCount := 0

	if origin != "" {
		argCount++
		query += fmt.Sprintf(" AND s.origin ILIKE $%d", argCount)
		args = append(args, "%"+origin+"%")
	}

	if destination != "" {
		argCount++
		query += fmt.Sprintf(" AND s.destination ILIKE $%d", argCount)
		args = append(args, "%"+destination+"%")
	}

	if departureDate != "" {
		argCount++
		query += fmt.Sprintf(" AND DATE(s.departure_time) = $%d", argCount)
		args = append(args, departureDate)
	}

	query += " ORDER BY s.departure_time ASC"

	argCount++
	query += fmt.Sprintf(" LIMIT $%d", argCount)
	args = append(args, limit)

	argCount++
	query += fmt.Sprintf(" OFFSET $%d", argCount)
	args = append(args, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var schedules []*pb.Schedule
	for rows.Next() {
		var schedule pb.Schedule
		var trainName string

		err := rows.Scan(&schedule.Id, &schedule.TrainId, &trainName, &schedule.Origin, &schedule.Destination,
			&schedule.DepartureTime, &schedule.ArrivalTime, &schedule.Price, &schedule.AvailableSeats)
		if err != nil {
			return nil, 0, err
		}

		schedule.TrainName = trainName
		schedules = append(schedules, &schedule)
	}

	// Count total
	countQuery := `SELECT COUNT(*) FROM schedules s WHERE s.deleted_at IS NULL`
	countArgs := []interface{}{}
	countArgCount := 0

	if origin != "" {
		countArgCount++
		countQuery += fmt.Sprintf(" AND s.origin ILIKE $%d", countArgCount)
		countArgs = append(countArgs, "%"+origin+"%")
	}

	if destination != "" {
		countArgCount++
		countQuery += fmt.Sprintf(" AND s.destination ILIKE $%d", countArgCount)
		countArgs = append(countArgs, "%"+destination+"%")
	}

	if departureDate != "" {
		countArgCount++
		countQuery += fmt.Sprintf(" AND DATE(s.departure_time) = $%d", countArgCount)
		countArgs = append(countArgs, departureDate)
	}

	var total int32
	err = r.db.QueryRow(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return schedules, total, nil
}
