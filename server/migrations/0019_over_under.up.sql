start transaction;

create table fp_deviations(
  id int not null primary key,
  k  double
);

commit;
