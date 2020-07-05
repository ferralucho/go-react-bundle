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
-You should now navigate to your cmd/projectname in the terminal and run the go.rice build command rice embed-go which will find the asset references and compile them so that they can be bundled alongside the final binary. This will create a rice-box.go file in the same folder.
rice embed-go
go build

-go build . to create a binary in the current location. It should create a binary with the project name which you can run by typing ./projectname in your terminal which should serve the application.

DOCKER IMAGE

-The first step is to spin up a node container that has access to npm and yarn, copy over the react code including the package.json and yarn.lock file so that we persist package versions, we run a yarn to pull all the packages and finally run a yarn build to build a static version of the site

-The second step spins up a go server and copies all the local code to the equivalent path on the go server. It then copies the frontend build folder in to the /web/frontend/ folder ready for us to compile it. We then change the WORKDIR to the cmd/golang-react folder and run the rice embed-go and GOOS=linux GOARCH=amd64 go build -o golang-react.linux.x86_64 . to create a linux binary.

-The final step creates a very simple alpine server. We copy the binary to the new container and set the entry path.

run docker-compose up --build to build and spin up the new container. 