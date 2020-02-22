package db

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/magiconair/properties/assert"
	"regexp"
	"strings"
	"testing"
)

const (
	dbName            = "go-admin-type-test"
	tableName         = "all_types"
	postgresCreateSql = `CREATE TABLE public.%s
(
    id integer NOT NULL,
    type_1 smallint,
    type_2 bigint,
    type_3 numeric,
    type_4 real,
    type_5 double precision,
    type_6 smallint NOT NULL DEFAULT nextval('all_types_type_6_seq'::regclass),
    type_7 integer NOT NULL DEFAULT nextval('all_types_type_7_seq'::regclass),
    type_8 bigint NOT NULL DEFAULT nextval('all_types_type_8_seq'::regclass),
    type_9 money,
    type_10 character varying COLLATE pg_catalog."default",
    type_11 character(1) COLLATE pg_catalog."default",
    type_12 text COLLATE pg_catalog."default",
    type_13 timestamp with time zone,
    type_14 time with time zone,
    type_15 date,
    type_16 timestamp without time zone,
    type_17 interval,
    type_18 point,
    type_19 line,
    type_20 lseg,
    type_21 box,
    type_22 path,
    type_23 polygon,
    type_24 circle,
    type_25 cidr,
    type_26 inet,
    type_27 macaddr,
    CONSTRAINT all_types_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.all_types
    OWNER to postgres;`
)

func TestMysqlGetTypeFromString(t *testing.T) {

	conn := connection(DriverMysql, fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/%s", dbName))
	_, err := conn.Exec(fmt.Sprintf("create database if not exists `%s`", dbName))
	assert.Equal(t, err, nil)
	_, err = conn.Exec(fmt.Sprintf("DROP TABLE IF EXISTS `%s`;", tableName))
	assert.Equal(t, err, nil)
	_, err = conn.Exec(fmt.Sprintf(`CREATE TABLE `+"`"+`%s`+"`"+` (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  `+delimiter("type_1")+` tinyint(11) DEFAULT NULL,
  `+delimiter("type_2")+` smallint(11) DEFAULT NULL,
  `+delimiter("type_3")+` mediumint(11) DEFAULT NULL,
  `+delimiter("type_4")+` bigint(11) DEFAULT NULL,
  `+delimiter("type_5")+` float DEFAULT NULL,
  `+delimiter("type_6")+` double(5,3) DEFAULT NULL,
  `+delimiter("type_7")+` double DEFAULT NULL,
  `+delimiter("type_8")+` double(5,3) DEFAULT NULL,
  `+delimiter("type_9")+` decimal(11,0) DEFAULT NULL,
  `+delimiter("type_10")+` bit(11) DEFAULT NULL,
  `+delimiter("type_11")+` tinyint(1) DEFAULT NULL,
  `+delimiter("type_12")+` tinyint(1) DEFAULT NULL,
  `+delimiter("type_13")+` decimal(10,5) DEFAULT NULL,
  `+delimiter("type_14")+` decimal(10,0) DEFAULT NULL,
  `+delimiter("type_15")+` decimal(10,0) DEFAULT NULL,
  `+delimiter("type_16")+` char(11) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `+delimiter("type_17")+` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `+delimiter("type_18")+` tinytext COLLATE utf8mb4_unicode_ci,
  `+delimiter("type_19")+` text COLLATE utf8mb4_unicode_ci,
  `+delimiter("type_20")+` mediumtext COLLATE utf8mb4_unicode_ci,
  `+delimiter("type_21")+` longtext COLLATE utf8mb4_unicode_ci,
  `+delimiter("type_22")+` tinyblob,
  `+delimiter("type_23")+` mediumblob,
  `+delimiter("type_24")+` blob,
  `+delimiter("type_25")+` longblob,
  `+delimiter("type_26")+` binary(1) DEFAULT NULL,
  `+delimiter("type_27")+` varbinary(1) DEFAULT NULL,
  `+delimiter("type_28")+` enum('RED','GREEN','BLUE') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `+delimiter("type_29")+` set('RED','GREEN','BLUE') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `+delimiter("type_30")+` date DEFAULT NULL,
  `+delimiter("type_31")+` datetime DEFAULT NULL,
  `+delimiter("type_32")+` timestamp NULL DEFAULT NULL,
  `+delimiter("type_33")+` time DEFAULT NULL,
  `+delimiter("type_34")+` year(4) DEFAULT NULL,
  `+delimiter("type_35")+` geometry DEFAULT NULL,
  `+delimiter("type_36")+` point DEFAULT NULL,
  `+delimiter("type_39")+` multilinestring DEFAULT NULL,
  `+delimiter("type_41")+` multipolygon DEFAULT NULL,
  `+delimiter("type_37")+` linestring DEFAULT NULL,
  `+delimiter("type_38")+` polygon DEFAULT NULL,
  `+delimiter("type_40")+` multipoint DEFAULT NULL,
  `+delimiter("type_42")+` geometrycollection DEFAULT NULL,
  `+delimiter("type_50")+` double(5,2) DEFAULT NULL,
  `+delimiter("type_51")+` json DEFAULT NULL,
  PRIMARY KEY (`+"`"+`id`+"`"+`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`, tableName))

	assert.Equal(t, err, nil)

	_, err = conn.Exec(`INSERT INTO ` + delimiter(tableName) + ` (id, type_1, type_2, type_3, type_4, type_5, type_6, type_7, type_8, type_9, type_10, type_11, type_12, type_13, type_14, type_15, type_16, type_17, type_18, type_19, type_20, type_21, type_22, type_23, type_24, type_25, type_26, type_27, type_28, type_29, type_30, type_31, type_32, type_33, type_34, type_35, type_36, type_39, type_41, type_37, type_38, type_40, type_42, type_50, type_51)
VALUES
	(1, 1, 1, 1, 1, 1, 1.000, 1, 1.000, 1, 0, 1, 1, 1.00000, 1, 1, '1', '1', '1', '1', '1', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2001', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);`)

	assert.Equal(t, err, nil)

	typeField := "Type"
	fieldField := "Field"

	conn = connection(DriverMysql, fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/%s?charset=utf8mb4", dbName))

	config.Set(config.Config{
		SqlLog: true,
	})

	columnsModel, _ := WithDriver(conn).Table(tableName).ShowColumns()
	item, err := WithDriver(conn).Table(tableName).First()

	for _, model := range columnsModel {
		fieldTypeName := strings.ToUpper(getType(model[typeField].(string)))
		GetDTAndCheck(fieldTypeName)
		GetValueFromSQLOfDatabaseType(DatabaseType(fieldTypeName), item[model[fieldField].(string)])
	}
	assert.Equal(t, err, nil)
}

func TestPostgresqlGetTypeFromString(t *testing.T) {

	// pg 11
	testPG(t, "5433")
	// pg 12
	//testPG(t, "5434")
}

func testPG(t *testing.T, port string) {
	connStatement := "host=127.0.0.1 port=" + port + " user=postgres password=root dbname=%s sslmode=disable"

	conn := connection(DriverPostgresql, fmt.Sprintf(connStatement, dbName))
	fmt.Println("creating database")
	_, err := conn.Exec(fmt.Sprintf(`SELECT 'CREATE DATABASE %s' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '%s')`, dbName, dbName))
	assert.Equal(t, err, nil)
	fmt.Println("drop table")
	_, err = conn.Exec(fmt.Sprintf("DROP TABLE IF EXISTS `%s`;", tableName))
	assert.Equal(t, err, nil)
	fmt.Println("create sequence all_types_type_6_seq")
	_, err = conn.Exec(`CREATE SEQUENCE IF NOT EXISTS public.all_types_type_6_seq START 1;`)
	assert.Equal(t, err, nil)
	fmt.Println("create sequence all_types_type_7_seq")
	_, err = conn.Exec(`CREATE SEQUENCE IF NOT EXISTS public.all_types_type_7_seq START 1;`)
	assert.Equal(t, err, nil)
	fmt.Println("create sequence all_types_type_8_seq")
	_, err = conn.Exec(`CREATE SEQUENCE IF NOT EXISTS public.all_types_type_8_seq START 1;`)
	assert.Equal(t, err, nil)
	fmt.Println("create table")
	_, err = conn.Exec(fmt.Sprintf(postgresCreateSql, tableName))

	assert.Equal(t, err, nil)

	fmt.Println("insert data")
	_, err = conn.Exec(`INSERT INTO public.` + tableName + `(
	id, type_1, type_2, type_3, type_4, type_5, type_6, type_7, type_8, type_9, type_10, type_11, type_12, type_13, type_14, type_15, type_16, type_17, type_18, type_19, type_20, type_21, type_22, type_23, type_24, type_25, type_26, type_27)
	VALUES (1, 1, 1, 0.3, 1, 1, 1, 1, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);`)

	assert.Equal(t, err, nil)

	typeField := "udt_name"
	fieldField := "column_name"

	conn = connection(DriverPostgresql, fmt.Sprintf(connStatement, dbName))

	config.Set(config.Config{
		SqlLog: true,
	})

	columnsModel, _ := WithDriver(conn).Table(tableName).ShowColumns()
	item, err := WithDriver(conn).Table(tableName).First()

	for _, model := range columnsModel {
		fieldTypeName := strings.ToUpper(getType(model[typeField].(string)))
		GetDTAndCheck(fieldTypeName)
		fmt.Println(model[fieldField].(string), GetValueFromSQLOfDatabaseType(DatabaseType(fieldTypeName), item[model[fieldField].(string)]))
	}

	assert.Equal(t, err, nil)
}

// *******************************
// helper method
// *******************************

func getType(typeName string) string {
	r, _ := regexp.Compile(`\((.*?)\)`)
	typeName = r.ReplaceAllString(typeName, "")
	return strings.ToLower(strings.Replace(typeName, " unsigned", "", -1))
}

func connection(driver, dsn string) Connection {
	return GetConnectionByDriver(driver).InitDB(map[string]config.Database{
		"default": {Dsn: dsn},
	})
}

func delimiter(s string) string {
	return "`" + s + "`"
}
