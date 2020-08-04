ALTER TABLE goadmin_menu
ADD COLUMN `uuid` varchar(150) DEFAULT NULL,
ADD COLUMN `plugin_name` varchar(150) NOT NULL DEFAULT '';