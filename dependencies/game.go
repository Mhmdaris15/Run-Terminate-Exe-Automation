package dependencies

import (
	"fmt"
	"os/exec"
)

// F:\NUC_Aris\Nusantara-2\Compile\PixelStreaming\1.9.8-prod\Pixel Jelajah 1.9.8

// func main() {
//     # Specify the path to the EXE file you want to execute
//     exePath := "C:\\path\\to\\example.exe"

//     # Create an instance of the Cmd struct
//     cmd := exec.Command(exePath)

//     # Start the EXE file
//     err := cmd.Start()
//     if err != nil {
//         fmt.Println("Error starting EXE:", err)
//         return
//     }

//     # Wait for the EXE file to finish executing
//     err = cmd.Wait()
//     if err != nil {
//         fmt.Println("Error waiting for EXE:", err)
//         return
//     }

//     fmt.Println("EXE file executed successfully")

//     # Terminate the EXE file
//     err = cmd.Process.Kill()
//     if err != nil {
//         fmt.Println("Error terminating EXE:", err)
//         return
//     }

//     fmt.Println("EXE file terminated successfully")
// }

// Create a map with key path and value exec.Cmd
var RunningExes = make(map[string]*exec.Cmd)

func RunningExe(path string) error {
	cmd := exec.Command(path)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting EXE:", err)
		return err
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for EXE:", err)
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

	// Terminate the cmd
	err := cmd.Process.Kill()
	if err != nil {
		fmt.Println("Error terminating EXE:", err)
		return err
	}

	fmt.Println("EXE file terminated successfully")

	// Remove the path and cmd from the map
	delete(RunningExes, path)

	// Log the map
	fmt.Println(RunningExes)

	return nil
}
