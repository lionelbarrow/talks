package interfaces

import "encoding/json"

type JSONer interface {
	ToJSON() []byte
}

type User struct {
	Username string
	password string
}

func (u User) ToJSON() []byte {
	bytes, _ := json.Marshal(u)
	return bytes
}

type Post struct {
	Title string
	Body  string
}

func (p Post) ToJSON() []byte {
	bytes, _ := json.Marshal(p)
	return bytes
}

func StoreInDB(j JSONer) {
	bytes := j.ToJSON
	// ...
}
