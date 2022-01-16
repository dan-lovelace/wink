package configs

type dbConn struct {
	Driver   string
	Location string
}

var DBConn = dbConn{
	Driver:   "sqlite3",
	Location: "./test.db",
}
