/*

    RenderMan bindings for golang.  For example:

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

*/
package rman

import "C"
