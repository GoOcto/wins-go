package main

import (
	"fmt"
	Wins "github.com/GoOcto/wins-go"

	"os"
	"image"
	_ "image/jpeg"
)






func main() {

	fmt.Println("About to go Stellar!!")


	// Load an image
	pwd, _ := os.Getwd()
	fImg, err := os.Open(pwd+"/Koala.jpg")
	if err!=nil {
		fmt.Println("Open error",err)
	}
	defer fImg.Close()

	img, str, err := image.Decode(fImg)
	if err!=nil {
		fmt.Println("Decode error",err)
	}
	if str!="" {
		fmt.Println("Image format ",str)
	}

	rect := img.Bounds()
	wid := rect.Max.X
	hgt := rect.Max.Y


	// create a Window and use it to show the image
	Wins.Init()
	win := Wins.CreateWindow("My Go Window",wid,hgt,true)
	//_ = win
	Wins.FillWindow(win,img)
	Wins.ExecMain()
}
