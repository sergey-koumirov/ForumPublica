START TRANSACTION;

    create view sys_digits as
        select 0 as digit union all
        select 1 union all
        select 2 union all
        select 3 union all
        select 4 union all
        select 5 union all
        select 6 union all
        select 7 union all
        select 8 union all
        select 9;

    create view sys_numbers as
      select ones.digit + tens.digit * 10 + hundreds.digit * 100 + thousands.digit * 1000 as number
        from sys_digits as ones,
             sys_digits as tens,
             sys_digits as hundreds,
             sys_digits as thousands;

    create view sys_dates as
      select subdate(current_date(), sys_numbers.number) as date
        from sys_numbers;

commit;
