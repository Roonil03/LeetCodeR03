-- Write your PostgreSQL query statement below
select name from Employee where id in(
    select e1.managerId from Employee e1 left join Employee e2 on e1.managerId = e2.id
    group by e1.managerId having count(e1.managerId) >= 5
);