use user_data;

DROP TABLE IF EXISTS `users` ;

CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT NOT NULL Auto_increment,
  `shard_id` tinyint unsigned NOT NULL,
  `uuid_hash` varchar(255) NOT NULL,
  `register_date` DATETIME NOT NULL,
  `login_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

CREATE INDEX `IDX-user-login_at` ON `users` (`login_at` DESC);
