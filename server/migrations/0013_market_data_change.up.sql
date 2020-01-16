START TRANSACTION;

alter table fp_market_data drop column market_location_id;
alter table fp_market_data add column market_item_id bigint;

commit;
