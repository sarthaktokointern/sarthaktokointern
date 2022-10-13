Select :-

-> Select statement is like switch statement
-> Each case specifies communication
-> All channel operations are considered simultaneously
-> Each case has an equal chance of being selected
-> select waits until some case is ready to proceed
-> if no case is ready then the entire select statement gets blocked and waits until some case gets ready
-> if multiple cases are ready then it picks one of them randomly
-> default case is added to prevent blocking
-> empty select statement will block forever
-> select on nil channel will block forever
