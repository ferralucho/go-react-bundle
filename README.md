# go-react-bundle
go with react in one bundle

Serve the static react code inside of a Go server and produce a single binary that would be placed 
inside a docker container and can be run at any point to serve the site.

The basic idea here is that using the rice box package to find the react build folder and generates a single Go source file called rice-box.go. 
The generated go file contains all assets. The Go tool compiles this into the binary and allows us to serve our web application as a single binary.

npx create-react-app .
