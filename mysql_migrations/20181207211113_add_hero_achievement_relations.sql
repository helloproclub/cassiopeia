-- +mig Up
CREATE TABLE hero_achievement_relations(
  hero_id INT UNSIGNED NOT NULL,
  achievement_id INT UNSIGNED NOT NULL,
  FOREIGN KEY(hero_id) REFERENCES heros(id)
  FOREIGN KEY(achievement_id) REFERENCES achievements(id)
);

-- +mig Down
DROP TABLE hero_achievement_relations;
