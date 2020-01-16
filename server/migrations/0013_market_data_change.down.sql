START TRANSACTION;

alter table fp_market_data drop column market_item_id;
alter table fp_market_data add column market_location_id  bigint;

commit;

