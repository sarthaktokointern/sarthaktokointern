M:N Scheduler :-

-> The Go Scheduler is part of the Go runtime, Go runtime is part of the executable, it is built into the executable of the 
application.  It is also known as M:N Scheduler 

-> it runs in the user space

-> uses od threads to schedule goroutines for execution 

-> goroutines runs in the context of os threads 

-> go runtime creates the number of worker os threads, equal to GOMAXPROCS

-> GOMAXPROCS has default value equal to the number of processors on the machine

-> Go scheduler distributes runnable goroutines over multiple worker OS threads

-> At anytime, N goroutines could be scheduled on M OS threads hat run on at most GOMAXPROCS number of processors

Asynchronous Preemption

-> As of GO 1.14, Go scheduler implements asynchronous preemption

-> This prevents long-running Goroutines from hogging onto the CPU, that could block other Goroutines

-> The asynchronous preemption is triggered based on a time condition. When a goroutine is running for more than 10 ms 
, Go will try to preempt

Different elements involved in go scheduling :-

-> For a CPU core, Go runtime creates an OS thread which is represented by letter M, OS thread works pretty much like POSIX 
thread. Go runtime also creates a logical processor depicted by letter P and associates that with the OS thread M. The Logical
processor holds the context for scheduling, which can be seen as a local scheduler running on a thread. G represents a goroutine
running on the OS thread. Each logical processor P has a local run queue where runnable goroutines are queued. There is a global local
queue, once the local run queue gets exhausted, the logical processor will pull out goroutines from here. When new goroutines
are created they are added to the end of the global run queue. 

Context Switching due to synchronous system call :-

Suppose G1 is a goroutine at the front of the local queue attached to the logical processor P. Suppose G1 makes a synchronous system call,
like reading on a file. This will make OS thread M1 to block. Go scheduler identifies that G1 has blocked OS thread M1, so it brings in a
new OS thread M2, either from the thread pull cache or it creates a new OS thread if a thread is not available in the thread pool. Then 
the Go scheduler will detach the logical processor P from M1 and move it to M2. G1 is still attached to M1. P can now schedule other
go routines in its local queue for execution on M2. Once the synchronous system call made by G1 is complete it is moved to the back of the 
local run queue at P and M1 is put to sleep and placed back at the thread pull cache so that it can be utilised again.

Context Switching due to asynchronous system call :-

Suppose G1 is executing on OS thread M1. G1 opens a network connection with net.Dial. The file descriptor used for the connection is set 
to non-blocking mode. When the goroutine tries to read or write to the connection  the networking code will do the operation until it
receives an error, then it calls into the netpoller. Then the scheduler will move the goroutine G1 out of the os thread to the netpoller thread.
Another goroutine in the local run queue, in the case of G2 gets scheduled to run on the OS thread M1. The netpoller uses the interface
provided by the operating system to poll on the file descriptor. When the netpoller receives the notification from the operating
system that it can perform an I/O operation on the File descriptor, then it looks through its internal data structure to see if there are any
go routines that are blocked on that file descriptor. Then it notifies that goroutine and then that goroutine can retry I/O operation.
Once the I/O operation is complete, G1 will be moved back at the end of local run  queue to be processed by the OS thread M1.