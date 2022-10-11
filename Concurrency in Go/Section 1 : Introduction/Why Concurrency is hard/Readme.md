Why Concurrency is hard?

-> shared memory :- threads communicate by sharing memory. They all operate on the same address space. So if 2 or more threads are running
concurrently, and they try to access the same address space with one trying to write and others trying to read, then this
will lead to Data Race condition and the output will be nondeterministic


How to avoid Data Race conditon :- 


<img width="1058" alt="Screenshot 2022-10-11 at 12 00 05 PM" src="https://user-images.githubusercontent.com/99721005/195013053-0959d7f2-c859-47fe-a6d6-5acc62f7587f.png">


<img width="652" alt="Screenshot 2022-10-11 at 12 00 31 PM" src="https://user-images.githubusercontent.com/99721005/195013101-80f23e5b-e7ed-4383-9c33-85f2e8a191c5.png">
