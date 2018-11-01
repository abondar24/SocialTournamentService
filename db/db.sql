CREATE DATABASE social_tournament;

USE social_tournament;



CREATE TABLE player(
  id BIGINT NOT NULL PRIMARY KEY  AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL ,
  points INT UNSIGNED NOT NULL
);

CREATE TABLE tournament(
  id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL ,
  deposit INT UNSIGNED NOT NULL
);

CREATE TABLE tournament_player(
  id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  player_id BIGINT,
  tournament_id BIGINT,
  prize INT UNSIGNED NOT NULL
);

CREATE TABLE backer(
  id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  player_id BIGINT,
  backer_id BIGINT,
  sum INT UNSIGNED NOT NULL
);


ALTER TABLE tournament_player ADD CONSTRAINT fk_player_tournament_player
FOREIGN KEY (player_id) REFERENCES player (id);

ALTER TABLE tournament_player ADD CONSTRAINT fk_player_tournament_tournament
FOREIGN KEY (tournament_id) REFERENCES tournament (id);

ALTER TABLE backer ADD CONSTRAINT fk_backer_player
FOREIGN KEY (player_id) REFERENCES player(id);

ALTER TABLE backer ADD CONSTRAINT fk_backer_backer
FOREIGN KEY (backer_id) REFERENCES player(id);
