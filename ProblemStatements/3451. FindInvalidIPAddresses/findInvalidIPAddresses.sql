/* Write your PL/SQL query statement below */
-- with InvalidIPs as(
--     select ip, COUNT(*) as invalid_count, case
--         when REGEXP_LIKE(ip, '^(\d{1,3}\.){3}\d{1,3}$') = 0 THEN 1
--         when REGEXP_LIKE(ip, '(^|\.)0\d') THEN 1
--         when REGEXP_LIKE(ip, '(^|\.)[0-9]{4,}') THEN 1
--         when ip like '%.%.%.%' and (
--                 CAST(REGEXP_SUBSTR(ip, '[^.]+', 1, 1) AS INTEGER) > 255 or
--                 CAST(REGEXP_SUBSTR(ip, '[^.]+', 1, 2) AS INTEGER) > 255 or
--                 CAST(REGEXP_SUBSTR(ip, '[^.]+', 1, 3) AS INTEGER) > 255 or
--                 CAST(REGEXP_SUBSTR(ip, '[^.]+', 1, 4) AS INTEGER) > 255
--             ) then 1
--         else 0
--     End as is_invalid
--     from logs group by ip having is_invalid = 1
-- ) 
-- select ip, invalid_count from InvalidIPs order by invalid_count desc, ip desc;

WITH ParsedIP AS (
    SELECT 
        ip,
        COUNT(*) as invalid_count,
        LENGTH(ip) - LENGTH(REPLACE(ip, '.', '')) as dot_count,
        SUBSTR(ip, 1, INSTR(ip, '.', 1, 1) - 1) as octet1,
        SUBSTR(ip, 
            INSTR(ip, '.', 1, 1) + 1,
            INSTR(ip, '.', 1, 2) - INSTR(ip, '.', 1, 1) - 1
        ) as octet2,
        SUBSTR(ip,
            INSTR(ip, '.', 1, 2) + 1,
            INSTR(ip, '.', 1, 3) - INSTR(ip, '.', 1, 2) - 1
        ) as octet3,
        SUBSTR(ip, INSTR(ip, '.', 1, 3) + 1) as octet4
    FROM logs
    GROUP BY ip
)
SELECT 
    ip,
    invalid_count
FROM ParsedIP
WHERE 
    dot_count != 3
    OR SUBSTR(octet1, 1, 1) = '0' AND LENGTH(octet1) > 1
    OR SUBSTR(octet2, 1, 1) = '0' AND LENGTH(octet2) > 1
    OR SUBSTR(octet3, 1, 1) = '0' AND LENGTH(octet3) > 1
    OR SUBSTR(octet4, 1, 1) = '0' AND LENGTH(octet4) > 1
    OR TO_NUMBER(octet1) > 255
    OR TO_NUMBER(octet2) > 255
    OR TO_NUMBER(octet3) > 255
    OR TO_NUMBER(octet4) > 255
ORDER BY invalid_count DESC, ip DESC;