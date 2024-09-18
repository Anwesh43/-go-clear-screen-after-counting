package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
	time.Sleep(time.Second)
}

func (c *Counter) IsTenth() bool {
	return c.count%10 == 0
}

func (c *Counter) Display() {
	fmt.Println(c.count)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func startingDisplay() {
	fmt.Println("Printing the numbers")
}

func main() {
	c := Counter{
		count: 0,
	}
	ch := make(chan bool)
	clearScreen()
	startingDisplay()
	go func() {
		i := 3
		for i > 0 {
			c.Increment()
			c.Display()
			if c.IsTenth() {
				clearScreen()
				startingDisplay()
				i--
			}
		}
		ch <- true
	}()
	<-ch
}
