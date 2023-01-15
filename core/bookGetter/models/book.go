package models

type Book struct {
	ID     interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string      `json:"title" bson:"title"`
	Author string      `json:"author" bson:"author"`
}

type BookUpdate struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        Book
}

type BookDelete struct {
	DeletedCount int64 `json:"deletedCount"`
}
