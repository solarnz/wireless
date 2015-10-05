package main

import (
	"fmt"

	"github.com/solarnz/wireless"
)

func main() {
	interfaces, err := wireless.NetworkInterfaces()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(interfaces)
}
