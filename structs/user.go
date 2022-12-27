package structs

import (
	"encoding/json"
	"errors"
	"fmt"
	"main/functions/mongo"
	"main/functions/security"
	"main/functions/snowflake"
	"regexp"
	"strings"

	// "github.com/gorilla/securecookie"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const (
	/*
	jX4:k357Q-,XK=!:hf4,r8RpiXVBvXzL
	jX4:k357Q-,XK=!:hf4,r8RpiXVBvXzL
	jX4:k357Q-,XK=!:hf4,r8RpiXVBvXzL
	jX4:k357Q-,XK=!:hf4,r8RpiXVBvXzL
	jX4:k357Q-,XK=!:hf4,r8RpiXVBvXzL
	*/
	key = "jX4:k357Q-,XK=!:hf4,r8RpiXVBvXzL"
	
)
var (
	NAME_REGEX = regexp.MustCompile("^[a-zA-Z0-9_]{3,16}$")
	PASSWORD_REGEX = regexp.MustCompile(`^.{8,32}$`)
)
//DIsplayName feature

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	ID string `json:"id"`
}
func (user *User) Key() string {
	return key
}
func (user *User) IsValidPass() bool {
	return PASSWORD_REGEX.MatchString(user.Password) 
}
func (user *User) IsValidName() bool {
	return NAME_REGEX.MatchString(user.Name) 
}
func (user *User) DecryptPassword(changeSTR bool) (string, error) {
	pass, err := security.Decrypt(user.Password, key)
	if changeSTR {
		user.Password = pass
	}
	return pass, err
}
func (user *User) EncryptPassword(changeSTR bool) (string, error) {
	pass, err := security.Encrypt(user.Password, key)
	if changeSTR {
		user.Password = pass
	}
	return pass, err
}
func (user *User) AccountExists(db *mongo.Database) (bool) {
	_, errs := user.EncryptPassword(true)

	data, _ := mongof.FindOne(bson.M{
		"name": strings.ToLower(user.Name),
		"password": user.Password,
	}, options.FindOne(), db, "user")
	if data != nil {
		return true
	}
	if errs != nil {
		return false
	}
	return false
}
func (user *User) FetchData(db *mongo.Database) (error) {
	_, errs := user.EncryptPassword(true)
	if errs != nil {
		return errs
	}
	data, errsS := mongof.FindOne(bson.M{
		"name": strings.ToLower(user.Name),
		"password": user.Password,
	}, options.FindOne(), db, "user")
	if errsS != nil {
		return errsS
	}
	user.MapToUser(data)
	_, err := user.DecryptPassword(true)
	return err
}
func (user *User) RegisterAccount(username, password string, db *mongo.Database) (error) {
	user.Name = username
	user.Password = password
	_, errs := user.EncryptPassword(true)
	if errs != nil {
		fmt.Println(errs)
		return errors.New("Failed Account Creation: password could not be encrypted")
	}
	if !user.AccountExists(db) {
		user.ID = user.CreateID()
		if !user.ValidID() {
			return errors.New("Failed Account Creation: Invalid id generated")
		}
		user.Name = strings.ToLower(user.Name)
		_, err := mongof.InsertOne(user.ToMap(), options.InsertOne(), db, "user")
		if err != nil {
			return errors.New("Failed Account Creation: Account could not be created")
		}
		
	} else {
		return errors.New("Failed Account Creation: Account already exists")
	}
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
