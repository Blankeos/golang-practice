### Building a Go server with Gin

Writing as of 9/9/2023, I really like it so far. I configured toml and created my own GET handler for album. Extra stuff I learned here:

1. **Air** like nodemon for go. I like it. (Guide in [Air docs](https://github.com/cosmtrek/air)):

   ```sh
       go install github.com/cosmtrek/air@latest # install it.
       air -c .air.toml # Tell air to use it.
       air init # intiializes a .air.toml
       air # run air based ont he .air.toml
   ```

2. If you get in trouble with air just (on windows):

   ```sh
       tasklist | findstr "8080" # find the port
       tasklist | findstr web-service-gin.exe # partial text search for processes

       # kill processes
       taskkill /F /PID 6146 # 6146 is the PID
   ```
