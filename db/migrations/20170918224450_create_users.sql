
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE USERS (
	ID INT PRIMARY KEY NOT NULL,
	NAME TEXT NOT NULL,
	AGE INT,
	ADDRESS CHAR(50),
	CITY CHAR(40),
	STATE CHAR(40),
	COUNTRY CHAR(40),
	ROLE INT,
	USERID INT, 
	PHOTO VARBINARY,
	CREATEDBYUSERID INT,
	MODIFIEDDATE DATETIME,
	CREATEDDATE DATETIME,
	ACTIVE BIT
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE USERS;
