# Weekly Contest 504
## Ranking: 
## Question 1:
*Digit Frequency Score*
### Difficulty: Easy
### Points: 3
### Description:
You are given an integer n.

The score of n is defined as the sum of d * freq(d) over all distinct digits d, where freq(d) denotes the number of times the digit d appears in n.

Return an integer denoting the score of n.
### Examples:
#### Example 1:

Input: n = 122

Output: 5

Explanation:

- The digit 1 appears 1 time, contributing 1 * 1 = 1.
- The digit 2 appears 2 times, contributing 2 * 2 = 4.
- Thus, the score of n is 1 + 4 = 5.

#### Example 2:

Input: n = 101

Output: 2

Explanation:

- The digit 0 appears 1 time, contributing 0 * 1 = 0.
- The digit 1 appears 2 times, contributing 1 * 2 = 2.
- Thus, the score of n is 2.

### Constraints:

- 1 <= n <= 10<sup>9</sup>

## Question 2:
*Maximum Number of Items From Sale I*
### Difficulty: Medium
### Points: 5
### Description:
You are given a 2D integer array items, where `items[i]` = `[factor`<sub>`i`</sub>`, price`<sub>`i`</sub>`]` represents the ith item. You are also given an integer budget.

There are unlimited copies of each item available for purchase.You may buy any number of copies of any items such that the total cost of the purchased copies is at most budget.

After buying items, you may receive free copies according to the following rules:
- For each item i that you bought at least one copy of, you receive one free copy of every item j such that `j != i` and `factor`<sub>`i`</sub> divides `factor`<sub>`j`</sub>.
- Buying multiple copies of the same item i does not give additional free copies through item i.
- The same item j can be received multiple times for free if it is received from purchases of different item types.

Return the maximum total number of item copies you can obtain, including both purchased copies and free copies, while spending at most budget on purchased items.

### Examples:
#### Example 1:

Input: items = `[[6,2],[2,6],[3,4]]`, budget = 9

Output: 4

Explanation:

- You can buy 2 copies of item 0 and 1 copy of item 2 for a total cost of 2 * 2 + 4 = 8, which is not greater than budget = 9.
- Buying item 2 gives 1 free copy of item 0, because factor<sub>2</sub> = 3 divides factor<sub>0</sub> = 6.
- You leave with 3 purchased copies and 1 free copy, for a total of 4 item copies.

#### Example 2:

Input: items = `[[2,4],[3,2],[4,1],[6,4],[12,4]]`, budget = 8

Output: 10

Explanation:

- You can buy 1 copy of item 0, 1 copy of item 1, and 2 copies of item 2 for a total cost of 4 + 2 + 2 * 1 = 8.
- Buying item 0 gives 1 free copy of items 2, 3, and 4.
- Buying item 1 gives 1 free copy of items 3 and 4.
- Buying item 2 gives 1 free copy of item 4.
- Thus, you receive 6 free copies. You leave with 4 purchased copies and 6 free copies, for a total of 10 item copies.

### Constraints:

- 1 <= `items.length` <= 1000
- `items[i]` = `[factor`<sub>`i`</sub>`, price`<sub>`i`</sub>`]`
- 1 <= `factor`<sub>`i`</sub>, `price`<sub>`i`</sub> <= 1500
- 1 <= `budget` <= 1500

## Question 3:
*Maximum Number of Items From Sale II*
### Difficulty: Medium
### Points: 5
### Description:
You are given a 2D integer array items, where `items[i]` = `[factor`<sub>`i`</sub>`, price`<sub>`i`</sub>`]` represents the ith item. You are also given an integer budget.

There are unlimited copies of each item available for purchase. You may buy any number of copies of any items such that the total cost of the purchased copies is at most budget.

After buying items, you may receive free copies according to the following rules:

- Each purchased copy of item i can give you at most one free copy of another item j.
- The free item must satisfy i != j and factor<sub>i</sub> divides factor<sub>j</sub>.
- For each ordered pair (i, j), you can receive a free copy of item j from purchases of item i at most once, regardless of how many copies of item i you buy.
- The same item j can be received multiple times for free if it is received from purchases of different item types.

Return the maximum total number of item copies you can obtain, including both purchased copies and free copies, while spending at most budget on purchased items.

### Examples:
#### Example 1:

Input: items = `[[1,6],[2,4],[3,5]]`, budget = 19

Output: 5

Explanation:

- You can buy 2 copies of item 0 and 1 copy of item 1 for a total cost of 2 * 6 + 4 = 16, which is not greater than budget = 19.
- One purchased copy of item 0 gives 1 free copy of item 1, because `factor`<sub>`0`</sub> = 1 divides `factor`<sub>`1`</sub> = 2.
- The other purchased copy of item 0 gives 1 free copy of item 2, because `factor`<sub>`0`</sub> = 1 divides `factor`<sub>`2`</sub> = 3.
- You leave with 3 purchased copies and 2 free copies, for a total of 5 item copies.

#### Example 2:

Input: items = `[[2,8],[1,10],[6,6],[4,12],[5,20],[5,17]]`, budget = 35

Output: 7

Explanation:

- You can buy 2 copies of item 0, 1 copy of item 1, and 1 copy of item 2 for a total cost of 2 * 8 + 10 + 6 = 32, which is not greater than budget = 35.
- One purchased copy of item 0 gives 1 free copy of item 2, because `factor`<sub>`0`</sub> = 2 divides `factor`<sub>`2`</sub> = 6.
- The other purchased copy of item 0 gives 1 free copy of item 3, because `factor`<sub>`0`</sub> = 2 divides `factor`<sub>`3`</sub> = 4.
- The purchased copy of item 1 gives 1 free copy of item 2, because `factor`<sub>`1`</sub> = 1 divides `factor`<sub>`2`</sub> = 6.
- Buying item 2 gives no free copy, because `factor`<sub>`2`</sub> = 6 does not divide the factor of any other item.
- You leave with 4 purchased copies and 3 free copies, for a total of 7 item copies.

### Constraints:

- 1 <= `items.length` <= 10<sup>5</sup>
- `items[i]` = `[factor`<sub>`i`</sub>`, price`<sub>`i`</sub>`]`
- 1 <= `factor`<sub>`i`</sub> <= `items.length`
- 1 <= `price`<sub>`i`</sub> <= 10<sup>9</sup>
- 1 <= `budget` <= 10<sup>9</sup>

## Question 4:
*Lexicographically Maximum MEX Array*
### Difficulty: Hard
### Points: 6
### Description:
You are given an integer array nums.

You want to construct an array result by repeatedly performing the following operation until nums becomes empty:

- Choose an integer k such that 1 <= `k` <= `len(nums)`.
- Compute the MEX of the first k elements of nums.
- Append this MEX to result.
- Remove the first k elements from nums.

Return the lexicographically maximum array result that can be obtained after performing the operations.

The MEX of an array is the smallest non-negative integer not present in the array.

An array a is lexicographically greater than an array b if in the first position where a and b differ, array a has an element that is greater than the corresponding element in b. If the first min(a.length, b.length) elements do not differ, then the longer array is the lexicographically greater one.

### Examples:
#### Example 1:

Input: nums = `[0,1,0]`

Output: `[2,1]`

Explanation:

- Take the first k = 2 elements `[0, 1]` which has MEX = 2. Current result = `[2]`.
- Remaining array `[0]` has MEX = 1. Thus, the final result = `x`.

#### Example 2:

Input: nums = `[1,0,2]`

Output: `[3]`

Explanation:

- Take the first k = 3 elements `[1, 0, 2]` which has MEX = 3.
- `nums` is now empty. Thus, the final result = `[3]`.

#### Example 3:

Input: nums = `[3,1]`

Output: `[0,0]`

Explanation:​​​​​​​

- Take k = 1, first element `[3]` has MEX = 0. Current result = `[0]`.
- Remaining array `[1]` has MEX = 0. Thus, the final result = `[0, 0]`.

### Constraints:

- 1 <= `nums.length` <= 10<sup>5</sup>
- 0 <= `nums[i]` <= 10<sup>5</sup>

