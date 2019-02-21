use master_data;

DROP TABLE IF EXISTS `mantenance_infos` ;

CREATE TABLE IF NOT EXISTS `mantenance_infos` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `state` TINYINT NOT NULL, 
  `start_dt` DATETIME NULL,
  `end_dt` DATETIME NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;
