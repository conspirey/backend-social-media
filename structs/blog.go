package structs

import (
	"fmt"
	"main/functions/snowflake"
	"time"
)

type Blog struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Author    *Author   `json:"author"`
	Tags      []string  `json:"tags"`
}
type Author struct {
	Name string `json:"name"`
	ID   string `json:"id"`
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
