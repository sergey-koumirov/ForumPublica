start transaction;

  alter table fp_market_locations add column expiration varchar(32);

commit;
