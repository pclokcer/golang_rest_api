package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclokcer/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentController interface {
	All(c *gin.Context)
	SetComment(c *gin.Context)
}

type commentController struct {
	connection *mongo.Database
}

func NewCommentController(db *mongo.Database) CommentController {
	return &commentController{
		connection: db,
	}
}

func (comment commentController) All(c *gin.Context) {
	commentsData, err := comment.connection.Collection("comment").Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}

	var comments []entity.Comment

	if err = commentsData.All(context.TODO(), &comments); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, comments)
}

func (new commentController) SetComment(c *gin.Context) {

	var comment entity.Comment
	err := c.BindJSON(&comment)

	if comment.ID != "" {
		comment.ID = ""
	}

	if err != nil {
		panic(err)
	}

	res, err := new.connection.Collection("comment").InsertOne(context.Background(), comment)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"id": res.InsertedID})
}
