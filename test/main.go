package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"

	"github.com/gaigepr/lazylist"
)

type message string

var done = false

func containsWorker(list *lazylist.LazyList, mailbox chan message) {
	var a uint64
	for {
		a = uint64(rand.Int63())
		if done {
			break
		} else if list.Contains(a) {
			// list contains a
			mailbox <- "membership"
		} else {
			// list does not contain a
			mailbox <- "failedMembership"
		}
	}
	fmt.Println("Exiting contains")
}

func addWorker(list *lazylist.LazyList, mailbox chan message) {
	var a uint64
	for {
		a = uint64(rand.Int63())
		if done {
			break
		} else if list.Add(a) {
			// now list contains a
			mailbox <- "addition"
		} else {
			// list already contains a, collision
			mailbox <- "collision"
		}
	}
	fmt.Println("Exiting add")
}

func removeWorker(list *lazylist.LazyList, mailbox chan message) {
	var a uint64
	for {
		a = uint64(rand.Int63())
		if done {
			break
		} else if list.Remove(a) {
			// now list does not contains a
			mailbox <- "removal"
		} else {
			// list already did not contains a, anti-collision(?)
			mailbox <- "antiCollision"
		}
	}
	fmt.Println("Exiting remove")
}

// Woah, this isn't idiomatic go at all!
func main() {

	// Command line flags
	var (
		containsThreads = flag.Int("contains-threads", 2, "An int")
		addThreads      = flag.Int("add-threads", 2, "An int")
		removeThreads   = flag.Int("remove-threads", 2, "An int")
		listStartSize   = flag.Int("list-size", 32, "A large int")
		//threadWaitTime  = flag.Int("thread-wait", 0, "An int representing milliseconds")
		//numGoProcs      = flag.Int("go-procs", 3 /*TODO: numCPUs*/, "An int")
		//maxRuntime      = flag.Int("max-runtime", 20, "An int representing seconds")
	)

	flag.Parse()

	// Channel for threads to report things
	mailbox := make(chan message)
	// Make and fill a new lazylist. Single threaded.
	list := lazylist.NewLazyList()
	for i := uint64(0); i < uint64(*listStartSize); i++ {
		list.Add(i)
	}

	// Spin up all the contains, add, and remove threads
	for i := 0; i < *containsThreads; i++ {
		go containsWorker(list, mailbox)
	}
	for i := 0; i < *addThreads; i++ {
		go addWorker(list, mailbox)
	}
	for i := 0; i < *removeThreads; i++ {
		go removeWorker(list, mailbox)
	}

	var collisions uint64 = 0
	var antiCollisions uint64 = 0
	var addition uint64 = 0
	var removal uint64 = 0
	var membership = 0
	var failedMembership = 0

	// Channel to handle interupts
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for {
		select {
		case <-c:
			fmt.Println("Got interupt")
			fmt.Printf("Collisions: %d\t antiCollisions: %d\t membership (total | failed): %d | %d\t additions: %d\t removals: %d\n",
				collisions, antiCollisions, membership, failedMembership, addition, removal)
			return

		case m := <-mailbox:
			switch {
			case m == "collision":
				collisions += 1

			case m == "antiCollision":
				antiCollisions += 1

			case m == "membership":
				membership += 1

			case m == "failedMembership":
				failedMembership += 1

			case m == "addition":
				addition += 1

			case m == "removal":
				removal += 1

			}
		}
	}

}
