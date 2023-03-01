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

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CategoryId  primitive.ObjectID `bson:"categoryId,omitempty" json:"categoryId"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Price       float32            `bson:"price,omitempty" json:"price"`
	CreatedTime primitive.DateTime `bson:"createdDate,omitempty" json:"createdDate"`
	UpdatedTime primitive.DateTime `bson:"updatedTime,omitempty" json:"updatedTime"`
	CreatedBy   string             `bson:"createdBy,omitempty" json:"createdBy"`
	UpdatedBy   string             `bson:"updatedBy,omitempty" json:"updatedBy"`
}
