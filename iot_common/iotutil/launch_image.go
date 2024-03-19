package iotutil

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

// //生成APP启动图
// //将logo设置为原图宽度的30%
// //将logo设置未在原图高度35%展示
// func GenLaunchImageFromIcon(srcFile, outFile string, width, height int) error {
// 	img, err := imgio.Open(srcFile)
// 	if err != nil {
// 		return err
// 	}
// 	logowidth := int(float64(width) * 0.3)
// 	logoImage := transform.Resize(img, logowidth, logowidth, transform.Lanczos)
// 	bkgImage := image.NewRGBA(image.Rect(0, 0, width, height))
// 	for x := 0; x < bkgImage.Bounds().Dx(); x++ {
// 		for y := 0; y < bkgImage.Bounds().Dy(); y++ {
// 			bkgImage.Set(x, y, color.White)
// 		}
// 	}
// 	offset := image.Pt((bkgImage.Bounds().Max.X-logoImage.Bounds().Max.X)/2, int(float64(bkgImage.Bounds().Max.Y-logoImage.Bounds().Max.Y)*0.35))
// 	draw.Draw(bkgImage, bkgImage.Bounds().Add(offset), logoImage, image.Pt(0, 0), draw.Over)
// 	ext := strings.ToLower(filepath.Ext(srcFile))
// 	var encoder imgio.Encoder
// 	switch ext {
// 	case ".png":
// 		encoder = imgio.PNGEncoder()
// 	case ".jpg", ".jpeg":
// 		encoder = imgio.JPEGEncoder(jpeg.DefaultQuality)
// 	case ".bmp":
// 		encoder = imgio.BMPEncoder()
// 	}
// 	if err = imgio.Save(outFile, bkgImage, encoder); err != nil {
// 		return err
// 	}
// 	return nil
// }

//生成APP启动图
//将logo设置为原图宽度的30%
//将logo设置未在原图高度35%展示
//radius 圆角尺寸
func GenLaunchImageFromIcon(srcFile, outFile string, width, height int, radius int) error {
	srcImage := DrawRadiusPic(srcFile, radius)
	logowidth := int(float64(width) * 0.3)
	//logoImage := transform.Resize(img, logowidth, logowidth, transform.Lanczos)
	logoImage := transform.Resize(srcImage, logowidth, logowidth, transform.Lanczos)
	bkgImage := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < bkgImage.Bounds().Dx(); x++ {
		for y := 0; y < bkgImage.Bounds().Dy(); y++ {
			bkgImage.Set(x, y, color.White)
		}
	}
	offset := image.Pt((bkgImage.Bounds().Max.X-logoImage.Bounds().Max.X)/2, int(float64(bkgImage.Bounds().Max.Y-logoImage.Bounds().Max.Y)*0.35))
	draw.Draw(bkgImage, bkgImage.Bounds().Add(offset), logoImage, image.Pt(0, 0), draw.Over)
	ext := strings.ToLower(filepath.Ext(srcFile))
	var encoder imgio.Encoder
	switch ext {
	case ".png":
		encoder = imgio.PNGEncoder()
	case ".jpg", ".jpeg":
		encoder = imgio.JPEGEncoder(jpeg.DefaultQuality)
	case ".bmp":
		encoder = imgio.BMPEncoder()
	}
	if err := imgio.Save(outFile, bkgImage, encoder); err != nil {
		return err
	}
	return nil
}

func ImageResize(srcFile, outFile string, width, height int) error {
	img, err := imgio.Open(srcFile)
	if err != nil {
		return err
	}
	logoImage := transform.Resize(img, width, height, transform.Lanczos)

	ext := strings.ToLower(filepath.Ext(srcFile))
	var encoder imgio.Encoder
	switch ext {
	case ".png":
		encoder = imgio.PNGEncoder()
	case ".jpg", ".jpeg":
		encoder = imgio.JPEGEncoder(jpeg.DefaultQuality)
	case ".bmp":
		encoder = imgio.BMPEncoder()
	}
	if err := imgio.Save(outFile, logoImage, encoder); err != nil {
		return err
	}
	return nil

}

//给图片做圆角
func DrawRadiusPic(imgsrc string, rad int) *image.RGBA {
	f, err := os.Open(imgsrc)
	if err != nil {
		panic(err)
	}
	gopherImg, _, err := image.Decode(f)
	w := gopherImg.Bounds().Dx()
	h := gopherImg.Bounds().Dy()
	//c := radius{p: image.Point{X: w, Y: h}, r: int(40)}
	c := radius{p: image.Point{X: w, Y: h}, r: rad}
	radiusImg := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.DrawMask(radiusImg, radiusImg.Bounds(), gopherImg, image.Point{}, &c, image.Point{}, draw.Over)
	return radiusImg
}

type radius struct {
	p image.Point // 矩形右下角位置
	r int
}

func (c *radius) ColorModel() color.Model {
	return color.AlphaModel
}
func (c *radius) Bounds() image.Rectangle {
	return image.Rect(0, 0, c.p.X, c.p.Y)
}

// 对每个像素点进行色值设置，分别处理矩形的四个角，在四个角的内切圆的外侧，色值设置为全透明，其他区域不透明
func (c *radius) At(x, y int) color.Color {
	var xx, yy, rr float64
	var inArea bool
	// left up
	if x <= c.r && y <= c.r {
		xx, yy, rr = float64(c.r-x)+0.5, float64(y-c.r)+0.5, float64(c.r)
		inArea = true
	}
	// right up
	if x >= (c.p.X-c.r) && y <= c.r {
		xx, yy, rr = float64(x-(c.p.X-c.r))+0.5, float64(y-c.r)+0.5, float64(c.r)
		inArea = true
	}
	// left bottom
	if x <= c.r && y >= (c.p.Y-c.r) {
		xx, yy, rr = float64(c.r-x)+0.5, float64(y-(c.p.Y-c.r))+0.5, float64(c.r)
		inArea = true
	}
	// right bottom
	if x >= (c.p.X-c.r) && y >= (c.p.Y-c.r) {
		xx, yy, rr = float64(x-(c.p.X-c.r))+0.5, float64(y-(c.p.Y-c.r))+0.5, float64(c.r)
		inArea = true
	}
	if inArea && xx*xx+yy*yy >= rr*rr {
		return color.Alpha{}
	}
	return color.Alpha{A: 255}
}
