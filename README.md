# Yako Chat

## Overview
Yako Chat is a lightweight, local chat server built with Go. It enables users on the same local network (LAN) to communicate in real time through a simple web-based interface. The project is designed for easy setup, privacy, and efficient LAN messaging without relying on external services.

## Features
- Local chat server for LAN communication.
- Web-based frontend served via static files.
- Simple and fast Go application.
- Modular and maintainable codebase.

## Technologies Used
- **Go**: Backend development and chat server implementation.
- **HTML**: Web-based frontend.
- **Zsh**: Default shell for running commands.

## Prerequisites
Before you begin, ensure you have the following installed:
- [Go](https://golang.org/dl/) (version 1.20 or later)
- A terminal with `zsh` shell (default for Linux users).

## Installation
1. Clone the repository:
    ```zsh
    git clone https://github.com/yourusername/yako-chat.git
    ```
2. Navigate to the project directory:
    ```zsh
    cd yako-chat
    ```
3. Install dependencies:
    ```zsh
    go mod tidy
    ```

## Usage
1. Run the application:
    ```zsh
    go run main.go
    ```
2. Open your browser and navigate to `http://localhost:8080` (or your server's LAN IP) to access the chat interface.

## Project Structure
```
├── go.mod
├── go.sum
├── yako-chat
├── main.go
├── static
│   └── index.html
```
- `main.go`: Entry point of the application.
- `static/`: Contains static files for the web-based frontend.
- `go.mod` and `go.sum`: Dependency management files.

## Getting Started
To get started with Yako Chat, follow the installation steps above. Once the server is running, users on the same LAN can join the chat by accessing the server's IP address in their browsers.

## Future Plans
- Add support for chat rooms and private messaging.
- Implement user authentication for secure access.
- Enhance the frontend with a modern UI framework.

## Contributing
Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a new branch:
    ```zsh
    git checkout -b feature-name
    ```
3. Make your changes and commit them:
    ```zsh
    git commit -m "Add feature-name"
    ```
4. Push to your branch:
    ```zsh
    git push origin feature-name
    ```
5. Create a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact
For questions or feedback, please contact [Kirebyte](mailto:me@kirebyte.dev).
