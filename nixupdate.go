package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func update() {
	fmt.Println("Updating channel...")
	updatechanel := exec.Command("sudo", "nix-channel", "--update")

	fmt.Println("Updating packages...")
	updatepackages := exec.Command("sudo", "nixos-rebuild", "switch")

	updatechanel.Run()
	updatepackages.Run()
}

func collect() {
	fmt.Println("Collecting garbage...")
	garbagecollect := exec.Command("sudo", "nix-collect-garbage", "-d")
	garbagecollect.Run()
}

func main() {
	var userinput string

	fmt.Print("Do you want to collect garbage? (Y/N/Exit): ")
	fmt.Scanln(&userinput)

	switch strings.ToLower(userinput) {
	case "y", "yes":
		update()
		collect()
	case "n", "no":
		update()
	case "exit":
		return

	default:
		fmt.Println("Invalid input. Please enter Y, N, or Exit.")
	}
}
