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
    `room_hash` varchar(255) primary key,
    `owner_name` varchar(255) NOT NULL,
    `is_broad_cast` bool default false,
    `before_broad_cast` timestamp default 0,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX `auth_index_1` ON `auth` (`id`);
CREATE INDEX `room_index_1` ON `room` (`room_hash`);

ALTER TABLE `room` ADD FOREIGN KEY (`owner_name`) REFERENCES `auth` (`name`);
```