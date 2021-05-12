CREATE DATABASE db1;
USE db1;

CREATE TABLE compromised (
  userid     INT(11) NOT NULL AUTO_INCREMENT,
  firstname  VARCHAR(255) NOT NULL,
  lastname   VARCHAR(255) NOT NULL,
  email      VARCHAR(255) NOT NULL,
  PRIMARY KEY (userid)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

INSERT INTO compromised (firstname, lastname, email) VALUES ('Alex',  'One',   'alex@one.com');
INSERT INTO compromised (firstname, lastname, email) VALUES ('Bob',   'Two',   'bob@two.com');
INSERT INTO compromised (firstname, lastname, email) VALUES ('Chris', 'Three', 'chris@three.com');
INSERT INTO compromised (firstname, lastname, email) VALUES ('Don',   'Four',  'don@four.com');
INSERT INTO compromised (firstname, lastname, email) VALUES ('Emily', 'Five',  'emily@five.com');


