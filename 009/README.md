
# 007

## Problem statement

Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.

For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5]should return 10, since we pick 5 and 5.

Follow-up: Can you do this in O(N) time and constant space?

## Solution: Attempt 1

Try doing cascading iterations like 'Shell Sort'.  Not the right solution, though. Does not work with variable gaps between non-adjacent numbers.

## Solution: Attempt 2

A Dynamic Programming variant that will account for the variable and irregular adjacent numbers. Needs to be recursive, though. 