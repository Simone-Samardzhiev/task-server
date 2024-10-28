package main

import "task-server/config"

func main() {
	config.LoadEnvironmentFiles("../../config/.env")
}
