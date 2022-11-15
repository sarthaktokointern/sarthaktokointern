Go Race Detector :-

-> Go provides race detector tool for finding race conditions in Go code

-> go test -race mypkg //test the package
-> go run -race main.go //compile and run the program
-> go build -race mycmd //build the command
-> go install -race mypkg //install the package

-> Binary needs to be race enabled
-> when racy behaviour is detected a warning is printed
-> race enabled binary will be 10 times slower and consumes 10 times more memory
-> Integration tests and load tests are good candidates for testing with the binary that is race enabled