package main

import (
	"example.com/m/printer"
)

type Dog struct {
	Name       string
	Color      string
	PassportId int
	Owner      string
}

func (d *Dog) GetName() string {
	return d.Name
}

func (d *Dog) GetColor() string {
	return d.Color
}

func (d *Dog) GetOwner() string {
	return d.Owner
}

func (d *Dog) GetPassportId() int {
	return d.PassportId
}

func main() {
	printer.PrintDogProfile(&Dog{Name: "John", Color: "White"})
}
