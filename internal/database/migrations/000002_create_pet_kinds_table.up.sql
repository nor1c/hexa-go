CREATE TABLE `pet_kinds` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(100) NOT NULL COLLATE 'latin1_swedish_ci',
	`leg` INT(1) NULL DEFAULT '0',
	`can_swim` INT(1) NULL DEFAULT '0',
	`can_fly` INT(1) NULL DEFAULT '0',
	`can_run` INT(1) NULL DEFAULT '0',
	`is_dangerous` INT(1) NULL DEFAULT '0',
	`is_venomous` INT(1) NULL DEFAULT '0',
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB;