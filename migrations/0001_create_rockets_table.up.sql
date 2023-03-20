CREATE TABLE IF NOT EXISTS rockets(
  id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
  rkt_type varchar(50),
  rkt_name varchar(50),
  flights int
);