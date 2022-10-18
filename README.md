# Go-dev
practicing & learning go 
https://go.dev/doc/tutorial/getting-started

## Learnings
* [go mod and sum](https://golangbyexample.com/go-mod-sum-module/) 
* go mod init: initializes and writes a new go.mod file in the current directory

* go mod tidy: ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.

* Declare a main package. In Go, code executed as an application must be in a main package. 

* init sets initial values for variables used in the function.
  
* modules -> lib dependencies, go.mod enables tracking of such dependencies in your module.

* module authentication: making sure downloaded direct and indirect dependencies are correct, go.sum checks the cryptographic hash with checksum -> data integrity.


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

* go mod edit -replace example.com/greetings=../greetings: this command creates a redirection to local directory. The command found the local code in the greetings directory, then added a require directive to specify that example.com/hello requires example.com/greetings. You created this dependency when you imported the greetings package in hello.go. To reference a published module, a go.mod file would typically omit the replace directive and use a require directive with a tagged version number at the end. To summarize, local -> replace (redirect), require, published -> just require. 
  
* Error handling and loggingg: use package "errors" and "log". Note: go automatically import the packages necessary in the code when saved.

* Ending a file's name with _test.go tells the go test command that this file contains test functions.

* T is a type passed to Test functions to manage test state and support formatted test logs.

* The go build command compiles the packages, along with their dependencies, but it doesn't install the results.
* The go install command compiles and installs the packages.
* The go run command is a useful shortcut for compiling and running a program when you're making frequent changes, it doesn't generate a binary executable.
* $ go list -f '{{.Target}}': You can discover the install path by running the go list command

## Questions 
* Difference between %v vs %w vs %s
* log formats
* error handling in depth
* backward compatibility
* slice vs array: A slice is like an array, except that its size changes dynamically as you add and remove items. The slice is one of Go's most useful types. The basic difference between a slice and an array is that a slice is a reference to a contiguous segment of an array. Unlike an array, which is a value-type, slice is a reference type.  
* built-in map keys, values? 
* built-in unit tests??? wow 
* project structure: core, modules, testing -> how to organize files effectively? Is there a convention 
* what is goroutine
* Difference in ` vs ' or ": 
  * ``` 
    a := "\n" // This is one character, a line break.
    b := `\n` // These are two characters, backslash followed by letter n.
    ```

