/* Write your PL/SQL query statement below */
delete from Person p where id not in (
    select min(id) from person group by email
)