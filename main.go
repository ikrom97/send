package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type UserDTO struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
func RequestType(user UserDTO) {
	client := &http.Client{}
	URL := "http://127.0.0.1:8888/api/Sign-in"
	marshal, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("cannot parse")
	}
	r := bytes.NewReader(marshal)
	newRequest, err := http.NewRequest("POST", URL, r)
	if err != nil {
		log.Println("can't create request", err)
		return
	}
	res, err := client.Do(newRequest)
	if err != nil {
		fmt.Println("Cant send")
		return
	}
	log.Printf("", res.StatusCode)
}
func main() {
	for {
		time.Sleep(time.Second *1)
		user := UserDTO{
			Name:     "1",
			Surname:  "1",
			Age:      1,
			Gender:   "1",
			Login:    "1",
			Password: "1",
		}
		RequestType(user)
	}
}
