/* Write your PL/SQL query statement below */
SELECT 
    s1.student_id,
    s1.subject,
    s1.score AS first_score,
    s2.score AS latest_score
FROM Scores s1
JOIN (
    SELECT student_id, subject, MIN(exam_date) as first_date
    FROM Scores
    GROUP BY student_id, subject
) first ON s1.student_id = first.student_id 
    AND s1.subject = first.subject 
    AND s1.exam_date = first.first_date
JOIN (
    SELECT student_id, subject, MAX(exam_date) as last_date
    FROM Scores
    GROUP BY student_id, subject
) last ON s1.student_id = last.student_id 
    AND s1.subject = last.subject
JOIN Scores s2 ON s2.student_id = s1.student_id 
    AND s2.subject = s1.subject 
    AND s2.exam_date = last.last_date
WHERE s2.score > s1.score
ORDER BY s1.student_id, s1.subject;