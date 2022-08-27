package core

import "fmt"

func SayHello(name string) (output string) {
	output = fmt.Sprintf("Hello %s! You are awesome", name)
	return
}
