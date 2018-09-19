START TRANSACTION;

alter table fp_construction_bpo_runs add column construction_id int;
alter table fp_construction_bpo_runs add column type_id int;

commit;
