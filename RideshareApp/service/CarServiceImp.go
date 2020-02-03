package service

import "rideshareapp/model"
import "rideshareapp/repository"

type CarServiceImp struct {
	carRepository repository.CarRepository
}

var carServiceImp = new(CarServiceImp)

func (carService CarServiceImp) FindCar(plate string) model.Car {
	return carServiceImp.carRepository.FindCar(plate)
}