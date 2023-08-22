### Beginner Notes:

#### 1. `go mod init example/hello`

kinda creates your module (like npm init)

#### 2. `go run .`

runs your project

#### 3. `go get <package_url>`

- example: `go get rsc.io/quote/v4`

- Their package manager: https://pkg.go.dev/

#### 4. `go mod tidy`

- if you have your package_url as import "<package_url>", it's possible to do go mod tidy to automatically get it.

### 5. go mod edit -replace example.com/greetings=../greetings

replaces the current example.com/greetings module (which is currently not in a repository) to a local path.
