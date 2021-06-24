package main

import (
	"flag"
	"fmt"
	"image/gif"
	"os"
	"path/filepath"
)

func main() {
	fileName := flag.String("file", "image1.gif", "Name of gif file")
	flag.Parse()
	fileExtension := filepath.Ext(*fileName)
	if fileExtension != ".gif" {
		fmt.Println("File extension must be .gif")
		return
	}
	originalGif := &gif.GIF{}
	reversedGif := &gif.GIF{}
	fmt.Println("Starting...")
	f, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("Error")
		return
	}
	originalGif, _ = gif.DecodeAll(f)

	images := originalGif.Image
	delay := originalGif.Delay
	for i := len(images) - 1; i > 0; i-- {
		reversedGif.Image = append(reversedGif.Image, images[i])
		reversedGif.Delay = append(reversedGif.Delay, delay[i])
	}
	
	f, _ = os.OpenFile("reversed.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, reversedGif)
	fmt.Printf("Reversed GIF has been generated\n")
}
