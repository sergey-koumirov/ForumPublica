START TRANSACTION;

alter table fp_market_locations drop column store_location_id;
alter table fp_market_locations drop column store_location_type;
alter table fp_market_locations drop column store_qty;

create table fp_market_stores(
  id int not null auto_increment primary key,  
  market_item_id int not null,
  location_type varchar(32),
  location_id bigint,
  esi_character_id bigint,
  store_qty bigint
);

alter table fp_market_data drop column sold_vol;
alter table fp_market_data drop column bought_vol;

alter table fp_market_data add column lower_vol bigint;
alter table fp_market_data add column my_vol bigint;
alter table fp_market_data add column greater_vol bigint;

drop table fp_market_deciles;

create table fp_market_screenshots(
  id int not null auto_increment primary key,
  market_data_id int not null,
  vol bigint not null,
  price double not null
);

commit;
