package main

import (
	"time"

	"github.com/democratic-coin/dcoin-go/vendor/src/github.com/c-darwin/go-thrust/thrust"
	"github.com/democratic-coin/dcoin-go/vendor/src/github.com/c-darwin/go-thrust/tutorials/provisioner"
	"fmt"
)

func main() {

	thrust.InitLogger()
	// Set any Custom Provisioners before Start
	thrust.SetProvisioner(tutorial.NewTutorialProvisioner())
	// thrust.Start() must always come before any bindings are created.
	thrust.Start()

	thrustWindow := thrust.NewWindow(thrust.WindowOptions{
		RootUrl:  "http://breach.cc/",
		HasFrame: true,
	})
	thrustWindow.Show()
	thrustWindow.Focus()
	thrustWindow.HandleEvent("closed", func(cr commands.EventResult) {
		fmt.Println("Close Event Occured")
	})

	// Lets do a window timeout
	go func() {
		<-time.After(time.Second * 5)
		thrustWindow.Close()
		thrust.Exit()
	}()

	// In lieu of something like an http server, we need to lock this thread
	// in order to keep it open, and keep the process running.
	// Dont worry we use runtime.Gosched :)
	thrust.LockThread()
}
