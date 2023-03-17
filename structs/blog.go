package structs

import (
	"encoding/json"
	"fmt"
	mongof "main/functions/mongo"
	"main/functions/snowflake"
	

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Blog struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	Author    *Author   `json:"author"`
	Tags      []string  `json:"tags"`
	Description string `json:"description"`
}
type Author struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func GetBlogs(filter any, option *options.FindOptions, db *mongo.Database) ([]primitive.M, error) {
	data, err := mongof.Find(option, filter, db, "blog")
	if err != nil {
		return []primitive.M{}, err
	}
	return data, nil
}
func (blog *Blog) Create(db *mongo.Database) (error) {
	_, err := mongof.InsertOne(blog, options.InsertOne(), db, "blog")
	return err
}
func (blog Blog) Exists(db *mongo.Database) (exists bool) {
	data, err := mongof.FindOne(bson.M{
		"id": blog.ID,
	}, options.FindOne(), db, "blog")
	if err != nil {
		return false
	}
	if data["id"] != nil {
		return true
	}
	return false
}
func (blog Blog) ValidID() bool {
	return blog.ID != ""
}

func (blog *Blog) GenerateID() string {
	snowflake.NewNode(1)
	node, err := snowflake.NewNode(10)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	id := node.Generate().String()

	blog.ID = id
	return id
}
func (blog *Blog) MapToBlog(umap map[string]any) {
	by, err := json.Marshal(umap)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(by, &blog); err != nil {
		panic(err)
	}
}
func (blog *Blog) ToMap() map[string]any {
	by, err := json.Marshal(blog)
	if err != nil {
		panic(err)
	}
	UMap := map[string]any{}
	if err := json.Unmarshal(by, &UMap); err != nil {
		panic(err)
	}
	return UMap
}
