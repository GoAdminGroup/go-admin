--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.14
-- Dumped by pg_dump version 10.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'EUC_CN';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- Name: goadmin_menu_myid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.goadmin_menu_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_menu_myid_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: goadmin_menu; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goadmin_menu (
    id integer DEFAULT nextval('public.goadmin_menu_myid_seq'::regclass) NOT NULL,
    parent_id integer DEFAULT 0 NOT NULL,
    type integer DEFAULT 0,
    "order" integer DEFAULT 0 NOT NULL,
    title character varying(50) NOT NULL,
    header character varying(100),
    icon character varying(50) NOT NULL,
    uri character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_menu OWNER TO postgres;

--
-- Name: goadmin_operation_log_myid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.goadmin_operation_log_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_operation_log_myid_seq OWNER TO postgres;

--
-- Name: goadmin_operation_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goadmin_operation_log (
    id integer DEFAULT nextval('public.goadmin_operation_log_myid_seq'::regclass) NOT NULL,
    user_id integer NOT NULL,
    path character varying(255) NOT NULL,
    method character varying(10) NOT NULL,
    ip character varying(15) NOT NULL,
    input text NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_operation_log OWNER TO postgres;

--
-- Name: goadmin_permissions_myid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.goadmin_permissions_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_permissions_myid_seq OWNER TO postgres;

--
-- Name: goadmin_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goadmin_permissions (
    id integer DEFAULT nextval('public.goadmin_permissions_myid_seq'::regclass) NOT NULL,
    name character varying(50) NOT NULL,
    slug character varying(50) NOT NULL,
    http_method character varying(255),
    http_path text NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_permissions OWNER TO postgres;

--
-- Name: goadmin_role_menu; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goadmin_role_menu (
    role_id integer NOT NULL,
    menu_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_role_menu OWNER TO postgres;

--
-- Name: goadmin_role_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goadmin_role_permissions (
    role_id integer NOT NULL,
    permission_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_role_permissions OWNER TO postgres;

--
-- Name: goadmin_role_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goadmin_role_users (
    role_id integer NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_role_users OWNER TO postgres;

--
-- Name: goadmin_roles_myid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.goadmin_roles_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_roles_myid_seq OWNER TO postgres;

--
-- Name: goadmin_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goadmin_roles (
    id integer DEFAULT nextval('public.goadmin_roles_myid_seq'::regclass) NOT NULL,
    name character varying NOT NULL,
    slug character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_roles OWNER TO postgres;

--
-- Name: goadmin_session_myid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.goadmin_session_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_session_myid_seq OWNER TO postgres;

--
-- Name: goadmin_session; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goadmin_session (
    id integer DEFAULT nextval('public.goadmin_session_myid_seq'::regclass) NOT NULL,
    sid character varying(50) NOT NULL,
    "values" character varying(3000) NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_session OWNER TO postgres;

--
-- Name: goadmin_user_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goadmin_user_permissions (
    user_id integer NOT NULL,
    permission_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_user_permissions OWNER TO postgres;

--
-- Name: goadmin_users_myid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.goadmin_users_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_users_myid_seq OWNER TO postgres;

--
-- Name: goadmin_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goadmin_users (
    id integer DEFAULT nextval('public.goadmin_users_myid_seq'::regclass) NOT NULL,
    username character varying(190) NOT NULL,
    password character varying(80) NOT NULL,
    name character varying(255) NOT NULL,
    avatar character varying(255),
    remember_token character varying(100),
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_users OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(100),
    homepage character varying(3000),
    email character varying(100),
    birthday timestamp with time zone,
    country character varying(50),
    city character varying(50),
    password character varying(100),
    ip character varying(20),
    certificate character varying(300),
    money integer,
    resume text,
    gender smallint,
    fruit character varying(200),
    drink character varying(200),
    experience smallint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Data for Name: goadmin_menu; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goadmin_menu (id, parent_id, type, "order", title, header, icon, uri, created_at, updated_at) FROM stdin;
1	0	1	2	Admin	\N	fa-tasks		2019-09-10 00:00:00	2019-09-10 00:00:00
2	1	1	2	Users	\N	fa-users	/info/manager	2019-09-10 00:00:00	2019-09-10 00:00:00
4	1	1	4	Permission	\N	fa-ban	/info/permission	2019-09-10 00:00:00	2019-09-10 00:00:00
5	1	1	5	Menu	\N	fa-bars	/menu	2019-09-10 00:00:00	2019-09-10 00:00:00
6	1	1	6	Operation log	\N	fa-history	/info/op	2019-09-10 00:00:00	2019-09-10 00:00:00
7	0	1	1	Dashboard	\N	fa-bar-chart	/	2019-09-10 00:00:00	2019-09-10 00:00:00
8	0	0	2	test menu		fa-angellist	/example/test	2019-11-27 22:26:12.849014	2019-11-27 22:26:12.849014
3	0	1	3	test2 menu		fa-angellist	/example/test	2019-09-10 00:00:00	2019-09-10 00:00:00
\.


--
-- Data for Name: goadmin_operation_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goadmin_operation_log (id, user_id, path, method, ip, input, created_at, updated_at) FROM stdin;
1	1	/admin/logout	GET	127.0.0.1		2019-11-27 22:26:11.753822	2019-11-27 22:26:11.753822
2	1	/admin/info/permission	GET	127.0.0.1		2019-11-27 22:26:11.955262	2019-11-27 22:26:11.955262
3	1	/admin/info/permission/new	GET	127.0.0.1		2019-11-27 22:26:11.977105	2019-11-27 22:26:11.977105
4	1	/admin/new/permission	POST	127.0.0.1	{"_previous_":["/admin/info/permission?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc"],"_t":["99777d7f-52a4-4b7e-af01-e56e15a495c3"],"http_method[]":["GET"],"http_path":["/\\n/admin/info/op"],"id":["3"],"name":["tester"],"slug":["tester"]}	2019-11-27 22:26:12.0068	2019-11-27 22:26:12.0068
5	1	/admin/info/permission/edit	GET	127.0.0.1		2019-11-27 22:26:12.020841	2019-11-27 22:26:12.020841
6	1	/admin/info/permission/edit	GET	127.0.0.1		2019-11-27 22:26:12.041507	2019-11-27 22:26:12.041507
7	1	/admin/edit/permission	POST	127.0.0.1	{"_previous_":["/admin/info/permission?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc"],"_t":["5cf493a5-f548-4bfd-9dd9-3e80eb7057b3"],"http_method[]":["GET","POST"],"http_path":["/\\n/admin/info/op"],"id":["3"],"name":["tester"],"slug":["tester"]}	2019-11-27 22:26:12.070701	2019-11-27 22:26:12.070701
8	1	/admin/info/permission/new	GET	127.0.0.1		2019-11-27 22:26:12.090218	2019-11-27 22:26:12.090218
9	1	/admin/new/permission	POST	127.0.0.1	{"_previous_":["/admin/info/permission?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc"],"_t":["39ee2f7b-836c-4d7e-80fa-8da41ab017ff"],"http_method[]":["GET"],"http_path":["/\\n/admin/info/op"],"id":["4"],"name":["tester2"],"slug":["tester2"]}	2019-11-27 22:26:12.11848	2019-11-27 22:26:12.11848
10	1	/admin/delete/permission	POST	127.0.0.1	{"id":["4"]}	2019-11-27 22:26:12.126336	2019-11-27 22:26:12.126336
11	1	/admin/info/roles	GET	127.0.0.1		2019-11-27 22:26:12.141837	2019-11-27 22:26:12.141837
12	1	/admin/info/roles/new	GET	127.0.0.1		2019-11-27 22:26:12.158497	2019-11-27 22:26:12.158497
13	1	/admin/new/roles	POST	127.0.0.1	{"_previous_":["/admin/info/roles?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc"],"_t":["1d4a2838-85e1-479e-ae8b-b22081bd4472"],"name":["tester"],"permission_id[]":["3"],"slug":["tester"]}	2019-11-27 22:26:12.180896	2019-11-27 22:26:12.180896
14	1	/admin/info/roles/edit	GET	127.0.0.1		2019-11-27 22:26:12.193585	2019-11-27 22:26:12.193585
15	1	/admin/info/roles/edit	GET	127.0.0.1		2019-11-27 22:26:12.211173	2019-11-27 22:26:12.211173
16	1	/admin/edit/roles	POST	127.0.0.1	{"_previous_":["/admin/info/roles?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc"],"_t":["06e48f62-b223-40d5-a0b0-be4c8c2e6ccb"],"id":["3"],"name":["tester"],"permission_id[]":["3","2"],"slug":["tester"]}	2019-11-27 22:26:12.237336	2019-11-27 22:26:12.237336
17	1	/admin/info/roles/new	GET	127.0.0.1		2019-11-27 22:26:12.257681	2019-11-27 22:26:12.257681
18	1	/admin/new/roles	POST	127.0.0.1	{"_previous_":["/admin/info/roles?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc"],"_t":["ce5bfea2-72d0-4a32-97cb-5ea07497f7ad"],"name":["tester2"],"permission_id[]":["3"],"slug":["tester2"]}	2019-11-27 22:26:12.27778	2019-11-27 22:26:12.27778
19	1	/admin/delete/roles	POST	127.0.0.1	{"id":["3"]}	2019-11-27 22:26:12.283768	2019-11-27 22:26:12.283768
20	1	/admin/info/manager	GET	127.0.0.1		2019-11-27 22:26:12.300918	2019-11-27 22:26:12.300918
21	1	/admin/edit/manager	POST	127.0.0.1	{"_previous_":["/admin/info/manager?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc"],"_t":["123"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}	2019-11-27 22:26:12.311142	2019-11-27 22:26:12.311142
22	1	/admin/info/manager/edit	GET	127.0.0.1		2019-11-27 22:26:12.318903	2019-11-27 22:26:12.318903
23	1	/admin/info/manager/edit	GET	127.0.0.1		2019-11-27 22:26:12.337205	2019-11-27 22:26:12.337205
24	1	/admin/edit/manager	POST	127.0.0.1	{"_previous_":["/admin/info/manager?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc"],"_t":["11a2e42e-b4be-4acc-92fe-a018a1c6475e"],"avatar":[""],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}	2019-11-27 22:26:12.438473	2019-11-27 22:26:12.438473
25	1	/admin/info/manager/new	GET	127.0.0.1		2019-11-27 22:26:12.459245	2019-11-27 22:26:12.459245
26	1	/admin/new/manager	POST	127.0.0.1	{"_previous_":["/admin/info/manager?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc"],"_t":["b42e934b-59d8-416c-9377-5173c354a1c3"],"avatar":[""],"id":["1"],"name":["tester"],"password":["tester"],"password_again":["tester"],"permission_id[]":["1"],"role_id[]":["1"],"username":["tester"]}	2019-11-27 22:26:12.586842	2019-11-27 22:26:12.586842
27	1	/admin/menu	GET	127.0.0.1		2019-11-27 22:26:12.842035	2019-11-27 22:26:12.842035
28	1	/admin/menu/new	POST	127.0.0.1	{"_previous_":["/admin/menu"],"_t":["0b2d5e83-204d-4ff0-bcba-a1f4aa96b0bd"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test menu"],"uri":["/example/test"]}	2019-11-27 22:26:12.881608	2019-11-27 22:26:12.881608
29	1	/admin/menu/edit/show	GET	127.0.0.1		2019-11-27 22:26:12.908306	2019-11-27 22:26:12.908306
30	1	/admin/menu/edit/show	GET	127.0.0.1		2019-11-27 22:26:12.936097	2019-11-27 22:26:12.936097
31	1	/admin/menu/edit	POST	127.0.0.1	{"_previous_":["/admin/menu"],"_t":["934c2003-acb8-471e-859d-a1e5bd7f4ae3"],"header":[""],"icon":["fa-angellist"],"id":["3"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test"]}	2019-11-27 22:26:12.97449	2019-11-27 22:26:12.97449
32	1	/admin/menu/new	POST	127.0.0.1	{"_previous_":["/admin/menu"],"_t":["934c2003-acb8-471e-859d-a1e5bd7f4ae3"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test2"]}	2019-11-27 22:26:13.005189	2019-11-27 22:26:13.005189
33	1	/admin/menu/delete	POST	127.0.0.1		2019-11-27 22:26:13.020411	2019-11-27 22:26:13.020411
34	1	/admin/info/op	GET	127.0.0.1		2019-11-27 22:26:13.045159	2019-11-27 22:26:13.045159
\.


--
-- Data for Name: goadmin_permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goadmin_permissions (id, name, slug, http_method, http_path, created_at, updated_at) FROM stdin;
1	All permission	*		*	2019-09-10 00:00:00	2019-09-10 00:00:00
2	Dashboard	dashboard	GET,PUT,POST,DELETE	/	2019-09-10 00:00:00	2019-09-10 00:00:00
3	tester	tester	GET,POST	/\n/admin/info/op	2019-11-27 22:26:11.992207	2019-11-27 22:26:11.992207
\.


--
-- Data for Name: goadmin_role_menu; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goadmin_role_menu (role_id, menu_id, created_at, updated_at) FROM stdin;
1	1	2019-09-10 00:00:00	2019-09-10 00:00:00
1	7	2019-09-10 00:00:00	2019-09-10 00:00:00
2	7	2019-09-10 00:00:00	2019-09-10 00:00:00
1	8	2019-11-27 22:26:12.850341	2019-11-27 22:26:12.850341
1	3	2019-11-27 22:26:12.942714	2019-11-27 22:26:12.942714
\.


--
-- Data for Name: goadmin_role_permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goadmin_role_permissions (role_id, permission_id, created_at, updated_at) FROM stdin;
1	1	2019-09-10 00:00:00	2019-09-10 00:00:00
1	2	2019-09-10 00:00:00	2019-09-10 00:00:00
2	2	2019-09-10 00:00:00	2019-09-10 00:00:00
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
4	3	2019-11-27 22:26:12.265461	2019-11-27 22:26:12.265461
\.


--
-- Data for Name: goadmin_role_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goadmin_role_users (role_id, user_id, created_at, updated_at) FROM stdin;
2	2	2019-09-10 00:00:00	2019-09-10 00:00:00
1	1	2019-11-27 22:26:12.424216	2019-11-27 22:26:12.424216
1	3	2019-11-27 22:26:12.561067	2019-11-27 22:26:12.561067
\.


--
-- Data for Name: goadmin_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goadmin_roles (id, name, slug, created_at, updated_at) FROM stdin;
1	Administrator	administrator	2019-09-10 00:00:00	2019-09-10 00:00:00
2	Operator	operator	2019-09-10 00:00:00	2019-09-10 00:00:00
4	tester2	tester2	2019-11-27 22:26:12.264253	2019-11-27 22:26:12.264253
\.


--
-- Data for Name: goadmin_session; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goadmin_session (id, sid, "values", created_at, updated_at) FROM stdin;
2	f5a99916-36c8-4fd6-8873-6f2be8845cd0	{"user_id":1}	2019-11-27 22:26:11.917665	2019-11-27 22:26:11.917665
3	03263ffc-0043-4b89-a02f-3aa616bbf857	{"user_id":3}	2019-11-27 22:26:12.819931	2019-11-27 22:26:12.819931
\.


--
-- Data for Name: goadmin_user_permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goadmin_user_permissions (user_id, permission_id, created_at, updated_at) FROM stdin;
2	2	2019-09-10 00:00:00	2019-09-10 00:00:00
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
1	1	2019-11-27 22:26:12.425769	2019-11-27 22:26:12.425769
3	1	2019-11-27 22:26:12.572997	2019-11-27 22:26:12.572997
\.


--
-- Data for Name: goadmin_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goadmin_users (id, username, password, name, avatar, remember_token, created_at, updated_at) FROM stdin;
2	operator	$2a$10$rVqkOzHjN2MdlEprRflb1eGP0oZXuSrbJLOmJagFsCd81YZm0bsh.	Operator		\N	2019-09-10 00:00:00	2019-09-10 00:00:00
1	admin	$2a$10$wOkJ8fygob.TbIJCOqCaaOfb9y8YbJtWpfvCicydBPKsc1su0GPN6	admin1		tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh	2019-09-10 00:00:00	2019-09-10 00:00:00
3	tester	$2a$10$IKq3UBcm6ZRLVY0gLGWFKuB4Japc9SX0HeHBQXfRYmgdRHjguDUWu	tester		\N	2019-11-27 22:26:12.554012	2019-11-27 22:26:12.554012
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, homepage, email, birthday, country, city, password, ip, certificate, money, resume, gender, fruit, drink, experience, created_at, updated_at) FROM stdin;
1	jack	http://jack.me	jack@163.com	1993-10-21 00:00:00+08	china	guangzhou	123456	127.0.0.1	\N	10	<h1>Jacks Resume</h1>	1	apple	water	0	2020-03-09 15:24:00+08	2020-03-09 15:24:00+08
\.


--
-- Name: goadmin_menu_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.goadmin_menu_myid_seq', 8, true);


--
-- Name: goadmin_operation_log_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.goadmin_operation_log_myid_seq', 34, true);


--
-- Name: goadmin_permissions_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.goadmin_permissions_myid_seq', 4, true);


--
-- Name: goadmin_roles_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.goadmin_roles_myid_seq', 4, true);


--
-- Name: goadmin_session_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.goadmin_session_myid_seq', 3, true);


--
-- Name: goadmin_users_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.goadmin_users_myid_seq', 3, true);


--
-- Name: goadmin_menu goadmin_menu_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goadmin_menu
    ADD CONSTRAINT goadmin_menu_pkey PRIMARY KEY (id);


--
-- Name: goadmin_operation_log goadmin_operation_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goadmin_operation_log
    ADD CONSTRAINT goadmin_operation_log_pkey PRIMARY KEY (id);


--
-- Name: goadmin_permissions goadmin_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goadmin_permissions
    ADD CONSTRAINT goadmin_permissions_pkey PRIMARY KEY (id);


--
-- Name: goadmin_roles goadmin_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goadmin_roles
    ADD CONSTRAINT goadmin_roles_pkey PRIMARY KEY (id);


--
-- Name: goadmin_session goadmin_session_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goadmin_session
    ADD CONSTRAINT goadmin_session_pkey PRIMARY KEY (id);


--
-- Name: goadmin_users goadmin_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goadmin_users
    ADD CONSTRAINT goadmin_users_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

