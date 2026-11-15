# Write your MySQL query statement below

select e.name as Employee
from Employee e join Employee manager
on e.managerId = manager.id and e.salary > manager.salary 