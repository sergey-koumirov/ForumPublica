START TRANSACTION;

alter table fp_market_locations add column store_location_id varchar(32);
alter table fp_market_locations add column store_location_type bigint;
alter table fp_market_locations add column store_qty bigint;

drop table fp_market_stores;

alter table fp_market_data add column sold_vol bigint;
alter table fp_market_data add column bought_vol bigint;

alter table fp_market_data drop column lower_vol;
alter table fp_market_data drop column my_vol;
alter table fp_market_data drop column greater_vol;

create table fp_market_deciles(
  id int not null auto_increment primary key,
  market_data_id int not null,
  decile int not null,
  kind  varchar(32) not null,
  average_price double not null,
  decile_vol int not null
);

drop table fp_market_screenshots;

commit;
