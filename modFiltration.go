package main

import "log"

type ModFiltrationStage struct{}

func (mfs ModFiltrationStage) Process(exit <-chan bool, data <-chan int) <-chan int {
	res := make(chan int)

	go func() {
		defer close(res)
		for {
			select {
			case <-exit:
				log.Println("ModFiltration: exit. breaking...")
				return
			case i, isChannelOpen := <-data:
				if !isChannelOpen {
					log.Printf("ModFiltration: data channel is closed\n")
					return
				}
				//if mod 3 case break
				if i%3 == 0 {
					log.Printf("ModFiltration: --- %d\n", i)
					break
				}

				//sending filtered
				select {
				case <-exit:
					log.Println("ModFiltration: exit. breaking...")
					return
				case res <- i:
					log.Printf("ModFiltration: -> %d\n", i)
				}

			}
		}

	}()
	return res

}
