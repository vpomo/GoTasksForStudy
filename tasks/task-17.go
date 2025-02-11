package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Port struct {
	number, value int
}

type InPortJob struct {
	id   int
	port *Port
}

type OutPortJob struct {
	id   int
	port *Port
}

func (in *InPortJob) read(portNum int) int {
	time.Sleep(3 * time.Microsecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano() * int64(portNum) / int64(1000)))
	val := r.Intn(2)
	fmt.Println("Read API for job number ", in.id, " port number = ", portNum, " write value = ", val)
	return val
}

func (out *OutPortJob) write(portNum int, val int) {
	fmt.Println("Write API for job number ", out.id, " write value = ", val)
	out.port.number = portNum
	out.port.value = val
}

func main() {
	fmt.Println("Started")

	var inCount = 0
	var outCount = 0
	var inMaxJobs = 0
	var outMaxJobs = 0

	flag.IntVar(&inCount, "IN", 4, "IN ports count")
	flag.IntVar(&outCount, "OUT", 6, "OUT ports count")
	flag.IntVar(&inMaxJobs, "IN MAX JOBS", 10, "for IN ports max jobs")
	flag.IntVar(&outMaxJobs, "OUT MAX JOBS", 12, "for OUT ports max jobs")
	flag.Parse()

	var inWg sync.WaitGroup
	var outWg sync.WaitGroup

	inCh := make(chan InPortJob)
	outCh := make(chan OutPortJob)

	r := rand.New(rand.NewSource(time.Now().UnixNano() / int64(1000)))

	fmt.Println("Working with API ...")

	go func() {
		defer close(inCh)
		for i := 1; i < inMaxJobs+1; i++ {
			inWg.Add(1)
			go func(j int, wg *sync.WaitGroup) {
				defer wg.Done()
				inPortNum := r.Intn(inCount) + 1
				inPortJob := new(InPortJob)
				inPortJob.id = j
				inPortJob.port.number = inPortNum
				inPortJob.port.value = inPortJob.read(inPortNum)
				inCh <- *inPortJob
			}(i, &inWg)
		}
		inWg.Wait()
	}()

	go func() {
		defer close(outCh)
		for i := 1; i < outMaxJobs+1; i++ {
			outWg.Add(1)
			go func(j int, wg *sync.WaitGroup) {
				defer outWg.Done()
				outPortNum := r.Intn(outCount) + 1
				outPortJob := new(OutPortJob)
				outPortJob.id = j
				outPortJob.write(outPortNum, r.Intn(2))
				outCh <- *outPortJob
			}(i, &outWg)
		}
		outWg.Wait()
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Results:")

	for inCurrent := range inCh {
		fmt.Println("IN job id = ", inCurrent.id, ", port number = ", inCurrent.port.number, ", value = ", inCurrent.port.value)
	}
	for outCurrent := range outCh {
		fmt.Println("OUT job id = ", outCurrent.id, ", port number = ", outCurrent.port.number, ", value = ", outCurrent.port.value)
	}

	fmt.Println("Finished")
}
