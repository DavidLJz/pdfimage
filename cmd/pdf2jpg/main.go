package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/DavidLJz/pdfimage"
)

var (
	heigth int
	width  int
	gray   bool
)

func init() {
	flag.BoolVar(&gray, "gray", false, "Apply grayscale on output images")
	flag.IntVar(&heigth, "h", 0, "Height of images")
	flag.IntVar(&width, "w", 0, "Width of images")

	flag.Usage = func() {
		fmt.Printf("Usage: %s myfile.pdf parent_dir/output_dir", os.Args[0])
	}
}

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Printf("Input/Output arguments not given\n")
		flag.Usage()
		os.Exit(1)
	}

	input := flag.Arg(0)
	output := flag.Arg(1)

	fmt.Println(gray)
	os.Exit(0)

	opts := pdfimage.ConvertRequestOptions{
		Grayscale: gray,
		Size: pdfimage.ImageResizeRequest{
			H: heigth,
			W: width,
		},
	}

	filelist, err := pdfimage.PdfToImage(input, output, opts)
	if err != nil {
		log.Fatalf("%s", err)
	}

	if len(filelist) == 0 {
		panic("filelist is empty")
	}

	for _, file := range filelist {
		fmt.Printf("%s\n", file)
	}
}