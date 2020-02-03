package service

import "rideshareapp/model"

type CarService interface {
	FindCar(string) model.Car
}