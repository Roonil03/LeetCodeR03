# Weekly Contest 428
## Ranking: 6234
## Question 1:
<i>Button with Longest Push Time</i>
### Difficult: Easy
### Points: 3
### Description:
You are given a 2D array events which represents a sequence of events where a child pushes a series of buttons on a keyboard.

Each `events[i]` = `[indexi, timei]` indicates that the button at index indexi was pressed at time timei.
- The array is sorted in increasing order of time.
- The time taken to press a button is the difference in time between consecutive button presses. The time for the first button is simply the time at which it was pressed.

Return the index of the button that took the longest time to push. If multiple buttons have the same longest time, return the button with the smallest index.

### Test Cases:
#### Example 1:

Input: events = `[[1,2],[2,5],[3,9],[1,15]]`<br>
Output: 1<br>
Explanation:
- Button with index 1 is pressed at time 2.
- Button with index 2 is pressed at time 5, so it took 5 - 2 = 3 units of time.
- Button with index 3 is pressed at time 9, so it took 9 - 5 = 4 units of time.
- Button with index 1 is pressed again at time 15, so it took 15 - 9 = 6 units of time.
#### Example 2:

Input: events = `[[10,5],[1,7]]`<br>
Output: 10<br>
Explanation:
- Button with index 10 is pressed at time 5.
- Button with index 1 is pressed at time 7, so it took 7 - 5 = 2 units of time.

### Constraints:

- 1 <= `events.length` <= 1000
- `events[i]` == `[indexi, timei]`
- 1 <= `indexi`, `timei` <= 10^5
- The input is generated such that events is sorted in increasing order of `timei`.

## Question 2:
<i>Maximize Amount After Two Days of Conversions</i>
### Difficulty: Medium
### Points: 4
### Description:
You are given a string `initialCurrency`, and you start with 1.0 of `initialCurrency`.

You are also given four arrays with currency pairs (strings) and rates (real numbers):
- `pairs1[i]` = `[startCurrencyi, targetCurrencyi]` denotes that you can convert from `startCurrencyi` to `targetCurrencyi` at a rate of `rates1[i]` on day 1.
- `pairs2[i]` = `[startCurrencyi, targetCurrencyi]` denotes that you can convert from `startCurrencyi` to `targetCurrencyi` at a rate of `rates2[i]` on day 2.
- Also, each targetCurrency can be converted back to its corresponding startCurrency at a rate of `1 / rate`.

You can perform any number of conversions, including zero, using rates1 on day 1, followed by any number of additional conversions, including zero, using rates2 on day 2.

Return the maximum amount of `initialCurrency` you can have after performing any number of conversions on both days in order.

<i>Note: Conversion rates are valid, and there will be no contradictions in the rates for either day. The rates for the days are independent of each other.</i>

### Test Cases:
#### Example 1:
Input: initialCurrency = `"EUR"`, pairs1 = `[["EUR","USD"],["USD","JPY"]]`, rates1 = `[2.0,3.0]`, pairs2 = `[["JPY","USD"],["USD","CHF"],["CHF","EUR"]]`, rates2 = `[4.0,5.0,6.0]`<br>
Output: 720.00000<br>
Explanation:<br>
To get the maximum amount of EUR, starting with 1.0 EUR:
- On Day 1:
    - Convert EUR to USD to get 2.0 USD.
    - Convert USD to JPY to get 6.0 JPY.
- On Day 2:
    - Convert JPY to USD to get 24.0 USD.
    - Convert USD to CHF to get 120.0 CHF.
- Finally, convert CHF to EUR to get 720.0 EUR.

#### Example 2:
Input: initialCurrency = `"NGN"`, pairs1 = `[["NGN","EUR"]]`, rates1 = `[9.0]`, pairs2 = `[["NGN","EUR"]]`, rates2 = `[6.0]`<br>
Output: 1.50000<br>
Explanation:<br>
- Converting NGN to EUR on day 1 and EUR to NGN using the inverse rate on day 2 gives the maximum amount.

#### Example 3:
Input: initialCurrency = `"USD"`, pairs1 = `[["USD","EUR"]]`, rates1 = `[1.0]`, pairs2 = `[["EUR","JPY"]]`, rates2 = `[10.0]`<br>
Output: 1.00000<br>
Explanation:
- In this example, there is no need to make any conversions on either day.

### Constraints:

- 1 <= `initialCurrency.length` <= 3
- `initialCurrency` consists only of uppercase English letters.
- 1 <= `n` == `pairs1.length` <= 10
- 1 <= `m` == `pairs2.length` <= 10
- `pairs1[i]` == `[startCurrencyi, targetCurrencyi]`
- `pairs2[i]` == `[startCurrencyi, targetCurrencyi]`
- 1 <= `startCurrencyi.length`, `targetCurrencyi.length` <= 3
- `startCurrencyi` and `targetCurrencyi` consist only of uppercase English letters.
- `rates1.length` == n
- `rates2.length` == m
- 1.0 <= `rates1[i]`, `rates2[i]` <= 10.0
- The input is generated such that there are no contradictions or cycles in the conversion graphs for either day.

## Question 3:
<i>Count Beautiful Splits in an Array</i>
### Difficulty: Medium
### Points: 5
### Description:
You are given an array `nums`.

A split of an array `nums` is beautiful if:
- The array `nums` is split into three non-empty subarrays: `nums1`, `nums2`, and `nums3`, such that nums can be formed by concatenating `nums1`, `nums2`, and `nums3` in that order.
- The subarray nums1 is a prefix of `nums2` OR `nums2` is a prefix of `nums3`.

Return the number of ways you can make this split.

- A subarray is a contiguous non-empty sequence of elements within an array.
- A prefix of an array is a subarray that starts from the beginning of the array and extends to any point within it.

### Test Cases:
#### Example 1:
Input: nums = `[1,1,2,1]`
Output: 2
Explanation:
The beautiful splits are:
- A split with `nums1` = `[1]`, `nums2` = `[1,2]`, `nums3` = `[1]`.
- A split with `nums1` = `[1]`, `nums2` = `[1]`, `nums3` = `[2,1]`.

#### Example 2:
Input: nums = `[1,2,3,4]`
Output: 0
Explanation:
- There are 0 beautiful splits.

### Constraints:
- 1 <= `nums.length` <= 5000
- 0 <= `nums[i]` <= 50

## Question 4:
<i>Minimum Operations to Make Character Frequencies Equal</i>
### Difficulty: Hard
### Points: 8
### Description:
You are given a string s.

A string `t` is called good if all characters of `t` occur the same number of times.

You can perform the following operations any number of times:
- Delete a character from s.
- Insert a character in s.
- Change a character in s to its next letter in the alphabet.<br>
<i>Note that you cannot change 'z' to 'a' using the third operation.</i>

Return the minimum number of operations required to make s good.

### Test Cases:
#### Example 1:
Input: s = `"acab"`<br>
Output: 1<br>
Explanation:<br>
- We can make s good by deleting one occurrence of character `'a'`.

#### Example 2:
Input: s = `"wddw"`<br>
Output: 0<br>
Explanation:
- We do not need to perform any operations since `s` is initially good.

#### Example 3:
Input: s = `"aaabc"`<br>
Output: 2<br>
Explanation:<br>
We can make s good by applying these operations:
- Change one occurrence of `'a'` to `'b'`
- Insert one occurrence of `'c'` into `s`

### Constraints:

- 3 <= `s.length` <= 2 * 10^4
- `s` contains only lowercase English letters.