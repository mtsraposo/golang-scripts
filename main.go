package main

import (
	"fmt"
	. "github.com/mtsraposo/xp-strats/mathematics"
)

func main() {
	fmt.Printf("P(last ball removed is blue) = %f%%", ProbLastRemovedIsBlue(54, 18))
}
