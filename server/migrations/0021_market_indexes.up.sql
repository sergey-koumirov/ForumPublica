start transaction;

  create index fp_market_data_s001 on fp_market_data(market_item_id, dt);

  create index fp_market_screenshots_s001 on fp_market_screenshots(market_data_id);

commit;
