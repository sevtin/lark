package openai

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"os"
)

func pictureMultipart(request *ImageVariationReq, w *multipart.Writer) (err error) {
	var (
		file   *os.File
		writer io.Writer
	)
	file, err = os.Open(request.Image)
	if err != nil {
		return
	}
	if writer, err = w.CreateFormFile("image", file.Name()); err != nil {
		return
	}
	if _, err = io.Copy(writer, file); err != nil {
		return
	}
	if err = w.WriteField("size", request.Size); err != nil {
		return
	}
	if err = w.WriteField("n", fmt.Sprintf("%d", request.N)); err != nil {
		return
	}
	if err = w.WriteField("response_format", request.ResponseFormat); err != nil {
		return
	}
	w.Close()
	return
}

func VerifyPngs(paths []string) (err error) {
	var (
		path           string
		found          bool
		expectedWidth  int
		expectedHeight int
		f              *os.File
		fi             os.FileInfo
		image          image.Image
		width          int
		height         int
	)

	for _, path = range paths {
		if f, err = os.Open(path); err != nil {
			return
		}
		if fi, err = f.Stat(); err != nil {
			return
		}
		if fi.Size() > 4*1024*1024 {
			return
		}
		if image, err = png.Decode(f); err != nil {
			return
		}
		width = image.Bounds().Dx()
		height = image.Bounds().Dy()
		if width != height {
			return
		}

		if found == false {
			found = true
			expectedWidth = width
			expectedHeight = height
		} else {
			if width != expectedWidth || height != expectedHeight {
				return
			}
		}
	}
	return
}

func ConvertToRGBA(inputFilePath string, outputFilePath string) (err error) {
	var (
		inputFile  *os.File
		img        image.Image
		rgba       *image.RGBA
		x          int
		y          int
		outputFile *os.File
	)
	// 打开输入文件
	if inputFile, err = os.Open(inputFilePath); err != nil {
		return
	}
	defer inputFile.Close()

	// 解码图像
	if img, _, err = image.Decode(inputFile); err != nil {
		return
	}

	// 将图像转换为RGBA模式
	rgba = image.NewRGBA(img.Bounds())
	for x = 0; x < img.Bounds().Max.X; x++ {
		for y = 0; y < img.Bounds().Max.Y; y++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	// 创建输出文件
	if outputFile, err = os.Create(outputFilePath); err != nil {
		return
	}
	defer outputFile.Close()

	// 编码图像为 PNG 格式并写入输出文件
	if err = png.Encode(outputFile, rgba); err != nil {
		return
	}
	return
}

func ConvertJpegToPNG(path string) (err error) {
	var (
		f    *os.File
		out  *os.File
		name string
		img  image.Image
	)
	// Open the JPEG file for reading
	if f, err = os.Open(path); err != nil {
		return
	}
	defer f.Close()

	// Check if the file is a JPEG image
	if _, err = jpeg.Decode(f); err != nil {
		// The file is not a JPEG image, no need to convert it
		return
	}

	// Reset the file pointer to the beginning of the file
	if _, err = f.Seek(0, 0); err != nil {
		return err
	}

	// Create a new PNG file for writing
	name = path[:len(path)-4] + ".png" // replace .jpg extension with .png
	if out, err = os.Create(name); err != nil {
		return
	}
	defer out.Close()

	// Decode the JPEG image and encode it as PNG
	if img, err = jpeg.Decode(f); err != nil {
		return err
	}
	if err = png.Encode(out, img); err != nil {
		return err
	}
	return
}

func GetImageCompressionType(path string) (format string, err error) {
	var (
		file   *os.File
		reader *bufio.Reader
	)
	// 打开文件
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	// 创建 bufio.Reader
	reader = bufio.NewReader(file)

	// 解码图像
	_, format, err = image.DecodeConfig(reader)
	if err != nil {
		return
	}
	return
}
