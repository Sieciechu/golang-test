package main

import (
	"fmt"
	"io"
	"os"
)

// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
type Chan chan int

func (ch Chan) sendInt(w io.Writer) {
	ch <- 1
	_, _ = fmt.Fprint(w, "notification sent")
}

func main() {
	var zz = make(Chan)
	go zz.sendInt(os.Stdout)

	fmt.Println(<-zz)

}