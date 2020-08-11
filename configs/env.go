package configs

import (
	"os"
)

// SetENV ..
func SetENV() {
	/*
		set ENV
	*/
	os.Setenv("MONGO_HOST", "192.168.253.225")
	os.Setenv("MONGO_DBNAME", "demoGolang")
	os.Setenv("MONGO_USERNAME", "demoAdmin")
	os.Setenv("MONGO_PASSWORD", "demopwd01")
}
