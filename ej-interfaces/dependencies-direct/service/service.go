package service

import "fmt"

var data = make(map[string]string)

func Get(key string) string {
	fmt.Println("SELECT in DB")
	return data[key]
}

func Save(key string, value string) {
	fmt.Println("INSERT in DB")
	data[key] = value
}
