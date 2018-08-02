START TRANSACTION;

create table fp_constructions(
  id int not null auto_increment primary key,
  user_id int,
  name varchar(255),
  citadel_type varchar(32),
  rig_factor varchar(32),
  space_type varchar(32)
);

create table fp_construction_bpos(
  id int not null auto_increment primary key,
  construction_id int,
  transaction_id int,
  kind varchar(32),
  type_id int,

  me int,
  te int,
  qty bigint
);

create table fp_construction_bpo_runs(
  id int not null auto_increment primary key,
  construction_bpo_id int,

  repeats int,
  me int,
  te int,
  qty bigint,

  citadel_type varchar(32),
  rig_factor varchar(32),
  space_type varchar(32)
);

create table fp_construction_expenses(
  id int not null auto_increment primary key,
  construction_bpo_id int,
  description varchar(255),
  exvalue double
);

commit;
