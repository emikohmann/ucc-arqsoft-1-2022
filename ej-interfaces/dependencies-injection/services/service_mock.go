package services

import "fmt"

type ServiceMock struct{}

func (ServiceMock) Get(key string) string {
	fmt.Println("MOCK GET")
	return "Mocked string"
}

func (ServiceMock) Save(key string, value string) {
	fmt.Println("MOCK SAVE")
	fmt.Println("Mocked save")
}
