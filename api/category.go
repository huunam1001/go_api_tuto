package api

import (
	"context"
	"encoding/json"
	"fmt"
	"go_api_tuto/db/mongo"
	"go_api_tuto/util"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (server *Server) GetListCategory(ctx *gin.Context) {

	apiMongoDb := server.mongo.Database(util.MONGO_DATA_BASE)
	categoryCollection := apiMongoDb.Collection("category")

	filter := bson.M{}

	cursor, err := categoryCollection.Find(context.TODO(), filter)

	if err != nil {

		print(err.Error())
		util.SendValidationError(ctx)

		return
	}

	var results []mongo.Category
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
		// fmt.Printf("%s\n", output)
	}

	util.SendApiSuccess(ctx, results, "")
}

type createCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) AddCategory(ctx *gin.Context) {

	var req createCategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {

		util.SendValidationError(ctx)

		return
	}

	me, success := util.GetUserFromRequest(ctx)

	if !success {

		util.SendInternalServerError(ctx)
		return
	}

	category := mongo.Category{
		Name:        req.Name,
		CreatedTime: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedTime: primitive.NewDateTimeFromTime(time.Now()),
		CreatedBy:   me.Username,
		UpdatedBy:   me.Username,
	}

	/// insert data to mongo atlas
	apiMongoDb := server.mongo.Database(util.MONGO_DATA_BASE)
	categoryCollection := apiMongoDb.Collection("category")

	result, err := categoryCollection.InsertOne(ctx, category)

	if err != nil {
		util.SendInternalServerError(ctx)
		return
	}
	util.SendApiSuccess(ctx, result, "")
}

type updateCategoryRequest struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (server *Server) UpdateCategory(ctx *gin.Context) {

	var req updateCategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {

		util.SendValidationError(ctx)

		return
	}

	me, success := util.GetUserFromRequest(ctx)

	if !success {

		util.SendInternalServerError(ctx)
		return
	}

	categoryId, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		util.SendInternalServerError(ctx)
		return
	}

	filter := bson.M{"_id": bson.M{"$eq": categoryId}}

	category := mongo.Category{
		Name:        req.Name,
		UpdatedTime: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedBy:   me.Username,
	}

	update := bson.M{"$set": category}

	/// update data to mongo atlas
	apiMongoDb := server.mongo.Database(util.MONGO_DATA_BASE)
	categoryCollection := apiMongoDb.Collection("category")

	result, err := categoryCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		util.SendInternalServerError(ctx)
		fmt.Println(err.Error())
		return
	}
	util.SendApiSuccess(ctx, result, "")
}
