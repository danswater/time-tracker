
-- Read https://stackoverflow.com/questions/21757722/how-to-use-sqlite-decimal-precision-notation
-- for an interesting explanation of SQLites datatypes, in particular affinities.

CREATE TABLE IF NOT EXISTS users (
	Id          INTEGER PRIMARY KEY AUTOINCREMENT,
	Username    TEXT
);

CREATE TABLE IF NOT EXISTS transactions (
	Id               INTEGER PRIMARY KEY,
	UID              INTEGER,
	Active           INTEGER,
	StartTimestamp   DATETIME DEFAULT CURRENT_TIMESTAMP,
	EndTimestamp     DATETIME DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY(UID) REFERENCES users(Id));

-- Add default user.
INSERT INTO users (`Username`) VALUES ('user');
