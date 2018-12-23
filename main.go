package main

import (
	"fmt"
	"sync"
)

const (
	elfSize      = 9
	reindeerSize = 27
)

type Streams struct {
	ElfChan                  chan int
	ReindeerChan             chan int
	NotifElfChan      chan int
	NotifReindeerChan chan int
}

func main() {
	wg := sync.WaitGroup{}

	streams := Streams{
		ElfChan:                  make(chan int),
		ReindeerChan:             make(chan int),
		NotifElfChan:      make(chan int),
		NotifReindeerChan: make(chan int),
	}

	wg.Add(1)
	go santa(&wg, streams)

	for i := 0; i < elfSize; i++ {
		wg.Add(1)
		go elf(&wg, i, streams.ElfChan, streams.NotifElfChan)
	}

	for i := 0; i < reindeerSize; i++ {
		wg.Add(1)
		go reindeer(&wg, i, streams.ReindeerChan, streams.NotifReindeerChan)
	}

	wg.Wait()
}

func santa(wg *sync.WaitGroup, streams Streams) {
	defer wg.Done()

	elfCount := 0
	reindeerCount := 0

	developCount := 0
	shipCount := 0

	for {
		select {
		case t := <-streams.ElfChan:
			elfCount += t
		case t := <-streams.ReindeerChan:
			reindeerCount += t
		default:
			fmt.Println("Santa starts to sleep.")
		}

		if elfCount >= 3 {
			fmt.Println("Santa starts to develop.")
			for i := 0; i < 3; i++ {
				streams.NotifElfChan <- 1
			}
			elfCount = 0
			developCount++
		}

		if reindeerCount >= 9 {
			fmt.Println("Santa starts to ship.")
			for i := 0; i < 9; i++ {
				streams.NotifReindeerChan <- 1
			}
			reindeerCount = 0
			shipCount++
		}

		if developCount == elfSize/3 && shipCount == reindeerSize/9 {
			break
		}
	}
}

func elf(wg *sync.WaitGroup, num int, elfChan chan int, Notif chan int) {
	defer wg.Done()
	fmt.Println("start elf-", num)

	elfChan <- 1

	<-Notif

	fmt.Println("end elf-", num)
}

func reindeer(wg *sync.WaitGroup, num int, reindeerChan chan int, Notif chan int) {
	defer wg.Done()
	fmt.Println("start reindeer-", num)

	reindeerChan <- 1

	<-Notif

	fmt.Println("end reindeer-", num)
}
