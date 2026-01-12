package main

import (
	"fmt"
	app2 "gwi/app2/app"
)

func WhoAmI() string {
	return "I'am APP1"
}

func WhoIsApp2() string {
	return "App2 said: " + app2.WhoAmI()
}

func main() {
	fmt.Println(WhoAmI())
	fmt.Println(app2.WhoAmI())
}
