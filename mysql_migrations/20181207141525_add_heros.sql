-- +mig Up
CREATE TABLE heros(
  id INT UNSIGNED AUTO_INCREMENT,
  updated DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  created DATETIME DEFAULT CURRENT_TIMESTAMP, 
  username VARCHAR(30) UNIQUE,
  fullname VARCHAR(70) UNIQUE,
  power ENUM('programmer',  'business analyst', 'designer'),
  bio VARCHAR(160) DEFAULT '',
  proclub_year INT(4),
  facebook_username VARCHAR(50),
  twitter_username VARCHAR(15),
  instagram_username VARCHAR(30),
  telegram_username VARCHAR(32),
  line_id VARCHAR(50),
  PRIMARY KEY(id),
  INDEX(username)
);

-- +mig Down
DROP TABLE heros;
