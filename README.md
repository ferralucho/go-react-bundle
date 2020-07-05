# go-react-bundle
go with react in one bundle

Serve the static react code inside of a Go server and produce a single binary that would be placed 
inside a docker container and can be run at any point to serve the site.

The basic idea here is that using the rice box package to find the react build folder and generates a single Go source file called rice-box.go. 
The generated go file contains all assets. The Go tool compiles this into the binary and allow to serve the web application as a single binary.

TO RUN THE PROJECT: ./go-react-bundle in go-react-bundle/cmd/go-react-bundle

HOW I CREATED THE PROJECT:

npx create-react-app .

-Navigate to folder in the terminal so you can just run yarn build which will produce a static bundle.

https://github.com/GeertJohan/go.rice
-You should now navigate to your cmd/projectname in our terminal and run our go.rice build command rice embed-go which will find our asset references and compile them so that they can be bundled alongside our final binary. This will create a rice-box.go file in the same folder.
rice embed-go
go build

-go build . to create a binary in the current location. It should create a binary with the project name which you can run by typing ./projectname in your terminal which should serve the application.