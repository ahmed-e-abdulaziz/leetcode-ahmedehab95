# Write your MySQL query statement below
SELECT d.name Department, e.name Employee, e.salary Salary From Department d
JOIN Employee e on e.DepartmentId = d.id where (DepartmentId, Salary)
in 
(
select DepartmentId, max(Salary) from Employee group by DepartmentId
)

