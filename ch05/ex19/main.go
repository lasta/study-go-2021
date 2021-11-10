package main

import "fmt"

func p() (str string){
	defer func() {
		if p := recover(); p != nil {
			str = fmt.Sprintf("%v", p)
		}
	}()
	panic("panic!")
}
