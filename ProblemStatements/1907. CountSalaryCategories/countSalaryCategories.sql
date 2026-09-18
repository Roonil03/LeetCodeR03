-- Write your PostgreSQL query statement below
-- SELECT UNNEST(ARRAY['Low Salary', 'Average Salary', 'High Salary']) AS category,
-- UNNEST(ARRAY[COUNT(*) FILTER (WHERE income < 20000), COUNT(*) FILTER (WHERE income BETWEEN 20000 AND 50000), COUNT(*) FILTER(WHERE income > 50000)]) AS account_count
-- FROM Accounts;

SELECT 'Low Salary' AS category, COUNT(*) AS accounts_count FROM Accounts WHERE income < 20000
UNION ALL
SELECT 'Average Salary', COUNT(*) FROM Accounts WHERE income BETWEEN 20000 AND 50000
UNION ALL
SELECT 'High Salary', COUNT(*) FROM Accounts WHERE income > 50000;