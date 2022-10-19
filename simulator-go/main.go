package main

import (
	"fmt"
	"simulator/application/route"
)

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}
	err := route.LoadPositions()
	if err != nil {
		fmt.Println(err.Error())
	}
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson)
}
