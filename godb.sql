create database godb;
use godb;
CREATE TABLE `userinfo` (
        `uid` INT(10) NOT NULL AUTO_INCREMENT,
        `username` VARCHAR(64) NULL DEFAULT NULL,
        `departname` VARCHAR(64) NULL DEFAULT NULL,
        `created` DATE NULL DEFAULT NULL,
        PRIMARY KEY (`uid`)
    );

INSERT userinfo SET username="demoUser",departname="demoDept",created="2018-11-11";
