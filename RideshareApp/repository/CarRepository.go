package repository

import "rideshareapp/model"

type CarRepository interface {

	SaveCar(model.Car) model.Car
	FindCar(string) model.Car
	DeleteCar(string)

}