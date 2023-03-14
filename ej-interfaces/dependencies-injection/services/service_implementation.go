package services

import "fmt"

var data = make(map[string]string)

type ServiceImplementation struct{}

func (ServiceImplementation) Get(key string) string {
	fmt.Println("SELECT in DB")
	return data[key]
}

func (ServiceImplementation) Save(key string, value string) {
	fmt.Println("INSERT in DB")
	data[key] = value
}
