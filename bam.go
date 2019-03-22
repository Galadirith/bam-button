// Copyright (c) 2014-2015 Edward Fauchon-Jones
// Use of this source code is governed by a MIT license that can be found in
// the LICENSE.md file.
//
// ------------------------------------------------------------------------
//
// This file incorporates code from
// [gorilla/websocket](https://github.com/gorilla/websocket/tree/master/examples/echo)
// covered by the following terms:
//
// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
	"math/rand"

	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
)

func bam() {
  // Grab the current data on the clipboard
  data, _ := robotgo.ReadAll()

  // Update clipboard and paste BAM emoji
  BAM_array := []string{`👖`, `🍎`, `🍆`, `🍠`, `🍾`, `🐧`, `🎂`}
  rand.Seed(time.Now().UnixNano())
  BAM := BAM_array[rand.Intn(len(BAM_array)-1)]
  robotgo.WriteAll(BAM)
  robotgo.PasteStr(BAM)

  // Resotre clipboard after allowing paste time
  time.Sleep(1000 * time.Millisecond)
  robotgo.WriteAll(data)
}

func main() {
	// Observe interruption signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Open websocket connection to bam-button server
	u := url.URL{Scheme: "ws", Host: "bam-button.glitch.me", Path: "/echo"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// Handle websocket messages and call `bam` if message is `"bam"`
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
			if string(message) == "bam" {
				bam()
			}
		}
	}()

	for {
		select {

		// There has been an error from the bam-button server
		case <-done:
			return

		// The user has interruted execution
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
