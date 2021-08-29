package main

import "github.com/denis-shcherbinin/gora-studio-test-task/internal/api"

const configPath = "configs/config"

func main() {
	api.Run(configPath)
}
