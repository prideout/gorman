package main

import (
    ri "rman"
)

func main() {
    launch := "launch:prman? -t -ctrl $ctrlin $ctrlout"
    ri.Begin(launch)
    ri.Format(512, 320, 1)
    ri.Display("grasshopper", "framebuffer", "rgba")
    ri.ShadingRate(4)
    ri.Option("limits", "int threads", 2)
    ri.PixelSamples(4, 4)
    ri.Projection("perspective", "fov", 30.)
    ri.WorldBegin()
	ri.Color(0.2, 0.3, 0.9)
	ri.Translate(0, 0, 5)
    ri.Sphere(1.0, -1.0, 1.0, 360.0)
    ri.WorldEnd()
    ri.End()
}
