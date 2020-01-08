START TRANSACTION;

create table fp_prices(
  id int not null auto_increment primary key,
  type_id int not null,
  source varchar(32) not null,
  buy_price decimal(18,2) not null,
  sell_price decimal(18,2) not null,
  dt varchar(32) not null,
  market_volume bigint
);

create index idx_fp_prices_type_id_source on fp_prices(type_id,source);

create index idx_fp_prices_dt on fp_prices(dt);

commit;
