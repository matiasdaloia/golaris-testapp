package main

import (
	"github.com/matiasdaloia/golaris"
)

func main() {
	g := golaris.Golaris{
		AppName: "TestApp",
		Debug:   true,
		Version: "1",
	}
	g.New("./")
}
