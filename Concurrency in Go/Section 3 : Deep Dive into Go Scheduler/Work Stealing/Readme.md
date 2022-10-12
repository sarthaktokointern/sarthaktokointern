Work Stealing :-

-> helps to balance go routines across all logical processors
-> work gets better distributed and gets done more efficiently

Work Stealing Rules :-

-> If there are no goroutines in the local run queue then try to steal from other logical processors, if not found check the global
run queue for the goroutines if not found check the netpoller  