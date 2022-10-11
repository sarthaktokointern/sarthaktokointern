Operating System -> it's job is to provide all processes, fair access to CPU, Memory and other resources

Process :-

-> it is an instance of a running program
-> it provides environment for the program to execute

<img width="611" alt="Screenshot 2022-10-11 at 10 32 12 AM" src="https://user-images.githubusercontent.com/99721005/195001193-c5b3bf8e-130c-4533-90b4-d40e54f4f52f.png">

Threads :-

-> They are the smallest unit of execution that CPU accepts
-> Process has atleast one thread :- main thread
-> Process can have multiple threads
-> Threads share same address space
-> They run independent of each other
-> OS scheduler makes scheduling decisions at the thread level and not process level
-> Threads can run sequentially or in parallel

Thread States -> 

When a process is created, the main thread is put in the ready queue. It is in a runnable state. Once the CPU is available, the thread
starts to execute and is allotted a time slice. If the time slice expires, the thread is preempted and put back into the queue.
If the thread gets blocked due to I/O operation like read/write on disk or network operation or waiting for an event from other processes
then it is put in the waiting queue until I/O operation is complete. Once it is complete it is placed back onto the ready queue.  



Can we divide our application into processes and threads and achieve concurrency?

Yes but there are limitations :-

-> Context Switching : It is expensive and the CPU has to spend time saving the context of current thread into memory and then 
restoring context of the next chosen thread. Context switching between threads of the same process is cheaper as compared to 
context switching of threads belonging to different processes

Can we scale the number of threads per process?

-> If we scale the number of threads per process too high then we may hit the C10K problem

C10K problem -> 
