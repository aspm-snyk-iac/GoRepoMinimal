package main

import "fmt"

// Hardcoded secrets (for testing Gitleaks detection)

const awsSecretKey = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"

// GitHub token pattern
const githubToken = "ghp_1234567890abcdefghijklmnopqrstuv"

// Generic API key
const apiKey = "api_key=12345-ABCDE-SECRET-67890"

func main() {
	fmt.Println("Demo app with hardcoded secrets")
	fmt.Println(awsAccessKey, awsSecretKey, githubToken, apiKey)
}
