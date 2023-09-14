package formater

type DogNameInterfaces interface {
	GetName() string
}

func GetNameFormat(dogName DogNameInterfaces) string {
	return "Name is:" + dogName.GetName()
}

type DogColorInterfaces interface {
	GetColor() string
}

func GetColorFormat(dogName DogColorInterfaces) string {
	return "Color is:" + dogName.GetColor()
}
