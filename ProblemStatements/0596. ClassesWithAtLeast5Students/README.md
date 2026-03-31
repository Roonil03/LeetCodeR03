# 596. Classes With at Least 5 Students
## Question Level: Easy
### Description:
```SQL
Table: Courses

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| student     | varchar |
| class       | varchar |
+-------------+---------+
(student, class) is the primary key (combination of columns with unique values) for this table.
Each row of this table indicates the name of a student and the class in which they are enrolled.
```

Write a solution to find all the classes that have at least five students.

Return the result table in any order.

The result format is in the following example.

### Examples:
#### Example 1:

Input:
```SQL
Courses table:
+---------+----------+
| student | class    |
+---------+----------+
| A       | Math     |
| B       | English  |
| C       | Math     |
| D       | Biology  |
| E       | Math     |
| F       | Computer |
| G       | Math     |
| H       | Math     |
| I       | Math     |
+---------+----------+
```
Output:
```SQL
+---------+
| class   |
+---------+
| Math    |
+---------+
```
Explanation:  
- Math has 6 students, so we include it.
- English has 1 student, so we do not include it.
- Biology has 1 student, so we do not include it.
- Computer has 1 student, so we do not include it.

### <i>Concepts Used:
- Database</i>