start transaction;

create table vol30d(
  market_item_id bigint not null,
  dt varchar(32) not null,
  is_my int not null,
  group_num int not null,
  vol bigint not null
);

create index vol30d_001_idx on vol30d(market_item_id);

commit;