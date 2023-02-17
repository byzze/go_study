package main

import (
	"fmt"

	"github.com/disiqueira/gotree"
)

func main() {
	artist := gotree.New("Pantera")
	album := artist.Add("Far Beyond Driven")
	album.Add("5 minutes Alone")

	fmt.Println(artist.Print())
}

