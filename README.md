# silly-golang-monad

*Disclosure: I don't really know Go expert, I was just experimenting!*

This is the `main.go` module:

```go
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

```

In the `Work()` function, there is a nasty trap that 
causes a division-by-zero based on the length of the input. This is
intentional to demonstrate the short-circuting of the `Err` variant
of the `Result` type.

The value `"hi there!"` will succeed, but `"hello"` will fail.
This is the output when run:

```
checking value: hi there!  result:  9
checking value: hello  result:  runtime error: integer divide by zero
```
