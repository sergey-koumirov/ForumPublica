START TRANSACTION;

  create table esi_transactions (
    id bigint(20) not null,
    client_id bigint(20) not null,
    dt varchar(32) not null,
    is_buy tinyint(1) not null,
    is_personal tinyint(1) not null,
    journal_ref_id bigint(20) not null,
    location_id bigint(20) not null,
    quantity bigint(20) not null,
    type_id	integer not null,
    unit_price double not null,

    esi_character_id bigint(20) not null,

    primary key (id),
    key idx_esi_transactions_dt (dt),
    key idx_multi (type_id,esi_character_id,location_id)
  );

  create table client_names (
    id bigint(20) not null,
    name varchar(255) not null,
    primary key (id)
  );

commit;
