package controller

import "rideshareapp/service"
import "net/http"
import "fmt"

type CarController struct {
	carService service.CarService
}

var carController = new(CarController)

func GetCar(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hi there, I love %s!", carController.carService.FindCar(request.URL.Path[1:]))
}