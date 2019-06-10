START TRANSACTION;

ALTER TABLE fp_market_locations add COLUMN esi_character_id int;

commit;
