--
-- PostgreSQL database dump
--

-- Dumped from database version 10.7 (Ubuntu 10.7-1.pgdg18.04+1)
-- Dumped by pg_dump version 10.7 (Ubuntu 10.7-1.pgdg18.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

CREATE TABLE public.goadmin_menu (
    id integer NOT NULL,
    parent_id integer DEFAULT 0 NOT NULL,
    "order" integer DEFAULT 0 NOT NULL,
    title character varying(50) DEFAULT ''::character varying NOT NULL,
    icon character varying(50) DEFAULT ''::character varying NOT NULL,
    uri character varying(50) DEFAULT ''::character varying NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    type integer DEFAULT 0 NOT NULL
);


CREATE TABLE public.goadmin_permissions (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    slug character varying(50) NOT NULL,
    http_method character varying(255) NOT NULL,
    http_path text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE SEQUENCE public.goadmin_permissions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


CREATE TABLE public.goadmin_role_menu (
    role_id integer NOT NULL,
    menu_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE TABLE public.goadmin_role_permissions (
    role_id integer NOT NULL,
    permission_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE TABLE public.goadmin_role_users (
    role_id integer NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE TABLE public.goadmin_roles (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    slug character varying(50) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE SEQUENCE public.goadmin_roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.goadmin_roles_id_seq OWNED BY public.goadmin_roles.id;


CREATE TABLE public.goadmin_users (
    id integer NOT NULL,
    username character varying(190) DEFAULT ''::character varying NOT NULL,
    password character varying(80) DEFAULT ''::character varying NOT NULL,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    avatar character varying(255) DEFAULT ''::character varying NOT NULL,
    remember_token character varying(100) DEFAULT ''::character varying NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


COPY public.goadmin_menu (id, parent_id, "order", title, icon, uri, created_at, updated_at, type) FROM stdin;
2	0	2	Admin	fa-tasks		2019-03-18 16:39:51.668419+00	2019-03-18 16:39:51.668419+00	1
3	2	2	Users	fa-users	/info/manager	2019-03-18 16:39:51.668419+00	2019-03-18 16:39:51.668419+00	1
4	2	3	Roles	fa-user	/info/roles	2019-03-18 16:39:51.668419+00	2019-03-18 16:39:51.668419+00	1
5	2	4	Permission	fa-ban	/info/permission	2019-03-18 16:39:51.668419+00	2019-03-18 16:39:51.668419+00	1
6	2	5	Menu	fa-bars	/menu	2019-03-18 16:39:51.668419+00	2019-03-18 16:39:51.668419+00	1
7	2	6	Operation log	fa-history	/info/op	2019-03-18 16:39:51.668419+00	2019-03-18 16:39:51.668419+00	1
12	0	7	Users	fa-user	/info/user	2019-03-18 16:39:51.668419+00	2019-03-18 16:39:51.668419+00	0
13	0	1	Dashboard	fa-bar-chart	/	2019-03-18 16:39:51.668419+00	2019-03-18 16:39:51.668419+00	1
16	0	8	Posts	fa-file-powerpoint-o	/info/posts	2019-03-18 16:39:51.668419+00	2019-03-18 16:39:51.668419+00	0
17	0	9	Authors	fa-users	/info/authors	2019-03-18 16:39:51.668419+00	2019-03-18 16:39:51.668419+00	0
18	0	10	Example Plugin	fa-plug	/example	2019-03-18 16:39:51.668419+00	2019-03-18 16:39:51.668419+00	0
\.


COPY public.goadmin_permissions (id, name, slug, http_method, http_path, created_at, updated_at) FROM stdin;
1	All permission	*		*	2019-03-18 16:17:59.490537+00	2019-03-18 16:17:59.490537+00
2	Dashboard	dashboard	GET,PUT,POST,DELETE	/	2019-03-18 16:17:59.490537+00	2019-03-18 16:17:59.490537+00
3	Login	auth.login		/auth/login\\r\\n/auth/logout	2019-03-18 16:17:59.490537+00	2019-03-18 16:17:59.490537+00
4	User setting	auth.setting	GET,PUT	/auth/setting	2019-03-18 16:17:59.490537+00	2019-03-18 16:17:59.490537+00
5	Auth management	auth.management		/auth/roles\\r\\n/auth/permissions\\r\\n/auth/menu\\r\\n/auth/logs	2019-03-18 16:17:59.490537+00	2019-03-18 16:17:59.490537+00
\.


COPY public.goadmin_role_menu (role_id, menu_id, created_at, updated_at) FROM stdin;
1	1	2019-03-18 18:27:02.664973+00	2019-03-18 18:27:02.664973+00
1	11	2019-03-18 18:27:02.664973+00	2019-03-18 18:27:02.664973+00
\.


COPY public.goadmin_role_permissions (role_id, permission_id, created_at, updated_at) FROM stdin;
1	1	2019-03-18 18:21:32.596272+00	2019-03-18 18:21:32.596272+00
1	2	2019-03-18 18:21:32.596272+00	2019-03-18 18:21:32.596272+00
\.


COPY public.goadmin_role_users (role_id, user_id, created_at, updated_at) FROM stdin;
1	1	2019-03-18 18:02:02.210694+00	2019-03-18 18:02:02.210694+00
\.


COPY public.goadmin_roles (id, name, slug, created_at, updated_at) FROM stdin;
1	Administrator	administrator	2019-03-18 14:46:43.812101+00	2019-03-18 14:46:43.812101+00
\.


COPY public.goadmin_users (id, username, password, name, avatar, remember_token, created_at, updated_at) FROM stdin;
1	admin	$2a$10$FgN6YzUL1UZ4IjgDj/Fb5uaKb6zRXISTkM7/vonco1RxpvxMIrzdS	admin		tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh	2019-03-18 17:27:38.106738+00	2019-03-18 17:27:38.106738+00
\.


SELECT pg_catalog.setval('public.goadmin_permissions_id_seq', 1, false);


SELECT pg_catalog.setval('public.goadmin_roles_id_seq', 1, false);


--
-- PostgreSQL database dump complete
--

