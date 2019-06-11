START TRANSACTION;

create table fp_market_items(
  id int not null auto_increment primary key,
  type_id int not null,
  user_id int not null
);

create table esi_locations(
  id bigint primary key not null,
  name varchar(255)
);

create table fp_market_locations(
  id int not null auto_increment primary key,  
  market_item_id int not null,
  location_type varchar(32),
  location_id bigint,  
  store_location_type varchar(32),
  store_location_id bigint,
  store_qty bigint
);

create table fp_market_data(
  id int not null auto_increment primary key,  
  market_location_id int not null,
  dt varchar(32) not null,
  sell_vol int not null,
  buy_vol int not null,
  sold_vol int not null,
  bought_vol int not null,
  sell_lowest_price double,
  buy_highest_price double
);


create table fp_market_deciles(
  id int not null auto_increment primary key,
  market_data_id int not null,
  decile int not null,
  kind  varchar(32) not null,
  average_price double not null,
  decile_vol int not null
);

commit;
