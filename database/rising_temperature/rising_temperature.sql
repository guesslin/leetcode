select d2.Id from Weather d1, Weather d2 where d2.Date = DATE_ADD(d1.Date, INTERVAL 1 Day) and d1.Temperature < d2.Temperature;
