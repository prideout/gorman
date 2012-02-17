Usage
-----

RenderMan bindings for golang.  Makes extensive use of go's nice support for variadic functions.  For example:

    import ri "rman"
    void main() {
        ri.Begin("")
        // Draw a plastic teapot?
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
