CREATE TABLE esi_skills (
  id bigint NOT NULL AUTO_INCREMENT,
  esi_character_id bigint DEFAULT NULL,
  skill_id int DEFAULT NULL,
  level int DEFAULT NULL,
  name varchar(255),
  PRIMARY KEY (id)
)
