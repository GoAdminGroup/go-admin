package db

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/magiconair/properties/assert"
)

const (
	typeTestdbName            = "go-admin-type-test"
	typeTesttableName         = "all_types"
	typeTestpostgresCreateSql = `CREATE TABLE public.%s
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
	type_28 boolean,
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

	conn := testConnDSN(DriverMysql, fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/%s", typeTestdbName))
	_, err := conn.Exec(fmt.Sprintf("create database if not exists `%s`", typeTestdbName))
	assert.Equal(t, err, nil)
	_, err = conn.Exec(fmt.Sprintf("DROP TABLE IF EXISTS `%s`;", typeTesttableName))
	assert.Equal(t, err, nil)
	_, err = conn.Exec(fmt.Sprintf(`CREATE TABLE `+"`"+`%s`+"`"+` (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  `+testDelimiter("type_1")+` tinyint(11) DEFAULT NULL,
  `+testDelimiter("type_2")+` smallint(11) DEFAULT NULL,
  `+testDelimiter("type_3")+` mediumint(11) DEFAULT NULL,
  `+testDelimiter("type_4")+` bigint(11) DEFAULT NULL,
  `+testDelimiter("type_5")+` float DEFAULT NULL,
  `+testDelimiter("type_6")+` double(5,3) DEFAULT NULL,
  `+testDelimiter("type_7")+` double DEFAULT NULL,
  `+testDelimiter("type_8")+` double(5,3) DEFAULT NULL,
  `+testDelimiter("type_9")+` decimal(11,0) DEFAULT NULL,
  `+testDelimiter("type_10")+` bit(11) DEFAULT NULL,
  `+testDelimiter("type_11")+` tinyint(1) DEFAULT NULL,
  `+testDelimiter("type_12")+` tinyint(1) DEFAULT NULL,
  `+testDelimiter("type_13")+` decimal(10,5) DEFAULT NULL,
  `+testDelimiter("type_14")+` decimal(10,0) DEFAULT NULL,
  `+testDelimiter("type_15")+` decimal(10,0) DEFAULT NULL,
  `+testDelimiter("type_16")+` char(11) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `+testDelimiter("type_17")+` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `+testDelimiter("type_18")+` tinytext COLLATE utf8mb4_unicode_ci,
  `+testDelimiter("type_19")+` text COLLATE utf8mb4_unicode_ci,
  `+testDelimiter("type_20")+` mediumtext COLLATE utf8mb4_unicode_ci,
  `+testDelimiter("type_21")+` longtext COLLATE utf8mb4_unicode_ci,
  `+testDelimiter("type_22")+` tinyblob,
  `+testDelimiter("type_23")+` mediumblob,
  `+testDelimiter("type_24")+` blob,
  `+testDelimiter("type_25")+` longblob,
  `+testDelimiter("type_26")+` binary(1) DEFAULT NULL,
  `+testDelimiter("type_27")+` varbinary(1) DEFAULT NULL,
  `+testDelimiter("type_28")+` enum('RED','GREEN','BLUE') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `+testDelimiter("type_29")+` set('RED','GREEN','BLUE') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `+testDelimiter("type_30")+` date DEFAULT NULL,
  `+testDelimiter("type_31")+` datetime DEFAULT NULL,
  `+testDelimiter("type_32")+` timestamp NULL DEFAULT NULL,
  `+testDelimiter("type_33")+` time DEFAULT NULL,
  `+testDelimiter("type_34")+` year(4) DEFAULT NULL,
  `+testDelimiter("type_35")+` geometry DEFAULT NULL,
  `+testDelimiter("type_36")+` point DEFAULT NULL,
  `+testDelimiter("type_39")+` multilinestring DEFAULT NULL,
  `+testDelimiter("type_41")+` multipolygon DEFAULT NULL,
  `+testDelimiter("type_37")+` linestring DEFAULT NULL,
  `+testDelimiter("type_38")+` polygon DEFAULT NULL,
  `+testDelimiter("type_40")+` multipoint DEFAULT NULL,
  `+testDelimiter("type_42")+` geometrycollection DEFAULT NULL,
  `+testDelimiter("type_50")+` double(5,2) DEFAULT NULL,
  `+testDelimiter("type_51")+` json DEFAULT NULL,
  PRIMARY KEY (`+"`"+`id`+"`"+`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`, typeTesttableName))

	assert.Equal(t, err, nil)

	_, err = conn.Exec(`INSERT INTO ` + testDelimiter(typeTesttableName) + ` (id, type_1, type_2, type_3, type_4, type_5, type_6, type_7, type_8, type_9, type_10, type_11, type_12, type_13, type_14, type_15, type_16, type_17, type_18, type_19, type_20, type_21, type_22, type_23, type_24, type_25, type_26, type_27, type_28, type_29, type_30, type_31, type_32, type_33, type_34, type_35, type_36, type_39, type_41, type_37, type_38, type_40, type_42, type_50, type_51)
VALUES
	(1, 1, 1, 1, 1, 1, 1.000, 1, 1.000, 1, 0, 1, 1, 1.00000, 1, 1, '1', '1', '1', '1', '1', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2001', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);`)

	assert.Equal(t, err, nil)

	typeField := "Type"
	fieldField := "Field"

	conn = testConnDSN(DriverMysql, fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/%s?charset=utf8mb4", typeTestdbName))

	config.Initialize(&config.Config{
		SqlLog: true,
	})

	columnsModel, _ := WithDriver(conn).Table(typeTesttableName).ShowColumns()
	item, err := WithDriver(conn).Table(typeTesttableName).First()

	for _, model := range columnsModel {
		fieldTypeName := strings.ToUpper(testGetType(model[typeField].(string)))
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

	conn := testConnDSN(DriverPostgresql, fmt.Sprintf(connStatement, typeTestdbName))
	fmt.Println("creating database")
	_, err := conn.Exec(fmt.Sprintf(`SELECT 'CREATE DATABASE %s' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '%s')`, typeTestdbName, typeTestdbName))
	assert.Equal(t, err, nil)
	fmt.Println("drop table")
	_, err = conn.Exec(fmt.Sprintf("DROP TABLE IF EXISTS `%s`;", typeTesttableName))
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
	_, err = conn.Exec(fmt.Sprintf(typeTestpostgresCreateSql, typeTesttableName))

	assert.Equal(t, err, nil)

	fmt.Println("insert data")
	_, err = conn.Exec(`INSERT INTO public.` + typeTesttableName + `(
	id, type_1, type_2, type_3, type_4, type_5, type_6, type_7, type_8, type_9, type_10, type_11, type_12, type_13, type_14, type_15, type_16, type_17, type_18, type_19, type_20, type_21, type_22, type_23, type_24, type_25, type_26, type_27, type_28)
	VALUES (1, 1, 1, 0.3, 1, 1, 1, 1, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'n');`)

	assert.Equal(t, err, nil)

	typeField := "udt_name"
	fieldField := "column_name"

	conn = testConnDSN(DriverPostgresql, fmt.Sprintf(connStatement, typeTestdbName))

	config.Initialize(&config.Config{
		SqlLog: true,
	})

	columnsModel, _ := WithDriver(conn).Table(typeTesttableName).ShowColumns()
	item, err := WithDriver(conn).Table(typeTesttableName).First()

	for _, model := range columnsModel {
		fieldTypeName := strings.ToUpper(testGetType(model[typeField].(string)))
		fmt.Println("fieldTypeName", fieldTypeName)
		GetDTAndCheck(fieldTypeName)
		fmt.Println(model[fieldField].(string), GetValueFromSQLOfDatabaseType(DatabaseType(fieldTypeName), item[model[fieldField].(string)]))
	}

	assert.Equal(t, err, nil)
}

// *******************************
// test helper methods
// *******************************

func testGetType(typeName string) string {
	r, _ := regexp.Compile(`\((.*?)\)`)
	typeName = r.ReplaceAllString(typeName, "")
	return strings.ToLower(strings.ReplaceAll(typeName, " unsigned", ""))
}

func testConnDSN(driver, dsn string) Connection {
	return GetConnectionByDriver(driver).InitDB(map[string]config.Database{
		"default": {Dsn: dsn},
	})
}

func testConn(driver string, cfg config.Database) Connection {
	cfg.Driver = driver
	cfg.MaxIdleCon = 10
	cfg.MaxOpenCon = 80
	return GetConnectionByDriver(driver).InitDB(map[string]config.Database{
		"default": cfg,
	})
}

func testDelimiter(s string) string {
	return "`" + s + "`"
}

func testCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
