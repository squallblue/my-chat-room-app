package main

import (
	"bytes"
	"chatroomapp/mq"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-stomp/stomp"
	"log"
	"os"
	"os/exec"
	"strings"
)

var conn *stomp.Conn
var destination = "/topic/my-chat"
var name string

func main() {
	go func() {
		for true {
			subscribe()
		}
	}()

	os.Setenv("FYNE_FONT", "msyhl.ttc")
	cmd := exec.Command("git", "config", "--global",  "user.name")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Run()
	out := stdout.Bytes()
	name = strings.Trim(string(out), "\n")
	name = strings.Trim(string(out), "\r\n")
	if name == "" {
		 log.Println("please init your git name first:\ngit config --global user.name \"yourName\"")
	}

	a := app.New()
	w := a.NewWindow("room")
	w.Resize(fyne.NewSize(400, 2))
	entry := widget.NewEntry()
	entry.SetPlaceHolder("your context")
	w.SetContent(container.NewVBox(
		//layout.NewGridLayout(12),
		entry,
		widget.NewButton("SEND", func() {
			if entry.Text != "" {
				conn = mq.GetConnection()
				defer conn.Disconnect()
				mq.SendMessage([]byte(name + " : " + entry.Text), destination, conn)
				entry.SetText("")
			}
		}),
	))
	w.ShowAndRun()
	os.Unsetenv("FYNE_FONT")
}

func subscribe() error {
	conn = mq.GetConnection()

	subscription, err := conn.Subscribe(destination, stomp.AckClient)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case v := <-subscription.C:
			if v != nil {
				msg := string(v.Body)
				log.Println(fmt.Sprintf(msg))
				err := conn.Ack(v)
				if err != nil {
					return err
				}
			}
		}
	}
}
