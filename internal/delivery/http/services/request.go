package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetIdFromPath(c *gin.Context, paramName string) (primitive.ObjectID, error) {
	idString := c.Param(paramName)
	if idString == "" {
		return primitive.ObjectID{}, errors.New("empty parameter")
	}

	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return primitive.ObjectID{}, errors.New("invalid parameter")
	}

	return id, nil
}
