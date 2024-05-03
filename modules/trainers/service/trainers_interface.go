package service

import (
	"context"
	"github.com/b3liv3r/trainers-for-gym/modules/trainers/models"
)

type Trainer interface {
	GetTrainersForGym(ctx context.Context, gymID int) ([]models.Trainer, error)
	AvailableBookingListForTrainers(ctx context.Context, trainerID int) ([]models.Booking, error)
	CurrentBookingListForUser(ctx context.Context, userID int) ([]models.Booking, error)
	Booking(ctx context.Context, bookingID, userID int) (string, error)
	UnBooking(ctx context.Context, bookingID int) (string, error)
}
