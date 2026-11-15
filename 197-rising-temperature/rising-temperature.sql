# Write your MySQL query statement below

select next.id
from Weather prev join Weather next
on prev.recordDate = DATE_ADD(next.recordDate, interval -1 day) 
and next.temperature > prev.temperature