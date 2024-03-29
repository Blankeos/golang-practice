### 🚀 Getting Started

This getting started guide assumes you literally have no binaries on your computer besides Go.

You need the following to run this project:

- [SQLC](https://docs.sqlc.dev/en/stable/overview/install.html) (Static SQL Builder CLI)
  ```sh
  go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
  ```
- [Goose](https://github.com/pressly/goose) (SQL Migrations CLI)
  ```sh
  go install https://github.com/pressly/goose
  ```
- For Windows Only: [MSYS2 (GCC and Mingw32Make)](https://www.msys2.org/) - (You need `gcc` binary to run sqlite3 on Windows. Follow the instructions on the link to install `gcc` This seems to be on GNU which isn't on Windows by default, not sure with MACOS and Linux though since they seem to have GNU by default.). Here is VSCode's instructions to [install MingW](https://code.visualstudio.com/docs/cpp/config-mingw)

  ```sh
  # tl;dr of the instructions:
  # 1. Install MSYS2. Then run the following on msys terminal UCRT64 environemnt:
  pacman -S mingw-w64-ucrt-x86_64-gcc
  # confirm you installed it with:
  gcc --version
  # Install the tool chain, Press Enter with (default=all), and then Y to install.
  pacman -S --needed base-devel mingw-w64-ucrt-x86_64-toolchain
  # Go to environment variables, add this to Path:
  C:\msys64\ucrt64\bin

  # This should also have installed mingw32-make.exe (go to the C:\msys64\ucrt64\bin to find this )
  # Then, just rename `mingw32-make.exe` to `make.exe`, so we can easily use it on Windows as well.
  ```

#### 1. Install Needed Binaries/Modules

```sh
# Install dependencies
go mod tidy

# Install SQLC (Static SQL Builder CLI)
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Install Goose (SQL Migrations CLI)
go install github.com/pressly/goose/v3/cmd/goose@latest
```

#### 2. Copy Environment Variables

```sh
cp .env.example .env
```

#### 3. Migrate & Seed

```sh
# Migrate (Establish Columns in DB)
make migrate # or sh scripts/migrate
```

#### 4. Run the app

```sh
# This basically builds and runs go runs the project in `/bin/gql-server`
make run # or go build -o bin/gql-server, then go run bin/gql-server
```

### 💿 Working with the Database

### 👾 Cheatsheet of Commands (Notes)

This project was initialized with the command at the bottom:

Follow this [gqlgen getting started](https://gqlgen.com/getting-started/#building-the-server).

```sh
# creates a project skeleton of a todolist app
go run github.com/99designs/gqlgen init

# This generated the schema.graphqls, resolver.go, generated.go, model folder, schema.resolvers.go (and everything inside the `graph` folder)
go run github.com/99designs/gqlgen generate
```
