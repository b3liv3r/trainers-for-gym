package repository

import (
	"context"
	"errors"
	"github.com/b3liv3r/trainers-for-gym/modules/trainers/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type TrainersRepositoryDB struct {
	db *sqlx.DB
}

func NewTrainersRepository(db *sqlx.DB) TrainerRepository {
	return &TrainersRepositoryDB{db: db}
}

func (repo *TrainersRepositoryDB) GetTrainersForGym(ctx context.Context, gymID int) ([]models.Trainer, error) {
	var trainers []models.Trainer
	err := repo.db.SelectContext(ctx, &trainers, "SELECT * FROM trainers WHERE gym_id = $1", gymID)
	if err != nil {
		return nil, err
	}
	return trainers, nil
}

func (repo *TrainersRepositoryDB) AvailableBookingListForTrainers(ctx context.Context, trainerID int) ([]models.Booking, error) {
	var bookings []models.Booking
	err := repo.db.SelectContext(ctx, &bookings, "SELECT * FROM available_bookings WHERE trainer_id = $1 AND start_time >= $2", trainerID, time.Now())
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (repo *TrainersRepositoryDB) CurrentBookingListForUser(ctx context.Context, userID int) ([]models.Booking, error) {
	var bookings []models.Booking
	err := repo.db.SelectContext(ctx, &bookings, "SELECT * FROM current_bookings WHERE user_id = $1 AND start_time >= $2", userID, time.Now())
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (repo *TrainersRepositoryDB) Booking(ctx context.Context, bookingID, userID int) error {
	tx, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	var booking models.Booking

	err = tx.GetContext(ctx, &booking, "SELECT * FROM available_bookings WHERE booking_id = $1 FOR UPDATE", bookingID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if time.Now().After(booking.StartTime) {
		tx.Rollback()
		return errors.New("невозможно забронировать прошедший слот")
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO current_bookings (user_id, trainer_id, start_time, end_time, activity) VALUES ($1, $2, $3, $4, $5)", userID, booking.TrainerID, booking.StartTime, booking.EndTime, booking.Activity)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM available_bookings WHERE booking_id = $1", bookingID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *TrainersRepositoryDB) UnBooking(ctx context.Context, bookingID int) error {
	tx, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	var booking models.Booking

	err = tx.GetContext(ctx, &booking, "SELECT * FROM current_bookings WHERE booking_id = $1 FOR UPDATE", bookingID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if time.Now().Add(24 * time.Hour).After(booking.StartTime) {
		return errors.New("невозможно отменить бронирование, так как до начала слота остается 24 часа или менее")
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO available_bookings (trainer_id, start_time, end_time, activity) VALUES ($1, $2, $3, $4)", booking.TrainerID, booking.StartTime, booking.EndTime, booking.Activity)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM current_bookings WHERE booking_id = $1", bookingID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
