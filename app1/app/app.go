package app

import "gwi/common"

func WhoAmI() string {
	return "I'am APP1"
}

func WhoIsApp2() string {
	return "App2 said: " + common.WhoIsApp2()
}
