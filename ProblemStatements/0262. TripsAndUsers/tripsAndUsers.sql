/* Write your PL/SQL query statement below */
select request_at as "Day", ROUND(
    SUM(case when status != 'completed' then 1 else 0 end) * 1.0 / COUNT(*), 2) as "Cancellation Rate" 
from Trips t join users u on t.client_id = u.users_id and u.banned = 'No'
join Users d on t.driver_id = d.users_id and d.banned = 'No'
where request_at between '2013-10-01' and '2013-10-03'
group by request_at having  COUNT(*) > 0 order by request_at;