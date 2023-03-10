package mongodb

import "go.mongodb.org/mongo-driver/bson/primitive"

type TokenUserInfo struct {
	Username string `bson:"username,omitempty" json:"username"`
	FullName string `bson:"fullName,omitempty" json:"fullName"`
	Email    string `bson:"email,omitempty" json:"email"`
}

type LoginTokenInfo struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Token     string              `bson:"token,omitempty" json:"token"`
	LoginDate primitive.DateTime  `bson:"loginDate,omitempty" json:"loginDate"`
	User      TokenUserInfo       `bson:"user,omitempty" json:"user"`
}

type Category struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	CreatedTime primitive.DateTime `bson:"createdDate,omitempty" json:"createdDate"`
	UpdatedTime primitive.DateTime `bson:"updatedTime,omitempty" json:"updatedTime"`
	CreatedBy   string             `bson:"createdBy,omitempty" json:"createdBy"`
	UpdatedBy   string             `bson:"updatedBy,omitempty" json:"updatedBy"`
	IsDeleted   bool               `bson:"isDeleted" json:"isDeleted"`
	DeletedBy   string             `bson:"deletedBy" json:"deletedBy"`
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
	IsDeleted   bool               `bson:"isDeleted" json:"isDeleted"`
	DeletedBy   string             `bson:"deletedBy" json:"deletedBy"`
	Category    Category           `bson:"category,omitempty" json:"category"`
}

type MongGoListProductResponse struct {
	Paging   Paging    `bson:"paging,omitempty" json:"paging"`
	Products []Product `bson:"products,omitempty" json:"products"`
}

type Paging struct {
	Total int `bson:"total,omitempty" json:"total"`
	Page  int `bson:"page,omitempty" json:"page"`
	Limit int `bson:"limit,omitempty" json:"limit"`
}
