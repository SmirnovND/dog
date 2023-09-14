package printer

import (
	"example.com/m/formater"
	"example.com/m/logger"
	"fmt"
)

type DogPrinterInterfaces interface {
	GetName() string
	GetColor() string
}

// хотел бы тут принимать dog DogPrinterInterfaces и не зависеть от вложенного стека
func PrintDogProfile(dog logger.DogLoggerInterfaces) {
	fmt.Println(formater.GetNameFormat(dog), formater.GetColorFormat(dog))
	logger.WriteLog(dog)
}
