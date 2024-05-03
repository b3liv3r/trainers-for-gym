package server

import (
	"context"
	trainerv1 "github.com/b3liv3r/protos-for-gym/gen/go/trainer"
	"github.com/b3liv3r/trainers-for-gym/modules/trainers/models"
	"github.com/b3liv3r/trainers-for-gym/modules/trainers/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TrainerRPCServer struct {
	trainerv1.UnimplementedTrainerServer
	srv service.Trainer
}

func NewTrainerRPCServer(srv service.Trainer) trainerv1.TrainerServer {
	return &TrainerRPCServer{srv: srv}
}

func (s *TrainerRPCServer) ListForGym(ctx context.Context, req *trainerv1.ListForGymRequest) (*trainerv1.ListForGymResponse, error) {
	trainers, err := s.srv.GetTrainersForGym(ctx, int(req.GymId))
	if err != nil {
		return nil, err
	}
	return &trainerv1.ListForGymResponse{Trainers: convertTrainers(trainers)}, nil
}

func (s *TrainerRPCServer) AvailableBookingList(ctx context.Context, req *trainerv1.AvailableBookingListRequest) (*trainerv1.AvailableBookingListResponse, error) {
	bookings, err := s.srv.AvailableBookingListForTrainers(ctx, int(req.TrainerId))
	if err != nil {
		return nil, err
	}
	return &trainerv1.AvailableBookingListResponse{Bookings: convertBookings(bookings)}, nil
}

func (s *TrainerRPCServer) CurrentBookingList(ctx context.Context, req *trainerv1.CurrentBookingListRequest) (*trainerv1.CurrentBookingListResponse, error) {
	bookings, err := s.srv.CurrentBookingListForUser(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &trainerv1.CurrentBookingListResponse{Bookings: convertBookings(bookings)}, nil
}

func (s *TrainerRPCServer) Booking(ctx context.Context, req *trainerv1.BookingRequest) (*trainerv1.BookingResponse, error) {
	message, err := s.srv.Booking(ctx, int(req.AvailableBookingId), int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &trainerv1.BookingResponse{Message: message}, nil
}

func (s *TrainerRPCServer) UnBooking(ctx context.Context, req *trainerv1.UnBookingRequest) (*trainerv1.UnBookingResponse, error) {
	message, err := s.srv.UnBooking(ctx, int(req.CurrentBookingId))
	if err != nil {
		return nil, err
	}
	return &trainerv1.UnBookingResponse{Message: message}, nil
}

func convertTrainers(trainers []models.Trainer) []*trainerv1.Trainers {
	var pbTrainers []*trainerv1.Trainers
	for _, trainer := range trainers {
		pbTrainers = append(pbTrainers, &trainerv1.Trainers{
			TrainerId:  int64(trainer.ID),
			GymId:      int64(trainer.GymId),
			Name:       trainer.Name,
			Speciality: trainer.Speciality,
		})
	}
	return pbTrainers
}

func convertBookings(bookings []models.Booking) []*trainerv1.Bookings {
	var pbBookings []*trainerv1.Bookings
	for _, booking := range bookings {
		startTime := timestamppb.New(booking.StartTime)
		endTime := timestamppb.New(booking.EndTime)
		pbBookings = append(pbBookings, &trainerv1.Bookings{
			BookingId: int64(booking.ID),
			Activity:  booking.Activity,
			StartTime: startTime,
			EndTime:   endTime,
		})
	}
	return pbBookings
}
