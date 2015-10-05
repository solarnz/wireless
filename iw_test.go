package wireless_test

import (
	"fmt"

	"github.com/solarnz/wireless"
)

func ExampleNetworkInterfaces() {
	interfaces, _ := wireless.NetworkInterfaces()
	fmt.Println(interfaces)
	// Output: [wl0]
}
