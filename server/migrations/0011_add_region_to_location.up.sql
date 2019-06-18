START TRANSACTION;

ALTER TABLE esi_locations add COLUMN region_id bigint;
ALTER TABLE esi_locations add COLUMN solar_system_id bigint;

commit;
