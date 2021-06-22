package main

import (
	"fmt"
	"image/gif"
	"os"
)

func main() {
	originalGif := &gif.GIF{}
	reversedGif := &gif.GIF{}
	fmt.Println("Starting...")
	f, err := os.Open("image1.gif")
	if err != nil {
		fmt.Printf("Error")
	}
	originalGif, _ = gif.DecodeAll(f)

	images := originalGif.Image
	delay := originalGif.Delay
	for i := len(images) - 1; i > 0; i-- {
		reversedGif.Image = append(reversedGif.Image, images[i])
		reversedGif.Delay = append(reversedGif.Delay, delay[i])
	}
	// for i, image := range reversedGif.Image {
	// 	iamgeName := strconv.Itoa(i) + ".gif"
	// 	f, _ := os.Create(iamgeName)
	// 	opt := gif.Options{
	// 		NumColors: 256,
	// 	}
	// 	gif.Encode(f, image, &opt)
	// }
	f, _ = os.OpenFile("reversed.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, reversedGif)
}
