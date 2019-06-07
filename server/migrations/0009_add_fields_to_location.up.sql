START TRANSACTION;

alter table esi_locations add column last_check_at varchar(32);

commit;
