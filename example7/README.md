# example7

let's look at how we'd use the decorator pattern for logging things...

the example is basic, but you see
- a custmo logger
- a struct with a string
- a method to that struct that returns a "prefixed string" (and delays returning by 5s)
- a decorator that uses the custom logger to log the start... and end... of when the method starts and returns

as always, please leave feedback / comments 

