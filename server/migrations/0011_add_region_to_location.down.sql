START TRANSACTION;

ALTER TABLE esi_locations drop COLUMN region_id bigint;
ALTER TABLE esi_locations drop COLUMN solar_system_id bigint;

commit;
