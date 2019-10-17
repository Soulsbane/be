# Description
Dead simple application launcher using Lua for scripting.

# Building
This program is built like any other go program using go build.
The only caveat is that the go version must be 1.13 or above.

# Creating a Command
* Make a new file called myscripts.lua in your config directory. The config directories location can be found here: https://golang.org/pkg/os/#UserConfigDir
* Open myscripts.lua and add

  ```lua
  Commands:addCommand("godoc", "godoc -http=:6060")
  ```
   Now build launch using this command
   ```
   go build -o launch
   ```
* You should be able to launch the command you created using
  ```
	./launch godoc
  ```
* Open your browser and navigate to http://localhost:6060/
* You should now be able to browse the Golang documentation locally.
