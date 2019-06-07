START TRANSACTION;

ALTER TABLE esi_locations DROP COLUMN last_check_at;

commit;
