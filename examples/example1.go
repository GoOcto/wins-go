package main

import (
	"fmt"
//	"goocto.com/grafx"
//
//	"github.com/BurntSushi/xgbutil"
//	"github.com/BurntSushi/xgbutil/xevent"
//	"github.com/BurntSushi/xgbutil/xgraphics"
//
//	"github.com/BurntSushi/xgbutil/xwindow"
//	"github.com/BurntSushi/xgbutil/ewmh"
//	"github.com/BurntSushi/xgbutil/icccm"
//	"github.com/BurntSushi/xgbutil/keybind"
//	"github.com/BurntSushi/xgbutil/mousebind"
	Wins "github.com/GoOcto/wins-go"
)






func main() {

	fmt.Println("About to go Stellar!!")

	//img := grafx.First()

	Wins.Init()
	Wins.CreateWindow("My Go Window",400,400,true)

	// output it to X Windows
	//X, _ := xgbutil.NewConn()

	// Now convert it into an X image.
	//ximg := xgraphics.NewConvert(X, img)

	//ximg.XShowExtra("Window", true)  // basic default win behaviour
	//showImage( ximg, "Window", true )  // window with more features

	//xevent.Main(X)

	Wins.ExecWindow()
}















