package main

import (
	"fmt"
	app1 "gwi/app1/app"
	app2 "gwi/app2/app"
)

func main() {
	fmt.Println(app1.WhoAmI())
	fmt.Println(app2.WhoAmI())
	fmt.Println(app1.WhoIsApp2())
}
