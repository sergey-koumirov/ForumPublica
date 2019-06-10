START TRANSACTION;

ALTER TABLE fp_market_locations DROP COLUMN esi_character_id;

commit;
