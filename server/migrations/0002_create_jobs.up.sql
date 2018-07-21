create table esi_jobs(
  id bigint not null primary key,

  esi_character_id bigint not null,

  activity_id int not null,
  blueprint_id bigint not null,
  blueprint_location_id bigint not null,
  blueprint_type_id int not null,
  cost bigint not null,
  duration bigint not null,
  end_date varchar(32) not null,
  facility_id bigint not null,
  installer_id bigint not null,
  licensed_runs int not null,
  output_location_id bigint not null,
  probability float not null,
  product_type_id int not null,
  runs int not null,
  start_date varchar(32) not null,
  station_id bigint not null,
  status varchar(32) not null
);

create index esi_jobs_esi_character_id on esi_jobs(esi_character_id);
