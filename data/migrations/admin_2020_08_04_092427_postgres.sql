ALTER TABLE goadmin_menu
ADD COLUMN plugin_name character varying(150) NOT NULL DEFAULT '',
ADD COLUMN uuid character varying(150) NOT NULL DEFAULT '';