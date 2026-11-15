# Write your MySQL query statement below

select p.email as Email
from Person p join Person p2 
on p.email = p2.email and p.id <> p2.id
group by p.email