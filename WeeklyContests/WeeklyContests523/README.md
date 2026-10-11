# Weekly Contest 523
## Ranking: 
## Question 1:
*Three Fibonacci Sum*
### Difficulty: Easy
### Points: 3
### Description:
You are given an integer n.

The Fibonacci sequence starts with 0 and 1, and each subsequent number is the sum of the previous two numbers. Its first few terms are `0, 1, 1, 2, 3, 5, 8, 13, ....`

Return true if n can be expressed as the sum of three consecutive terms in the Fibonacci sequence, and false otherwise.

### Examples:
#### Example 1:

Input: n = 16

Output: true

Explanation:

The three consecutive terms 3, 5, and 8 have a sum of 3 + 5 + 8 = 16.

#### Example 2:

Input: n = 8

Output: false

Explanation:

Although 1 + 2 + 5 = 8, these terms are not consecutive in the Fibonacci sequence. No three consecutive terms have a sum of 8.

#### Example 3:

Input: n = 2

Output: true

Explanation:

The first three terms have a sum of 0 + 1 + 1 = 2.

### Constraints:

- 1 <= n <= 10<i>9</i>

## Question 2:
*Prime Subset Selection I*
### Difficulty: Medium
### Points: 5
### Description:
You are given two integers n and s.

Consider all prime numbers less than or equal to n. Select some of these primes such that:
- The sum of the selected primes is less than or equal to s.
- The number of selected primes is maximized.
- Each prime can be selected at most once.

Return an integer array containing the selected primes in increasing order. If multiple optimal selections exist, return any of them.

If no prime can be selected, return an empty array.

### Examples:
#### Example 1:

Input: n = 7, s = 20

Output: `[2,3,5,7]`

Explanation:
- The prime numbers less than or equal to n = 7 are 2, 3, 5, and 7.  
- The sum of all these primes is 2 + 3 + 5 + 7 = 17, which is less than or equal to s = 20.
- Thus, the answer is `[2, 3, 5, 7]`.
#### Example 2:

Input: n = 15, s = 15

Output: `[3,5,7]`

Explanation:
- The prime numbers less than or equal to n = 15 are 2, 3, 5, 7, 11, and 13.
- One possible selection is `[3, 5, 7]`, whose sum is 3 + 5 + 7 = 15.
- The four smallest available primes have sum 2 + 3 + 5 + 7 = 17, which exceeds s = 15. Therefore, selecting four or more primes is impossible, and the maximum number of selected primes is three.
- Thus, `[3, 5, 7]` is one possible answer. Other valid answers include `[2, 3, 5]`, `[2, 3, 7]`, and `[2, 5, 7]`.
#### Example 3:

Input: n = 5, s = 1

Output: `[]`

Explanation:

- The smallest prime number is 2, which is greater than s = 1.
- Thus, no prime can be selected, so the answer is an empty array.

### Constraints:

- 1 <= n <= 10<sup>6</sup>
- 1 <= s <= 10<sup>9</sup>

## Question 3:
*Design Bus Seat Booking System*
### Difficulty: Medium
### Points: 5
### Description:
You are asked to design a seat booking system for a bus.

The bus has n rows. Each row has four seats labeled "`A`", "`B`", "`C`", and "`D`".

Seats "`A`" and "`D`" are window seats; seats "`B`" and "`C`" are middle seats. In each row, seat "`A`" is adjacent to seat "`B`", and seat "`D`" is adjacent to seat "`C`".

Booking rules:
- Booking a middle seat ("`B`" or "`C`") always takes 1 second.
- Booking a window seat ("`A`" or "`D`") takes 1 second if its adjacent middle seat is not booked, or 3 seconds (an extra delay of 2 seconds) if its adjacent middle seat is already booked.

A seat is identified by a string such as "`A1`" or "`D20`" (one letter from `A` to `D` followed by its row number).

Implement the BusBooking class:
- `BusBooking(int n)`: Initializes the system with n empty rows.
- `void toggle(string seat)`: Books seat if it is free, or frees it if it is already booked.
- `int getTotalTime()`: Returns the total time (in seconds) taken to book all currently booked seats, where the seats are booked in the order their toggle calls were made. If no seats are currently booked, return 0.
- `int getMinTime()`: Returns the smallest possible total time (in seconds) to book all currently booked seats, over every possible order of booking them. If no seats are currently booked, return 0.

Note: The time cost of a seat is fixed at the moment it is booked, based on the seats booked at that time. Booking or freeing other seats afterward does not change the cost of a seat that is already booked.

### Examples:
#### Example 1:

Input:
```cpp
["BusBooking", "toggle", "toggle", "getTotalTime", "getMinTime", "toggle", "toggle", "getTotalTime", "getMinTime"]
[[20], ["B3"], ["A3"], [], [], ["C3"], ["D3"], [], []]
```
Output:
```cpp
[null, null, null, 4, 2, null, null, 8, 4]
```
Explanation
```cpp
BusBooking busBooking = new BusBooking(20); // start with 20 empty rows
busBooking.toggle("B3"); // book "B3", a middle seat, which costs 1 second. Total time is now 1.
busBooking.toggle("A3"); // book "A3", a window seat. Its adjacent middle "B3" is already booked, so it costs 3 seconds. Total time is now 1 + 3 = 4.
busBooking.getTotalTime(); // return 4, since in toggle order "B3" costs 1 and "A3" costs 3.
busBooking.getMinTime(); // return 2. Booking "A3" before "B3" avoids the delay, so 2 seats takes 2 second to book.
busBooking.toggle("C3"); // book "C3", a middle seat, which costs 1 second. Total time is now 1 + 4 = 5.
busBooking.toggle("D3"); // book "D3", a window seat. Its adjacent middle "C3" is already booked, so it costs 3 seconds. Total time is now 5 + 3 = 8.
busBooking.getTotalTime(); // return 8, since in toggle order the seat costs are 1, 3, 1, 3.
busBooking.getMinTime(); // return 4. With the best possible booking order, no delay is added, so the 4 booked seats take 4 * 1 = 4 seconds.
```
#### Example 2:

Input:
```cpp
["BusBooking", "toggle", "toggle", "getTotalTime", "toggle", "toggle", "toggle", "getTotalTime", "getMinTime"]
[[10], ["C7"], ["D7"], [], ["D7"], ["C7"], ["D7"], [], []]
```
Output:
```cpp
[null, null, null, 4, null, null, null, 1, 1]
```
Explanation
```cpp
BusBooking busBooking = new BusBooking(10); // start with 10 empty rows
busBooking.toggle("C7"); // book "C7", a middle seat, which costs 1 second. Total time is now 1.
busBooking.toggle("D7"); // book "D7", a window seat. Its adjacent middle "C7" is already booked, so it costs 3 seconds. Total time is now 1 + 3 = 4.
busBooking.getTotalTime(); // return 4, since in toggle order "C7" costs 1 and "D7" costs 3.
busBooking.toggle("D7"); // "D7" is already booked, so free it. This removes its cost of 3. Total time is now 4 - 3 = 1.
busBooking.toggle("C7"); // "C7" is already booked, so free it. This removes its cost of 1. Total time is now 1 - 1 = 0, and row 7 is empty.
busBooking.toggle("D7"); // book "D7" again, a window seat. Its adjacent middle "C7" is not booked this time, so it costs only 1 second. Total time is now 0 + 1 = 1.
busBooking.getTotalTime(); // return 1, since only "D7" is booked and it cost 1 this time.
busBooking.getMinTime(); // return 1. With the best possible booking order, the 1 booked seat takes 1 * 1 = 1 second.
```

### Constraints:

- 1 <= `n` <= 10<sup>5​​​​​​​</sup>
- `seat[0]` is one of '`A`', '`B`', '`C`', '`D`'.
- The remaining characters of seat form an integer row with 1 <= `row` <= n.
- The total number of calls to `toggle`, `getTotalTime`, and `getMinTime` does not exceed 10<sup>5</sup>.

## Question 4:
*Prime Subset Selection II*
### Difficulty: Hard
### Points: 6
### Description:
You are given two integers n and s.

Consider all prime numbers less than or equal to n. Select a subset of these primes such that:
- Each prime can be selected at most once.
- The sum of the selected primes is as large as possible while remaining less than or equal to s.
- Among subsets with the maximum sum, select one containing the fewest primes.
- If multiple subsets remain tied, select the one whose elements, written in increasing order, form the lexicographically smallest array.

Return an integer array containing the selected primes in increasing order. If no valid non-empty subset exists, return an empty array.

### Examples:
#### Example 1:

Input: n = 7, s = 20

Output: `[2, 3, 5, 7]`

Explanation:
- The prime numbers less than or equal to n = 7 are 2, 3, 5, and 7.
- The sum of all these primes is 2 + 3 + 5 + 7 = 17, which is less than or equal to s = 20.
- Thus, the answer is `[2, 3, 5, 7]`.
#### Example 2:

Input: n = 15, s = 18

Output: `[5,13]`

Explanation:

- The prime numbers less than or equal to 15 are 2, 3, 5, 7, 11, and 13.
- The subsets with sum 18 are `[5, 13]`, `[7, 11]`, `[2, 3, 13]`, and `[2, 5, 11]`. Since this sum equals s, it is the maximum possible sum.
- Among these subsets, `[5, 13]` and `[7, 11]` contain the fewest primes: two.
- Between these two subsets, `[5, 13]` is lexicographically smaller because its first element, 5, is smaller than 7. Thus, the answer is `[5, 13]`.
#### Example 3:

Input: n = 5, s = 1

Output: `[]`

Explanation:

- The smallest prime number is 2, which is greater than s = 1.  
- Thus, no prime can be selected, so the answer is an empty array.

### Constraints:

- 1 <= `n` <= 1000
- 1 <= `s` <= 1000


