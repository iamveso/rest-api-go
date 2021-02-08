package controllers

import (
	"company-api/db"
	"company-api/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePersonEndpoint(c echo.Context) error {
	var newPerson models.Person
	json.NewDecoder(c.Request().Body).Decode(&newPerson)
	collection := db.Client.Database("people").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, newPerson)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, result)
}

func GetPeopleEndpoint(c echo.Context) error {
	var people []models.Person
	collection := db.Client.Database("people").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person models.Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, people)
}

func GetPersonEndpoint(c echo.Context) error {
	params := c.Param("id")
	id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var person models.Person
	collection := db.Client.Database("people").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, models.Person{ID: id}).Decode(&person)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, person)
}
