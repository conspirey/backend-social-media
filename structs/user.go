package structs

import (
	"encoding/json"
	"fmt"
	"main/functions/snowflake"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	ID string `json:"id"`
}
func (user *User) FetchData(username, password string, db *mongo.Database) (error) {
	return nil
}
func (user *User) RegisterAccount(username, password string, db *mongo.Database) (error) {
	return nil
}
func (user User) ValidID() bool {
	return user.ID != ""
}
func (user *User) CreateID() string {
	snowflake.NewNode(1)
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	id := node.Generate().String()
	
	user.ID = id;
	return id
}
func (user *User) MapToUser(umap map[string]any) {
	by, err := json.Marshal(umap)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(by, &user); err != nil {
		panic(err)
	}
}
func (user *User) ToMap() map[string]any {
	by, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	UMap := map[string]any{}
	if err := json.Unmarshal(by, &UMap); err != nil {
		panic(err)
	}
	return UMap
}