Select e1.Name FROM Employee e1, Employee e2 WHERE e1.ManagerId = e2.Id and e1.Salary > e2.Salary;
