// Lissajous генерирует анимированный GIF из случайных
// фигур Лиссажу.
// Измените палитру программы lissajous так,
// чтобы она генерировала изображения разных цветов,
// добавляя в палитру palette больше значений,
// а затем выводя их путем изменения третьего аргумента
// функции SetColorlndex некоторым нетривиальным способом.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White,
	color.RGBA{0x15, 0x4d, 0x0f, 0xff}, // зеленый
	color.RGBA{0xeb, 0x05, 0x05, 0xff}, // красный
	color.RGBA{0xeb, 0xdf, 0x05, 0xff}, // желтый
	color.RGBA{0x18, 0x05, 0xeb, 0xff}} // синий

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0                // Относительная частота колебаний y
	colorIndex := uint8(rand.Float64()*3.0 + 1) // индекс цвета
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Примечание: игнорируем ошибки
}

// сбилдить, затем результат работы сохранить в gif файл:
//$ go build lissajous_hw1.6.go
//$ lissajous_hw1.6.exe >out2.gif
