package dependencies

import (
	"fmt"
	"log"
	"os/exec"
)

// Create a map with key path and value exec.Cmd
var RunningExes = make(map[string]*exec.Cmd)

func RunningExe(path string) error {
	cmd := exec.Command(path)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting EXE:", err)
		return err
	}

	fmt.Println("EXE file executed successfully")

	// Add the path and cmd to the map
	RunningExes[path] = cmd

	// Log the map
	fmt.Println(RunningExes)

	return nil
}

func TerminateExe(path string) error {
	// Get the cmd from the map
	cmd := RunningExes[path]
	log.Printf("Terminating %s", cmd)

	// Terminate the Game
	Pid := cmd.Process.Pid

	exec.Command("taskkill", "/F", "/T", "/PID", fmt.Sprint(Pid)).Run()

	// Remove the path and cmd from the map
	delete(RunningExes, path)

	// Log the map
	fmt.Println(RunningExes)

	return nil
}
