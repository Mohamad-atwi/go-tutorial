### Channels and goroutines

In this task you will write a "monitoring system" which will
periodically check the status of a list of URLs and report it to the console.

1. Look at the checkLink function, it is missing some arguments and two lines of code. Fix it (hint: channel variable is initialized for a reason).
2. Call the checkLink function appropriately in line 19.
3. Write a nested goroutine (https://www.techieindoor.com/go-how-to-use-nested-goroutines-in-go/) which will sleep for 5 seconds and call checkLink function again (hint: an elegant solution is 4 lines long).
4. Run the program and see if it works as expected.