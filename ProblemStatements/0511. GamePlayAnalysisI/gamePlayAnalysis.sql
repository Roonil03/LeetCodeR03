/* Write your PL/SQL query statement below */
select player_id, TO_CHAR(MIN(event_date), 'YYYY-MM-DD') as "first_login" from Activity group by player_id;