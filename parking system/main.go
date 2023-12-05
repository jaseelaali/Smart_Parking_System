package main

import (
	"fmt"

	"github.com/jaseelaali/Smart_Parking_System/config"
	"github.com/jaseelaali/Smart_Parking_System/dep"
)

func mainI() {
	// LOAD CONNFIG FILE
	config, err := config.LoadConfig()
	 fmt.Println(config, ":", err)
	if err != nil {
		fmt.Println("error in here is 1:", err)
	}
	//INTITIALIZE API
	server,errr := dep.InitializeApi(config)
	if errr != nil{
		fmt.Println("error is here 2:",errr)
	}
}
