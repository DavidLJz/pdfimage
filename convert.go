package pdfimage

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/karmdip-mi/go-fitz"
)

func CheckInputError(input string, output string) error {
	if input == "" || output == "" {
		return errors.New("Input / Output parameters are needed")
	}

	filext := strings.ToLower(filepath.Ext(filepath.Base(input)))

	if filext != ".pdf" {
		return errors.New("Input parameter must be path of a file with PDF extension")
	}

	fi, err := os.Lstat(output)
	if err != nil {
		return err
	}

	if fi.Mode().IsDir() == false {
		return errors.New("Output parameter is not a valid directory")
	}

	if _, err := os.Stat(output); err != nil {
		if errors.Is(err, os.ErrNotExist) == true {
			if err := os.Mkdir(output, 0666); err != nil {
				return errors.New(fmt.Sprintf("Can't create output directory (%s): %s", output, err))
			}
		} else {
			return err
		}
	}

	return nil
}

func GetPdfImageData(input string, opt ConvertRequestOptions) ([]image.Image, error) {
	doc, err := fitz.New(input)
	if err != nil {
		return nil, err
	}

	var images []image.Image

	for n := 0; n < doc.NumPage(); n++ {
		imgdata, err := doc.Image(n)
		if err != nil {
			panic(err)
		}

		if opt.Size.H != 0 || opt.Size.W != 0 {
			imgdata = imaging.Resize(imgdata, opt.Size.W, opt.Size.H, imaging.Linear)
		}

		if opt.Grayscale == true {
			imgdata = imaging.Grayscale(imgdata)
		}

		images = append(images, imgdata)
	}

	return images, nil
}

func PdfToImage(input string, output string, opt ConvertRequestOptions) ([]string, error) {
	if err := CheckInputError(input, output); err != nil {
		return nil, err
	}

	var filelist []string

	pdfimgdata, err := GetPdfImageData(input, opt)
	if err != nil {
		return nil, err
	}

	imgprefix := strings.TrimSuffix(strings.ToLower(filepath.Base(input)), ".pdf")

	for n, img := range pdfimgdata {
		filename := fmt.Sprintf("%s/%s_%03d.jpg", output, imgprefix, n)

		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}

		err = jpeg.Encode(file, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			panic(err)
		}

		file.Close()

		filelist = append(filelist, filename)
	}

	return filelist, nil
}
