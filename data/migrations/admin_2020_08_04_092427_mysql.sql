ALTER TABLE goadmin_menu
ADD COLUMN `uuid` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
ADD COLUMN `plugin_name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '';