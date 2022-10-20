CREATE DATABASE IF NOT EXISTS mydb;
USE mydb;
CREATE TABLE IF NOT EXISTS USER (
    USER_ID  MEDIUMINT NOT NULL AUTO_INCREMENT
    , USERNAME VARCHAR(32)
    , PASSWORD_HASH VARCHAR(60)
    , POLICY BLOB
    , CREATED_DATE TIMESTAMP
    , UPDATED_DATE TIMESTAMP
    , PRIMARY KEY (USER_ID)
    ,  CONSTRAINT UC_Person UNIQUE (USERNAME)
);