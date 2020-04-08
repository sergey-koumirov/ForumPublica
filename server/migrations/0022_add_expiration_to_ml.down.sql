start transaction;

  alter table fp_market_locations drop column expiration;

commit;
