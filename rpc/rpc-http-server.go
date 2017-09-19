/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"fmt"
	"net/http"
	"net/rpc"
	//"os"
	"sync"
	"time"
)

var (
	stringMemory string
	iCount       = 0
	mapMemory    = map[int]string{}

	Mux = struct {
		sync.RWMutex
		m map[int]string
	}{m: make(map[int]string)}
)

// Method Multiply arguments
type Args struct {
	A, B int
}

// Kind for my method
type Matt int

// My method Multiply
func (t *Matt) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Method StopServer arguments
type Args2 struct {
	A string
}

// type stop
type Stop string

// My method StopServer
func (s *Stop) StopServer(args *Args2, replys *string) error {

	*replys = args.A + " ok! "
	fmt.Println("Stopping the server by rpc!")

	var count = 5
	for i := 0; i < count; i++ {

		stringMemory = "service[" + fmt.Sprintf("%d", iCount) + "] map "

		Mux.Lock()
		mapMemory[iCount] = stringMemory
		Mux.Unlock()

		fmt.Println(stringMemory)

		//fmt.Println("service[", i, "]", "stop")
		time.Sleep(2 * time.Second)

		iCount++
	}

	// fmt.Println("iCount: ", iCount)
	// time.Sleep(time.Second * 1)
	//fmt.Println(mapMemory)

	//os.Exit(1)
	return nil
}

func WriteMemory() {

	for {

		time.Sleep(3 * time.Second)

		fmt.Println("Read map in Memory")

		//Mux.RLock()
		for j, val := range Mux.m {

			//view := Mux.m[iCount]

			fmt.Println("Map[", j, "] = ", val)

			time.Sleep(1 * time.Second)
		}

		//Mux.RUnlock()
	}
}

func ReadMemory() {

	for {

		time.Sleep(2 * time.Second)

		fmt.Println("Read Memory")

		for j, val := range mapMemory {

			fmt.Println("Map[", j, "] = ", val)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	// Recording the method Matt
	matt := new(Matt)
	rpc.Register(matt)

	// Recording the method Stop
	stop := new(Stop)
	rpc.Register(stop)

	// Start handler
	rpc.HandleHTTP()

	go WriteMemory()

	// Opening the port for communication
	err := http.ListenAndServe(":1234", nil)

	if err != nil {

		fmt.Println(err.Error())
	}

}
