/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"delivery/cmd"
	_ "delivery/internal/repository"
)

func main() {
	cmd.Execute()
}
