package main

import "fmt"

func main() {
	values := [2]string{"hi there!", "hello"}
	for _, v := range values {
		var data Result = Ok{v}

		var r = Work(data)
		switch r.(type) {
		case Ok:
			fmt.Println("checking value:", v, " result: ", r.(Ok).value)
		case Err:
			fmt.Println("checking value:", v, " result: ", r.(Err).err)
		}
	}
}
