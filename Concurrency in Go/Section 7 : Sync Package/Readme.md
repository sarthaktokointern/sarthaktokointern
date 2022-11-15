When to use channels and when to use Mutex :-

-> channels are great for communication between goroutines but what if we have caches, registry and states which are big to be sent over the channels
and we want only one goroutine to have access over this data at a particular given frame of time. So this where mutex comes into play.

Mutex :-

-> used to protect shared resources

RWMutex :-

-> allows multiple readers, writers get exclusive lock

Atomic :-

-> used to perform low level atomic operations on memory
-> lockless operation
-> used for atomic operations on counter

Conditional Variable :-

-> it is one of the synchronisation mechanisms
-> it is basically a container of goroutines that are waiting for a certain condition
-> they are of type var c *Sync.Cond
-> we use constructor method sync.NewCond() to create a conditional variable
-> it takes sync.Locker interface as input, which is usually sync.Mutex
-> it contains 3 methods
a.) wait()
-> it suspends execution of the calling goroutine
-> automatically unlocks c.L
-> Wait cannot return unless woken by Broadcast or Signal
-> Wait locks c.L before returning
-> Because c.L is not locked when Wait first resumes, the caller typically cannot resume that the condition is true when Wait returns
Instead the caller should wait in a loop

b.) signal()
-> wakes one goroutine waiting on c, if there is any
-> signal finds that goroutine that has been waiting for the longest and notifies that goroutine
-> it is allowed but not required for the caller to hold c.L  during the call

c.) broadcast()
-> wakes up all the goroutines waiting on c
->  it is allowed but not required for the caller to hold c.L  during the call

Sync.Once :- 
-> Run one-time initialisation functions
-> ensures that only one call to Do ever calls the function passed in-even on different goroutines

Sync.Pool :-
-> create and make available pool of things for use