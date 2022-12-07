package main

import (
	"fmt"
	"os"
	"time"
	//"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"janp/janplib"
)

//Make sure to create ~/.janp.json before running.
func main() {
	gtk.Init(nil)
	onload := func(w *gtk.Window) {
		fmt.Println("Goroutine is exiting")
	}
	var alert =  "Test alert"
	if (len(os.Args) > 1) {
		if (os.Args[1] == "-v" || os.Args[1] == "--version") {
			fmt.Println("janp version 1.0")
			return;
		}
		alert = os.Args[1]
	}
	janplib.Notifyf(1, time.Second*3, onload, alert)
	fmt.Println("Test")
	//go gtk.Main()
}
