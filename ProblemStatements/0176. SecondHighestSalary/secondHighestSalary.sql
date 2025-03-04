# Write your MySQL query statement below
select MAX(Salary) as SecondHighestSalary from Employee where salary < (Select MAX(salary) from Employee);