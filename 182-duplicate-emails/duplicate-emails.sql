# Write your MySQL query statement below

select Person.email as Email
from Person
group by email
having count(*) > 1