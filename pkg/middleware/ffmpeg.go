package middleware

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"os"
)

func ReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

func SavePicture(video string, cover string) error {
	reader := ReadFrameAsJpeg("/public/"+video, 5)
	img, err := imaging.Decode(reader)
	if err != nil {
		return err
	}
	err = imaging.Save(img, "/public/"+cover)
	if err != nil {
		return err
	}
	return nil
}
