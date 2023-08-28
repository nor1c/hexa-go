CREATE TABLE `pets` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`kind_id` INT(11) UNSIGNED NOT NULL,
	`owner_id` INT(11) UNSIGNED NOT NULL,
	`name` VARCHAR(100) NOT NULL COLLATE 'latin1_swedish_ci',
	`age` INT(2) NULL DEFAULT NULL,
	`adoption_date` DATE NOT NULL,
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `FK_pets_pet_kinds` (`kind_id`) USING BTREE,
	INDEX `FK_pets_owners` (`owner_id`) USING BTREE,
	CONSTRAINT `FK_pets_owners` FOREIGN KEY (`owner_id`) REFERENCES `go-cardio-hexa-go`.`owners` (`id`) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT `FK_pets_pet_kinds` FOREIGN KEY (`kind_id`) REFERENCES `go-cardio-hexa-go`.`pet_kinds` (`id`) ON UPDATE CASCADE ON DELETE NO ACTION
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB;
