# Go-dev
practicing & learning go 
https://go.dev/doc/tutorial/getting-started

## Learnings
* [go mod and sum](https://golangbyexample.com/go-mod-sum-module/) 
* go mod init: initializes and writes a new go.mod file in the current directory

* go mod tidy: ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.

* modules -> lib dependencies, go.mod enables tracking of such dependencies in your module 

* module authentication: making sure downloaded direct and indirect dependencies are correct, go.sum checks the cryptographic hash with checksum -> data integrity

* [Exported Names](https://go.dev/tour/basics/3): In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name.

* := Operator: In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type). 

    * Taking the long way, you might have written this as:
        ```
        var message string
        message = fmt.Sprintf("Hi, %v. Welcome!", name)
        ```

    * Shorter way: 
      ```
        message := fmt.Sprintf("hi, %v. Welcome!", name)
      ```