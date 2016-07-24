select
    Score,
    @rank := @rank + (@prev != (@prev := Score)) Rank
from
    Scores,
    (select @rank := 0, @prev := -1) init
order by Score DESC
