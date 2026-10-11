# 2887. Fill Missing Data
## Question Level: Easy
### Description:
```sql
DataFrame products
+-------------+--------+
| Column Name | Type   |
+-------------+--------+
| name        | object |
| quantity    | int    |
| price       | int    |
+-------------+--------+
```
Write a solution to fill in the missing value as 0 in the quantity column.

The result format is in the following example.

### Examples:
#### Example 1:
Input:
```sql
+-----------------+----------+-------+
| name            | quantity | price |
+-----------------+----------+-------+
| Wristwatch      | None     | 135   |
| WirelessEarbuds | None     | 821   |
| GolfClubs       | 779      | 9319  |
| Printer         | 849      | 3051  |
+-----------------+----------+-------+
```
Output:
```sql
+-----------------+----------+-------+
| name            | quantity | price |
+-----------------+----------+-------+
| Wristwatch      | 0        | 135   |
| WirelessEarbuds | 0        | 821   |
| GolfClubs       | 779      | 9319  |
| Printer         | 849      | 3051  |
+-----------------+----------+-------+
```
Explanation:   
The quantity for Wristwatch and WirelessEarbuds are filled by 0.  

### <i>Concepts Used:
- Pandas</i>