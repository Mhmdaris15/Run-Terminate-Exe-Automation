# Run-Terminate-Exe-Automation

🚵 Program for Automating Process Running and Terminating Exe Program Remotely with WebSocket Architecture built with GoLang 🏗️

## Overview

This program allows you to run and terminate executable programs remotely using a WebSocket architecture built with GoLang. It includes a server and a client component, allowing you to control the execution of processes on a remote machine from a web interface.

## Features

- Run executables remotely using a WebSocket connection
- Terminate running processes remotely
- Supports both Windows and Linux operating systems
- Written in GoLang for efficiency and performance

## Getting Started

1. Clone this repository to your local machine: `git clone https://github.com/Mhmdaris15/Run-Terminate-Exe-Automation/`
2. Build the project by navigating to the root directory and running `go build -o main .` (assuming you have Go installed)
3. Run the server component by executing `./main -server`
4. Open a web browser and navigate to `http://localhost:3000` to access the web interface
5. Use the web interface to upload an executable file and start the process on the remote machine
6. Once the process is running, you can use the web interface to terminate it

## Usage

### Server Component

The server component listens for incoming connections on port 3000 by default. You can change the port number by passing a flag when starting the server: `./main -server -port 3001`

### Client Component

The client component connects to the server using a WebSocket connection. You can specify the IP address and port number of the server when starting the client: `./main -client <ip_address>:<port>`

For example, if the server is running on `http://localhost:3000`, you can start the client like this: ./main -client localhost:3000

## Running Executables

To run an executable on the remote machine, follow these steps:

1. Upload the executable file using the web interface
2. Enter the command line arguments and environment variables (if any)
3. Click the "Start" button to initiate the process

Once the process is running, you can monitor its progress and terminate it using the web interface.

## Terminating Processes

To terminate a running process, find the corresponding process ID on the web interface and click the "Terminate" button next to it.

## Troubleshooting

If you encounter issues while running the program, check the logs in the server and client directories for error messages. You can also use tools like Postman or cURL to test the WebSocket API endpoints.

## Contributing

Pull requests are welcome! If you'd like to contribute to this project, please fork the repository, make changes, and submit a pull request with a clear description of your modifications.

## License

This project is licensed under the MIT License. See LICENSE.md for details.

## Acknowledgments

Thanks to the GoLang community for creating such a powerful and efficient programming language. Special thanks to the authors of the following libraries used in this project:

- GoWebSocket: A lightweight WebSocket library for GoLang
- GoExec: A simple library for executing shell commands in GoLang
