package main

import "log"

type NegativeFiltrationStage struct{}

func (nfs NegativeFiltrationStage) Process(exit <-chan bool, data <-chan int) <-chan int {
	res := make(chan int)

	go func() {
		defer close(res)
		for {
			select {
			case <-exit:
				log.Println("NegativeFiltration: exit. breaking...")
				return
			case i, isChannelOpen := <-data:
				if !isChannelOpen {
					log.Printf("NegativeFiltration: data channel is closed\n")
					return
				}
				//if negative case break
				if i < 0 {
					log.Printf("NegativeFiltration: --- %d\n", i)
					break
				}

				//sending filtered
				select {
				case <-exit:
					log.Println("NegativeFiltration: exit. breaking...")
					return
				case res <- i:
					log.Printf("NegativeFiltration: -> %d\n", i)
				}

			}

		}

	}()
	return res

}
