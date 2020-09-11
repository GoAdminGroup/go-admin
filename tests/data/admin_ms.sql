/*
 Navicat Premium Data Transfer

 Source Server         : mssql
 Source Server Type    : SQL Server
 Source Server Version : 14003281
 Source Host           : localhost
 Source Database       : goadmin
 Source Schema         : dbo

 Target Server Type    : SQL Server
 Target Server Version : 14003281
 File Encoding         : utf-8

 Date: 04/08/2020 01:11:22 AM
*/

-- ----------------------------
--  Table structure for goadmin_menu
-- ----------------------------
IF EXISTS (SELECT * FROM sys.all_objects WHERE object_id = OBJECT_ID('[dbo].[goadmin_menu]') AND type IN ('U'))
	DROP TABLE [dbo].[goadmin_menu]
GO
CREATE TABLE [dbo].[goadmin_menu] (
	[id] int IDENTITY(1,1) NOT NULL,
	[parent_id] int NOT NULL DEFAULT ((0)),
	[type] tinyint NOT NULL DEFAULT ((0)),
	[order] int NOT NULL DEFAULT ('0'),
	[title] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	[icon] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	[uri] varchar(3000) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL DEFAULT '',
	[header] varchar(150) COLLATE SQL_Latin1_General_CP1_CI_AS NULL DEFAULT NULL,
	[uuid] varchar(150) COLLATE SQL_Latin1_General_CP1_CI_AS NULL DEFAULT NULL,
	[plugin_name] varchar(150) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL DEFAULT '',
	[created_at] datetime NULL DEFAULT (getdate()),
	[updated_at] datetime NULL DEFAULT (getdate())
)
ON [PRIMARY]
GO

-- ----------------------------
--  Records of goadmin_menu
-- ----------------------------
BEGIN TRANSACTION
GO
SET IDENTITY_INSERT [dbo].[goadmin_menu] ON
GO
INSERT INTO [dbo].[goadmin_menu] ([id], [parent_id], [type], [order], [title], [icon], [uri], [header], [created_at], [updated_at]) VALUES ('1', '0', '1', '2', 'Admin', 'fa-tasks', '', null, '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
INSERT INTO [dbo].[goadmin_menu] ([id], [parent_id], [type], [order], [title], [icon], [uri], [header], [created_at], [updated_at]) VALUES ('2', '1', '1', '2', 'Users', 'fa-users', '/info/manager', null, '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
INSERT INTO [dbo].[goadmin_menu] ([id], [parent_id], [type], [order], [title], [icon], [uri], [header], [created_at], [updated_at]) VALUES ('3', '0', '1', '3', 'test2 menu', 'fa-angellist', '/example/test', '', '2019-09-10 00:00:00.000', '2020-03-15 12:24:47.000');
INSERT INTO [dbo].[goadmin_menu] ([id], [parent_id], [type], [order], [title], [icon], [uri], [header], [created_at], [updated_at]) VALUES ('4', '1', '1', '4', 'Permission', 'fa-ban', '/info/permission', null, '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
INSERT INTO [dbo].[goadmin_menu] ([id], [parent_id], [type], [order], [title], [icon], [uri], [header], [created_at], [updated_at]) VALUES ('5', '1', '1', '5', 'Menu', 'fa-bars', '/menu', null, '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
INSERT INTO [dbo].[goadmin_menu] ([id], [parent_id], [type], [order], [title], [icon], [uri], [header], [created_at], [updated_at]) VALUES ('6', '1', '1', '6', 'Operation log', 'fa-history', '/info/op', null, '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
INSERT INTO [dbo].[goadmin_menu] ([id], [parent_id], [type], [order], [title], [icon], [uri], [header], [created_at], [updated_at]) VALUES ('7', '0', '1', '1', 'Dashboard', 'fa-bar-chart', '/', null, '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
INSERT INTO [dbo].[goadmin_menu] ([id], [parent_id], [type], [order], [title], [icon], [uri], [header], [created_at], [updated_at]) VALUES ('8', '0', '1', '7', 'User', 'fa-users', '/info/user', '', '2020-03-15 03:09:14.810', '2020-03-15 03:09:14.810');
GO
SET IDENTITY_INSERT [dbo].[goadmin_menu] OFF
GO
COMMIT
GO

-- ----------------------------
--  Table structure for goadmin_operation_log
-- ----------------------------
IF EXISTS (SELECT * FROM sys.all_objects WHERE object_id = OBJECT_ID('[dbo].[goadmin_operation_log]') AND type IN ('U'))
	DROP TABLE [dbo].[goadmin_operation_log]
GO
CREATE TABLE [dbo].[goadmin_operation_log] (
	[id] int IDENTITY(1,1) NOT NULL,
	[user_id] int NOT NULL,
	[path] varchar(255) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	[method] varchar(10) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	[ip] varchar(15) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	[input] text COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	[created_at] datetime NULL DEFAULT (getdate()),
	[updated_at] datetime NULL DEFAULT (getdate())
)
ON [PRIMARY]
TEXTIMAGE_ON [PRIMARY]
GO

-- ----------------------------
--  Records of goadmin_operation_log
-- ----------------------------
BEGIN TRANSACTION
GO
SET IDENTITY_INSERT [dbo].[goadmin_operation_log] ON
GO
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('1', '1', '/admin/logout', 'GET', '127.0.0.1', '', '2020-03-15 03:09:12.760', '2020-03-15 03:09:12.760');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('2', '1', '/admin/info/permission', 'GET', '127.0.0.1', '', '2020-03-15 03:09:13.177', '2020-03-15 03:09:13.177');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('3', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 03:09:13.210', '2020-03-15 03:09:13.210');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('4', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["5f810995-beca-4c0f-8802-b327bef75743"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester"],"slug":["tester"]}', '2020-03-15 03:09:13.253', '2020-03-15 03:09:13.253');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('5', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:09:13.283', '2020-03-15 03:09:13.283');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('6', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:09:13.317', '2020-03-15 03:09:13.317');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('7', '1', '/admin/edit/permission', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["11114d9c-3bf2-4670-8638-09ee291d0233"],"http_method[]":["GET","POST"],"http_path":["/\n/admin/info/op"],"id":["3"],"name":["tester"],"slug":["tester"]}', '2020-03-15 03:09:13.363', '2020-03-15 03:09:13.363');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('8', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 03:09:13.400', '2020-03-15 03:09:13.400');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('9', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["9ae81396-6250-4e53-85f6-a459471f1bbf"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester2"],"slug":["tester2"]}', '2020-03-15 03:09:13.450', '2020-03-15 03:09:13.450');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('10', '1', '/admin/delete/permission', 'POST', '127.0.0.1', '{"id":["4"]}', '2020-03-15 03:09:13.487', '2020-03-15 03:09:13.487');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('11', '1', '/admin/info/roles', 'GET', '127.0.0.1', '', '2020-03-15 03:09:13.547', '2020-03-15 03:09:13.547');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('12', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 03:09:13.590', '2020-03-15 03:09:13.590');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('13', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["c3061c30-e384-4938-b5f7-6f07a11d04ac"],"name":["tester"],"permission_id[]":["3"],"slug":["tester"]}', '2020-03-15 03:09:13.660', '2020-03-15 03:09:13.660');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('14', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:09:13.690', '2020-03-15 03:09:13.690');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('15', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:09:13.733', '2020-03-15 03:09:13.733');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('16', '1', '/admin/edit/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["1e69d7a2-6e01-42f6-95a2-a5fba51e29a8"],"id":["3"],"name":["tester"],"permission_id[]":["3","2"],"slug":["tester"]}', '2020-03-15 03:09:13.800', '2020-03-15 03:09:13.800');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('17', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 03:09:13.833', '2020-03-15 03:09:13.833');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('18', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["79d55f58-03f2-4ad3-940d-b88174d3783e"],"name":["tester2"],"permission_id[]":["3"],"slug":["tester2"]}', '2020-03-15 03:09:13.893', '2020-03-15 03:09:13.893');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('19', '1', '/admin/delete/roles', 'POST', '127.0.0.1', '{"id":["3"]}', '2020-03-15 03:09:13.937', '2020-03-15 03:09:13.937');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('20', '1', '/admin/info/manager', 'GET', '127.0.0.1', '', '2020-03-15 03:09:14.010', '2020-03-15 03:09:14.010');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('21', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["123"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 03:09:14.040', '2020-03-15 03:09:14.040');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('22', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:09:14.070', '2020-03-15 03:09:14.070');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('23', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:09:14.130', '2020-03-15 03:09:14.130');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('24', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["44274eaa-9b15-43f1-9573-ec5c322f8e5f"],"avatar__delete_flag":["0"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 03:09:14.280', '2020-03-15 03:09:14.280');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('25', '1', '/admin/info/manager/new', 'GET', '127.0.0.1', '', '2020-03-15 03:09:14.317', '2020-03-15 03:09:14.317');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('26', '1', '/admin/new/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["71687e2d-cad3-48b5-be5a-d90b65a2891a"],"avatar__delete_flag":["0"],"id":["1"],"name":["tester"],"password":["tester"],"password_again":["tester"],"permission_id[]":["1"],"role_id[]":["1"],"username":["tester"]}', '2020-03-15 03:09:14.487', '2020-03-15 03:09:14.487');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('27', '1', '/admin/menu', 'GET', '127.0.0.1', '', '2020-03-15 03:09:14.780', '2020-03-15 03:09:14.780');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('28', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["e1c920bc-9131-4486-87c6-a332120c9efe"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test menu"],"uri":["/example/test"]}', '2020-03-15 03:09:14.863', '2020-03-15 03:09:14.863');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('29', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 03:09:14.910', '2020-03-15 03:09:14.910');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('30', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 03:09:14.967', '2020-03-15 03:09:14.967');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('31', '1', '/admin/menu/edit', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["d7eebcbe-a6c3-45ea-a5d0-d3af5bebe37f"],"header":[""],"icon":["fa-angellist"],"id":["3"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test"]}', '2020-03-15 03:09:15.047', '2020-03-15 03:09:15.047');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('32', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["d7eebcbe-a6c3-45ea-a5d0-d3af5bebe37f"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test2"]}', '2020-03-15 03:09:15.103', '2020-03-15 03:09:15.103');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('33', '1', '/admin/menu/delete', 'POST', '127.0.0.1', '{}', '2020-03-15 03:09:15.137', '2020-03-15 03:09:15.137');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('34', '1', '/admin/info/op', 'GET', '127.0.0.1', '', '2020-03-15 03:09:15.217', '2020-03-15 03:09:15.217');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('35', '1', '/admin/info/external', 'GET', '127.0.0.1', '', '2020-03-15 03:09:15.247', '2020-03-15 03:09:15.247');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('36', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:09:15.280', '2020-03-15 03:09:15.280');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('37', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:09:15.320', '2020-03-15 03:09:15.320');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('38', '1', '/admin/info/external/new', 'GET', '127.0.0.1', '', '2020-03-15 03:09:15.353', '2020-03-15 03:09:15.353');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('39', '1', '/admin/info/user', 'GET', '127.0.0.1', '', '2020-03-15 03:09:15.400', '2020-03-15 03:09:15.400');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('40', '1', '/admin/export/user', 'POST', '127.0.0.1', '{"id":["1"]}', '2020-03-15 03:09:15.443', '2020-03-15 03:09:15.443');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('41', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:09:15.483', '2020-03-15 03:09:15.483');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('42', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:09:15.520', '2020-03-15 03:09:15.520');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('43', '1', '/admin/info/user/new', 'GET', '127.0.0.1', '', '2020-03-15 03:09:15.550', '2020-03-15 03:09:15.550');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('44', '1', '/admin/logout', 'GET', '127.0.0.1', '', '2020-03-15 03:14:36.897', '2020-03-15 03:14:36.897');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('45', '1', '/admin/info/permission', 'GET', '127.0.0.1', '', '2020-03-15 03:14:37.297', '2020-03-15 03:14:37.297');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('46', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 03:14:37.330', '2020-03-15 03:14:37.330');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('47', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["d66311db-d66a-403b-9154-60710429a1cb"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester"],"slug":["tester"]}', '2020-03-15 03:14:37.363', '2020-03-15 03:14:37.363');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('48', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:14:37.400', '2020-03-15 03:14:37.400');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('49', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:14:37.430', '2020-03-15 03:14:37.430');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('50', '1', '/admin/edit/permission', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["a95d2d24-fb3e-4020-b1a2-e7e50a94ff19"],"http_method[]":["GET","POST"],"http_path":["/\n/admin/info/op"],"id":["3"],"name":["tester"],"slug":["tester"]}', '2020-03-15 03:14:37.483', '2020-03-15 03:14:37.483');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('51', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 03:14:37.517', '2020-03-15 03:14:37.517');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('52', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["5c3ebaad-2459-43d5-9536-bb8b69d5ded1"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester2"],"slug":["tester2"]}', '2020-03-15 03:14:37.563', '2020-03-15 03:14:37.563');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('53', '1', '/admin/delete/permission', 'POST', '127.0.0.1', '{"id":["4"]}', '2020-03-15 03:14:37.593', '2020-03-15 03:14:37.593');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('54', '1', '/admin/info/roles', 'GET', '127.0.0.1', '', '2020-03-15 03:14:37.640', '2020-03-15 03:14:37.640');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('55', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 03:14:37.700', '2020-03-15 03:14:37.700');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('56', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["c8c64eb9-e8c7-4f38-be00-7c817e3ed14e"],"name":["tester"],"permission_id[]":["3"],"slug":["tester"]}', '2020-03-15 03:14:37.760', '2020-03-15 03:14:37.760');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('57', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:14:37.790', '2020-03-15 03:14:37.790');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('58', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:14:37.830', '2020-03-15 03:14:37.830');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('59', '1', '/admin/edit/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["3ffbad00-9670-4262-b637-f88026b52bfc"],"id":["3"],"name":["tester"],"permission_id[]":["3","2"],"slug":["tester"]}', '2020-03-15 03:14:37.873', '2020-03-15 03:14:37.873');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('60', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 03:14:37.913', '2020-03-15 03:14:37.913');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('61', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["3ba7d233-1bc2-4257-95a6-544eedf18981"],"name":["tester2"],"permission_id[]":["3"],"slug":["tester2"]}', '2020-03-15 03:14:37.940', '2020-03-15 03:14:37.940');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('62', '1', '/admin/delete/roles', 'POST', '127.0.0.1', '{"id":["3"]}', '2020-03-15 03:14:37.980', '2020-03-15 03:14:37.980');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('63', '1', '/admin/info/manager', 'GET', '127.0.0.1', '', '2020-03-15 03:14:38.030', '2020-03-15 03:14:38.030');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('64', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["123"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 03:14:38.063', '2020-03-15 03:14:38.063');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('65', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:14:38.097', '2020-03-15 03:14:38.097');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('66', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:14:38.150', '2020-03-15 03:14:38.150');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('67', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["8c0004a3-e902-46ad-a80a-ccb01a5b90c7"],"avatar__delete_flag":["0"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 03:14:38.283', '2020-03-15 03:14:38.283');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('68', '1', '/admin/info/manager/new', 'GET', '127.0.0.1', '', '2020-03-15 03:14:38.327', '2020-03-15 03:14:38.327');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('69', '1', '/admin/new/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["3c7ae974-1e74-48fd-a70d-6f3a502d4767"],"avatar__delete_flag":["0"],"id":["1"],"name":["tester"],"password":["tester"],"password_again":["tester"],"permission_id[]":["1"],"role_id[]":["1"],"username":["tester"]}', '2020-03-15 03:14:38.460', '2020-03-15 03:14:38.460');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('70', '1', '/admin/menu', 'GET', '127.0.0.1', '', '2020-03-15 03:14:38.730', '2020-03-15 03:14:38.730');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('71', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["c8b82f1c-45c3-4a9e-b92d-e9055133af08"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test menu"],"uri":["/example/test"]}', '2020-03-15 03:14:38.803', '2020-03-15 03:14:38.803');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('72', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 03:14:38.853', '2020-03-15 03:14:38.853');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('73', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 03:14:38.910', '2020-03-15 03:14:38.910');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('74', '1', '/admin/menu/edit', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["3d6ae122-32dc-40d5-9930-84aff50aa97f"],"header":[""],"icon":["fa-angellist"],"id":["3"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test"]}', '2020-03-15 03:14:38.990', '2020-03-15 03:14:38.990');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('75', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["3d6ae122-32dc-40d5-9930-84aff50aa97f"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test2"]}', '2020-03-15 03:14:39.043', '2020-03-15 03:14:39.043');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('76', '1', '/admin/menu/delete', 'POST', '127.0.0.1', '{}', '2020-03-15 03:14:39.083', '2020-03-15 03:14:39.083');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('77', '1', '/admin/info/op', 'GET', '127.0.0.1', '', '2020-03-15 03:14:39.133', '2020-03-15 03:14:39.133');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('78', '1', '/admin/info/external', 'GET', '127.0.0.1', '', '2020-03-15 03:14:39.163', '2020-03-15 03:14:39.163');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('79', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:14:39.193', '2020-03-15 03:14:39.193');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('80', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:14:39.233', '2020-03-15 03:14:39.233');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('81', '1', '/admin/info/external/new', 'GET', '127.0.0.1', '', '2020-03-15 03:14:39.280', '2020-03-15 03:14:39.280');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('82', '1', '/admin/info/user', 'GET', '127.0.0.1', '', '2020-03-15 03:14:39.310', '2020-03-15 03:14:39.310');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('83', '1', '/admin/export/user', 'POST', '127.0.0.1', '{"id":["1"]}', '2020-03-15 03:14:39.360', '2020-03-15 03:14:39.360');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('84', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:14:39.403', '2020-03-15 03:14:39.403');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('85', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:14:39.443', '2020-03-15 03:14:39.443');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('86', '1', '/admin/info/user/new', 'GET', '127.0.0.1', '', '2020-03-15 03:14:39.487', '2020-03-15 03:14:39.487');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('87', '1', '/admin/logout', 'GET', '127.0.0.1', '', '2020-03-15 03:39:21.340', '2020-03-15 03:39:21.340');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('88', '1', '/admin/info/permission', 'GET', '127.0.0.1', '', '2020-03-15 03:39:21.750', '2020-03-15 03:39:21.750');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('89', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 03:39:21.787', '2020-03-15 03:39:21.787');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('90', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["8bee7649-6223-4fe6-a505-ea375cd0b49c"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester"],"slug":["tester"]}', '2020-03-15 03:39:21.820', '2020-03-15 03:39:21.820');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('91', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:39:21.850', '2020-03-15 03:39:21.850');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('92', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:39:21.887', '2020-03-15 03:39:21.887');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('93', '1', '/admin/edit/permission', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["087d600a-86a5-4fe2-9ff3-fd4e6b0176c4"],"http_method[]":["GET","POST"],"http_path":["/\n/admin/info/op"],"id":["3"],"name":["tester"],"slug":["tester"]}', '2020-03-15 03:39:21.947', '2020-03-15 03:39:21.947');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('94', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 03:39:21.987', '2020-03-15 03:39:21.987');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('95', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["8c2c3b2a-108e-42c7-8de8-55dd101d44e3"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester2"],"slug":["tester2"]}', '2020-03-15 03:39:22.020', '2020-03-15 03:39:22.020');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('96', '1', '/admin/delete/permission', 'POST', '127.0.0.1', '{"id":["4"]}', '2020-03-15 03:39:22.053', '2020-03-15 03:39:22.053');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('97', '1', '/admin/info/roles', 'GET', '127.0.0.1', '', '2020-03-15 03:39:22.100', '2020-03-15 03:39:22.100');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('98', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 03:39:22.130', '2020-03-15 03:39:22.130');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('99', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["2485611e-bd7e-4ae2-acc9-4ce72a601c44"],"name":["tester"],"permission_id[]":["3"],"slug":["tester"]}', '2020-03-15 03:39:22.167', '2020-03-15 03:39:22.167');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('100', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:39:22.207', '2020-03-15 03:39:22.207');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('101', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:39:22.473', '2020-03-15 03:39:22.473');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('102', '1', '/admin/edit/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["42723e0f-6198-4a96-b13b-83e8406b9ba7"],"id":["3"],"name":["tester"],"permission_id[]":["3","2"],"slug":["tester"]}', '2020-03-15 03:39:22.537', '2020-03-15 03:39:22.537');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('103', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 03:39:22.577', '2020-03-15 03:39:22.577');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('104', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["1687c542-94e9-45b3-b362-16aeb6d92169"],"name":["tester2"],"permission_id[]":["3"],"slug":["tester2"]}', '2020-03-15 03:39:22.617', '2020-03-15 03:39:22.617');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('105', '1', '/admin/delete/roles', 'POST', '127.0.0.1', '{"id":["3"]}', '2020-03-15 03:39:22.647', '2020-03-15 03:39:22.647');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('106', '1', '/admin/info/manager', 'GET', '127.0.0.1', '', '2020-03-15 03:39:22.690', '2020-03-15 03:39:22.690');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('107', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["123"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 03:39:22.713', '2020-03-15 03:39:22.713');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('108', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:39:22.740', '2020-03-15 03:39:22.740');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('109', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:39:22.780', '2020-03-15 03:39:22.780');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('110', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["f37d07a2-cff6-42dc-8004-8bd174da141d"],"avatar__delete_flag":["0"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 03:39:22.927', '2020-03-15 03:39:22.927');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('111', '1', '/admin/info/manager/new', 'GET', '127.0.0.1', '', '2020-03-15 03:39:22.970', '2020-03-15 03:39:22.970');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('112', '1', '/admin/new/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["04a35b51-52e5-4dfb-8a9f-fa2211709d3e"],"avatar__delete_flag":["0"],"id":["1"],"name":["tester"],"password":["tester"],"password_again":["tester"],"permission_id[]":["1"],"role_id[]":["1"],"username":["tester"]}', '2020-03-15 03:39:23.093', '2020-03-15 03:39:23.093');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('113', '1', '/admin/menu', 'GET', '127.0.0.1', '', '2020-03-15 03:39:23.387', '2020-03-15 03:39:23.387');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('114', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["ab730b2e-7c8b-4c55-b0bb-f5ba0cae4e5b"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test menu"],"uri":["/example/test"]}', '2020-03-15 03:39:23.457', '2020-03-15 03:39:23.457');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('115', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 03:39:23.517', '2020-03-15 03:39:23.517');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('116', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 03:39:23.573', '2020-03-15 03:39:23.573');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('117', '1', '/admin/menu/edit', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["402e0b89-7179-472f-af66-424e02f090b5"],"header":[""],"icon":["fa-angellist"],"id":["3"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test"]}', '2020-03-15 03:39:23.663', '2020-03-15 03:39:23.663');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('118', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["402e0b89-7179-472f-af66-424e02f090b5"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test2"]}', '2020-03-15 03:39:23.723', '2020-03-15 03:39:23.723');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('119', '1', '/admin/menu/delete', 'POST', '127.0.0.1', '{}', '2020-03-15 03:39:23.757', '2020-03-15 03:39:23.757');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('120', '1', '/admin/info/op', 'GET', '127.0.0.1', '', '2020-03-15 03:39:23.813', '2020-03-15 03:39:23.813');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('121', '1', '/admin/info/external', 'GET', '127.0.0.1', '', '2020-03-15 03:39:23.847', '2020-03-15 03:39:23.847');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('122', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:39:23.877', '2020-03-15 03:39:23.877');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('123', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:39:23.910', '2020-03-15 03:39:23.910');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('124', '1', '/admin/info/external/new', 'GET', '127.0.0.1', '', '2020-03-15 03:39:23.947', '2020-03-15 03:39:23.947');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('125', '1', '/admin/info/user', 'GET', '127.0.0.1', '', '2020-03-15 03:39:23.970', '2020-03-15 03:39:23.970');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('126', '1', '/admin/export/user', 'POST', '127.0.0.1', '{"id":["1"]}', '2020-03-15 03:39:24.010', '2020-03-15 03:39:24.010');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('127', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:39:24.050', '2020-03-15 03:39:24.050');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('128', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:39:24.093', '2020-03-15 03:39:24.093');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('129', '1', '/admin/info/user/new', 'GET', '127.0.0.1', '', '2020-03-15 03:39:24.137', '2020-03-15 03:39:24.137');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('130', '1', '/admin/logout', 'GET', '127.0.0.1', '', '2020-03-15 03:49:24.813', '2020-03-15 03:49:24.813');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('131', '1', '/admin/info/permission', 'GET', '127.0.0.1', '', '2020-03-15 03:49:25.190', '2020-03-15 03:49:25.190');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('132', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 03:49:25.233', '2020-03-15 03:49:25.233');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('133', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["c3a60aba-926d-424a-8cbd-a999668753fe"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester"],"slug":["tester"]}', '2020-03-15 03:49:25.277', '2020-03-15 03:49:25.277');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('134', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:49:25.310', '2020-03-15 03:49:25.310');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('135', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:49:25.357', '2020-03-15 03:49:25.357');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('136', '1', '/admin/edit/permission', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["68d8bb95-b37a-4f35-ba19-efea7b75a2f2"],"http_method[]":["GET","POST"],"http_path":["/\n/admin/info/op"],"id":["3"],"name":["tester"],"slug":["tester"]}', '2020-03-15 03:49:25.410', '2020-03-15 03:49:25.410');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('137', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 03:49:25.443', '2020-03-15 03:49:25.443');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('138', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["ff95fcec-ad5a-4604-ba83-b51cb9b18a89"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester2"],"slug":["tester2"]}', '2020-03-15 03:49:25.480', '2020-03-15 03:49:25.480');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('139', '1', '/admin/delete/permission', 'POST', '127.0.0.1', '{"id":["4"]}', '2020-03-15 03:49:25.507', '2020-03-15 03:49:25.507');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('140', '1', '/admin/info/roles', 'GET', '127.0.0.1', '', '2020-03-15 03:49:25.527', '2020-03-15 03:49:25.527');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('141', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 03:49:25.573', '2020-03-15 03:49:25.573');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('142', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["fc0ec315-3254-4145-bd90-fa54dc115d6f"],"name":["tester"],"permission_id[]":["3"],"slug":["tester"]}', '2020-03-15 03:49:25.613', '2020-03-15 03:49:25.613');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('143', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:49:25.643', '2020-03-15 03:49:25.643');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('144', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:49:25.680', '2020-03-15 03:49:25.680');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('145', '1', '/admin/edit/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["d539d703-d4eb-4551-83b3-4b41049e884c"],"id":["3"],"name":["tester"],"permission_id[]":["3","2"],"slug":["tester"]}', '2020-03-15 03:49:25.723', '2020-03-15 03:49:25.723');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('146', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 03:49:25.763', '2020-03-15 03:49:25.763');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('147', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["a62ae5e8-66c8-44c2-9bc4-be92c75f255d"],"name":["tester2"],"permission_id[]":["3"],"slug":["tester2"]}', '2020-03-15 03:49:25.800', '2020-03-15 03:49:25.800');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('148', '1', '/admin/delete/roles', 'POST', '127.0.0.1', '{"id":["3"]}', '2020-03-15 03:49:25.830', '2020-03-15 03:49:25.830');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('149', '1', '/admin/info/manager', 'GET', '127.0.0.1', '', '2020-03-15 03:49:25.860', '2020-03-15 03:49:25.860');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('150', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["123"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 03:49:25.903', '2020-03-15 03:49:25.903');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('151', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:49:25.937', '2020-03-15 03:49:25.937');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('152', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:49:25.983', '2020-03-15 03:49:25.983');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('153', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["a8b6b05c-f420-4c33-852f-d673599b47a4"],"avatar__delete_flag":["0"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 03:49:26.143', '2020-03-15 03:49:26.143');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('154', '1', '/admin/info/manager/new', 'GET', '127.0.0.1', '', '2020-03-15 03:49:26.190', '2020-03-15 03:49:26.190');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('155', '1', '/admin/new/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["b8a8e809-9672-462e-937a-3fd9a516a35b"],"avatar__delete_flag":["0"],"id":["1"],"name":["tester"],"password":["tester"],"password_again":["tester"],"permission_id[]":["1"],"role_id[]":["1"],"username":["tester"]}', '2020-03-15 03:49:26.293', '2020-03-15 03:49:26.293');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('156', '1', '/admin/menu', 'GET', '127.0.0.1', '', '2020-03-15 03:49:26.600', '2020-03-15 03:49:26.600');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('157', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["8441b1f4-d6de-4c34-a65c-6bb9969e2974"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test menu"],"uri":["/example/test"]}', '2020-03-15 03:49:26.673', '2020-03-15 03:49:26.673');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('158', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 03:49:26.727', '2020-03-15 03:49:26.727');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('159', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 03:49:26.773', '2020-03-15 03:49:26.773');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('160', '1', '/admin/menu/edit', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["a981fd88-2911-4b32-ab74-990842814dac"],"header":[""],"icon":["fa-angellist"],"id":["3"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test"]}', '2020-03-15 03:49:26.847', '2020-03-15 03:49:26.847');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('161', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["a981fd88-2911-4b32-ab74-990842814dac"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test2"]}', '2020-03-15 03:49:26.893', '2020-03-15 03:49:26.893');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('162', '1', '/admin/menu/delete', 'POST', '127.0.0.1', '{}', '2020-03-15 03:49:26.927', '2020-03-15 03:49:26.927');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('163', '1', '/admin/info/op', 'GET', '127.0.0.1', '', '2020-03-15 03:49:26.957', '2020-03-15 03:49:26.957');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('164', '1', '/admin/info/external', 'GET', '127.0.0.1', '', '2020-03-15 03:49:26.997', '2020-03-15 03:49:26.997');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('165', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:49:27.027', '2020-03-15 03:49:27.027');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('166', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:49:27.063', '2020-03-15 03:49:27.063');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('167', '1', '/admin/info/external/new', 'GET', '127.0.0.1', '', '2020-03-15 03:49:27.103', '2020-03-15 03:49:27.103');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('168', '1', '/admin/info/user', 'GET', '127.0.0.1', '', '2020-03-15 03:49:27.133', '2020-03-15 03:49:27.133');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('169', '1', '/admin/export/user', 'POST', '127.0.0.1', '{"id":["1"]}', '2020-03-15 03:49:27.177', '2020-03-15 03:49:27.177');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('170', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:49:27.220', '2020-03-15 03:49:27.220');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('171', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:49:27.260', '2020-03-15 03:49:27.260');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('172', '1', '/admin/info/user/new', 'GET', '127.0.0.1', '', '2020-03-15 03:49:27.293', '2020-03-15 03:49:27.293');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('173', '1', '/admin/logout', 'GET', '127.0.0.1', '', '2020-03-15 03:55:28.027', '2020-03-15 03:55:28.027');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('174', '1', '/admin/info/permission', 'GET', '127.0.0.1', '', '2020-03-15 03:55:28.430', '2020-03-15 03:55:28.430');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('175', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 03:55:28.477', '2020-03-15 03:55:28.477');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('176', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["418ee4ee-3a05-49a2-b666-830c719fd776"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester"],"slug":["tester"]}', '2020-03-15 03:55:28.513', '2020-03-15 03:55:28.513');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('177', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:55:28.550', '2020-03-15 03:55:28.550');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('178', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:55:28.580', '2020-03-15 03:55:28.580');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('179', '1', '/admin/edit/permission', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["515bd79c-e610-4f1e-a600-b46ff6d902af"],"http_method[]":["GET","POST"],"http_path":["/\n/admin/info/op"],"id":["3"],"name":["tester"],"slug":["tester"]}', '2020-03-15 03:55:28.637', '2020-03-15 03:55:28.637');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('180', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 03:55:28.677', '2020-03-15 03:55:28.677');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('181', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["9738e849-c22c-4c38-a1b8-1d7d1741efe2"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester2"],"slug":["tester2"]}', '2020-03-15 03:55:28.713', '2020-03-15 03:55:28.713');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('182', '1', '/admin/delete/permission', 'POST', '127.0.0.1', '{"id":["4"]}', '2020-03-15 03:55:28.743', '2020-03-15 03:55:28.743');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('183', '1', '/admin/info/roles', 'GET', '127.0.0.1', '', '2020-03-15 03:55:28.763', '2020-03-15 03:55:28.763');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('184', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 03:55:28.813', '2020-03-15 03:55:28.813');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('185', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["e6a37010-b56a-4f8f-9d88-396059440744"],"name":["tester"],"permission_id[]":["3"],"slug":["tester"]}', '2020-03-15 03:55:28.857', '2020-03-15 03:55:28.857');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('186', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:55:28.890', '2020-03-15 03:55:28.890');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('187', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:55:28.947', '2020-03-15 03:55:28.947');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('188', '1', '/admin/edit/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["6560686e-ac53-499b-8a8c-9813787446f5"],"id":["3"],"name":["tester"],"permission_id[]":["3","2"],"slug":["tester"]}', '2020-03-15 03:55:28.993', '2020-03-15 03:55:28.993');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('189', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 03:55:29.037', '2020-03-15 03:55:29.037');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('190', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["1844756f-0a44-4fbf-8019-cd3f97a6db70"],"name":["tester2"],"permission_id[]":["3"],"slug":["tester2"]}', '2020-03-15 03:55:29.080', '2020-03-15 03:55:29.080');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('191', '1', '/admin/delete/roles', 'POST', '127.0.0.1', '{"id":["3"]}', '2020-03-15 03:55:29.117', '2020-03-15 03:55:29.117');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('192', '1', '/admin/info/manager', 'GET', '127.0.0.1', '', '2020-03-15 03:55:29.147', '2020-03-15 03:55:29.147');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('193', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["123"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 03:55:29.190', '2020-03-15 03:55:29.190');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('194', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:55:29.227', '2020-03-15 03:55:29.227');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('195', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:55:29.290', '2020-03-15 03:55:29.290');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('196', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["4e465c13-ee6f-4b95-9187-f46a231caae9"],"avatar__delete_flag":["0"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 03:55:29.463', '2020-03-15 03:55:29.463');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('197', '1', '/admin/info/manager/new', 'GET', '127.0.0.1', '', '2020-03-15 03:55:29.517', '2020-03-15 03:55:29.517');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('198', '1', '/admin/new/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["b0b0d014-f418-4281-82fd-da8e5be8f43d"],"avatar__delete_flag":["0"],"id":["1"],"name":["tester"],"password":["tester"],"password_again":["tester"],"permission_id[]":["1"],"role_id[]":["1"],"username":["tester"]}', '2020-03-15 03:55:29.623', '2020-03-15 03:55:29.623');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('199', '1', '/admin/menu', 'GET', '127.0.0.1', '', '2020-03-15 03:55:29.943', '2020-03-15 03:55:29.943');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('200', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["63980a38-e643-46b1-aefa-e49da95de6e0"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test menu"],"uri":["/example/test"]}', '2020-03-15 03:55:30.057', '2020-03-15 03:55:30.057');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('201', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 03:55:31.580', '2020-03-15 03:55:31.580');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('202', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 03:55:31.623', '2020-03-15 03:55:31.623');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('203', '1', '/admin/menu/edit', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["e95cab98-2517-44ed-b59c-4b65b8d2837d"],"header":[""],"icon":["fa-angellist"],"id":["3"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test"]}', '2020-03-15 03:55:31.730', '2020-03-15 03:55:31.730');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('204', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["e95cab98-2517-44ed-b59c-4b65b8d2837d"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test2"]}', '2020-03-15 03:55:31.770', '2020-03-15 03:55:31.770');
GO
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('205', '1', '/admin/menu/delete', 'POST', '127.0.0.1', '{}', '2020-03-15 03:55:31.807', '2020-03-15 03:55:31.807');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('206', '1', '/admin/info/op', 'GET', '127.0.0.1', '', '2020-03-15 03:55:31.830', '2020-03-15 03:55:31.830');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('207', '1', '/admin/info/external', 'GET', '127.0.0.1', '', '2020-03-15 03:55:31.880', '2020-03-15 03:55:31.880');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('208', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:55:31.913', '2020-03-15 03:55:31.913');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('209', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:55:31.947', '2020-03-15 03:55:31.947');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('210', '1', '/admin/info/external/new', 'GET', '127.0.0.1', '', '2020-03-15 03:55:31.980', '2020-03-15 03:55:31.980');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('211', '1', '/admin/info/user', 'GET', '127.0.0.1', '', '2020-03-15 03:55:32.010', '2020-03-15 03:55:32.010');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('212', '1', '/admin/export/user', 'POST', '127.0.0.1', '{"id":["1"]}', '2020-03-15 03:55:32.053', '2020-03-15 03:55:32.053');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('213', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:55:32.097', '2020-03-15 03:55:32.097');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('214', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:55:32.143', '2020-03-15 03:55:32.143');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('215', '1', '/admin/info/user/new', 'GET', '127.0.0.1', '', '2020-03-15 03:55:32.180', '2020-03-15 03:55:32.180');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('216', '1', '/admin/logout', 'GET', '127.0.0.1', '', '2020-03-15 03:57:01.097', '2020-03-15 03:57:01.097');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('217', '1', '/admin/info/permission', 'GET', '127.0.0.1', '', '2020-03-15 03:57:01.477', '2020-03-15 03:57:01.477');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('218', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 03:57:01.530', '2020-03-15 03:57:01.530');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('219', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["51757bf6-b3c4-46f4-bf9f-6e91f347dc2d"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester"],"slug":["tester"]}', '2020-03-15 03:57:01.563', '2020-03-15 03:57:01.563');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('220', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:57:01.593', '2020-03-15 03:57:01.593');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('221', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:57:01.640', '2020-03-15 03:57:01.640');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('222', '1', '/admin/edit/permission', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["9762d55b-c444-405e-b0b6-5feb851c0e40"],"http_method[]":["GET","POST"],"http_path":["/\n/admin/info/op"],"id":["3"],"name":["tester"],"slug":["tester"]}', '2020-03-15 03:57:01.690', '2020-03-15 03:57:01.690');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('223', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 03:57:01.730', '2020-03-15 03:57:01.730');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('224', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["64218a2f-eb41-4ca3-99d9-c925d9d11700"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester2"],"slug":["tester2"]}', '2020-03-15 03:57:01.763', '2020-03-15 03:57:01.763');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('225', '1', '/admin/delete/permission', 'POST', '127.0.0.1', '{"id":["4"]}', '2020-03-15 03:57:01.787', '2020-03-15 03:57:01.787');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('226', '1', '/admin/info/roles', 'GET', '127.0.0.1', '', '2020-03-15 03:57:01.810', '2020-03-15 03:57:01.810');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('227', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 03:57:01.847', '2020-03-15 03:57:01.847');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('228', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["cfa68542-545d-4cd4-bc1c-fef5a4005b0d"],"name":["tester"],"permission_id[]":["3"],"slug":["tester"]}', '2020-03-15 03:57:01.880', '2020-03-15 03:57:01.880');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('229', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:57:01.907', '2020-03-15 03:57:01.907');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('230', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:57:01.950', '2020-03-15 03:57:01.950');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('231', '1', '/admin/edit/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["7ef4b394-e721-4012-a24a-9296233d0053"],"id":["3"],"name":["tester"],"permission_id[]":["3","2"],"slug":["tester"]}', '2020-03-15 03:57:01.987', '2020-03-15 03:57:01.987');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('232', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 03:57:02.030', '2020-03-15 03:57:02.030');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('233', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["6a6318cd-89be-4a37-8e59-fb63a30c0d03"],"name":["tester2"],"permission_id[]":["3"],"slug":["tester2"]}', '2020-03-15 03:57:02.067', '2020-03-15 03:57:02.067');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('234', '1', '/admin/delete/roles', 'POST', '127.0.0.1', '{"id":["3"]}', '2020-03-15 03:57:02.107', '2020-03-15 03:57:02.107');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('235', '1', '/admin/info/manager', 'GET', '127.0.0.1', '', '2020-03-15 03:57:02.130', '2020-03-15 03:57:02.130');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('236', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["123"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 03:57:02.173', '2020-03-15 03:57:02.173');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('237', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:57:02.207', '2020-03-15 03:57:02.207');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('238', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:57:02.253', '2020-03-15 03:57:02.253');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('239', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["aa9b074b-c297-4aff-9a66-179b5020078c"],"avatar__delete_flag":["0"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 03:57:02.407', '2020-03-15 03:57:02.407');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('240', '1', '/admin/info/manager/new', 'GET', '127.0.0.1', '', '2020-03-15 03:57:02.453', '2020-03-15 03:57:02.453');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('241', '1', '/admin/new/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["062dc246-ad4b-44b8-899f-dfc3815eefd3"],"avatar__delete_flag":["0"],"id":["1"],"name":["tester"],"password":["tester"],"password_again":["tester"],"permission_id[]":["1"],"role_id[]":["1"],"username":["tester"]}', '2020-03-15 03:57:02.553', '2020-03-15 03:57:02.553');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('242', '1', '/admin/menu', 'GET', '127.0.0.1', '', '2020-03-15 03:57:02.850', '2020-03-15 03:57:02.850');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('243', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["168e4192-270a-4ac4-947d-1a7d73785474"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test menu"],"uri":["/example/test"]}', '2020-03-15 03:57:02.923', '2020-03-15 03:57:02.923');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('244', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 03:57:02.973', '2020-03-15 03:57:02.973');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('245', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 03:57:03.027', '2020-03-15 03:57:03.027');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('246', '1', '/admin/menu/edit', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["6697607f-8220-41c5-a5da-ba1c6a432195"],"header":[""],"icon":["fa-angellist"],"id":["3"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test"]}', '2020-03-15 03:57:03.093', '2020-03-15 03:57:03.093');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('247', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["6697607f-8220-41c5-a5da-ba1c6a432195"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test2"]}', '2020-03-15 03:57:03.143', '2020-03-15 03:57:03.143');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('248', '1', '/admin/menu/delete', 'POST', '127.0.0.1', '{}', '2020-03-15 03:57:03.170', '2020-03-15 03:57:03.170');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('249', '1', '/admin/info/op', 'GET', '127.0.0.1', '', '2020-03-15 03:57:03.200', '2020-03-15 03:57:03.200');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('250', '1', '/admin/info/external', 'GET', '127.0.0.1', '', '2020-03-15 03:57:03.237', '2020-03-15 03:57:03.237');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('251', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:57:03.267', '2020-03-15 03:57:03.267');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('252', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:57:03.300', '2020-03-15 03:57:03.300');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('253', '1', '/admin/info/external/new', 'GET', '127.0.0.1', '', '2020-03-15 03:57:03.330', '2020-03-15 03:57:03.330');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('254', '1', '/admin/info/user', 'GET', '127.0.0.1', '', '2020-03-15 03:57:03.350', '2020-03-15 03:57:03.350');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('255', '1', '/admin/export/user', 'POST', '127.0.0.1', '{"id":["1"]}', '2020-03-15 03:57:03.380', '2020-03-15 03:57:03.380');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('256', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:57:03.413', '2020-03-15 03:57:03.413');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('257', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 03:57:03.447', '2020-03-15 03:57:03.447');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('258', '1', '/admin/info/user/new', 'GET', '127.0.0.1', '', '2020-03-15 03:57:03.480', '2020-03-15 03:57:03.480');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('259', '1', '/admin/logout', 'GET', '127.0.0.1', '', '2020-03-15 04:08:41.627', '2020-03-15 04:08:41.627');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('260', '1', '/admin/info/permission', 'GET', '127.0.0.1', '', '2020-03-15 04:08:41.993', '2020-03-15 04:08:41.993');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('261', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 04:08:42.047', '2020-03-15 04:08:42.047');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('262', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["95001d14-ab8f-4d1e-a16e-78e85ce1830e"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester"],"slug":["tester"]}', '2020-03-15 04:08:42.083', '2020-03-15 04:08:42.083');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('263', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:08:42.120', '2020-03-15 04:08:42.120');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('264', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:08:42.160', '2020-03-15 04:08:42.160');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('265', '1', '/admin/edit/permission', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["c247b48f-fd2b-412f-adb0-c7ca03a96331"],"http_method[]":["GET","POST"],"http_path":["/\n/admin/info/op"],"id":["3"],"name":["tester"],"slug":["tester"]}', '2020-03-15 04:08:42.210', '2020-03-15 04:08:42.210');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('266', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 04:08:42.250', '2020-03-15 04:08:42.250');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('267', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["25f2de82-68e1-4810-830f-a8cf0eb8f5a0"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester2"],"slug":["tester2"]}', '2020-03-15 04:08:42.290', '2020-03-15 04:08:42.290');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('268', '1', '/admin/delete/permission', 'POST', '127.0.0.1', '{"id":["4"]}', '2020-03-15 04:08:42.320', '2020-03-15 04:08:42.320');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('269', '1', '/admin/info/roles', 'GET', '127.0.0.1', '', '2020-03-15 04:08:42.340', '2020-03-15 04:08:42.340');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('270', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 04:08:42.393', '2020-03-15 04:08:42.393');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('271', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["064639b2-d4eb-4c7e-9082-35580a9b65e2"],"name":["tester"],"permission_id[]":["3"],"slug":["tester"]}', '2020-03-15 04:08:42.433', '2020-03-15 04:08:42.433');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('272', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:08:42.463', '2020-03-15 04:08:42.463');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('273', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:08:42.503', '2020-03-15 04:08:42.503');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('274', '1', '/admin/edit/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["27b2084a-c9fd-4b76-b178-f298d1f42be6"],"id":["3"],"name":["tester"],"permission_id[]":["3","2"],"slug":["tester"]}', '2020-03-15 04:08:42.547', '2020-03-15 04:08:42.547');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('275', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 04:08:42.593', '2020-03-15 04:08:42.593');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('276', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["6f88027a-3bba-4e1d-8b03-bf5ff269e0c4"],"name":["tester2"],"permission_id[]":["3"],"slug":["tester2"]}', '2020-03-15 04:08:42.630', '2020-03-15 04:08:42.630');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('277', '1', '/admin/delete/roles', 'POST', '127.0.0.1', '{"id":["3"]}', '2020-03-15 04:08:42.660', '2020-03-15 04:08:42.660');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('278', '1', '/admin/info/manager', 'GET', '127.0.0.1', '', '2020-03-15 04:08:42.690', '2020-03-15 04:08:42.690');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('279', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["123"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 04:08:42.730', '2020-03-15 04:08:42.730');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('280', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:08:42.763', '2020-03-15 04:08:42.763');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('281', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:08:42.817', '2020-03-15 04:08:42.817');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('282', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["79b123f3-ea8b-465a-af6d-493bc5e91a4e"],"avatar__delete_flag":["0"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 04:08:42.967', '2020-03-15 04:08:42.967');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('283', '1', '/admin/info/manager/new', 'GET', '127.0.0.1', '', '2020-03-15 04:08:43.010', '2020-03-15 04:08:43.010');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('284', '1', '/admin/new/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["ec317734-967b-401f-ab45-48fc59bc9276"],"avatar__delete_flag":["0"],"id":["1"],"name":["tester"],"password":["tester"],"password_again":["tester"],"permission_id[]":["1"],"role_id[]":["1"],"username":["tester"]}', '2020-03-15 04:08:43.117', '2020-03-15 04:08:43.117');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('285', '1', '/admin/menu', 'GET', '127.0.0.1', '', '2020-03-15 04:08:43.403', '2020-03-15 04:08:43.403');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('286', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["e0284537-2c85-4562-a383-43d3d61dcfb6"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test menu"],"uri":["/example/test"]}', '2020-03-15 04:08:43.477', '2020-03-15 04:08:43.477');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('287', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 04:08:43.520', '2020-03-15 04:08:43.520');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('288', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 04:08:43.563', '2020-03-15 04:08:43.563');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('289', '1', '/admin/menu/edit', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["fdb68af2-d8c9-42b3-9c98-d2134fd67afd"],"header":[""],"icon":["fa-angellist"],"id":["3"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test"]}', '2020-03-15 04:08:43.627', '2020-03-15 04:08:43.627');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('290', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["fdb68af2-d8c9-42b3-9c98-d2134fd67afd"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test2"]}', '2020-03-15 04:08:43.673', '2020-03-15 04:08:43.673');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('291', '1', '/admin/menu/delete', 'POST', '127.0.0.1', '{}', '2020-03-15 04:08:43.703', '2020-03-15 04:08:43.703');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('292', '1', '/admin/info/op', 'GET', '127.0.0.1', '', '2020-03-15 04:08:43.723', '2020-03-15 04:08:43.723');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('293', '1', '/admin/info/external', 'GET', '127.0.0.1', '', '2020-03-15 04:08:43.767', '2020-03-15 04:08:43.767');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('294', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:08:43.790', '2020-03-15 04:08:43.790');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('295', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:08:43.820', '2020-03-15 04:08:43.820');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('296', '1', '/admin/info/external/new', 'GET', '127.0.0.1', '', '2020-03-15 04:08:43.853', '2020-03-15 04:08:43.853');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('297', '1', '/admin/info/user', 'GET', '127.0.0.1', '', '2020-03-15 04:08:43.877', '2020-03-15 04:08:43.877');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('298', '1', '/admin/export/user', 'POST', '127.0.0.1', '{"id":["1"]}', '2020-03-15 04:08:43.913', '2020-03-15 04:08:43.913');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('299', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:08:43.957', '2020-03-15 04:08:43.957');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('300', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:08:43.993', '2020-03-15 04:08:43.993');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('301', '1', '/admin/info/user/new', 'GET', '127.0.0.1', '', '2020-03-15 04:08:44.033', '2020-03-15 04:08:44.033');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('302', '1', '/admin/logout', 'GET', '127.0.0.1', '', '2020-03-15 04:09:20.960', '2020-03-15 04:09:20.960');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('303', '1', '/admin/info/permission', 'GET', '127.0.0.1', '', '2020-03-15 04:09:21.343', '2020-03-15 04:09:21.343');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('304', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 04:09:21.390', '2020-03-15 04:09:21.390');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('305', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["8343b2e9-05a6-4989-869b-76a5f4bf4c89"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester"],"slug":["tester"]}', '2020-03-15 04:09:21.430', '2020-03-15 04:09:21.430');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('306', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:09:21.460', '2020-03-15 04:09:21.460');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('307', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:09:21.500', '2020-03-15 04:09:21.500');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('308', '1', '/admin/edit/permission', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["1bebaf0e-a0f3-4f27-a18f-fe12ff7ec333"],"http_method[]":["GET","POST"],"http_path":["/\n/admin/info/op"],"id":["3"],"name":["tester"],"slug":["tester"]}', '2020-03-15 04:09:21.547', '2020-03-15 04:09:21.547');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('309', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 04:09:21.587', '2020-03-15 04:09:21.587');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('310', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["903a8026-fb39-468b-9748-c97e8a691f54"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester2"],"slug":["tester2"]}', '2020-03-15 04:09:21.627', '2020-03-15 04:09:21.627');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('311', '1', '/admin/delete/permission', 'POST', '127.0.0.1', '{"id":["4"]}', '2020-03-15 04:09:21.660', '2020-03-15 04:09:21.660');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('312', '1', '/admin/info/roles', 'GET', '127.0.0.1', '', '2020-03-15 04:09:21.683', '2020-03-15 04:09:21.683');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('313', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 04:09:21.737', '2020-03-15 04:09:21.737');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('314', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["52a1496c-d043-4e91-904f-7da25a0fd35d"],"name":["tester"],"permission_id[]":["3"],"slug":["tester"]}', '2020-03-15 04:09:21.780', '2020-03-15 04:09:21.780');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('315', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:09:21.827', '2020-03-15 04:09:21.827');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('316', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:09:21.877', '2020-03-15 04:09:21.877');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('317', '1', '/admin/edit/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["87f737d4-7e7d-4c4c-9ac3-e95c4e1a27b7"],"id":["3"],"name":["tester"],"permission_id[]":["3","2"],"slug":["tester"]}', '2020-03-15 04:09:21.923', '2020-03-15 04:09:21.923');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('318', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 04:09:21.973', '2020-03-15 04:09:21.973');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('319', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["df905cd3-b59a-4a50-8db9-802b599e7512"],"name":["tester2"],"permission_id[]":["3"],"slug":["tester2"]}', '2020-03-15 04:09:22.013', '2020-03-15 04:09:22.013');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('320', '1', '/admin/delete/roles', 'POST', '127.0.0.1', '{"id":["3"]}', '2020-03-15 04:09:22.043', '2020-03-15 04:09:22.043');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('321', '1', '/admin/info/manager', 'GET', '127.0.0.1', '', '2020-03-15 04:09:22.073', '2020-03-15 04:09:22.073');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('322', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["123"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 04:09:22.127', '2020-03-15 04:09:22.127');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('323', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:09:22.160', '2020-03-15 04:09:22.160');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('324', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:09:22.217', '2020-03-15 04:09:22.217');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('325', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["759907db-027c-45df-a564-9e3206f60722"],"avatar__delete_flag":["0"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 04:09:22.483', '2020-03-15 04:09:22.483');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('326', '1', '/admin/info/manager/new', 'GET', '127.0.0.1', '', '2020-03-15 04:09:22.523', '2020-03-15 04:09:22.523');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('327', '1', '/admin/new/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["44c84e40-d665-4fd4-9f50-c0984580b709"],"avatar__delete_flag":["0"],"id":["1"],"name":["tester"],"password":["tester"],"password_again":["tester"],"permission_id[]":["1"],"role_id[]":["1"],"username":["tester"]}', '2020-03-15 04:09:22.633', '2020-03-15 04:09:22.633');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('328', '1', '/admin/menu', 'GET', '127.0.0.1', '', '2020-03-15 04:09:22.937', '2020-03-15 04:09:22.937');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('329', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["2048c150-95ab-47ac-a253-ec05cab832bf"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test menu"],"uri":["/example/test"]}', '2020-03-15 04:09:23.007', '2020-03-15 04:09:23.007');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('330', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 04:09:23.043', '2020-03-15 04:09:23.043');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('331', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 04:09:23.097', '2020-03-15 04:09:23.097');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('332', '1', '/admin/menu/edit', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["37d32f19-3984-415d-8286-79b6e5e1c957"],"header":[""],"icon":["fa-angellist"],"id":["3"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test"]}', '2020-03-15 04:09:23.177', '2020-03-15 04:09:23.177');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('333', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["37d32f19-3984-415d-8286-79b6e5e1c957"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test2"]}', '2020-03-15 04:09:23.223', '2020-03-15 04:09:23.223');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('334', '1', '/admin/menu/delete', 'POST', '127.0.0.1', '{}', '2020-03-15 04:09:23.253', '2020-03-15 04:09:23.253');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('335', '1', '/admin/info/op', 'GET', '127.0.0.1', '', '2020-03-15 04:09:23.283', '2020-03-15 04:09:23.283');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('336', '1', '/admin/info/external', 'GET', '127.0.0.1', '', '2020-03-15 04:09:23.327', '2020-03-15 04:09:23.327');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('337', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:09:23.360', '2020-03-15 04:09:23.360');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('338', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:09:23.393', '2020-03-15 04:09:23.393');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('339', '1', '/admin/info/external/new', 'GET', '127.0.0.1', '', '2020-03-15 04:09:23.430', '2020-03-15 04:09:23.430');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('340', '1', '/admin/info/user', 'GET', '127.0.0.1', '', '2020-03-15 04:09:23.457', '2020-03-15 04:09:23.457');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('341', '1', '/admin/export/user', 'POST', '127.0.0.1', '{"id":["1"]}', '2020-03-15 04:09:23.513', '2020-03-15 04:09:23.513');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('342', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:09:23.553', '2020-03-15 04:09:23.553');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('343', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:09:23.607', '2020-03-15 04:09:23.607');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('344', '1', '/admin/info/user/new', 'GET', '127.0.0.1', '', '2020-03-15 04:09:23.660', '2020-03-15 04:09:23.660');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('345', '1', '/admin/logout', 'GET', '127.0.0.1', '', '2020-03-15 04:14:39.593', '2020-03-15 04:14:39.593');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('346', '1', '/admin/info/permission', 'GET', '127.0.0.1', '', '2020-03-15 04:14:40.010', '2020-03-15 04:14:40.010');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('347', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 04:14:40.057', '2020-03-15 04:14:40.057');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('348', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["a381041b-f969-4c25-adea-fad2bc5d674e"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester"],"slug":["tester"]}', '2020-03-15 04:14:40.103', '2020-03-15 04:14:40.103');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('349', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:14:40.137', '2020-03-15 04:14:40.137');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('350', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:14:40.180', '2020-03-15 04:14:40.180');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('351', '1', '/admin/edit/permission', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["04fd9af1-0977-4547-805c-6d27a4fd07f1"],"http_method[]":["GET","POST"],"http_path":["/\n/admin/info/op"],"id":["3"],"name":["tester"],"slug":["tester"]}', '2020-03-15 04:14:40.230', '2020-03-15 04:14:40.230');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('352', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 04:14:40.273', '2020-03-15 04:14:40.273');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('353', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["0178e9bd-1e8d-44d9-9fa4-c134cf84d0f1"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester2"],"slug":["tester2"]}', '2020-03-15 04:14:40.310', '2020-03-15 04:14:40.310');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('354', '1', '/admin/delete/permission', 'POST', '127.0.0.1', '{"id":["4"]}', '2020-03-15 04:14:40.340', '2020-03-15 04:14:40.340');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('355', '1', '/admin/info/roles', 'GET', '127.0.0.1', '', '2020-03-15 04:14:40.383', '2020-03-15 04:14:40.383');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('356', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 04:14:40.427', '2020-03-15 04:14:40.427');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('357', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["9c213d9d-8965-4cc0-840e-ac6e0491f10c"],"name":["tester"],"permission_id[]":["3"],"slug":["tester"]}', '2020-03-15 04:14:40.470', '2020-03-15 04:14:40.470');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('358', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:14:40.500', '2020-03-15 04:14:40.500');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('359', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:14:40.537', '2020-03-15 04:14:40.537');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('360', '1', '/admin/edit/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["596cb505-caf2-4aba-8a0d-e8bfd624c59f"],"id":["3"],"name":["tester"],"permission_id[]":["3","2"],"slug":["tester"]}', '2020-03-15 04:14:40.583', '2020-03-15 04:14:40.583');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('361', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 04:14:40.627', '2020-03-15 04:14:40.627');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('362', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["f9e11d14-9521-492d-a835-0c4f3179bb85"],"name":["tester2"],"permission_id[]":["3"],"slug":["tester2"]}', '2020-03-15 04:14:40.670', '2020-03-15 04:14:40.670');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('363', '1', '/admin/delete/roles', 'POST', '127.0.0.1', '{"id":["3"]}', '2020-03-15 04:14:40.710', '2020-03-15 04:14:40.710');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('364', '1', '/admin/info/manager', 'GET', '127.0.0.1', '', '2020-03-15 04:14:40.763', '2020-03-15 04:14:40.763');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('365', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["123"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 04:14:40.797', '2020-03-15 04:14:40.797');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('366', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:14:40.827', '2020-03-15 04:14:40.827');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('367', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:14:40.880', '2020-03-15 04:14:40.880');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('368', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["ccd05fd4-5bfe-498d-8e14-4fc08df1f25b"],"avatar__delete_flag":["0"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 04:14:41.040', '2020-03-15 04:14:41.040');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('369', '1', '/admin/info/manager/new', 'GET', '127.0.0.1', '', '2020-03-15 04:14:41.080', '2020-03-15 04:14:41.080');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('370', '1', '/admin/new/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["d98d5b45-45a9-410c-97e0-018c22dd2071"],"avatar__delete_flag":["0"],"id":["1"],"name":["tester"],"password":["tester"],"password_again":["tester"],"permission_id[]":["1"],"role_id[]":["1"],"username":["tester"]}', '2020-03-15 04:14:41.207', '2020-03-15 04:14:41.207');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('371', '1', '/admin/menu', 'GET', '127.0.0.1', '', '2020-03-15 04:14:41.480', '2020-03-15 04:14:41.480');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('372', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["67b1a2a6-bf25-435e-9d21-104aa17ff0a9"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test menu"],"uri":["/example/test"]}', '2020-03-15 04:14:41.547', '2020-03-15 04:14:41.547');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('373', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 04:14:41.600', '2020-03-15 04:14:41.600');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('374', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 04:14:41.650', '2020-03-15 04:14:41.650');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('375', '1', '/admin/menu/edit', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["208aab5e-8c58-4934-8737-a711caffeed3"],"header":[""],"icon":["fa-angellist"],"id":["3"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test"]}', '2020-03-15 04:14:41.713', '2020-03-15 04:14:41.713');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('376', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["208aab5e-8c58-4934-8737-a711caffeed3"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test2"]}', '2020-03-15 04:14:41.770', '2020-03-15 04:14:41.770');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('377', '1', '/admin/menu/delete', 'POST', '127.0.0.1', '{}', '2020-03-15 04:14:41.793', '2020-03-15 04:14:41.793');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('378', '1', '/admin/info/op', 'GET', '127.0.0.1', '', '2020-03-15 04:14:41.847', '2020-03-15 04:14:41.847');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('379', '1', '/admin/info/external', 'GET', '127.0.0.1', '', '2020-03-15 04:14:41.880', '2020-03-15 04:14:41.880');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('380', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:14:41.903', '2020-03-15 04:14:41.903');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('381', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:14:41.940', '2020-03-15 04:14:41.940');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('382', '1', '/admin/info/external/new', 'GET', '127.0.0.1', '', '2020-03-15 04:14:41.973', '2020-03-15 04:14:41.973');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('383', '1', '/admin/info/user', 'GET', '127.0.0.1', '', '2020-03-15 04:14:42.000', '2020-03-15 04:14:42.000');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('384', '1', '/admin/export/user', 'POST', '127.0.0.1', '{"id":["1"]}', '2020-03-15 04:14:42.037', '2020-03-15 04:14:42.037');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('385', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:14:42.057', '2020-03-15 04:14:42.057');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('386', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:14:42.090', '2020-03-15 04:14:42.090');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('387', '1', '/admin/info/user/new', 'GET', '127.0.0.1', '', '2020-03-15 04:14:42.127', '2020-03-15 04:14:42.127');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('388', '1', '/admin/logout', 'GET', '127.0.0.1', '', '2020-03-15 04:19:24.570', '2020-03-15 04:19:24.570');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('389', '1', '/admin/info/permission', 'GET', '127.0.0.1', '', '2020-03-15 04:19:24.953', '2020-03-15 04:19:24.953');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('390', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 04:19:25.000', '2020-03-15 04:19:25.000');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('391', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["44aaca60-ddd9-4b32-b81d-1f8af1b01c98"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester"],"slug":["tester"]}', '2020-03-15 04:19:25.033', '2020-03-15 04:19:25.033');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('392', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:19:25.060', '2020-03-15 04:19:25.060');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('393', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:19:25.093', '2020-03-15 04:19:25.093');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('394', '1', '/admin/edit/permission', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["124fad17-284e-4c29-87cd-020626b0a275"],"http_method[]":["GET","POST"],"http_path":["/\n/admin/info/op"],"id":["3"],"name":["tester"],"slug":["tester"]}', '2020-03-15 04:19:25.140', '2020-03-15 04:19:25.140');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('395', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 04:19:25.177', '2020-03-15 04:19:25.177');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('396', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["7f9624e7-86cf-4517-997e-57f9c176682e"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester2"],"slug":["tester2"]}', '2020-03-15 04:19:25.220', '2020-03-15 04:19:25.220');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('397', '1', '/admin/delete/permission', 'POST', '127.0.0.1', '{"id":["4"]}', '2020-03-15 04:19:25.250', '2020-03-15 04:19:25.250');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('398', '1', '/admin/info/roles', 'GET', '127.0.0.1', '', '2020-03-15 04:19:25.280', '2020-03-15 04:19:25.280');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('399', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 04:19:25.330', '2020-03-15 04:19:25.330');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('400', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["1f3ddf1a-931c-4889-b3da-6b6c82806bc0"],"name":["tester"],"permission_id[]":["3"],"slug":["tester"]}', '2020-03-15 04:19:25.377', '2020-03-15 04:19:25.377');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('401', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:19:25.407', '2020-03-15 04:19:25.407');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('402', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:19:25.443', '2020-03-15 04:19:25.443');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('403', '1', '/admin/edit/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["41b65d45-fefa-4587-8fc9-3901b8aa98ce"],"id":["3"],"name":["tester"],"permission_id[]":["3","2"],"slug":["tester"]}', '2020-03-15 04:19:25.487', '2020-03-15 04:19:25.487');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('404', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 04:19:25.533', '2020-03-15 04:19:25.533');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('405', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["6087c942-2c6d-44fb-ac65-0b8c72eb62ab"],"name":["tester2"],"permission_id[]":["3"],"slug":["tester2"]}', '2020-03-15 04:19:25.570', '2020-03-15 04:19:25.570');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('406', '1', '/admin/delete/roles', 'POST', '127.0.0.1', '{"id":["3"]}', '2020-03-15 04:19:25.610', '2020-03-15 04:19:25.610');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('407', '1', '/admin/info/manager', 'GET', '127.0.0.1', '', '2020-03-15 04:19:25.640', '2020-03-15 04:19:25.640');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('408', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["123"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 04:19:25.687', '2020-03-15 04:19:25.687');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('409', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:19:25.720', '2020-03-15 04:19:25.720');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('410', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:19:25.777', '2020-03-15 04:19:25.777');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('411', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["d6b57ee4-01fb-49e8-a154-bd29999b5933"],"avatar__delete_flag":["0"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 04:19:25.907', '2020-03-15 04:19:25.907');
GO
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('412', '1', '/admin/info/manager/new', 'GET', '127.0.0.1', '', '2020-03-15 04:19:25.987', '2020-03-15 04:19:25.987');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('413', '1', '/admin/new/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["26cdd12f-ea5a-4f13-83a1-2498a04025c3"],"avatar__delete_flag":["0"],"id":["1"],"name":["tester"],"password":["tester"],"password_again":["tester"],"permission_id[]":["1"],"role_id[]":["1"],"username":["tester"]}', '2020-03-15 04:19:26.100', '2020-03-15 04:19:26.100');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('414', '1', '/admin/menu', 'GET', '127.0.0.1', '', '2020-03-15 04:19:26.410', '2020-03-15 04:19:26.410');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('415', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["d0b6c9d7-81dc-46e7-8f5c-274ffec50de5"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test menu"],"uri":["/example/test"]}', '2020-03-15 04:19:26.483', '2020-03-15 04:19:26.483');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('416', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 04:19:26.540', '2020-03-15 04:19:26.540');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('417', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 04:19:26.590', '2020-03-15 04:19:26.590');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('418', '1', '/admin/menu/edit', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["d260d214-fcde-4f1e-b076-0abaf3ab1ba5"],"header":[""],"icon":["fa-angellist"],"id":["3"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test"]}', '2020-03-15 04:19:26.670', '2020-03-15 04:19:26.670');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('419', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["d260d214-fcde-4f1e-b076-0abaf3ab1ba5"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test2"]}', '2020-03-15 04:19:26.717', '2020-03-15 04:19:26.717');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('420', '1', '/admin/menu/delete', 'POST', '127.0.0.1', '{}', '2020-03-15 04:19:26.747', '2020-03-15 04:19:26.747');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('421', '1', '/admin/info/op', 'GET', '127.0.0.1', '', '2020-03-15 04:19:26.777', '2020-03-15 04:19:26.777');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('422', '1', '/admin/info/external', 'GET', '127.0.0.1', '', '2020-03-15 04:19:26.817', '2020-03-15 04:19:26.817');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('423', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:19:26.850', '2020-03-15 04:19:26.850');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('424', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:19:26.883', '2020-03-15 04:19:26.883');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('425', '1', '/admin/info/external/new', 'GET', '127.0.0.1', '', '2020-03-15 04:19:26.920', '2020-03-15 04:19:26.920');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('426', '1', '/admin/info/user', 'GET', '127.0.0.1', '', '2020-03-15 04:19:26.950', '2020-03-15 04:19:26.950');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('427', '1', '/admin/export/user', 'POST', '127.0.0.1', '{"id":["1"]}', '2020-03-15 04:19:26.990', '2020-03-15 04:19:26.990');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('428', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:19:27.027', '2020-03-15 04:19:27.027');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('429', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:19:27.060', '2020-03-15 04:19:27.060');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('430', '1', '/admin/info/user/new', 'GET', '127.0.0.1', '', '2020-03-15 04:19:27.100', '2020-03-15 04:19:27.100');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('431', '1', '/admin/logout', 'GET', '127.0.0.1', '', '2020-03-15 04:24:45.410', '2020-03-15 04:24:45.410');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('432', '1', '/admin/info/permission', 'GET', '127.0.0.1', '', '2020-03-15 04:24:45.827', '2020-03-15 04:24:45.827');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('433', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 04:24:45.860', '2020-03-15 04:24:45.860');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('434', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["092217f5-5f35-40a2-93aa-e9a46a739cec"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester"],"slug":["tester"]}', '2020-03-15 04:24:45.893', '2020-03-15 04:24:45.893');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('435', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:24:45.930', '2020-03-15 04:24:45.930');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('436', '1', '/admin/info/permission/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:24:45.967', '2020-03-15 04:24:45.967');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('437', '1', '/admin/edit/permission', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["26457988-35b7-4cfe-b1cc-6505498993c8"],"http_method[]":["GET","POST"],"http_path":["/\n/admin/info/op"],"id":["3"],"name":["tester"],"slug":["tester"]}', '2020-03-15 04:24:46.017', '2020-03-15 04:24:46.017');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('438', '1', '/admin/info/permission/new', 'GET', '127.0.0.1', '', '2020-03-15 04:24:46.057', '2020-03-15 04:24:46.057');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('439', '1', '/admin/new/permission', 'POST', '127.0.0.1', '{"__go_admin_post_type":["1"],"__go_admin_previous_":["/admin/info/permission?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["1cbe9682-8185-4a6a-ab5e-27c7986e00a4"],"http_method[]":["GET"],"http_path":["/\n/admin/info/op"],"name":["tester2"],"slug":["tester2"]}', '2020-03-15 04:24:46.093', '2020-03-15 04:24:46.093');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('440', '1', '/admin/delete/permission', 'POST', '127.0.0.1', '{"id":["4"]}', '2020-03-15 04:24:46.123', '2020-03-15 04:24:46.123');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('441', '1', '/admin/info/roles', 'GET', '127.0.0.1', '', '2020-03-15 04:24:46.170', '2020-03-15 04:24:46.170');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('442', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 04:24:46.210', '2020-03-15 04:24:46.210');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('443', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["e0809bf8-b0ea-4296-8343-188294f0d04d"],"name":["tester"],"permission_id[]":["3"],"slug":["tester"]}', '2020-03-15 04:24:46.250', '2020-03-15 04:24:46.250');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('444', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:24:46.277', '2020-03-15 04:24:46.277');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('445', '1', '/admin/info/roles/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:24:46.320', '2020-03-15 04:24:46.320');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('446', '1', '/admin/edit/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["d2bf536e-07cf-44ac-ad21-702a7a89bb6b"],"id":["3"],"name":["tester"],"permission_id[]":["3","2"],"slug":["tester"]}', '2020-03-15 04:24:46.360', '2020-03-15 04:24:46.360');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('447', '1', '/admin/info/roles/new', 'GET', '127.0.0.1', '', '2020-03-15 04:24:46.400', '2020-03-15 04:24:46.400');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('448', '1', '/admin/new/roles', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/roles?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["e47c72f5-88ba-420c-959e-da5701b25ef1"],"name":["tester2"],"permission_id[]":["3"],"slug":["tester2"]}', '2020-03-15 04:24:46.440', '2020-03-15 04:24:46.440');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('449', '1', '/admin/delete/roles', 'POST', '127.0.0.1', '{"id":["3"]}', '2020-03-15 04:24:46.477', '2020-03-15 04:24:46.477');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('450', '1', '/admin/info/manager', 'GET', '127.0.0.1', '', '2020-03-15 04:24:46.527', '2020-03-15 04:24:46.527');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('451', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["123"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 04:24:46.560', '2020-03-15 04:24:46.560');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('452', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:24:46.590', '2020-03-15 04:24:46.590');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('453', '1', '/admin/info/manager/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:24:46.643', '2020-03-15 04:24:46.643');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('454', '1', '/admin/edit/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["a14b2cb1-8403-4e9e-91e7-642eb43de615"],"avatar__delete_flag":["0"],"id":["1"],"name":["admin1"],"password":["admin"],"password_again":["admin"],"permission_id[]":["1"],"role_id[]":["1"],"username":["admin"]}', '2020-03-15 04:24:46.797', '2020-03-15 04:24:46.797');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('455', '1', '/admin/info/manager/new', 'GET', '127.0.0.1', '', '2020-03-15 04:24:46.840', '2020-03-15 04:24:46.840');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('456', '1', '/admin/new/manager', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/info/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc"],"__go_admin_t_":["bf003301-4b93-4d11-b5e1-117608b5b1a1"],"avatar__delete_flag":["0"],"id":["1"],"name":["tester"],"password":["tester"],"password_again":["tester"],"permission_id[]":["1"],"role_id[]":["1"],"username":["tester"]}', '2020-03-15 04:24:46.963', '2020-03-15 04:24:46.963');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('457', '1', '/admin/menu', 'GET', '127.0.0.1', '', '2020-03-15 04:24:47.243', '2020-03-15 04:24:47.243');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('458', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["83df8f18-88d5-4d47-8b8b-4db42e1f76f1"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test menu"],"uri":["/example/test"]}', '2020-03-15 04:24:47.313', '2020-03-15 04:24:47.313');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('459', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 04:24:47.370', '2020-03-15 04:24:47.370');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('460', '1', '/admin/menu/edit/show', 'GET', '127.0.0.1', '', '2020-03-15 04:24:47.420', '2020-03-15 04:24:47.420');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('461', '1', '/admin/menu/edit', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["916fc225-9d2f-4079-aec2-e53c0effaa41"],"header":[""],"icon":["fa-angellist"],"id":["3"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test"]}', '2020-03-15 04:24:47.493', '2020-03-15 04:24:47.493');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('462', '1', '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["916fc225-9d2f-4079-aec2-e53c0effaa41"],"header":[""],"icon":["fa-angellist"],"parent_id":[""],"roles[]":["1"],"title":["test2 menu"],"uri":["/example/test2"]}', '2020-03-15 04:24:47.547', '2020-03-15 04:24:47.547');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('463', '1', '/admin/menu/delete', 'POST', '127.0.0.1', '{}', '2020-03-15 04:24:47.573', '2020-03-15 04:24:47.573');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('464', '1', '/admin/info/op', 'GET', '127.0.0.1', '', '2020-03-15 04:24:47.633', '2020-03-15 04:24:47.633');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('465', '1', '/admin/info/external', 'GET', '127.0.0.1', '', '2020-03-15 04:24:47.670', '2020-03-15 04:24:47.670');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('466', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:24:47.700', '2020-03-15 04:24:47.700');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('467', '1', '/admin/info/external/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:24:47.740', '2020-03-15 04:24:47.740');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('468', '1', '/admin/info/external/new', 'GET', '127.0.0.1', '', '2020-03-15 04:24:47.777', '2020-03-15 04:24:47.777');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('469', '1', '/admin/info/user', 'GET', '127.0.0.1', '', '2020-03-15 04:24:47.837', '2020-03-15 04:24:47.837');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('470', '1', '/admin/export/user', 'POST', '127.0.0.1', '{"id":["1"]}', '2020-03-15 04:24:47.870', '2020-03-15 04:24:47.870');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('471', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:24:47.903', '2020-03-15 04:24:47.903');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('472', '1', '/admin/info/user/edit', 'GET', '127.0.0.1', '', '2020-03-15 04:24:47.940', '2020-03-15 04:24:47.940');
INSERT INTO [dbo].[goadmin_operation_log] ([id], [user_id], [path], [method], [ip], [input], [created_at], [updated_at]) VALUES ('473', '1', '/admin/info/user/new', 'GET', '127.0.0.1', '', '2020-03-15 04:24:47.980', '2020-03-15 04:24:47.980');
GO
SET IDENTITY_INSERT [dbo].[goadmin_operation_log] OFF
GO
COMMIT
GO

-- ----------------------------
--  Table structure for goadmin_permissions
-- ----------------------------
IF EXISTS (SELECT * FROM sys.all_objects WHERE object_id = OBJECT_ID('[dbo].[goadmin_permissions]') AND type IN ('U'))
	DROP TABLE [dbo].[goadmin_permissions]
GO
CREATE TABLE [dbo].[goadmin_permissions] (
	[id] int IDENTITY(1,1) NOT NULL,
	[name] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	[slug] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	[http_method] varchar(255) COLLATE SQL_Latin1_General_CP1_CI_AS NULL DEFAULT NULL,
	[http_path] text COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	[created_at] datetime NULL DEFAULT (getdate()),
	[updated_at] datetime NULL DEFAULT (getdate())
)
ON [PRIMARY]
TEXTIMAGE_ON [PRIMARY]
GO

-- ----------------------------
--  Records of goadmin_permissions
-- ----------------------------
BEGIN TRANSACTION
GO
SET IDENTITY_INSERT [dbo].[goadmin_permissions] ON
GO
INSERT INTO [dbo].[goadmin_permissions] ([id], [name], [slug], [http_method], [http_path], [created_at], [updated_at]) VALUES ('1', 'All permission', '*', '', '*', '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
INSERT INTO [dbo].[goadmin_permissions] ([id], [name], [slug], [http_method], [http_path], [created_at], [updated_at]) VALUES ('2', 'Dashboard', 'dashboard', 'GET,PUT,POST,DELETE', '/', '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
GO
SET IDENTITY_INSERT [dbo].[goadmin_permissions] OFF
GO
COMMIT
GO

-- ----------------------------
--  Table structure for goadmin_role_menu
-- ----------------------------
IF EXISTS (SELECT * FROM sys.all_objects WHERE object_id = OBJECT_ID('[dbo].[goadmin_role_menu]') AND type IN ('U'))
	DROP TABLE [dbo].[goadmin_role_menu]
GO
CREATE TABLE [dbo].[goadmin_role_menu] (
	[role_id] int NOT NULL,
	[menu_id] int NOT NULL,
	[created_at] datetime NULL DEFAULT (getdate()),
	[updated_at] datetime NULL DEFAULT (getdate())
)
ON [PRIMARY]
GO

-- ----------------------------
--  Records of goadmin_role_menu
-- ----------------------------
BEGIN TRANSACTION
GO
INSERT INTO [dbo].[goadmin_role_menu] VALUES ('1', '1', '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
INSERT INTO [dbo].[goadmin_role_menu] VALUES ('1', '7', '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
INSERT INTO [dbo].[goadmin_role_menu] VALUES ('1', '8', '2019-09-11 10:20:55.000', '2019-09-11 10:20:55.000');
INSERT INTO [dbo].[goadmin_role_menu] VALUES ('2', '7', '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
INSERT INTO [dbo].[goadmin_role_menu] VALUES ('2', '8', '2019-09-11 10:20:55.000', '2019-09-11 10:20:55.000');
GO
COMMIT
GO

-- ----------------------------
--  Table structure for goadmin_role_permissions
-- ----------------------------
IF EXISTS (SELECT * FROM sys.all_objects WHERE object_id = OBJECT_ID('[dbo].[goadmin_role_permissions]') AND type IN ('U'))
	DROP TABLE [dbo].[goadmin_role_permissions]
GO
CREATE TABLE [dbo].[goadmin_role_permissions] (
	[role_id] int NOT NULL,
	[permission_id] int NOT NULL,
	[created_at] datetime NULL DEFAULT (getdate()),
	[updated_at] datetime NULL DEFAULT (getdate())
)
ON [PRIMARY]
GO

-- ----------------------------
--  Records of goadmin_role_permissions
-- ----------------------------
BEGIN TRANSACTION
GO
INSERT INTO [dbo].[goadmin_role_permissions] VALUES ('1', '1', '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
INSERT INTO [dbo].[goadmin_role_permissions] VALUES ('1', '2', '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
INSERT INTO [dbo].[goadmin_role_permissions] VALUES ('2', '2', '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
GO
COMMIT
GO

-- ----------------------------
--  Table structure for goadmin_role_users
-- ----------------------------
IF EXISTS (SELECT * FROM sys.all_objects WHERE object_id = OBJECT_ID('[dbo].[goadmin_role_users]') AND type IN ('U'))
	DROP TABLE [dbo].[goadmin_role_users]
GO
CREATE TABLE [dbo].[goadmin_role_users] (
	[role_id] int NOT NULL,
	[user_id] int NOT NULL,
	[created_at] datetime NULL DEFAULT (getdate()),
	[updated_at] datetime NULL DEFAULT (getdate())
)
ON [PRIMARY]
GO

-- ----------------------------
--  Records of goadmin_role_users
-- ----------------------------
BEGIN TRANSACTION
GO
INSERT INTO [dbo].[goadmin_role_users] VALUES ('1', '1', '2020-03-15 04:24:46.757', '2020-03-15 04:24:46.757');
INSERT INTO [dbo].[goadmin_role_users] VALUES ('2', '2', '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
GO
COMMIT
GO

-- ----------------------------
--  Table structure for goadmin_roles
-- ----------------------------
IF EXISTS (SELECT * FROM sys.all_objects WHERE object_id = OBJECT_ID('[dbo].[goadmin_roles]') AND type IN ('U'))
	DROP TABLE [dbo].[goadmin_roles]
GO
CREATE TABLE [dbo].[goadmin_roles] (
	[id] int IDENTITY(1,1) NOT NULL,
	[name] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	[slug] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	[created_at] datetime NULL DEFAULT (getdate()),
	[updated_at] datetime NULL DEFAULT (getdate())
)
ON [PRIMARY]
GO

-- ----------------------------
--  Records of goadmin_roles
-- ----------------------------
BEGIN TRANSACTION
GO
SET IDENTITY_INSERT [dbo].[goadmin_roles] ON
GO
INSERT INTO [dbo].[goadmin_roles] ([id], [name], [slug], [created_at], [updated_at]) VALUES ('1', 'Administrator', 'administrator', '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
INSERT INTO [dbo].[goadmin_roles] ([id], [name], [slug], [created_at], [updated_at]) VALUES ('2', 'Operator', 'operator', '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
GO
SET IDENTITY_INSERT [dbo].[goadmin_roles] OFF
GO
COMMIT
GO

-- ----------------------------
--  Table structure for goadmin_session
-- ----------------------------
IF EXISTS (SELECT * FROM sys.all_objects WHERE object_id = OBJECT_ID('[dbo].[goadmin_session]') AND type IN ('U'))
	DROP TABLE [dbo].[goadmin_session]
GO
CREATE TABLE [dbo].[goadmin_session] (
	[id] int IDENTITY(1,1) NOT NULL,
	[sid] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS NULL DEFAULT '',
	[values] varchar(3000) COLLATE SQL_Latin1_General_CP1_CI_AS NULL DEFAULT '',
	[created_at] datetime NULL DEFAULT (getdate()),
	[updated_at] datetime NULL DEFAULT (getdate())
)
ON [PRIMARY]
GO

-- ----------------------------
--  Records of goadmin_session
-- ----------------------------
BEGIN TRANSACTION
GO
SET IDENTITY_INSERT [dbo].[goadmin_session] ON
GO
INSERT INTO [dbo].[goadmin_session] ([id], [sid], [values], [created_at], [updated_at]) VALUES ('43', '5ada7ffd-1025-4bfe-802a-43b7e6e79582', '{"user_id":1}', '2020-03-15 04:24:45.750', '2020-03-15 04:24:45.750');
INSERT INTO [dbo].[goadmin_session] ([id], [sid], [values], [created_at], [updated_at]) VALUES ('44', 'a01d8b11-083e-4075-8287-1bfb1865fe70', '{"user_id":3}', '2020-03-15 04:24:47.190', '2020-03-15 04:24:47.190');
GO
SET IDENTITY_INSERT [dbo].[goadmin_session] OFF
GO
COMMIT
GO

-- ----------------------------
--  Table structure for goadmin_site
-- ----------------------------
IF EXISTS (SELECT * FROM sys.all_objects WHERE object_id = OBJECT_ID('[dbo].[goadmin_site]') AND type IN ('U'))
	DROP TABLE [dbo].[goadmin_site]
GO
CREATE TABLE [dbo].[goadmin_site] (
	[id] int IDENTITY(1,1) NOT NULL,
	[key] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	[value] text COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	[state] tinyint NULL DEFAULT ((0)),
	[created_at] datetime NULL DEFAULT (getdate()),
	[updated_at] datetime NULL DEFAULT (getdate())
)
ON [PRIMARY]
TEXTIMAGE_ON [PRIMARY]
GO

-- ----------------------------
--  Table structure for goadmin_user_permissions
-- ----------------------------
IF EXISTS (SELECT * FROM sys.all_objects WHERE object_id = OBJECT_ID('[dbo].[goadmin_user_permissions]') AND type IN ('U'))
	DROP TABLE [dbo].[goadmin_user_permissions]
GO
CREATE TABLE [dbo].[goadmin_user_permissions] (
	[user_id] int NOT NULL,
	[permission_id] int NOT NULL,
	[created_at] datetime NULL DEFAULT (getdate()),
	[updated_at] datetime NULL DEFAULT (getdate())
)
ON [PRIMARY]
GO

-- ----------------------------
--  Records of goadmin_user_permissions
-- ----------------------------
BEGIN TRANSACTION
GO
INSERT INTO [dbo].[goadmin_user_permissions] VALUES ('1', '1', '2020-03-15 04:24:46.763', '2020-03-15 04:24:46.763');
INSERT INTO [dbo].[goadmin_user_permissions] VALUES ('2', '2', '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
GO
COMMIT
GO

-- ----------------------------
--  Table structure for goadmin_users
-- ----------------------------
IF EXISTS (SELECT * FROM sys.all_objects WHERE object_id = OBJECT_ID('[dbo].[goadmin_users]') AND type IN ('U'))
	DROP TABLE [dbo].[goadmin_users]
GO
CREATE TABLE [dbo].[goadmin_users] (
	[id] int IDENTITY(1,1) NOT NULL,
	[username] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	[password] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL DEFAULT '',
	[name] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	[avatar] varchar(255) COLLATE SQL_Latin1_General_CP1_CI_AS NULL DEFAULT NULL,
	[remember_token] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL DEFAULT NULL,
	[created_at] datetime NULL DEFAULT (getdate()),
	[updated_at] datetime NULL DEFAULT (getdate())
)
ON [PRIMARY]
GO

-- ----------------------------
--  Records of goadmin_users
-- ----------------------------
BEGIN TRANSACTION
GO
SET IDENTITY_INSERT [dbo].[goadmin_users] ON
GO
INSERT INTO [dbo].[goadmin_users] ([id], [username], [password], [name], [avatar], [remember_token], [created_at], [updated_at]) VALUES ('1', 'admin', '$2a$10$9Bb3g7FUC/dDhYISUWegBOHy0498EuV50rVCgxkvAbpjavrKwhjvm', 'admin1', '', 'tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh', '2019-09-10 00:00:00.000', '2020-03-15 12:24:46.000');
INSERT INTO [dbo].[goadmin_users] ([id], [username], [password], [name], [avatar], [remember_token], [created_at], [updated_at]) VALUES ('2', 'operator', '$2a$10$rVqkOzHjN2MdlEprRflb1eGP0oZXuSrbJLOmJagFsCd81YZm0bsh.', 'Operator', '', null, '2019-09-10 00:00:00.000', '2019-09-10 00:00:00.000');
GO
SET IDENTITY_INSERT [dbo].[goadmin_users] OFF
GO
COMMIT
GO

-- ----------------------------
--  Table structure for user_like_books
-- ----------------------------
IF EXISTS (SELECT * FROM sys.all_objects WHERE object_id = OBJECT_ID('[dbo].[user_like_books]') AND type IN ('U'))
	DROP TABLE [dbo].[user_like_books]
GO
CREATE TABLE [dbo].[user_like_books] (
	[id] int IDENTITY(1,1) NOT NULL,
	[user_id] int NULL,
	[name] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	[created_at] datetime NULL DEFAULT (getdate()),
	[updated_at] datetime NULL DEFAULT (getdate())
)
ON [PRIMARY]
GO

-- ----------------------------
--  Records of user_like_books
-- ----------------------------
BEGIN TRANSACTION
GO
SET IDENTITY_INSERT [dbo].[user_like_books] ON
GO
INSERT INTO [dbo].[user_like_books] ([id], [user_id], [name], [created_at], [updated_at]) VALUES ('1', '1', 'Robinson Crusoe', '2020-03-15 09:57:49.000', '2020-03-15 09:57:51.000');
GO
SET IDENTITY_INSERT [dbo].[user_like_books] OFF
GO
COMMIT
GO

-- ----------------------------
--  Table structure for users
-- ----------------------------
IF EXISTS (SELECT * FROM sys.all_objects WHERE object_id = OBJECT_ID('[dbo].[users]') AND type IN ('U'))
	DROP TABLE [dbo].[users]
GO
CREATE TABLE [dbo].[users] (
	[id] int IDENTITY(1,1) NOT NULL,
	[name] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	[homepage] varchar(3000) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	[email] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	[birthday] timestamp NULL,
	[country] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	[member_id] smallint NULL DEFAULT ((0)),
	[city] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	[password] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	[ip] varchar(20) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	[certificate] varchar(300) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	[money] int NULL,
	[age] int NULL DEFAULT ((0)),
	[resume] text COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	[gender] smallint NULL DEFAULT ((0)),
	[fruit] varchar(200) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	[drink] varchar(200) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	[experience] smallint NULL DEFAULT ((0)),
	[created_at] datetime NULL,
	[updated_at] datetime NULL
)
ON [PRIMARY]
TEXTIMAGE_ON [PRIMARY]
GO

-- ----------------------------
--  Records of users
-- ----------------------------
BEGIN TRANSACTION
GO
SET IDENTITY_INSERT [dbo].[users] ON
GO
INSERT INTO [dbo].[users] ([id], [name], [homepage], [email], [birthday], [country], [member_id], [city], [password], [ip], [certificate], [money], [age], [resume], [gender], [fruit], [drink], [experience], [created_at], [updated_at]) VALUES ('1', 'Jack', 'http://jack.me', 'jack@163.com', DEFAULT, 'china', '6', 'guangzhou', '123456', '127.0.0.1', null, '503', '25', '<h1>Jack`s Resume</h1>', '0', 'apple', 'water', '0', '2020-03-15 09:54:39.000', '2020-03-15 09:54:42.000');
GO
SET IDENTITY_INSERT [dbo].[users] OFF
GO
COMMIT
GO


-- ----------------------------
--  Primary key structure for table goadmin_menu
-- ----------------------------
ALTER TABLE [dbo].[goadmin_menu] ADD
	CONSTRAINT [PK__goadmin___3213E83FF5EBF8CC] PRIMARY KEY CLUSTERED ([id]) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [default]
GO

-- ----------------------------
--  Primary key structure for table goadmin_operation_log
-- ----------------------------
ALTER TABLE [dbo].[goadmin_operation_log] ADD
	CONSTRAINT [PK__goadmin___3213E83F6C7E1A25] PRIMARY KEY CLUSTERED ([id]) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [default]
GO

-- ----------------------------
--  Primary key structure for table goadmin_permissions
-- ----------------------------
ALTER TABLE [dbo].[goadmin_permissions] ADD
	CONSTRAINT [PK__goadmin___3213E83F6A86BC1A] PRIMARY KEY CLUSTERED ([id]) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [default]
GO

-- ----------------------------
--  Primary key structure for table goadmin_role_menu
-- ----------------------------
ALTER TABLE [dbo].[goadmin_role_menu] ADD
	CONSTRAINT [PK__goadmin___A2C36A610BBE2A9C] PRIMARY KEY CLUSTERED ([role_id],[menu_id]) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [default]
GO

-- ----------------------------
--  Primary key structure for table goadmin_role_permissions
-- ----------------------------
ALTER TABLE [dbo].[goadmin_role_permissions] ADD
	CONSTRAINT [PK__goadmin___C85A54636259324C] PRIMARY KEY CLUSTERED ([role_id],[permission_id]) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [default]
GO

-- ----------------------------
--  Primary key structure for table goadmin_role_users
-- ----------------------------
ALTER TABLE [dbo].[goadmin_role_users] ADD
	CONSTRAINT [PK__goadmin___9D9286BC2FC99BF2] PRIMARY KEY CLUSTERED ([role_id],[user_id]) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [default]
GO

-- ----------------------------
--  Primary key structure for table goadmin_roles
-- ----------------------------
ALTER TABLE [dbo].[goadmin_roles] ADD
	CONSTRAINT [PK__goadmin___3213E83F7D29D84E] PRIMARY KEY CLUSTERED ([id]) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [default]
GO

-- ----------------------------
--  Uniques structure for table goadmin_roles
-- ----------------------------
ALTER TABLE [dbo].[goadmin_roles] ADD
	CONSTRAINT [UQ__goadmin___72E12F1BF4C046D9] UNIQUE NONCLUSTERED ([name] ASC) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [PRIMARY]
GO

-- ----------------------------
--  Primary key structure for table goadmin_session
-- ----------------------------
ALTER TABLE [dbo].[goadmin_session] ADD
	CONSTRAINT [PK__goadmin___3213E83FD9B9AF30] PRIMARY KEY CLUSTERED ([id]) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [default]
GO

-- ----------------------------
--  Primary key structure for table goadmin_site
-- ----------------------------
ALTER TABLE [dbo].[goadmin_site] ADD
	CONSTRAINT [PK__goadmin___3213E83FBC0B9B22] PRIMARY KEY CLUSTERED ([id]) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [default]
GO

-- ----------------------------
--  Primary key structure for table goadmin_user_permissions
-- ----------------------------
ALTER TABLE [dbo].[goadmin_user_permissions] ADD
	CONSTRAINT [PK__goadmin___07ED06A0046D7D01] PRIMARY KEY CLUSTERED ([user_id],[permission_id]) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [default]
GO

-- ----------------------------
--  Primary key structure for table goadmin_users
-- ----------------------------
ALTER TABLE [dbo].[goadmin_users] ADD
	CONSTRAINT [PK__goadmin___3213E83F354C3AA1] PRIMARY KEY CLUSTERED ([id]) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [default]
GO

-- ----------------------------
--  Uniques structure for table goadmin_users
-- ----------------------------
ALTER TABLE [dbo].[goadmin_users] ADD
	CONSTRAINT [UQ__goadmin___F3DBC572BFCF4D97] UNIQUE NONCLUSTERED ([username] ASC) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [PRIMARY]
GO

-- ----------------------------
--  Primary key structure for table user_like_books
-- ----------------------------
ALTER TABLE [dbo].[user_like_books] ADD
	CONSTRAINT [PK__user_lik__3213E83F151C3FE8] PRIMARY KEY CLUSTERED ([id]) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [default]
GO

-- ----------------------------
--  Primary key structure for table users
-- ----------------------------
ALTER TABLE [dbo].[users] ADD
	CONSTRAINT [PK__users__3213E83F2F479EFF] PRIMARY KEY CLUSTERED ([id]) 
	WITH (PAD_INDEX = OFF,
		IGNORE_DUP_KEY = OFF,
		ALLOW_ROW_LOCKS = ON,
		ALLOW_PAGE_LOCKS = ON)
	ON [default]
GO

-- ----------------------------
--  Options for table goadmin_menu
-- ----------------------------
ALTER TABLE [dbo].[goadmin_menu] SET (LOCK_ESCALATION = TABLE)
GO
DBCC CHECKIDENT ('[dbo].[goadmin_menu]', RESEED, 9)
GO

-- ----------------------------
--  Options for table goadmin_operation_log
-- ----------------------------
ALTER TABLE [dbo].[goadmin_operation_log] SET (LOCK_ESCALATION = TABLE)
GO
DBCC CHECKIDENT ('[dbo].[goadmin_operation_log]', RESEED, 473)
GO

-- ----------------------------
--  Options for table goadmin_permissions
-- ----------------------------
ALTER TABLE [dbo].[goadmin_permissions] SET (LOCK_ESCALATION = TABLE)
GO
DBCC CHECKIDENT ('[dbo].[goadmin_permissions]', RESEED, 2)
GO

-- ----------------------------
--  Options for table goadmin_role_menu
-- ----------------------------
ALTER TABLE [dbo].[goadmin_role_menu] SET (LOCK_ESCALATION = TABLE)
GO

-- ----------------------------
--  Options for table goadmin_role_permissions
-- ----------------------------
ALTER TABLE [dbo].[goadmin_role_permissions] SET (LOCK_ESCALATION = TABLE)
GO

-- ----------------------------
--  Options for table goadmin_role_users
-- ----------------------------
ALTER TABLE [dbo].[goadmin_role_users] SET (LOCK_ESCALATION = TABLE)
GO

-- ----------------------------
--  Options for table goadmin_roles
-- ----------------------------
ALTER TABLE [dbo].[goadmin_roles] SET (LOCK_ESCALATION = TABLE)
GO
DBCC CHECKIDENT ('[dbo].[goadmin_roles]', RESEED, 2)
GO

-- ----------------------------
--  Options for table goadmin_session
-- ----------------------------
ALTER TABLE [dbo].[goadmin_session] SET (LOCK_ESCALATION = TABLE)
GO
DBCC CHECKIDENT ('[dbo].[goadmin_session]', RESEED, 44)
GO

-- ----------------------------
--  Options for table goadmin_site
-- ----------------------------
ALTER TABLE [dbo].[goadmin_site] SET (LOCK_ESCALATION = TABLE)
GO
DBCC CHECKIDENT ('[dbo].[goadmin_site]', RESEED, 1)
GO

-- ----------------------------
--  Options for table goadmin_user_permissions
-- ----------------------------
ALTER TABLE [dbo].[goadmin_user_permissions] SET (LOCK_ESCALATION = TABLE)
GO

-- ----------------------------
--  Options for table goadmin_users
-- ----------------------------
ALTER TABLE [dbo].[goadmin_users] SET (LOCK_ESCALATION = TABLE)
GO
DBCC CHECKIDENT ('[dbo].[goadmin_users]', RESEED, 3)
GO

-- ----------------------------
--  Options for table user_like_books
-- ----------------------------
ALTER TABLE [dbo].[user_like_books] SET (LOCK_ESCALATION = TABLE)
GO
DBCC CHECKIDENT ('[dbo].[user_like_books]', RESEED, 1)
GO

-- ----------------------------
--  Options for table users
-- ----------------------------
ALTER TABLE [dbo].[users] SET (LOCK_ESCALATION = TABLE)
GO
DBCC CHECKIDENT ('[dbo].[users]', RESEED, 1)
GO

