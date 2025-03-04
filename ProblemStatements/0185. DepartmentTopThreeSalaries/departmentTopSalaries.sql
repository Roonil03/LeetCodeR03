/* Write your PL/SQL query statement below */
select Department, Employee, Salary from(
    select d.name as Department, e.name as Employee, e.salary as Salary, DENSE_RANK() over (
        partition by e.departmentId order by e.salary desc) as rank from Employee e join Department d on e.departmentId = d.id
) where rank <= 3;