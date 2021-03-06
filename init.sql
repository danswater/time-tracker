
-- Read https://stackoverflow.com/questions/21757722/how-to-use-sqlite-decimal-precision-notation
-- for an interesting explanation of SQLites datatypes, in particular affinities.

CREATE TABLE IF NOT EXISTS intervals (
	Id          INTEGER PRIMARY KEY AUTOINCREMENT,
	StopwatchId INTEGER,
	StartTime   BIGINT,
	StopTime    BIGINT,

	FOREIGN KEY(StopwatchId) REFERENCES stopwatches(Id)
);

CREATE TABLE IF NOT EXISTS stopwatches (
	Id    INTEGER PRIMARY KEY,
	Color VARCHAR(255),
	Name  VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS pool_datas (
	Id           INTEGER PRIMARY KEY AUTOINCREMENT,
	CreationDate BIGINT,
	LastModDate  BIGINT,
	StopwatchId  INTEGER DEFAULT 0,

	FOREIGN KEY(StopwatchId) REFERENCES stopwatches(Id)
);

CREATE TABLE IF NOT EXISTS pools (
	Id              INTEGER PRIMARY KEY AUTOINCREMENT,
	EventName       VARCHAR(255),
	IsReadOnly      BOOLEAN,
	Message         TEXT,
	PoolDataId      INTEGER,
	PoolKey         VARCHAR(255),
	PoolKeyReadOnly VARCHAR(255),

	FOREIGN KEY(PoolDataId) REFERENCES pool_datas(Id)
);
