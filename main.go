package main

import (
	"fmt"
	"os"

	"ely.by/sessionserver/internal/cmd"
)

func main() {
	err := cmd.Serve()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
