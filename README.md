# Go-dev
practicing go 

## How Go works 
go mod init: initializes and writes a new go.mod file in the current directory
go mod tidy: ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.

modules -> lib dependencies, go.mod enables tracking of such dependencies in your module 
module authentication: making sure downloaded direct and indirect dependencies are correct, go.sum checks the cryptographic hask, checksum -> data integrity

