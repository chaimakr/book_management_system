package database

import (
	"context"
	//"encoding/json"
	"github.com/chaimakr/book_management_system/core/getter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookInterface interface {
	// Insert(models.Book) (models.Book, error)
	// Update(string, interface{}) (models.BookUpdate, error)
	// Delete(string) (models.BookDelete, error)
	Get(string) (models.Book, error)
	Search(interface{}) ([]models.Book, error)
}

type BookClient struct {
	Ctx context.Context
	Col *mongo.Collection
}

func (c *BookClient) Get(id string) (models.Book, error) {
	book := models.Book{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return book, err
	}

	err = c.Col.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&book)
	if err != nil {
		return book, err
	}

	return book, nil
}
func (c *BookClient) Search(filter interface{}) ([]models.Book, error) {
	books := []models.Book{}
	if filter == nil {
		filter = bson.M{}
	}

	cursor, err := c.Col.Find(c.Ctx, filter)
	if err != nil {
		return books, err
	}

	for cursor.Next(c.Ctx) {
		row := models.Book{}
		cursor.Decode(&row)
		books = append(books, row)
	}

	return books, nil
}
