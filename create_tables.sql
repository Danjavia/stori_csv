CREATE TABLE users (
  id VARCHAR(36) NOT NULL DEFAULT uuid(),
  email VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE summary (
  id INT NOT NULL AUTO_INCREMENT,
  user_id VARCHAR(36) NOT NULL,
  summary JSON,
  PRIMARY KEY (id),
  CONSTRAINT fk_summary_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);