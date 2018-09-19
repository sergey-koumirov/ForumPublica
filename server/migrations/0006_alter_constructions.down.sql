START TRANSACTION;

alter table fp_construction_bpo_runs drop column construction_id;
alter table fp_construction_bpo_runs drop column type_id;

commit;
