## Installation

To install Repo Scraper, clone the repository and install the dependencies:

```bash
git clone https://github.com/GadSAR/repo-scraper.git
cd repo-scraper
go mod tidy
```

Note: dont forget to use .env.example to create .env file

## Usage

To use Repo Scraper, run the following command:

```bash
go run main.go
```

To develop with hotReload, run the following command:

```bash
air
``` 

To build and run, run the following commands:

```bash
go build .
./repo-scraper.exe
```