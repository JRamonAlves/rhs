package infraestructure

import "os"

func GetServiceFilePath() string {
	return os.Getenv("SERVICE_PATH")
}
