package main

import "rideshareapp/controller"
import "net/http"

func main() {
    http.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
        controller.GetCar(res, req)
	})
	http.ListenAndServe(":9090", nil)
}