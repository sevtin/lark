package xresize

import (
	"errors"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func ResizePhotoFile(src, dst string, width, height, quality int) error {
	fIn, _ := os.Open(src)
	defer fIn.Close()

	fOut, _ := os.Create(dst)
	defer fOut.Close()

	if err := ResizePhoto(fIn, fOut, width, height, quality); err != nil {
		return err
	}
	return nil
}

func ResizePhoto(in io.Reader, out *os.File, width, height, quality int) error {
	origin, fm, err := image.Decode(in)
	if err != nil {
		return err
	}
	return CropPhoto(origin, fm, out, width, height, quality)
}

func CropPhoto(origin image.Image, fm string, out *os.File, width, height, quality int) error {
	if width == 0 || height == 0 {
		width = origin.Bounds().Max.X / 2
		height = origin.Bounds().Max.Y / 2
	}
	if quality == 0 {
		quality = 25
	}
	canvas := resize.Thumbnail(uint(width), uint(height), origin, resize.Lanczos3)

	switch fm {
	case "jpeg":
		return jpeg.Encode(out, canvas, &jpeg.Options{quality})
	case "png":
		return png.Encode(out, canvas)
	case "gif":
		return gif.Encode(out, canvas, &gif.Options{})
	case "bmp":
		return bmp.Encode(out, canvas)
	default:
		return errors.New("ERROR FORMAT")
	}
	return nil
}
