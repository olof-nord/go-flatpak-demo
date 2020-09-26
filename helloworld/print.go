package helloworld

import (
	"fmt"
)

func Hello() string {
    return "Hello World"
}

func Say() {
    fmt.Println(Hello())
}
