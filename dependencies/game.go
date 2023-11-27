package dependencies

import (
	"fmt"
	"log"
	"os/exec"
)

type SocketSignalling struct {
	// The socket that will be used to send messages to the client
	IPAddress string
	Port      uint32
}

// For running with arguments
var SocketMappings = map[string]SocketSignalling{
	"wss://s3bangka.nusantara.digital/": {
		IPAddress: "10.10.64.129",
		Port:      2085,
	},
	"wss://s4bangka.nusantara.digital/": {
		IPAddress: "10.10.64.129",
		Port:      2081,
	},
	"wss://signalling1.nusantara.digital/": {
		IPAddress: "10.11.60.213",
		Port:      2073,
	},
	"wss://signalling2.nusantara.digital/": {
		IPAddress: "10.11.60.213",
		Port:      2077,
	},
	"wss://signalling3.nusantara.digital/": {
		IPAddress: "10.11.60.213",
		Port:      2081,
	},
	"wss://signalling4.nusantara.digital/": {
		IPAddress: "10.11.60.213",
		Port:      2085,
	},
	"wss://signalling5.nusantara.digital/": {
		IPAddress: "10.11.60.213",
		Port:      2089,
	},
	"wss://signalling6.nusantara.digital/": {
		IPAddress: "10.11.60.213",
		Port:      2093,
	},
	"wss://signallingdev1.nusantara.digital/": {
		IPAddress: "10.11.60.213",
		Port:      61,
	},
	"wss://signallingdev2.nusantara.digital/": {
		IPAddress: "10.11.60.213",
		Port:      63,
	},
	"wss://signallingdev3.nusantara.digital/": {
		IPAddress: "10.11.60.213",
		Port:      65,
	},
}

// Create a map with key path and value exec.Cmd
var RunningExes = make(map[string]*exec.Cmd)

func RunningExe(path string, socketEndpoint string) error {
	chosenPath := SocketMappings[socketEndpoint]
	// Running exe with arguments -AudioMixer -PixelStreamingIP=10.11.60.213 -PixelStreamingPort=63 -PixelStreamingEncoder=H.264 -PixelStreamingEncoderMinQP=1 -PixelStreamingEncoderMaxQP=51 -RenderOffscreen -log
	cmd := exec.Command(path, fmt.Sprintf("-PixelStreamingIP=%s", chosenPath.IPAddress), fmt.Sprintf("-PixelStreamingPort=%d", chosenPath.Port), "-PixelStreamingEncoder=H.264", "-PixelStreamingEncoderMinQP=1", "-PixelStreamingEncoderMaxQP=51", "-RenderOffscreen", "-log")
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
