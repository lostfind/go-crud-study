DROP SCHEMA IF EXISTS go_study;
CREATE SCHEMA go_study;
USE go_study;

DROP TABLE IF EXISTS zipcode;

CREATE TABLE zipcode
(
  zipcode          VARCHAR(10),
  prefecture_id    INT(5),
  prefecture_name  VARCHAR(40),
  prefecture_kana  VARCHAR(40),
  city_id          INT(10),
  city_name        VARCHAR(40),
  city_kana        VARCHAR(40),
  town_name        VARCHAR(40),
  town_kana        VARCHAR(40)
);

INSERT INTO zipcode (zipcode, prefecture_id, prefecture_name, prefecture_kana, city_id, city_name, city_kana, town_name, town_kana)
VALUES ("4630062",23,"愛知県","アイチケン",23113,"名古屋市守山区","ナゴヤシモリヤマク","長栄","チョウエイ");

INSERT INTO zipcode (zipcode, prefecture_id, prefecture_name, prefecture_kana, city_id, city_name, city_kana, town_name, town_kana)
VALUES ("4630063",23,"愛知県","アイチケン",23113,"名古屋市守山区","ナゴヤシモリヤマク","八反","ハッタン");


