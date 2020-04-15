package Server

import (
	"fmt"
	"strings"
)

func CheckInfo(s string){
	info:=strings.Split(s," ")
	switch tmp := info[0]; tmp {
	case "1":
		loginRequest(info[1:])
	case "2":
		updateRequest(info[2:])
	}
}

func loginRequest(info []string){
	fmt.Printf("This is a login request\n")
}

func updateRequest(info []string){
	fmt.Printf("This is a update request\n")
}