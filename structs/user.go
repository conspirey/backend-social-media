package structs

import "encoding/json"

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
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