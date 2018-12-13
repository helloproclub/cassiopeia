-- +mig Up
CREATE TABLE achievements(
  id INT UNSIGNED AUTO_INCREMENT,
  published DATETIME,
  updated DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  created DATETIME DEFAULT CURRENT_TIMESTAMP,
  issuer VARCHAR(50),
  title VARCHAR(80), 
  attainment ENUM('juara', 'penghargaan', 'kompetisi', 'penelitian', 'beasiswa', 'lain-lain'),
  PRIMARY KEY(id)
);

-- +mig Down
DROP TABLE achievements;
