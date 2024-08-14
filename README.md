# Sober Roadmap (Working Title)

http://jonesrussell.github.io/sober

Sober Roadmap is intended to be a starting point for alcoholics, drug addicts, and family members of both. It provides an easy-to-follow roadmap to a better understanding of addiction, introduction to meetings, counseling, and more.

## Installation

1. Clone the repo:
    ```sh
    git clone https://github.com/jonesrussell/sober.git
    ```
2. Navigate to the project directory:
    ```sh
    cd sober
    ```
3. Install dependencies:
    ```sh
    task default
    ```

## Usage

### Running the Echo Server

You can run the Echo server using the `task` command or directly with `go run`:

1. Using `task`:
    ```sh
    task run
    ```

2. Using `go run`:
    ```sh
    go run ./main.go
    ```

### Generating Static HTML

You can generate static HTML using the `task` command or directly with `go run`:

1. Using `task`:
    ```sh
    task static
    ```

2. Using `go run`:
    ```sh
    go run ./main.go -generate
    ```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.