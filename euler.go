package euler

import (
	"fmt"
	"time"
)

func main() {

	go ProjectOne(1000)
	go ProjectTwo(4000000)
	go ProjectThree(600851475143)
	go ProjectFour()
	go ProjectFive(1, 20)
	go ProjectSix(100)
	go ProjectSeven(10001)

	// do something to block waiting
	time.Sleep(5 * time.Second)
	fmt.Println("done")
}
