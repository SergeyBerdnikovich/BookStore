-- +swan Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE books
(
  id         SERIAL NOT NULL,
  name       TEXT,
  quantity   INTEGER,
  created_at TIMESTAMP WITH TIME ZONE,
  updated_at TIMESTAMP WITH TIME ZONE
);
-- +swan Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE books;
