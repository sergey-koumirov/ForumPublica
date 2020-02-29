START TRANSACTION;

alter table fp_market_data drop column technical;

commit;
