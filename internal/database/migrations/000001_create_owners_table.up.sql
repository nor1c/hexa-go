CREATE TABLE `owners` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(100) NOT NULL COLLATE 'latin1_swedish_ci',
	`age` INT(3) NOT NULL,
	`address` TEXT NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB;
