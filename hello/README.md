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

### 6. go test and go test -v

built in unit with go testing by creating files ending with `_test.go`.

### 7. go list -f '{{.Target}}'

gets the binary installation directory. This is where all the binaries from `go install` will go.

In order to make ues of this, you can also set the path of the binary installation directory with the following command:

```sh
# mac
$ export PATH=$PATH:/path/to/your/install/directory

# windows
$ set PATH=%PATH%;C:\path\to\your\install\directory
```

example usage:

```sh
$ go list -f '{{.Target}}'
# 'C:\Users\CARLOBOT\go\bin\hello.exe'

# -- set the path --
$ set PATH=%PATH%;C:\Users\CARLOBOT\go\bin
```

### 8. go install

This essentially copies your `hello.exe` to the binary installation directory. After you setup your path, you can pretty much just do:

```sh
$ hello
# <output of your hello.exe here>
```

How to uninstall? You can just **delete** the binary in its installation directory. Theres no magic involved here since it's just in your %PATH% anyway.
