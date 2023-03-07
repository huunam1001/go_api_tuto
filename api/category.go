package api

import (
	"context"
	"fmt"
	"go_api_tuto/db/mongo"
	"go_api_tuto/util"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
