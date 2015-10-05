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

	for _, iface := range interfaces {
		c, err := iface.BasicConfig()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%+v\n", c)
	}
}
