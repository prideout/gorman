/*

    RenderMan bindings for golang.  Makes extensive use of go's nice support for variadic functions.  For example:

    import ri "rman"
    void main() {
        ri.Begin("")
        // Draw a plastic teapot?
        ri.End()
    }

*/
package rman

import "C"
