START TRANSACTION;

alter table fp_market_data add column technical bool default 0;

commit;
