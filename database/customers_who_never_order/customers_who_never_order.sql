select C.Name from Customers C where C.Id not in(select CustomerId from Orders);
