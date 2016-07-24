# Write your MySQL query statement below
select 
    d.Name as Department,
    e.Name as Employee,
    e.Salary as Salary
from
    Employee as e,
    Department as d,
    (select DepartmentId, MAX(Salary) AS Salary From Employee Group By DepartmentId) as x
where
    x.DepartmentId = e.DepartmentId
    and e.DepartmentId = d.Id
    and e.Salary = x.Salary
