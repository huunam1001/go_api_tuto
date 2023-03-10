package api

import (
	"context"
	"encoding/json"
	"go_api_tuto/db/mongodb"
	"go_api_tuto/util"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (server *Server) GetListProduct(ctx *gin.Context) {

	search := ctx.Query("search")

	page, limit := util.GetPagingFromRequest(ctx)

	apiMongoDb := server.mongo.Database(util.MONGO_DATA_BASE)
	productCollection := apiMongoDb.Collection("product")

	filter := bson.A{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "$text", Value: bson.D{
					{Key: "$search", Value: search},
				},
				},
			},
			},
		},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "category"},
					{Key: "localField", Value: "categoryId"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "category"},
				},
			},
		},
		bson.D{
			{Key: "$facet",
				Value: bson.D{
					{Key: "paging",
						Value: bson.A{
							bson.D{{Key: "$count", Value: "total"}},
							bson.D{{Key: "$addFields", Value: bson.D{{Key: "page", Value: page}}}},
							bson.D{{Key: "$addFields", Value: bson.D{{Key: "limit", Value: limit}}}},
						},
					},
					{Key: "products",
						Value: bson.A{
							bson.D{{Key: "$skip", Value: page * limit}},
							bson.D{{Key: "$limit", Value: limit}},
						},
					},
				},
			},
		},
	}

	cursor, err := productCollection.Aggregate(ctx, filter)

	if err != nil {
		util.SendInternalServerError(ctx)
		return
	}

	var results []mongodb.MongGoListProductResponse
	if err = cursor.All(context.TODO(), &results); err != nil {
		util.SendInternalServerError(ctx)
		return
	}

	for _, result := range results {
		cursor.Decode(&result)
		_, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
	}

	util.SendApiSuccess(ctx, results, "")
}

type createProductRequest struct {
	Name       string  `json:"name" binding:"required"`
	CategoryId string  `json:"categoryId" binding:"required"`
	Price      float32 `json:"price" binding:"required"`
}

func (server *Server) AddProduct(ctx *gin.Context) {

	var req createProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {

		util.SendValidationError(ctx)

		return
	}

	me, success := util.GetUserFromRequest(ctx)

	if !success {

		util.SendInternalServerError(ctx)
		return
	}

	categoryId, err := primitive.ObjectIDFromHex(req.CategoryId)

	if err != nil {
		util.SendInternalServerError(ctx)
		return
	}

	product := mongodb.Product{
		CategoryId:  categoryId,
		Name:        req.Name,
		Price:       req.Price,
		CreatedTime: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedTime: primitive.NewDateTimeFromTime(time.Now()),
		CreatedBy:   me.Username,
		UpdatedBy:   me.Username,
	}

	/// insert data to mongo atlas
	apiMongoDb := server.mongo.Database(util.MONGO_DATA_BASE)
	productCollection := apiMongoDb.Collection("product")

	result, err := productCollection.InsertOne(ctx, product)

	if err != nil {
		util.SendInternalServerError(ctx)
		return
	}

	// _, err = productCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
	// 	Keys: bson.M{"name": "text"},
	// })

	// if err != nil {
	// 	println("ERROR INDEX")
	// 	println(err.Error())
	// }

	util.SendApiSuccess(ctx, result, "")
}
