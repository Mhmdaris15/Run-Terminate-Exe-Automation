# Run-Terminate-Exe-Automation
🚵 Program for Automating Process Running and Terminating Exe Program Remotely with WebSocket Architecture built with GoLang 🏗️

## Overview
This program allows you to run and terminate executable programs remotely using a WebSocket architecture built with GoLang. It includes a server and a client component, allowing you to control the execution of processes on a remote machine from a web interface.

## Features
* Run executables remotely using a WebSocket connection
* Terminate running processes remotely
* Supports both Windows and Linux operating systems
* Written in GoLang for efficiency and performance

## Getting Started
1. Clone this repository to your local machine: `git clone https://github.com/your-username/remote-process-control.git`
2. Build the project by navigating to the root directory and running `go build -o main .` (assuming you have Go installed)
3. Run the server component by executing `./main -server`
4. Open a web browser and navigate to `http://localhost:3000` to access the web interface
5. Use the web interface to upload an executable file and start the process on the remote machine
6. Once the process is running, you can use the web interface to terminate it

## Usage
### Server Component
The server component listens for incoming connections on port 8080 by default. You can change the port number by passing a flag when starting the server: `./main -server -port 8081`

### Client Component
The client component connects to the server using a WebSocket connection. You can specify the IP address and port number of the server when starting the client: ./main -client <ip_address>:<port>

For example, if the server is running on http://localhost:8080, you can start the client like this: ./main -client localhost:8080
