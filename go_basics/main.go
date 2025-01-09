package main

import (
	"fmt"
	"sync"
	"time"
)

type Meow struct {
	id     string
	number int
	amount float32
}

func newMeow(id string, number int, amount float32) *Meow {
	return &Meow{
		id:     id,
		number: number,
		amount: amount,
	}
}

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("this is task no ", id)
}

type Deo struct {
	number int
	mu     sync.Mutex
}

func (d *Deo) kero(w *sync.WaitGroup) {
	defer func() {
		d.mu.Unlock()
		w.Done()
	}()
	d.mu.Lock()
	d.number += 1
}

func main() {
	var wg sync.WaitGroup
	mmm := Deo{number: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go mmm.kero(&wg)
	}
	wg.Wait()
	fmt.Println("meow ka baccha", mmm.number)
}

func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing number", num)
		time.Sleep(time.Second)
	}
}
