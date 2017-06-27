// +build xwin

package Wins



import (
	//"fmt"
	//"goocto.com/grafx"

	//"github.com/BurntSushi/xgb"
	//"github.com/BurntSushi/xgb/xinerama"
	//"github.com/BurntSushi/xgb/xproto"

	"image"
	"image/color"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"

	"github.com/BurntSushi/xgbutil/xwindow"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/icccm"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/mousebind"
)






var X *xgbutil.XUtil




func CreateWindow(title string, width int, height int, quit bool) *xwindow.Window {

	// output it to X Windows
	X, _ = xgbutil.NewConn()

	win, err := xwindow.Generate(X)
	if err != nil {
		xgbutil.Logger.Printf("Could not generate new window id: %s", err)
		return nil
	}

	// Create a very simple window
	win.Create(X.RootWin(), 0, 0, width, height, 0)

	// Make this window close gracefully.
	win.WMGracefulClose(func(w *xwindow.Window) {
		xevent.Detach(w.X, w.Id)
		keybind.Detach(w.X, w.Id)
		mousebind.Detach(w.X, w.Id)
		w.Destroy()

		if quit {
			xevent.Quit(w.X)
		}
	})

	// Set WM_STATE so it is interpreted as a top-level window.
	err = icccm.WmStateSet(X, win.Id, &icccm.WmState{
		State: icccm.StateNormal,
	})
	if err != nil { // not a fatal error
		xgbutil.Logger.Printf("Could not set WM_STATE: %s", err)
	}

	// Set _NET_WM_NAME so it looks nice.
	err = ewmh.WmNameSet(X, win.Id, title)
	if err != nil { // not a fatal error
		xgbutil.Logger.Printf("Could not set _NET_WM_NAME: %s", err)
	}

//	//im *xgraphics.Image
//	// a single pixel background color
//	bytes := []uint8{0x99,0x99,0xff,0xff}
////	bytes[0] = 0x99 // red
////	bytes[1] = 0x99 // green
////	bytes[2] = 0xff // blue
////	bytes[3] = 0xff // reserved
//	bg, err := xgraphics.NewBytes(X,bytes)
//	if err != nil {
//		xgbutil.Logger.Printf("Could not set New Bytes: %s", err)
//	}




	c := color.RGBA{0xff,0xff,0xff,0xff}
	bg := image.NewRGBA(image.Rect(0,0,1,1))
	bg.Set(0,0,c)

	// Paint our bgimage before mapping.
	bgx := xgraphics.NewConvert(X,bg)
	bgx.XSurfaceSet(win.Id)
	bgx.XDraw()
	bgx.XPaint(win.Id)

	// let the window manager take over
	win.Map()

	return win
}


func ExecWindow() {
	xevent.Main(X)
}







//// This is a slightly modified version of xgraphics.XShowExtra that does
//// not set any resize constraints on the window (so that it can go
//// fullscreen).
//func showImage(im *xgraphics.Image, name string, quit bool) *xwindow.Window {
//
//	if len(name) == 0 {
//		name = "xgbutil Image Window"
//	}
//	w, h := im.Rect.Dx(), im.Rect.Dy()
//
//	win, err := xwindow.Generate(im.X)
//	if err != nil {
//		xgbutil.Logger.Printf("Could not generate new window id: %s", err)
//		return nil
//	}
//
//	// Create a very simple window with dimensions equal to the image.
//	win.Create(im.X.RootWin(), 0, 0, w, h, 0)
//
//	// Make this window close gracefully.
//	win.WMGracefulClose(func(w *xwindow.Window) {
//		xevent.Detach(w.X, w.Id)
//		keybind.Detach(w.X, w.Id)
//		mousebind.Detach(w.X, w.Id)
//		w.Destroy()
//
//		if quit {
//			xevent.Quit(w.X)
//		}
//	})
//
//	// Set WM_STATE so it is interpreted as a top-level window.
//	err = icccm.WmStateSet(im.X, win.Id, &icccm.WmState{
//		State: icccm.StateNormal,
//	})
//	if err != nil { // not a fatal error
//		xgbutil.Logger.Printf("Could not set WM_STATE: %s", err)
//	}
//
//	// Set _NET_WM_NAME so it looks nice.
//	err = ewmh.WmNameSet(im.X, win.Id, name)
//	if err != nil { // not a fatal error
//		xgbutil.Logger.Printf("Could not set _NET_WM_NAME: %s", err)
//	}
//
//	// Paint our image before mapping.
//	im.XSurfaceSet(win.Id)
//	im.XDraw()
//	im.XPaint(win.Id)
//
//	// Now we can map, since we've set all our properties.
//	// (The initial map is when the window manager starts managing.)
//	win.Map()
//
//	return win
//}
