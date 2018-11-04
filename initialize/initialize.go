package initialize

func Initialize() {
	// do all of the setup for first time running the application
	createDB()
	createConfig()
	promptUserInfo()
}

func createDB() {
	// create tables etc in content/data/bla.db
	// with SQLite
	// NOTE: call to storage package Initialize() method
}

func createConfig() {
	// creates the config file for basic config info
	// in content/data/config.json
}

func promptUserInfo() {
	// get user info like owner username/password
}
