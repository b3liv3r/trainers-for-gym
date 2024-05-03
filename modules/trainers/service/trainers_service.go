package service

import (
	"context"
	"github.com/b3liv3r/trainers-for-gym/modules/trainers/models"
	"github.com/b3liv3r/trainers-for-gym/modules/trainers/repository"
	"go.uber.org/zap"
)

type TrainerService struct {
	repo repository.TrainerRepository
	log  *zap.Logger
}

func NewTrainerService(repo repository.TrainerRepository, log *zap.Logger) Trainer {
	return &TrainerService{repo: repo, log: log}
}

func (s *TrainerService) GetTrainersForGym(ctx context.Context, gymID int) ([]models.Trainer, error) {
	trainers, err := s.repo.GetTrainersForGym(ctx, gymID)
	if err != nil {
		s.log.Error("Error getting trainers", zap.Error(err))
		return nil, err
	}
	return trainers, nil
}

func (s *TrainerService) AvailableBookingListForTrainers(ctx context.Context, trainerID int) ([]models.Booking, error) {
	bookings, err := s.repo.AvailableBookingListForTrainers(ctx, trainerID)
	if err != nil {
		s.log.Error("Error getting available bookings", zap.Error(err))
		return nil, err
	}
	return bookings, nil
}

func (s *TrainerService) CurrentBookingListForUser(ctx context.Context, userID int) ([]models.Booking, error) {
	bookings, err := s.repo.CurrentBookingListForUser(ctx, userID)
	if err != nil {
		s.log.Error("Error getting current bookings", zap.Error(err))
		return nil, err
	}
	return bookings, nil
}

func (s *TrainerService) Booking(ctx context.Context, bookingID, userID int) (string, error) {
	err := s.repo.Booking(ctx, bookingID, userID)
	if err != nil {
		s.log.Error("failed to book", zap.Error(err))
		return "", err
	}
	return "Booking successful", nil
}

func (s *TrainerService) UnBooking(ctx context.Context, bookingID int) (string, error) {
	err := s.repo.UnBooking(ctx, bookingID)
	if err != nil {
		s.log.Error("failed to unbook", zap.Error(err))
		return "", err
	}
	return "Unbooking successful", nil
}
