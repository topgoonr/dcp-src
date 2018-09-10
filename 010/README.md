# Problem 10

Implement a job scheduler which takes in a function f and an integer n, and calls f after n milliseconds.

## Red Herring Implementations

There are strange implementations made by contributors around the world, who have just made a function call  that registers the callback function to the specified delay time. This is NOT a code for a scheduler. These functions are either part of the language, or a part of the standard library available with the language. You will find such *bad* implementations all over Github.

## Solution: The one coded here

Assumption: The job duration is 1 ms. This allows for a clean alignment with the for loop that mimics a *single* millisecond.

There are two different data structures of note:

1. The first is called  attachedFunction'. This allows for the attach-ing of the job to a specified system time (in milliseconds).

   ```Go
   type arbitraryFunction func(int) int
   ```

2. The second is a struct called InterruptType that ties in the function and the trigger time in millisecond together.

   ```Go
    type InterruptType struct {
    millisec         MilliSecType
    attachedFunction arbitraryFunction
    duration         MilliSecType
    }   ```

The entire logic can be divided among 'scheduler' and 'ticktock'.
'ticktock' simulates an actual system time, and will trigger the jobs that have been registered.
'scheduler' allows for registering the job at a pre-specified system time (in millisecond).

And finally, the main portion consists of declaring two functions that will be called at regular times during the time interval of a second.


Here the duration will be assigned a value of 1 for a more simplified understanding of the code.

## What would a solution with a job 'duration' > 1ms entail

If the duration is greater than 1 millisecond, than this will mean that a concept similar to a Ready Queue will have to be implemented.
The ticktock will be incremented in larger steps. These steaps will be equal to the duration.

There will be a check to see whether any one of the functions registered is already late to be scheduled, due to another function with a longer duration taking slightly longer.

