START TRANSACTION;

alter table fp_market_data add column lower_vol bigint(20);
alter table fp_market_data add column greater_vol bigint(20);

alter table fp_market_screenshots drop column is_my;

commit;

