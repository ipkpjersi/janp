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
			fmt.Println("janp version 1.2")
			return;
		}
		alert = os.Args[1]
	}
	cfg, err2 := janplib.LoadConfig()
	if (err2 != nil) {
		fmt.Println("Error loading cfg")
		fmt.Println(err2)
	}

	monitor := int(cfg.XineramaHead);
	notificationTime := int(cfg.NotificationTime);

	janplib.Notifyf(uint32(monitor), time.Second*time.Duration(notificationTime), onload, alert)
	fmt.Println("Test")
	//go gtk.Main()
}
