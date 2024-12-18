import pandas as pd

def department_highest_salary(employee: pd.DataFrame, department: pd.DataFrame) -> pd.DataFrame:
    if employee.empty or department.empty:
        return pd.DataFrame(columns=['Department','Employee', 'Salary'])
    
    merged_df = pd.merge(employee, department, left_on='departmentId', right_on='id', suffixes=('_employee', '_department'))
    max_salary = merged_df.groupby('departmentId')['salary'].max().reset_index()
    result = pd.merge(merged_df, max_salary, on='departmentId', suffixes=('', '_max'))
    result = result[result['salary'] == result['salary_max']]
    final_result = result[['name_department', 'name_employee', 'salary']]
    final_result.columns = ['Department', 'Employee', 'Salary']
    
    return final_result
