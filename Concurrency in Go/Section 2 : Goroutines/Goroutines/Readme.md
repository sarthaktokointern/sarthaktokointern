Concurrency in Go :-

-> it is achieved using CSP (Communicating Sequential Processes)

3 core principles of CSP :-

1.) Each process is built for sequential execution

2.) Data is communicated between processes i.e. a copy of data to created to be shared hence No shared memory therefore 
Data Race condition is prevented

3.) Scale by adding more of the same

Go's Concurrency toolset :-

-> Goroutines : concurrently executing functions
-> Channels : used to communicate data between goroutines
-> Select : used to multiplex the channels
-> Sync package : provides classical synchronisation tools like mutex and all 

Goroutines :- 

-> We can think of Goroutines as user space threads managed by go runtime
-> Goroutines are extremely lightweight, they start with 2kb of stack, which grows and shrinks as required
-> low cpu overhead : 3 instructions per function call
-> can create hundreds of thousands of goroutines in the same address space
-> Context switching is cheaper
-> Go runtime creates worker OS threads
-> Go routines run in the context of OS threads
-> Many goroutines execute in the context of single OS thread