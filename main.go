package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var (
    // cycle buffer drain interval
    bufferDrainInterval time.Duration

    // buffer max capacity
    bufferSize int
)

const (
    BREAK_WORD = "exit"
)

func main() {
	var (
        err             error
        intervalSeconds int
    )

	//logging output set
	log.SetOutput(os.Stdout)

	//enter buffer params block
	icr := NewIntConsoleReader()

    bufferSize, err = icr.Read("Enter integer (>=1) buffer size  or type '"+BREAK_WORD+"'", BREAK_WORD)
    if err != nil {
        fmt.Println(err)
        return
    }
    if bufferSize < 1 {
        fmt.Println("Too small buffer. Might be at least 1")
		return
    }

    intervalSeconds, err = icr.Read("Enter integer (>=1) buffer drain interval in seconds or type '"+BREAK_WORD+"'", BREAK_WORD)
    if err != nil {
        fmt.Println(err)
        return
    }
    if intervalSeconds < 1 {
        fmt.Println("Too small interval. Might be at least 1")
        return
    }
    bufferDrainInterval = time.Duration(intervalSeconds) * time.Second
	//--end buffer params block

	//create pipeline
	pl := NewPipeline(
		NewConsoleDataSource(),                                 // console data source (creates exit- and data- channels)
		NegativeFiltrationStage{},                              // negative filter
		ModFiltrationStage{},                                   // mod 3 filter
		NewBufferizationStage(bufferSize, bufferDrainInterval), // cycle buffer
	)

	exit, res := pl.Run()

	//begin consumer
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-exit:
				return
			case i := <-res:
				fmt.Println("Pipeline result: ", i)
			}
		}
	}()
	wg.Wait()
	//end consumer

	fmt.Println("Finish")

}
