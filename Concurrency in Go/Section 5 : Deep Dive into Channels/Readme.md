Internally channels are represented by the hchan structure. hchan struct comprises different fields. It has a mutex lock field 
Any goroutine doing any channel operation must first acquire the lock on the channel. buf is a circular ring buffer where the actual data is stored
and this is used only for the buffered channels. data queue size is the size of the buffer. qcount denotes total data elements in the
queue. sendx and recvx indicates the current index of the buffer from where it can send or receive data. sendq and recvq are the 
waiting queues which are used to store blocked goroutines. waitq is a linked list of goroutines. the elements in the linked list are 
represented by the sudog struct. In the sudog struct we have a field g which is a reference to the goroutine and elem field is a pointer
to memory which contains the value to be sent or to which the received value will be written.

When a channel is created using built-in function make, hchan struct is allocated in the heap and make returns a reference to 
the allocated memory. Since ch is a pointer it can be sent between functions which can perform send or receive operation on the channel.


First, the goroutine has to acquire the lock on the hchan struct.

Then it enqueues the element into the circular ring buffer.

Note that this is a memory copy.

The element is copied into the buffer.

OK.

Then it increments the value of the sendx to 1.

Then it releases the lock on the channel and proceed with its other computation.

Now G2 comes along and tries to receive the value from the channel.

First, it has to acquire the lock on the hchan struct,

then it dequeues the element from the buffer queue and copies the value to its variable, v.

And it increments the receive index by one and releases the lock on the channel struct and proceeds

with its other computation.

This is a simple send and receive an a buffered channel.

The points to note are, there is no memory sharing between the goroutines.

The goroutines copy elements to and from hchan struct and hchan struct is protected by the

mutex lock.

So this is where the Go's tag line comes from.

Do not communicate by sharing memory, but instead share memory by communicating.