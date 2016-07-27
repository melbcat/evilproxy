/* EvilProxy*/

package main

import (
	"evilproxy/core"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Only a new GO EVILPROXY for cSploit!")
}
