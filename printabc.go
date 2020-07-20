package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	c1 := sync.NewCond(new(sync.Mutex))
	c2 := sync.NewCond(new(sync.Mutex))
	c3 := sync.NewCond(new(sync.Mutex))

	p1 := &Printer{
		c1,
		c2,
		"a",
	}

	p2 := &Printer{
		c2,
		c3,
		"b",
	}

	p3 := &Printer{
		c3,
		c1,
		"c",
	}

	go func() {
		for {
			p1.print()
		}
	}()

	go func() {
		for {
			p2.print()
		}
	}()

	go func() {
		for {
			p3.print()
			time.Sleep(1 * time.Second)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	fmt.Println("finished")
}

type Printer struct {
	pre      *sync.Cond
	next     *sync.Cond
	printsth string
}

func (p *Printer) print() {
	pre := p.pre
	next := p.next

	next.L.Lock()
	next.Broadcast()
	next.L.Unlock()

	pre.L.Lock()
	pre.Wait()
	fmt.Println(p.printsth)
	pre.L.Unlock()
}
