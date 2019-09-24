# gopeloton
Just some basic golang stuff to grab onepeloton data

As a newbie go programmer I wanted to see if I could do something useful.   My starting point was the github api doodles from 
https://github.com/geudrik/peloton-api/blob/master/API_DOCS.md  

Having got a long way through the Udemy training course from https://github.com/GoesToEleven I decided to try and see how I could pull from the onepeloton API.   In the end I decided just to do a total count of the calories burned in all my workouts.

The structs documented here are fine at the time of me writing the code, but I expect this to change - I have only just started with peloton so for example in the WorkoutSummary struct you can see:

`Summary struct {
		Two01909 int `json:"2019-09"`
	} `json:"summary"``
  
As I have only used it in September I suspect I will need to do some further work once I go into October and beyond.
  
I only use a tiny bit of the data that gets returned but I have included the full structs in pelotonstructs.go as there may be something else there that is of interest.

I am almost certain that this code can be cleaned up "A LOT" so feel free to enhance it.

# To run

Update the user and password fields:  
`pelotonID := "fred" `   
and 
`pelotonPass := "abc123"`

then:

` go run main.go peletonstructs.go`
  
  


