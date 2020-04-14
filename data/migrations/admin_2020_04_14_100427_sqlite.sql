CREATE TABLE IF NOT EXISTS "goadmin_site" (
`id` integer PRIMARY KEY autoincrement,
`key` CHAR(100) COLLATE NOCASE NOT NULL,
`value` text COLLATE NOCASE NOT NULL,
`state` INT NOT NULL DEFAULT '0',
`description` CHAR(3000) COLLATE NOCASE,
`created_at` TIMESTAMP default CURRENT_TIMESTAMP,
`updated_at` TIMESTAMP default CURRENT_TIMESTAMP
);