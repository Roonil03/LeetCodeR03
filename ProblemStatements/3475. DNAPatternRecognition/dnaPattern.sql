/* Write your PL/SQL query statement below */
SELECT 
    sample_id,
    dna_sequence,
    species,
    CASE WHEN SUBSTR(dna_sequence, 1, 3) = 'ATG' THEN 1 ELSE 0 END as has_start,
    CASE WHEN dna_sequence LIKE '%TAA' 
         OR dna_sequence LIKE '%TAG' 
         OR dna_sequence LIKE '%TGA' THEN 1 ELSE 0 END as has_stop,
    CASE WHEN dna_sequence LIKE '%ATAT%' THEN 1 ELSE 0 END as has_atat,
    CASE WHEN REGEXP_LIKE(dna_sequence, 'G{3,}') THEN 1 ELSE 0 END as has_ggg
FROM Samples
ORDER BY sample_id;