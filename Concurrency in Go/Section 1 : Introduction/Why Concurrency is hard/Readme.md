Why Concurrency is hard?

-> shared memory :- threads communicate by sharing memory. They all operate on the same address space. So if 2 or more threads are running
concurrently, and they try to access the same address space with one trying to write and others trying to read, then this
will lead to Data Race condition and the output will be nondeterministic