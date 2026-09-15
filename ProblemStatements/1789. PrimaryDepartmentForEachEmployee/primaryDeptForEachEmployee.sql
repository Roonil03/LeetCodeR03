-- Write your PostgreSQL query statement below
SELECT DISTINCT ON (employee_id) employee_id, department_id FROM Employee ORDER BY employee_id, primary_flag DESC;