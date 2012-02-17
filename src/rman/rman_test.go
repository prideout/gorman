package rman_test

import (
    "bufio"
    "fmt"
    ri "rman"
    "os"
    "os/exec"
    "strings"
    "testing"
)

func TestTorus(t *testing.T) {

    compileShader(OCCLUSION_SHADER)
    compileShader(VIGNETTE_SHADER)

    launch := "launch:prman? -t -ctrl $ctrlin $ctrlout -capture debug.rib"
    ri.Begin(launch)
    ri.Format(512, 320, 1)
    ri.Display("grasshopper", "framebuffer", "rgba")
    ri.ShadingRate(4)
    ri.Option("limits", "int threads", 2)
    ri.Option("statistics",
        "xmlfilename", "stats.xml",
        "endofframe", true)
    ri.PixelSamples(4, 4)
    ri.Projection("perspective", "fov", 30.)
    ri.Translate(0, 0.25, 4)
    ri.Rotate(-20, 1, 0, 0)
    ri.Rotate(180, 1, 0, 0)
    ri.Imager("Vignette")
    ri.WorldBegin()
    ri.Declare("samples", "float")
    ri.Declare("em", "color")
    ri.Attribute("cull",
        "int backfacing", false,
        "int hidden", false)
    ri.Attribute("visibility",
        "int diffuse", true,
        "int specular", true)
    ri.Attribute("dice",
        "int rasterorient", false)

    // Floor
    ri.Attribute("identifier", "string name", "Floor")
    ri.Surface("Occlusion",
        "em", color(0, 0.65, 0.83),
        "samples", 64.)
    ri.TransformBegin()
    ri.Rotate(90, 1, 0, 0)
    ri.Disk(-0.7, 300, 360)
    ri.TransformEnd()

    // Sculpture
    ri.Attribute("identifier", "string name", "Sculpture")
    ri.Surface("Occlusion",
        "em", gray(1.1),
        "samples", 64.)
    ri.TransformBegin()
    ri.Translate(0, 0.75, 0)
    ri.Torus(1, 0.2, 0, 360, 360)
    ri.TransformEnd()
    ri.WorldEnd()

    ri.End()
}

func compileShader(content string) {
    cmd := exec.Command("shader")
    b := bufio.NewWriter(os.Stdout)
    cmd.Stdout = b
    cmd.Stderr = b
    cmd.Stdin = strings.NewReader(content)
    if nil != cmd.Run() {
        fmt.Print(ANSI_RED)
        b.Flush()
        fmt.Print(ANSI_RESET)
        os.Exit(1)
    }
}

func color(r, g, b float32) [3]float32 {
    return [3]float32{r, b, g}
}

func gray(x float32) [3]float32 {
    return [3]float32{x, x, x}
}

const (
    ANSI_BLACK   string = "\x1b[1;30m"
    ANSI_RED     string = "\x1b[1;31m"
    ANSI_GREEN   string = "\x1b[1;32m"
    ANSI_YELLOW  string = "\x1b[1;33m"
    ANSI_BLUE    string = "\x1b[1;34m"
    ANSI_MAGENTA string = "\x1b[1;35m"
    ANSI_CYAN    string = "\x1b[1;36m"
    ANSI_WHITE   string = "\x1b[1;37m"
    ANSI_RESET   string = "\x1b[0m"
)

const OCCLUSION_SHADER string = `
class Occlusion(float darkness = 1;
                color em = (1,0,1);
                float samples = 128;
                float maxdist = 100;
)
{
  public void surface(output color Ci, Oi)
  {
    normal Nf = normalize(faceforward(N, I));
    normal Nn = normalize(Nf);
    float occ = darkness * occlusion(P, Nn, samples, "maxdist", maxdist);
    Ci = (1 - occ) * Cs * em * Os;
    Oi = Os;
  }
}
`
const VIGNETTE_SHADER string = `
class Vignette(color background = (0,0,0))
{
   public void imager(
          output varying color Ci;
          output varying color Oi;)
   {
        Ci = Ci + (1-Oi)*background;

        float x = s - 0.5;
        float y = t - 0.5;
        float d = x * x + y * y;
        Ci *= 1.0 - d;

        float n = (cellnoise(s*1000, t*1000) - 0.5);
        Ci += n * 0.025;
   }
}
`