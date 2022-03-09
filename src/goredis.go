package goredis2

import "fmt"

func SayHi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}
