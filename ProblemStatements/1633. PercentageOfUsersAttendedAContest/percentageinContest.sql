-- Write your PostgreSQL query statement below
select r.contest_id, round(100.0 * count(distinct r.user_id) / (select count(*) from Users), 2) as percentage
from register r 
group by r.contest_id order by percentage desc, contest_id asc;