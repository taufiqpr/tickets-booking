package repository

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	pb "ticket-booking/proto/booking"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ScheduleDetails struct {
	Price         float64
	Origin        string
	Destination   string
	DepartureTime string
	ArrivalTime   string
	TrainName     string
}

type BookingRepository interface {
	Create(ctx context.Context, req *pb.CreateBookingRequest) (*pb.Booking, error)
	GetByID(ctx context.Context, id int64) (*pb.Booking, error)
	ListByUser(ctx context.Context, userId int64, page, limit int32) ([]*pb.Booking, int32, error)
	UpdateStatus(ctx context.Context, bookingId int64, status int) (*pb.Booking, error)
	Cancel(ctx context.Context, bookingId, userId int64) error
	ExpireBookings(ctx context.Context) error
}

type pgBookingRepo struct {
	pool *pgxpool.Pool
}

func NewBookingRepository(pool *pgxpool.Pool) BookingRepository {
	return &pgBookingRepo{pool: pool}
}

func (r *pgBookingRepo) Create(ctx context.Context, req *pb.CreateBookingRequest) (*pb.Booking, error) {
	var id int64
	expiredAt := time.Now().Add(10 * time.Minute)
	bookingCode := generateBookingCode()

	err := r.pool.QueryRow(ctx, `
		INSERT INTO bookings (user_id, schedule_id, seat_count, status, expires_at, booking_code, created_at) 
		VALUES ($1,$2,$3,$4,$5,$6,NOW()) RETURNING id`,
		req.UserId, req.ScheduleId, req.SeatCount, 1, expiredAt, bookingCode).Scan(&id)
	if err != nil {
		return nil, err
	}

	// Get schedule details for booking response
	scheduleDetails, err := r.getScheduleDetails(ctx, req.ScheduleId)
	if err != nil {
		return nil, err
	}

	b := &pb.Booking{
		Id:            id,
		UserId:        req.UserId,
		ScheduleId:    req.ScheduleId,
		SeatCount:     req.SeatCount,
		Status:        "pending",
		ExpiresAt:     expiredAt.Format(time.RFC3339),
		BookingCode:   bookingCode,
		CreatedAt:     time.Now().Format(time.RFC3339),
		TotalPrice:    scheduleDetails.Price * float64(req.SeatCount),
		Origin:        scheduleDetails.Origin,
		Destination:   scheduleDetails.Destination,
		DepartureTime: scheduleDetails.DepartureTime,
		ArrivalTime:   scheduleDetails.ArrivalTime,
		TrainName:     scheduleDetails.TrainName,
	}
	return b, nil
}

func (r *pgBookingRepo) GetByID(ctx context.Context, id int64) (*pb.Booking, error) {
	query := `
		SELECT b.id, b.user_id, b.schedule_id, b.seat_count, b.status, b.expires_at, b.created_at, 
		       b.booking_code, b.total_price,
		       s.origin, s.destination, s.departure_time, s.arrival_time, t.name as train_name
		FROM bookings b
		LEFT JOIN schedules s ON b.schedule_id = s.id
		LEFT JOIN trains t ON s.train_id = t.id
		WHERE b.id = $1 AND b.deleted_at IS NULL`

	var b pb.Booking
	var statusInt int
	var expiredAt, createdAt time.Time
	var totalPrice *float64
	var origin, destination, departureTime, arrivalTime, trainName *string

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&b.Id, &b.UserId, &b.ScheduleId, &b.SeatCount, &statusInt, &expiredAt, &createdAt,
		&b.BookingCode, &totalPrice,
		&origin, &destination, &departureTime, &arrivalTime, &trainName)
	if err != nil {
		return nil, err
	}

	b.Status = mapStatusIntToString(statusInt)
	b.ExpiresAt = expiredAt.Format(time.RFC3339)
	b.CreatedAt = createdAt.Format(time.RFC3339)

	if totalPrice != nil {
		b.TotalPrice = *totalPrice
	}
	if origin != nil {
		b.Origin = *origin
	}
	if destination != nil {
		b.Destination = *destination
	}
	if departureTime != nil {
		b.DepartureTime = *departureTime
	}
	if arrivalTime != nil {
		b.ArrivalTime = *arrivalTime
	}
	if trainName != nil {
		b.TrainName = *trainName
	}

	return &b, nil
}

func (r *pgBookingRepo) ListByUser(ctx context.Context, userId int64, page, limit int32) ([]*pb.Booking, int32, error) {
	offset := (int(page) - 1) * int(limit)
	query := `
		SELECT b.id, b.user_id, b.schedule_id, b.seat_count, b.status, b.expires_at, b.created_at,
		       b.booking_code, b.total_price,
		       s.origin, s.destination, s.departure_time, s.arrival_time, t.name as train_name
		FROM bookings b
		LEFT JOIN schedules s ON b.schedule_id = s.id
		LEFT JOIN trains t ON s.train_id = t.id
		WHERE b.user_id = $1 AND b.deleted_at IS NULL 
		ORDER BY b.id DESC LIMIT $2 OFFSET $3`

	rows, err := r.pool.Query(ctx, query, userId, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var res []*pb.Booking
	for rows.Next() {
		var b pb.Booking
		var statusInt int
		var expiredAt, createdAt time.Time
		var totalPrice *float64
		var origin, destination, departureTime, arrivalTime, trainName *string

		err := rows.Scan(&b.Id, &b.UserId, &b.ScheduleId, &b.SeatCount, &statusInt, &expiredAt, &createdAt,
			&b.BookingCode, &totalPrice,
			&origin, &destination, &departureTime, &arrivalTime, &trainName)
		if err != nil {
			return nil, 0, err
		}

		b.Status = mapStatusIntToString(statusInt)
		b.ExpiresAt = expiredAt.Format(time.RFC3339)
		b.CreatedAt = createdAt.Format(time.RFC3339)

		if totalPrice != nil {
			b.TotalPrice = *totalPrice
		}
		if origin != nil {
			b.Origin = *origin
		}
		if destination != nil {
			b.Destination = *destination
		}
		if departureTime != nil {
			b.DepartureTime = *departureTime
		}
		if arrivalTime != nil {
			b.ArrivalTime = *arrivalTime
		}
		if trainName != nil {
			b.TrainName = *trainName
		}

		res = append(res, &b)
	}

	// count
	var total int32
	if err := r.pool.QueryRow(ctx, `SELECT COUNT(1) FROM bookings WHERE user_id=$1 AND deleted_at IS NULL`, userId).Scan(&total); err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func (r *pgBookingRepo) UpdateStatus(ctx context.Context, bookingId int64, status int) (*pb.Booking, error) {
	_, err := r.pool.Exec(ctx, `UPDATE bookings SET status=$1, updated_at=NOW() WHERE id=$2`, status, bookingId)
	if err != nil {
		return nil, err
	}
	return r.GetByID(ctx, bookingId)
}

func (r *pgBookingRepo) Cancel(ctx context.Context, bookingId, userId int64) error {
	_, err := r.pool.Exec(ctx, `UPDATE bookings SET deleted_at=NOW(), status=$1 WHERE id=$2 AND user_id=$3`, 0, bookingId, userId)
	return err
}

func (r *pgBookingRepo) ExpireBookings(ctx context.Context) error {
	_, err := r.pool.Exec(ctx, `UPDATE bookings SET status=4 WHERE expires_at < NOW() AND status = 1`)
	return err
}

func (r *pgBookingRepo) getScheduleDetails(ctx context.Context, scheduleId int64) (*ScheduleDetails, error) {
	query := `
		SELECT s.price, s.origin, s.destination, s.departure_time, s.arrival_time, t.name as train_name
		FROM schedules s
		LEFT JOIN trains t ON s.train_id = t.id
		WHERE s.id = $1`

	var details ScheduleDetails
	err := r.pool.QueryRow(ctx, query, scheduleId).Scan(
		&details.Price, &details.Origin, &details.Destination,
		&details.DepartureTime, &details.ArrivalTime, &details.TrainName)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

func generateBookingCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return fmt.Sprintf("BK%s", string(b))
}

func mapStatusIntToString(s int) string {
	switch s {
	case 1:
		return "pending"
	case 2:
		return "success"
	case 3:
		return "failed"
	case 4:
		return "expired"
	default:
		return "unknown"
	}
}
