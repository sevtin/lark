package ximaging

import (
	"bytes"
	"github.com/disintegration/imaging"
	"image"
	"lark/pkg/common/xbytes"
	"lark/pkg/common/xlog"
	"strings"
)

const (
	JPEG imaging.Format = iota
	PNG
	GIF
	TIFF
	BMP
	WEBP
)

type resizeInfo struct {
	isWidth bool
	size    int
}

func makeResizeInfo(img image.Image) resizeInfo {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	if w >= h {
		return resizeInfo{
			isWidth: true,
			size:    w,
		}
	} else {
		return resizeInfo{
			isWidth: false,
			size:    h,
		}
	}
}

func getImageFormat(extName string) (int, error) {
	formats := map[string]imaging.Format{
		".jpg":  JPEG,
		".jpeg": JPEG,
		".png":  PNG,
		".tif":  TIFF,
		".tiff": TIFF,
		".bmp":  BMP,
		".gif":  GIF,
		// ".webp": WEBP,
	}

	ext := strings.ToLower(extName)
	f, ok := formats[ext]
	if !ok {
		return -1, imaging.ErrUnsupportedFormat
	}

	return int(f), nil
}

func ReSizeImage(rb []byte, extName string, isABC bool, cb func(szType string, localId int, w, h int32, b []byte) error) (err error) {
	var (
		img image.Image
		f   int
	)

	img, err = imaging.Decode(bytes.NewReader(rb))
	if err != nil {
		xlog.Warn(err.Error())
		return
	}
	imgSz := makeResizeInfo(img)

	var (
		szList    []ReSizeInfo
		willBreak = false
		rsz       int
	)

	if isABC {
		szList = ReSizeInfoABCList
	} else {
		szList = ReSizeInfoPhotoList
	}

	for _, sz := range szList {
		rsz = sz.Size
		if !isABC {
			if rsz >= imgSz.size {
				rsz = imgSz.size
				willBreak = true
			}
		}

		var dst *image.NRGBA
		if imgSz.isWidth {
			dst = imaging.Resize(img, rsz, 0, imaging.Lanczos)
		} else {
			dst = imaging.Resize(img, 0, rsz, imaging.Lanczos)
		}

		f, err = getImageFormat(extName)
		if err != nil {
			xlog.Warn(err.Error())
			return
		}

		o := xbytes.NewBuffer(make([]byte, 0, len(rb)))
		if f == int(imaging.JPEG) {
			err = imaging.Encode(o, dst, imaging.JPEG)
		} else {
			err = imaging.Encode(o, dst, imaging.Format(f))
		}

		if err != nil {
			xlog.Warn(err.Error())
			return
		}
		err = cb(sz.Type, sz.LocalId, int32(dst.Bounds().Dx()), int32(dst.Bounds().Dy()), o.Bytes())
		if err != nil {
			return
		}

		if willBreak {
			break
		}
	}
	return
}
