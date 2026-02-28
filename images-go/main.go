package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"image/png"
	_ "image/png"
)

func load(filePath string) (*image.NRGBA, string) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		log.Println("cannot read file: ", err)
	}
	defer imgFile.Close()

	img, format, err := image.Decode(imgFile)

	if err != nil {
		log.Println("cannot decode file: ", err)
	}

	return img.(*image.NRGBA), format
}

func save(filePath string, img *image.NRGBA) error {
	imgFile, err := os.Create(filePath)
	if err != nil {
		log.Println("cannot create file: ", err)
	}
	defer imgFile.Close()
	
	return png.Encode(imgFile, img.SubImage(image.Rect(1000,1000,2000,2000)))
}

func main() {
	img, format := load("image.png")

	bounds := img.Bounds()
	fmt.Println("width:", bounds.Dx())
	fmt.Println("height:", bounds.Dy())
	fmt.Println("format:", format)

	if	err := save("cropped.png", img); err != nil {
		log.Println("cannot save image:", err)
	}


}
