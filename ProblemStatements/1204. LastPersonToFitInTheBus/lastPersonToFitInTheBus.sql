-- Write your PostgreSQL query statement below
SELECT person_name FROM (
    SELECT person_name, SUM(weight) OVER(ORDER BY turn) AS w from Queue
) t where w <= 1000 ORDER BY w DESC LIMIT 1;