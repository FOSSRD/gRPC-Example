package main

import (
        "context"
        "fmt"
        "image"
        "image/color"
        "image/jpeg"
        "image/png"
        "log"
        "net"
        "bytes"

        pb "imager/imageconversionpb" // Replace with your actual Protobuf package path

        "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedImageConverterServer
}

func (s *server) BlackAndWhite(ctx context.Context, req *pb.BlackAndWhiteRequest) (*pb.Image, error) {
        img, err := decodeImage(req.GetImage())
        if err != nil {
                return nil, err
        }

        bounds := img.Bounds()
        bwImg := image.NewRGBA(bounds)

        threshold := req.GetThreshold() * 255.0

        for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
                for x := bounds.Min.X; x < bounds.Max.X; x++ {
                        r, g, b, _ := img.At(x, y).RGBA()
                        gray := float64(r+g+b) / (3 * 65535.0) * 255.0
                        var c color.RGBA
                        if gray > float64(threshold) {
                                c = color.RGBA{255, 255, 255, 255}
                        } else {
                                c = color.RGBA{0, 0, 0, 255}
                        }
                        bwImg.Set(x, y, c)
                }
        }

        return encodeImage(bwImg, req.GetImage().GetFormat())
}

func (s *server) Sepia(ctx context.Context, req *pb.SepiaRequest) (*pb.Image, error) {
        img, err := decodeImage(req.GetImage())
        if err != nil {
                return nil, err
        }
        bounds := img.Bounds()
        sepiaImg := image.NewRGBA(bounds)
        intensity := req.GetIntensity()

        for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
                for x := bounds.Min.X; x < bounds.Max.X; x++ {
                        r, g, b, a := img.At(x, y).RGBA()
                        fr, fg, fb, fa := float64(r)/65535.0, float64(g)/65535.0, float64(b)/65535.0, float64(a)/65535.0

                        sr := (fr*0.393 + fg*0.769 + fb*0.189) * float64(intensity) + fr*(1- float64(intensity))
                        sg := (fr*0.349 + fg*0.686 + fb*0.168) * float64(intensity) + fg*(1-float64(intensity))
                        sb := (fr*0.272 + fg*0.534 + fb*0.131) * float64(intensity) + fb*(1-float64(intensity))

                        sr = clamp(sr, 0, 1)
                        sg = clamp(sg, 0, 1)
                        sb = clamp(sb, 0, 1)

                        c := color.RGBA{uint8(sr * 255), uint8(sg * 255), uint8(sb * 255), uint8(fa * 255)}
                        sepiaImg.Set(x, y, c)
                }
        }
        return encodeImage(sepiaImg, req.GetImage().GetFormat())
}

func (s *server) Blur(ctx context.Context, req *pb.BlurRequest) (*pb.Image, error) {
        img, err := decodeImage(req.GetImage())
        if err != nil {
                return nil, err
        }

        bounds := img.Bounds()
        blurredImg := image.NewRGBA(bounds)
        kernelSize := int(req.GetKernelSize())
        if kernelSize%2 == 0 {
                kernelSize++
        }
        kernelRadius := kernelSize / 2

        for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
                for x := bounds.Min.X; x < bounds.Max.X; x++ {
                        var r, g, b, a float64
                        var count float64
                        for ky := -kernelRadius; ky <= kernelRadius; ky++ {
                                for kx := -kernelRadius; kx <= kernelRadius; kx++ {
                                        nx, ny := x+kx, y+ky
                                        if nx >= bounds.Min.X && nx < bounds.Max.X && ny >= bounds.Min.Y && ny < bounds.Max.Y {
                                                cr, cg, cb, ca := img.At(nx, ny).RGBA()
                                                r += float64(cr)
                                                g += float64(cg)
                                                b += float64(cb)
                                                a += float64(ca)
                                                count++
                                        }
                                }
                        }
                        if count > 0 {
                                r /= count
                                g /= count
                                b /= count
                                a /= count
                        }

                        c := color.RGBA{uint8(r / 256), uint8(g / 256), uint8(b / 256), uint8(a / 256)}
                        blurredImg.Set(x, y, c)
                }
        }
        return encodeImage(blurredImg, req.GetImage().GetFormat())
}

func decodeImage(imgData *pb.Image) (image.Image, error) {
        switch imgData.GetFormat() {
        case "jpeg", "jpg":
                return jpeg.Decode(bytes.NewReader(imgData.GetData()))
        case "png":
                return png.Decode(bytes.NewReader(imgData.GetData()))
        default:
                return nil, fmt.Errorf("unsupported image format: %s", imgData.GetFormat())
        }
}
func encodeImage(img image.Image, format string) (*pb.Image, error) {
        var buf bytes.Buffer

        switch format {
        case "jpeg", "jpg":
                err := jpeg.Encode(&buf, img, nil)
                if err != nil {
                        return nil, err
                }
        case "png":
                err := png.Encode(&buf, img)
                if err != nil {
                        return nil, err
                }
        default:
                return nil, fmt.Errorf("unsupported image format: %s", format)
        }

        return &pb.Image{Data: buf.Bytes(), Format: format}, nil
}

func clamp(v, min, max float64) float64 {
        if v < min {
                return min
        }
        if v > max {
                return max
        }
        return v
}

func main() {
        lis, err := net.Listen("tcp", ":50051")
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
        }
        s := grpc.NewServer()
        pb.RegisterImageConverterServer(s, &server{})

        if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
        }
}
