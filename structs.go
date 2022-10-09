package pdfimage

type ImageResizeRequest struct {
	H int
	W int
}

type ConvertRequestOptions struct {
	Size      ImageResizeRequest
	Grayscale bool
}