CREATE FUNCTION getNthHighestSalary(N IN NUMBER) RETURN NUMBER IS
result NUMBER;
BEGIN
    /* Write your PL/SQL query statement below */
    select salary into result from(
        select salary, ROWNUM as rn from(
            select distinct salary from Employee order by salary desc
        )
        where ROWNUM <= N
    )
    where rn = N;
    RETURN result;
EXCEPTION
   WHEN NO_DATA_FOUND THEN
       RETURN NULL;
END;