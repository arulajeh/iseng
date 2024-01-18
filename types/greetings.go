package types

import "time"

type Greetings struct {
	Name      string    `bson:"name" json:"name"`
	Message   string    `bson:"message" json:"message"`
	IsHadir   bool      `bson:"is_hadir" json:"is_hadir"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}
