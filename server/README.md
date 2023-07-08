```azure
use chat;

create table auth (
	`id` BIGINT AUTO_INCREMENT primary key,
	`name` varchar(255) UNIQUE NOT NULL,
    `verified_email` varchar(255) UNIQUE NOT NULL, 
    `google_id` varchar(20) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table room (
	`room_hash` varchar(10) primary key,
    `owner_name` varchar(255) NOT NULL,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table room_participant (
	`room_hash` varchar(10) primary key,
	`user_name` varchar(255) NOT NULL
);

CREATE INDEX `auth_index_1` ON `auth` (`id`);
CREATE INDEX `room_index_1` ON `room` (`room_hash`);
CREATE INDEX `room_participant_index_1` ON `room` (`room_hash`);

ALTER TABLE `room` ADD FOREIGN KEY (`owner_name`) REFERENCES `auth` (`name`);
ALTER TABLE `room_participant` ADD FOREIGN KEY (`room_hash`) REFERENCES `room` (`room_hash`);
```