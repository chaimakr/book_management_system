package database

import (
	"context"
	"encoding/json"

	"github.com/chaimakr/book_management_system/core/setter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookInterface interface {
	Insert(models.Book) (models.Book, error)
	Update(string, interface{}) (models.BookUpdate, error)
	Delete(string) (models.BookDelete, error)
	Get(string) (models.Book, error)
	Search(interface{}) ([]models.Book, error)
}

type BookClient struct {
	Ctx context.Context
	Col *mongo.Collection
}

func (c *BookClient) Insert(docs models.Book) (models.Book, error) {
	book := models.Book{}

	res, err := c.Col.InsertOne(c.Ctx, docs)
	if err != nil {
		return book, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return c.Get(id)
}
func (c *BookClient) Update(id string, update interface{}) (models.BookUpdate, error) {
	result := models.BookUpdate{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	book, err := c.Get(id)
	if err != nil {
		return result, err
	}
	var exist map[string]interface{}
	b, err := json.Marshal(book)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &exist)

	change := update.(map[string]interface{})
	for k := range change {
		if change[k] == exist[k] {
			delete(change, k)
		}
	}

	if len(change) == 0 {
		return result, nil
	}

	res, err := c.Col.UpdateOne(c.Ctx, bson.M{"_id": _id}, bson.M{"$set": change})
	if err != nil {
		return result, err
	}

	newBook, err := c.Get(id)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = res.ModifiedCount
	result.Result = newBook
	return result, nil
}
func (c *BookClient) Delete(id string) (models.BookDelete, error) {
	result := models.BookDelete{
		DeletedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	res, err := c.Col.DeleteOne(c.Ctx, bson.M{"_id": _id})
	if err != nil {
		return result, err
	}
	result.DeletedCount = res.DeletedCount
	return result, nil
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
