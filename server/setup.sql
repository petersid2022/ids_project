DROP DATABASE IF EXISTS Articles;
CREATE DATABASE IF NOT EXISTS Articles;
USE Articles;

CREATE TABLE IF NOT EXISTS Posts (
  PostID INT UNSIGNED NOT NULL AUTO_INCREMENT,
  Title varchar(255) NOT NULL DEFAULT '',
  Content text NOT NULL,
  CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (PostID)
);

INSERT INTO Posts (Title, Content) VALUES ('must read', 'this is a super important article');

SELECT * FROM Posts;
