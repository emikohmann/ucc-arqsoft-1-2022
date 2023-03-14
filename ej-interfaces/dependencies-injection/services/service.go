package services

type Service interface {
	Get(key string) string
	Save(key string, value string)
}
