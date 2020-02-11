START TRANSACTION;

alter table fp_market_data drop column lower_vol;
alter table fp_market_data drop column greater_vol;

alter table fp_market_data add column my_lowest_price double;

alter table fp_market_data modify column sell_vol bigint;
alter table fp_market_data modify column buy_vol bigint;

alter table fp_market_screenshots add column is_my tinyint(1);

commit;
