# Movio programming exercises

Interview Questions from Movio

### Exercise 1

Write a function that will merge two arrays of sorted integers in a linear fashion. The output should be a sorted array containing all the elements from both input arrays, including any duplicates.

For example, given these two arrays as arguments:

[1,1,3,5], [1,2,3,4]

Your function should return:

[1,1,1,2,3,3,4,5]

### Exercise 2

Write a library for generating random PIN codes. You probably know what a PIN code is; it’s a short sequence of numbers, often used as a passcode for bank cards.

Here are the requirements:

1. The library should export a function that returns a batch of 1,000 PIN codes in random order
2. Each PIN code in the batch should be unique
3. Each PIN should be:
   - 4 digits long
   - Two consecutive digits should not be the same (e.g. 1156 is invalid)
   - Three consecutive digits should not be incremental (e.g. 1236 is invalid)
4. The library should have automated tests.
