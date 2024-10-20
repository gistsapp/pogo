DROP TABLE IF EXISTS users;
CREATE TABLE users ( 
  user_id SERIAL,
  name TEXT
);

INSERT INTO users(name) values('test1');
INSERT INTO users(name) values('test2');
INSERT INTO users(name) values('test3');

