package api

import (
	"context"
	"encoding/json"
	"go_api_tuto/db/mongo"
	"go_api_tuto/util"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (server *Server) GetListProduct(ctx *gin.Context) {

	search := ctx.Query("search")

	apiMongoDb := server.mongo.Database(util.MONGO_DATA_BASE)
	productCollection := apiMongoDb.Collection("product")

	filter := bson.M{"name": bson.M{"$regex": search, "$options": "i"}}

	options := new(options.FindOptions)
	options.SetSkip(0)
	options.SetLimit(10)

	total, _ := productCollection.CountDocuments(context.TODO(), filter)

	cursor, err := productCollection.Find(context.TODO(), filter, options)

	var results []mongo.Product
	if err = cursor.All(context.TODO(), &results); err != nil {
		util.SendInternalServerError(ctx)
		return
	}

	println("PAGE COUNT")
	println(total)

	for _, result := range results {
		cursor.Decode(&result)
		_, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		// fmt.Printf("%s\n", output)
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

	product := mongo.Product{
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

	util.SendApiSuccess(ctx, result, "")
}
