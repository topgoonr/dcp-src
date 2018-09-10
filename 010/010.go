/*
Implement a job scheduler which takes in a function f and an integer n, and calls f after n milliseconds.
*/

package main

import "fmt"

// arbitraryFunction type that will be registered with the scheduler
type arbitraryFunction func(int) int

// MilliSecType to save millisecs
type MilliSecType int

// InterruptType -  a job to be scheduled looks like this
type InterruptType struct {
	millisec         MilliSecType
	attachedFunction arbitraryFunction
	duration         MilliSecType // only for the purpose of this simulation
}

func scheduler(ruptor InterruptType, interr map[MilliSecType]InterruptType) bool {
	interr[ruptor.millisec] = ruptor
	fmt.Println("registered...", interr)
	return true
}

// takes a map
func ticktock(interr map[MilliSecType]InterruptType) {
	// go through a second and call the interrupt handlers on every millisecond
	// simulating a 1000th of a second
	var ms MilliSecType
	var someparam int
	for i := 0; i < 1; i++ { // forever. This is one second
		for ms = 0; ms < 1000; ms++ {
			// check the scheduler registrations; or the 'schedule'
			_, ok := interr[ms]
			if ok {
				someparam = int(ms)
				fmt.Println(interr[ms].attachedFunction(someparam))
			}
			// assuming a duration of 1 ms, otherwise a duration check logic and
			//  addition to an equivalent of the 'ready' queue would need to be done here
			// ms += interr[ms].duration, and interr[ms].attachedFunction would be
			// called all over again

		}
	}
}

func main() {
	interruptMap := make(map[MilliSecType]InterruptType)
	var add arbitraryFunction = func(i int) int { return i + i }
	var mult arbitraryFunction = func(i int) int { return i * i }

	int1 := InterruptType{millisec: 10, attachedFunction: add, duration: 1}
	int2 := InterruptType{millisec: 25, attachedFunction: mult, duration: 1}

	scheduler(int1, interruptMap)
	scheduler(int2, interruptMap)

	ticktock(interruptMap)
}
