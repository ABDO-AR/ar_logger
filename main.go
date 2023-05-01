package main

import "ar_logger/logger"

// Main:
func main() {
	// Setup:
	logger.SetupLoggingSystem("./logs/logs.json") // Pass empty string will make it log into "./logs.json".
	// Logging:
	logger.E("Authentication API", "Failed to authenticate user(2131)!")
	logger.W("Authentication API", "Authentication is currently disabled, please turn it on!")
	logger.I("Authentication API", "10,000 Users was authenticated in this month.")
	logger.S("Authentication API", "User(2131) was authenticated successfully")
}
