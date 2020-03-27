package main

import (
	"fmt"

	xLog "pgxs.io/chassis/log"
)

var log = xLog.New().Category("cmd").Component("ngx")

func main() {
	printLogo()

	log.Info("ngx run")
}

func printLogo() {

	var logo = `
	 _ __   __ _ _ __
	| '_ \ / _` + "`" + ` | '_ \
	| | | | (_| | |_) |
	|_| |_|\__,_| .__/
                    |_|`
	fmt.Println(logo)
}
