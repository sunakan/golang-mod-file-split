package main

// NG "./hoge"
// OK "module名/hoge"
import (
	"fmt"

	"golang-mod-file-split/hoge"
)

func main() {
	fmt.Println("Hello")
	user := hoge.NewUser("suna")
	fmt.Println(user)
}
