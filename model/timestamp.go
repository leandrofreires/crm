package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DateTime struct {
	CreatedAt  primitive.DateTime `bson:"created_at" json:"created_at"`
	ModifiedAt primitive.DateTime `bson:"modified_at" json:"modified_at"`
}

func (d *DateTime) update() {
	d.ModifiedAt = primitive.NewDateTimeFromTime(time.Now())
}
func (d *DateTime) init() {
	now := primitive.NewDateTimeFromTime(time.Now())
	d.CreatedAt = now
	d.ModifiedAt = now
}
