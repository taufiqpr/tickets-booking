package repository

import (
	"context"
	pb "ticket-booking/proto/train"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TrainRepository interface {
	Create(ctx context.Context, req *pb.CreateTrainRequest) (*pb.Train, error)
	GetByID(ctx context.Context, id int64) (*pb.Train, error)
	List(ctx context.Context, page, limit int32) ([]*pb.Train, int32, error)
	Update(ctx context.Context, req *pb.UpdateTrainRequest) (*pb.Train, error)
	Delete(ctx context.Context, id int64) error
}

type trainRepository struct {
	db *pgxpool.Pool
}

func NewTrainRepository(db *pgxpool.Pool) TrainRepository {
	return &trainRepository{db: db}
}

func (r *trainRepository) Create(ctx context.Context, req *pb.CreateTrainRequest) (*pb.Train, error) {
	var id int64
	err := r.db.QueryRow(ctx, `
		INSERT INTO trains (name, type, capacity, status, created_at) 
		VALUES ($1, $2, $3, 'active', NOW()) RETURNING id`,
		req.Name, req.Type, req.Capacity).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.Train{
		Id:       id,
		Name:     req.Name,
		Type:     req.Type,
		Capacity: req.Capacity,
		Status:   "active",
	}, nil
}

func (r *trainRepository) GetByID(ctx context.Context, id int64) (*pb.Train, error) {
	var train pb.Train
	err := r.db.QueryRow(ctx, `
		SELECT id, name, type, capacity, status 
		FROM trains WHERE id = $1 AND deleted_at IS NULL`, id).Scan(
		&train.Id, &train.Name, &train.Type, &train.Capacity, &train.Status)
	if err != nil {
		return nil, err
	}
	return &train, nil
}

func (r *trainRepository) List(ctx context.Context, page, limit int32) ([]*pb.Train, int32, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Query(ctx, `
		SELECT id, name, type, capacity, status 
		FROM trains WHERE deleted_at IS NULL 
		ORDER BY id DESC LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var trains []*pb.Train
	for rows.Next() {
		var train pb.Train
		err := rows.Scan(&train.Id, &train.Name, &train.Type, &train.Capacity, &train.Status)
		if err != nil {
			return nil, 0, err
		}
		trains = append(trains, &train)
	}

	var total int32
	err = r.db.QueryRow(ctx, `SELECT COUNT(*) FROM trains WHERE deleted_at IS NULL`).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return trains, total, nil
}

func (r *trainRepository) Update(ctx context.Context, req *pb.UpdateTrainRequest) (*pb.Train, error) {
	_, err := r.db.Exec(ctx, `
		UPDATE trains SET name = $1, type = $2, capacity = $3, status = $4, updated_at = NOW() 
		WHERE id = $5 AND deleted_at IS NULL`,
		req.Name, req.Type, req.Capacity, req.Status, req.TrainId)
	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, req.TrainId)
}

func (r *trainRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, `
		UPDATE trains SET deleted_at = NOW() WHERE id = $1`, id)
	return err
}
