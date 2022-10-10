## Build
Requires the C compiler to be installed in your OS. For Windows, you can get it through [MinGW](https://sourceforge.net/projects/mingw/).

To build move to any directory under `./cmd` and use the `go build` command.
```
cd ./cmd/pdf2jpg
go build
```

## Usage

### PDF2JPG CLI Utility
 
Converts a PDF document and saves a JPEG image for every page in the PDF document.
Prints the list of saved images.

If the job fails on any of the pages being converted it will panic and fail.

```
pdf2jpg myfile.pdf ouput_dir
./myfile_000.jpg
./myfile_001.jpg
```

Use `gray` flag before parameters to output grayscale images.

```
pdf2jpg -gray myfile.pdf ouput_dir
```

Set custom height and width

```
pdf2jpg -h 100 -w 100 myfile.pdf ouput_dir
```