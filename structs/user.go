package structs

import (
	"encoding/json"
	"errors"
	"fmt"
	mongof "main/functions/mongo"
	"main/functions/security"
	"main/functions/snowflake"
	"os"
	"regexp"
	"strings"

	// "github.com/gorilla/securecookie"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	NAME_REGEX     = regexp.MustCompile("^[a-zA-Z0-9_]{3,16}$")
	PASSWORD_REGEX = regexp.MustCompile(`^.{8,32}$`)
	Key            = os.Getenv("ENCRYPTION_KEY")
)

//DIsplayName feature

type User struct {
	Name        string `json:"name"`
	Password    string `json:"password,omitempty"`
	ID          string `json:"id"`
	IP          string `json:"ip,omitempty"`
	Admin       bool   `json:"admin"`
	Banned      bool   `json:"banned"`
	BannedUntil int64  `json:"banned_until"`
}

func StripMapOfImportantInfo(u map[string]any) {
	delete(u, "ip")
	delete(u, "password")
	delete(u, "_id")
}
func (user *User) SetIP(ip string) {
	user.IP = strings.Split(ip, ":")[0]
}
func (user *User) Key() string {
	return os.Getenv("ENCRYPTION_KEY")
}
func (user *User) IsValidPass() bool {
	return PASSWORD_REGEX.MatchString(user.Password)
}
func (user *User) IsValidName() bool {
	return NAME_REGEX.MatchString(user.Name)
}
func (user *User) DecryptPassword(changeSTR bool) (string, error) {
	pass, err := security.Decrypt(user.Password, user.Key())
	if changeSTR {
		user.Password = pass
	}
	return pass, err
}

func (user *User) EncryptPassword(changeSTR bool) (string, error) {
	pass, err := security.Encrypt(user.Password, user.Key())
	if changeSTR {
		user.Password = pass
	}
	return pass, err
}
func (user *User) AccountExists(db *mongo.Database) bool {
	_, errs := user.EncryptPassword(true)
	// _, errsIP := user.EncryptIP(true)
	data, _ := mongof.FindOne(bson.M{
		"id":   strings.ToLower(user.ID),
		// "password": user.Password,
	}, options.FindOne(), db, "user")
	if data != nil {
		return true
	}
	if errs != nil /*|| errsIP !=nil*/ {
		return false
	}
	return false
}
func (user *User) IPExists(db *mongo.Database) bool {

	b, _ := mongof.FindOne(bson.M{
		"ip": user.IP,
	}, options.FindOne(), db, "user")
	// fmt.Println(b, user.IP)

	return len(b) > 0
}
func (user *User) NameExists(db *mongo.Database) bool {
	b, _ := mongof.FindOne(bson.M{
		"name": strings.ToLower(user.Name),
	}, options.FindOne(), db, "user")
	return len(b) > 0
}
func (user *User) FetchData(db *mongo.Database) error {
	_, errs := user.EncryptPassword(true)

	if errs != nil {
		return errs
	}
	data, errsS := mongof.FindOne(bson.M{
		"id": strings.ToLower(user.ID),
		//"password": user.Password,
	}, options.FindOne(), db, "user")
	if errsS != nil {
		return errsS
	}
	user.MapToUser(data)
	_, err := user.DecryptPassword(true)
	return err
}
func (user *User) Login(username, password string, db *mongo.Database) error {
	userM, _ := mongof.FindOne(bson.M{
		"name": strings.ToLower(username),
	}, options.FindOne(), db, "user")
	userMap := User{}
	userMap.MapToUser(userM)
	// fmt.Println(username, password, user)
	if userMap.Name == "" || username == "" {
		return NewErr("name is invalid", "invalid_name_1")
	}
	if password == "" {
		return NewErr("password is empty", "empty_password_2")
	}
	if _, err := userMap.DecryptPassword(true); err != nil {
		return NewErr("failed decrypting password", "depcr_failed_4")
	}
	if password != userMap.Password {
		return NewErr("password is not correct", "incorrect_password_2")
	}
	user.MapToUser(userM)
	return nil
}
func (user *User) RegisterAccount(username, password string, db *mongo.Database) error {
	user.Name = username
	user.Password = password

	// _, errs := user.EncryptPassword(true)
	// if errs != nil {
	// 	return errors.New("Failed Account Creation: password could not be encrypted")
	// }
	if user.IPExists(db) {
		return errors.New("1 account per IP : invalid_ip_3")
	}
	if user.NameExists(db) {
		return errors.New("name already exists : name_exists_1 ")
	}
	if !user.IsValidName() {
		return errors.New("name is invalid : invalid_name_1 ")
	}
	if !user.IsValidPass() {
		return errors.New("password is invalid : invalid_pass_2 ")
	}
	if !user.AccountExists(db) {
		user.ID = user.CreateID()
		if !user.ValidID() {
			return errors.New("invalid id generated : invalid_id_4")
		}
		user.Name = strings.ToLower(user.Name)
		_, err := mongof.InsertOne(user.ToMap(), options.InsertOne(), db, "user")
		if err != nil {
			return errors.New("account could not be created : failed_account_4")
		}

	} else {
		return NewErr("account already exists", "account_exists_4")
		// errors.New()
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

	user.ID = id
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
//Deprecated
func (user *User) ToMapCookie() map[string]any {
	user.Password = ""
	user.IP = ""
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
func NewErr(text, err string) error {
	return fmt.Errorf("%s : %s", text, err)
}
