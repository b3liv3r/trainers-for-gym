package repository

import (
	"context"
	"github.com/b3liv3r/trainers-for-gym/modules/trainers/models"
)

type TrainerRepository interface {
	GetTrainersForGym(context context.Context, GymID int) ([]models.Trainer, error)
	AvailableBookingListForTrainers(context context.Context, TrainerID int) ([]models.Booking, error)
	CurrentBookingListForUser(context context.Context, UserID int) ([]models.Booking, error)
	Booking(context context.Context, BookingID, UserID int) error
	UnBooking(context context.Context, BookingID int) error
}
