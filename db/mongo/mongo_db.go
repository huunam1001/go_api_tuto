package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	CreatedTime primitive.DateTime `bson:"createdDate,omitempty"`
	UpdatedTime primitive.DateTime `bson:"updatedTime,omitempty"`
	CreatedBy   string             `bson:"createdBy,omitempty"`
	UpdatedBy   string             `bson:"updatedBy,omitempty"`
}
