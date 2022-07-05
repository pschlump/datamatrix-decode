package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/datamatrix"
	"github.com/pschlump/filelib"
)

func main() {
	for ii, fn := range os.Args {
		if ii == 0 {
			continue
		}

		file, err := filelib.Fopen(fn, "r")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid file: %s error:%s\n", fn, err)
			continue
		}
		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid DataMatrix code, file: %s error:%s\n", fn, err)
			continue
		}

		// prepare BinaryBitmap
		bmp, err := gozxing.NewBinaryBitmapFromImage(img)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to convert to DataMatrix-code to bitmap, file: %s error:%s\n", fn, err)
			continue
		}

		// decode image
		datamatrixReader := datamatrix.NewDataMatrixReader()
		result, err := datamatrixReader.Decode(bmp, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to decode DataMatrix, file : %s error:%s\n", fn, err)
			continue
		}

		fmt.Printf("%s: %s\n", fn, result)
	}
}
