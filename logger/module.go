package logger

import "fmt"

// приходится реализовывать в интерфейсе не используемые методы GetColor GetName
// т.к. выше по стеку они используются а принимать свой интерфейс выше нельзя, т.к. в нем нет нужных методов для логера
type DogLoggerInterfaces interface {
	GetName() string
	GetColor() string
	GetOwner() string
	GetPassportId() int
}

func WriteLog(dog DogLoggerInterfaces) {
	fmt.Println("Log: ", dog.GetPassportId(), dog.GetOwner())
}
