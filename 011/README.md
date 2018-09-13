
# Problem 11

Implement an autocomplete system. That is, given a query string s and a set of all possible query strings, return all strings in the set that have s as a prefix.

For example, given the query string de and the set of strings [dog, deer, deal], return [deer, deal].

Hint: Try preprocessing the dictionary into a more efficient data structure to speed up queries.

## Solution: First Attempt

1. This approach relies on the fact that an ordering of the words is required, i.e., Sorting. I have used Quick Sort to do the sorting. However, this is expensive from a memory perspective, as the entire array has to be in memory.
2. Binary Search to search for the actual prefix that will allow for autocompletion.
3. Even after this, it just assures that it is at *a* match, and not the beginning at the sequence in the array which has a match with the prefix. So, this calls for travelling from this particular point in the array up towards the beginning until there is a match.

## Solution: Tries

A better solution is perhaps possible using Tries. Did not try that.

## Caution

Array out of bounds exception is quite possible here in Solution 1. Code carefully.