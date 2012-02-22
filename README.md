About
-----

RenderMan bindings for golang.  Work in progress.

Example
-------

    package main
    import ri "rman"
    func main() {
        ri.Begin("launch:prman? -t")
        ri.Display("sphere.tif", "file", "rgba")
        ri.Projection("perspective", "fov", 30.)
        ri.WorldBegin()
        ri.Color(0.9, 0.3, 0.1)
        ri.Translate(0, 0, 5)
        ri.Sphere(1.0, -1.0, 1.0, 360.0)
        ri.WorldEnd()
        ri.End()
    }

Take a look at rman_test.go for a bigger example.  Enjoy!

Installation
------------

First, clone the repo:

    git clone git@github.com:prideout/gorman

Next look at the two #cgo lines at the top of `src/rman/rman.go`, and change them if need be.

Next:

    cd gorman
    export GOPATH=$GOPATH:$PWD
    go install ./src/rman
    go test ./src/rman
    go build sphere.go; ./sphere
