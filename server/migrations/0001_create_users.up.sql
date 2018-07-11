create table fp_users(
  id int not null auto_increment primary key,
  role varchar(2) not null default "U"
);

create table esi_characters(
  id int not null primary key,
  name varchar(255) not null,

  user_id bigint not null,

  ver_expires_on varchar(255),
	ver_scopes varchar(255),
	ver_token_type varchar(255),
	ver_character_owner_hash varchar(255),

  tok_access_token varchar(255),
  tok_token_type varchar(255),
  tok_expires_in bigint,
  tok_refresh_token varchar(255)
);

create index esi_characters_user_id on esi_characters(user_id);
