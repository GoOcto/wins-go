// +build win32

package Wins

import (
	"github.com/AllenDang/w32"
	"syscall"
	"unsafe"
	"image"
	_ "image/jpeg"
	"image/draw"
	//"image/color"
	"fmt"
)




var hdcmem w32.HDC
var pixels unsafe.Pointer
var hbmp w32.HBITMAP

var wid int
var hgt int




func WndProc(hWnd w32.HWND, msg uint32, wParam, lParam uintptr) (uintptr) {

	switch msg {

//		case w32.WM_SIZE:
//			//var (*[4]bytes)by = lParam
//			wid = int( w32.LOWORD(uint32(lParam)) )
//			hgt = int( w32.HIWORD(uint32(lParam)) )

		case w32.WM_PAINT:
			if  hdcmem>0  {
				var ps w32.PAINTSTRUCT
				hdc := w32.BeginPaint(hWnd,&ps)

				hbmold := w32.SelectObject( hdcmem, w32.HGDIOBJ(hbmp) )
				w32.BitBlt(hdc, 0, 0, wid, hgt, hdcmem, 0, 0, w32.SRCCOPY )
				w32.SelectObject( hdcmem, hbmold )

				w32.EndPaint(hWnd,&ps)
			}

		case w32.WM_DESTROY:
			w32.PostQuitMessage(0)

		default:
			return w32.DefWindowProc(hWnd, msg, wParam, lParam)
	}

	return 0
}






func Init() {
}



func CreateWindow(title string, width int, height int, quit bool) w32.HWND {
	//_ = name
	//_ = width
	//_ = height
	_ = quit

	hInstance := w32.GetModuleHandle("")
	lpszClassName := syscall.StringToUTF16Ptr("WNDclass")

	var wcex w32.WNDCLASSEX
	wcex.Size       = uint32(unsafe.Sizeof(wcex))
	wcex.Style      = w32.CS_HREDRAW | w32.CS_VREDRAW
	wcex.WndProc    = syscall.NewCallback(WndProc)
	wcex.ClsExtra   = 0
	wcex.WndExtra   = 0
	wcex.Instance   = hInstance
	wcex.Icon       = w32.LoadIcon(hInstance, w32.MakeIntResource(w32.IDI_APPLICATION))
	wcex.Cursor     = w32.LoadCursor(0, w32.MakeIntResource(w32.IDC_ARROW))
	wcex.Background = w32.COLOR_WINDOW + 1
	wcex.MenuName   = nil
	wcex.ClassName  = lpszClassName
	wcex.IconSm     = w32.LoadIcon(hInstance, w32.MakeIntResource(w32.IDI_APPLICATION))

	w32.RegisterClassEx(&wcex)

	hWnd := w32.CreateWindowEx(
						0, lpszClassName, syscall.StringToUTF16Ptr(title), 
						w32.WS_OVERLAPPEDWINDOW | w32.WS_VISIBLE, 
						w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, width, height, 0, 0, hInstance, nil)

	//LoadWin(hWnd)

	//w32.ShowWindow( hWnd, w32.SW_SHOW ) // implied in CreateWindowEx

	//MakeSurface(hWnd)

	w32.ShowWindow(hWnd, w32.SW_SHOWDEFAULT)
	w32.UpdateWindow(hWnd)

	return hWnd
}



func ExecMain() {

	var msg w32.MSG
	for {
		if w32.GetMessage(&msg, 0, 0, 0) == 0 {
			break
		}
		w32.TranslateMessage(&msg)
		w32.DispatchMessage(&msg)
	}
	//w32.DeleteObject( w32.HGDIOBJ(hbmp) )
	//return msg.WParam
}




func MakeSurface(hWnd w32.HWND, wid int, hgt int, pixels *unsafe.Pointer) w32.HBITMAP {

	var bmi w32.BITMAPINFO
	bmi.BmiHeader.BiSize = uint32(unsafe.Sizeof(bmi.BmiHeader)) // 40
	bmi.BmiHeader.BiWidth  =  int32(wid)
	bmi.BmiHeader.BiHeight = -int32(hgt)
	bmi.BmiHeader.BiPlanes = 1
	bmi.BmiHeader.BiBitCount = 32
	bmi.BmiHeader.BiCompression = w32.BI_RGB

	hdc := w32.GetDC(hWnd)
	hbmp := w32.CreateDIBSection( hdc, &bmi, w32.DIB_RGB_COLORS, pixels, w32.HANDLE(0), 0 )
	w32.DeleteDC(hdc)

	fmt.Println("Pixels",*pixels)

	//go tickProc(hWnd)
	return hbmp
}




func FillWindow(hwnd w32.HWND, img image.Image) {

	rect := img.Bounds()
	wid = rect.Max.X
	hgt = rect.Max.Y


	//fmt.Println("Pixels",pixels)

	hdc := w32.GetDC( hwnd )
	hdcmem = w32.CreateCompatibleDC( hdc )

	hbmp = MakeSurface(hwnd,wid,hgt,&pixels)
	fmt.Println("Pixels",pixels)




	rgba := image.NewRGBA(rect)
	draw.Draw( rgba, rect, img, rect.Min, draw.Src )


	alt := 0 // for some reason alt = 1 causes Window to lock up

	for x:=0;x<wid;x++ {
		for y:=0;y<hgt;y++ {

			idx := 4*(y*wid + x)

			dstptr := unsafe.Pointer( uintptr(pixels) + uintptr(idx) )
			P := (*[4]uint8)( dstptr )

			if alt==0 {
				rgbapx := unsafe.Pointer(&rgba.Pix[0])
				srcptr := unsafe.Pointer( uintptr(rgbapx) + uintptr(idx) )
				R := (*[4]uint8)( srcptr )

				P[0] = R[2]
				P[1] = R[1]
				P[2] = R[0]
			} else {
				r,g,b,_ := img.At(x,y).RGBA()
				//_,_,_ = r,g,b
				//img.At(0,0)
				P[2] = uint8(r>>8)
				P[1] = uint8(g>>8)
				P[0] = uint8(b>>8)
			}

		}
	}





//	wrct := new(w32.RECT)
//	wrct.Right  = int32(wid)
//	wrct.Bottom = int32(hgt)
//	w32.InvalidateRect(hwnd,wrct,false)

	hbmold := w32.SelectObject( hdcmem, w32.HGDIOBJ(hbmp) )
	w32.BitBlt(hdc, 0, 0, wid, hgt, hdcmem, 0, 0, w32.SRCCOPY )
	w32.SelectObject( hdcmem, hbmold )


	// this never gets called
	w32.ReleaseDC( hwnd, hdc )
	//w32.DeleteDC( hdcmem )


}

