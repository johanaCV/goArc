package repository

import "rideshareapp/model"
import "os"
import "fmt"

type CarFileRepositoryImp struct {
}

func (c CarFileRepositoryImp) SaveCar(car model.Car) model.Car {
	file, err := os.Create("cars.txt")
	if err == nil {
		file.WriteString(fmt.Sprintf("%#v", car)) 
		file.Close()
		return car
	}
	return model.Car{}
}

func (c CarFileRepositoryImp) FindCar(plate string) model.Car {
	if plate == "mnz123" {
		return model.Car{Plate: "mnz123", Brand: "kia", Model: "rio", Color: "blanco" }
	}
	return model.Car{}
}