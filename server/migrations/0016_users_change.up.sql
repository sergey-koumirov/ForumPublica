START TRANSACTION;

  alter table esi_characters modify ver_scopes varchar(8000);

commit;
