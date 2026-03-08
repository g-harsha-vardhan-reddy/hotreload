# HotReload CLI Tool

HotReload is a CLI tool that automatically rebuilds and restarts a Go server when code changes are detected.

## Features

- Watches project folders for file changes
- Automatically rebuilds the project
- Automatically restarts the server
- Streams server logs in real time

## Project Structure

cmd/hotreload → CLI entry point  
internal/watcher → File change watcher  
internal/builder → Build command executor  
internal/process → Server process manager  
testserver → Demo HTTP server

## Running the Tool

Run the hotreload tool:

go run ./cmd/hotreload --root ./testserver --build "go build -o server.exe ./testserver" --exec "server.exe"

## Demo

1. Start the hotreload tool
2. Open http://localhost:8080
3. Edit testserver/main.go
4. Save the file
5. The server automatically rebuilds and restarts

## Technologies Used

- Go
- fsnotify
- exec package for process management