package util

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"www.github.com/kushalchg/DataEntryApis/global"
)

func AsciiConverter(filePath string) (htmlFile, textFile, imageFile string) {
	// retriving the file name value from filepath
	fileName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))

	global.Logger.Printf("with the use of filePath package %v", fileName)

	image := loadImage(filePath)
	// greater the image size more clear photo will produce
	resizeImage := resizeImage(image, 200)

	file, err := os.Create("output/resize.png")
	if err != nil {
		fmt.Printf("error while opening file %v\n", err)
	}
	defer file.Close()
	png.Encode(file, resizeImage)
	grayImage := convGrayScale(resizeImage)
	// for gray image
	grayFile, err := os.Create("output/gray.png")
	if err != nil {
		fmt.Printf("error while opening file %v\n", err)

	}
	defer file.Close()
	png.Encode(grayFile, grayImage)
	resultStr := mapAscii(grayImage)
	saveToText(resultStr, fileName)
	asciiToHTML(resultStr, fileName)
	asciiToImage(resultStr, fileName)

	return fmt.Sprintf("%v.html", fileName), fmt.Sprintf("%v.txt", fileName), fmt.Sprintf("%v.png", fileName)

}
func loadImage(filePath string) image.Image {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error while opening file %v\n", err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		fmt.Printf("error while decoding image %v\n", err)
	}
	return img
}

func resizeImage(img image.Image, width int) image.Image {
	bounds := img.Bounds()
	height := (bounds.Dy() * width) / bounds.Dx()
	newImage := image.NewRGBA(image.Rect(0, 0, width, height))
	// Resize the mask image to match the target dimensions
	draw.CatmullRom.Scale(newImage, newImage.Bounds(), img, bounds, draw.Over, nil)
	return newImage
}

func convGrayScale(img image.Image) image.Image {
	bound := img.Bounds()
	grayImage := image.NewRGBA(bound)

	for i := bound.Min.X; i < bound.Max.X; i++ {
		for j := bound.Min.Y; j < bound.Max.Y; j++ {
			oldPixel := img.At(i, j)
			color := color.GrayModel.Convert(oldPixel)
			// // fmt.Print(color)
			// r, g, b, _ := oldPixel.RGBA()
			// grayValue := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			// color := color.Gray{uint8(grayValue / 256)}
			grayImage.Set(i, j, color)
		}
	}
	return grayImage
}

func mapAscii(img image.Image) []string {
	asciiChar := "$@B%#*+=,.      "
	bound := img.Bounds()
	height, width := bound.Max.Y, bound.Max.X
	result := make([]string, height)

	for y := bound.Min.Y; y < height; y++ {
		line := ""
		for x := bound.Min.X; x < width; x++ {
			pixelValue := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			pixel := pixelValue.Y
			asciiIndex := int(pixel) * (len(asciiChar) - 1) / 255
			line += string(asciiChar[asciiIndex])
		}
		result[y] = line
	}
	return result
}

func saveToText(asciiArt []string, fileName string) error {

	file, err := os.Create(fmt.Sprintf("converted/text/%v.txt", fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range asciiArt {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func asciiToHTML(ascii []string, fileName string) {
	HtmlFile, err := os.Create(fmt.Sprintf("converted/html/%v.html", fileName))
	if err != nil {
		fmt.Println("error while creatig html file")
	}

	for lin, lines := range ascii {
		htmlString := `<!DOCTYPE html>
		<html lang="en"><head>
   	 	<meta charset="UTF-8">
    	<meta name="viewport" content="width=device-width, initial-scale=0.5">
    	<title>AsciiImage</title>
		</head>
		<body>
			<code>
		 		<span class="ascii" style="color: black;
		  		background: white;
		  		display:inline-block;
		  		white-space:pre;
		  		letter-spacing:0;
		  		line-height:0.9;
		  		font-family:'Consolas','BitstreamVeraSansMono','CourierNew',Courier,monospace;
		  		font-size:10px;
		  		border-width:1px;
		  		border-style:solid;
		  		border-color:lightgray;">`
		if lin == 0 {
			_, err := HtmlFile.WriteString(htmlString)
			if err != nil {
				fmt.Println("error while start writing into html file")
			}
		}

		for _, char := range lines {
			_, err := HtmlFile.WriteString(fmt.Sprintf("<span>%v</span>", string(char)))
			if err != nil {
				fmt.Println("error while writing into html file")
			}
		}
		_, err := HtmlFile.WriteString("<br>")
		if err != nil {
			fmt.Println("error while writing into html file")
		}
		if lin == len(ascii)-1 {
			_, err := HtmlFile.WriteString("</code></body></html>")
			if err != nil {
				fmt.Println("error while end writing into html file")
			}

		}
	}
}

func asciiToImage(strArray []string, fileName string) {
	// Create a larger image to fit the text
	fontImage := image.NewRGBA(image.Rect(0, 0, 1400, len(strArray)*11))
	// backgroundColor := color.RGBA{0, 0, 255, 255}
	draw.Draw(fontImage, fontImage.Bounds(), image.White, image.Point{}, draw.Src)
	// draw.Draw(fontImage, fontImage.Bounds(), image.NewUniform(backgroundColor), image.Point{}, draw.Src)

	// Create the font face
	drawconf := &font.Drawer{
		Dst:  fontImage,
		Src:  image.Black,
		Face: basicfont.Face7x13,
	}
	// Draw the string
	for i, line := range strArray {
		drawconf.Dot = fixed.Point26_6{
			X: fixed.Int26_6(10 * 64),          // 10 pixels from left
			Y: fixed.Int26_6((20 + i*11) * 64), // Start at 20 pixels from top, then move down by lineHeight for each line
		}

		drawconf.DrawString(line)

	}

	// Create the output file
	file, err := os.Create(fmt.Sprintf("converted/images/%v.png", fileName))
	if err != nil {
		fmt.Println("Error while creating file:", err)
		return
	}
	defer file.Close()

	// Encode and save the image
	err = png.Encode(file, fontImage)
	if err != nil {
		fmt.Println("Error encoding image:", err)
	}
}
