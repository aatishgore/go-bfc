package controller

import (
	service "BFC/modules/event/service"
	"fmt"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in side a event controller")
	service.GetEventList()
}
