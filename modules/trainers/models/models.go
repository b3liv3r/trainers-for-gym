package models

import "time"

type Trainer struct {
	ID         int    `db:"trainer_id"`
	GymId      int    `db:"gym_id"`
	Name       string `db:"trainer_name"`
	Speciality string `db:"speciality"`
}

type Booking struct {
	ID        int       `db:"booking_id"`
	UserID    int       `db:"user_id"`
	TrainerID int       `db:"trainer_id"`
	Activity  string    `db:"activity"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}
