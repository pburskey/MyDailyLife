SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL';
SHOW WARNINGS;

DROP SCHEMA IF EXISTS `mydailylife` ;
CREATE SCHEMA IF NOT EXISTS `mydailylife` ;
USE `mydailylife` ;

DROP TABLE IF EXISTS `mydailylife`.`person` ;

CREATE  TABLE IF NOT EXISTS `mydailylife`.`person` (
  `id` varchar(200) NOT NULL ,
  `first` VARCHAR(100) NOT NULL ,
  `last` VARCHAR(100) NOT NULL ,
  PRIMARY KEY (`id`) ,
  UNIQUE INDEX `person_id_UNIQUE` (`id` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = latin1;
SHOW WARNINGS;


DROP TABLE IF EXISTS `mydailylife`.`task` ;
CREATE  TABLE IF NOT EXISTS `mydailylife`.`task` (
    `id` varchar(200) NOT NULL ,
    `party_id` varchar(200) NOT NULL ,
    `name` VARCHAR(100) NOT NULL ,
  `description` VARCHAR(1024) NOT NULL ,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `task_id_UNIQUE` (`id` ASC),
    KEY `task_party_id` (`party_id`),
    CONSTRAINT `task_party_id` FOREIGN KEY (`party_id`) REFERENCES `person` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
    )

ENGINE = InnoDB
DEFAULT CHARACTER SET = latin1;
SHOW WARNINGS;



DROP TABLE IF EXISTS `mydailylife`.`status` ;
CREATE  TABLE IF NOT EXISTS `mydailylife`.`status` (
   `skey` int unsigned AUTO_INCREMENT NOT NULL ,
    `id` varchar(200) NOT NULL ,
    `ts` timestamp ,
    `status` int unsigned NOT NULL ,
    PRIMARY KEY (`skey`),
    UNIQUE INDEX `status_skey_UNIQUE` (`skey` ASC)
    )

    ENGINE = InnoDB
    DEFAULT CHARACTER SET = latin1;
SHOW WARNINGS;


DROP TABLE IF EXISTS `mydailylife`.`task_in_progress` ;
CREATE  TABLE IF NOT EXISTS `mydailylife`.`task_in_progress` (
    `id` varchar(200) NOT NULL ,
    `task_id` varchar(200) NOT NULL ,
    `creation` timestamp ,
    `description` VARCHAR(1024) NOT NULL ,
    `status_id` varchar(200) NOT NULL ,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `tip_id_UNIQUE` (`id` ASC),
    KEY `tip_task_id` (`task_id`),
    KEY `tip_status_id` (`status_id`),
    CONSTRAINT `tip_task_id` FOREIGN KEY (`task_id`) REFERENCES `task` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
    CONSTRAINT `tip_status_id` FOREIGN KEY (`status_id`) REFERENCES `status` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
    )

    ENGINE = InnoDB
    DEFAULT CHARACTER SET = latin1;
SHOW WARNINGS;



SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
