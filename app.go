package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/lifei6671/gocaptcha"
)

const (
	dx = 100
	dy = 50
)

func main() {

	err := gocaptcha.ReadFonts("fonts", ".ttf")

	if err != nil {
		fmt.Println(err)
		return
	}
	http.HandleFunc("/", Get)
	fmt.Println("服务已启动...4000")
	err = http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin")) //允许访问所有域
	w.Header().Set("Access-Control-Allow-Credentials", "true")            //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")        //header的类型
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(200)
	// dx =
	width := r.FormValue("width")
	var currWidth = dx
	if width != "" {
		_width, err := strconv.Atoi(width)
		if err != nil {
			res := map[string]interface{}{
				"code": 0,
				"data": "width 参数有误",
			}
			result, err := json.Marshal(res)
			if err != nil {
				log.Fatal("json转换失败")
			}
			w.Write([]byte(result))
			return
		}
		currWidth = _width
	}

	height := r.FormValue("height")
	var currHeight = dy
	if height != "" {
		_height, err := strconv.Atoi(height)
		if err != nil {
			res := map[string]interface{}{
				"code": 0,
				"data": "height 参数有误",
			}
			result, err := json.Marshal(res)
			if err != nil {
				log.Fatal("json转换失败")
			}
			w.Write([]byte(result))
			return
		}
		currHeight = _height
	}

	captchaImage, err := gocaptcha.NewCaptchaImage(currWidth, currHeight, gocaptcha.RandLightColor())
	if err != nil {
		log.Fatal(err)
	}

	//captchaImage.DrawNoise(gocaptcha.CaptchaComplexLower)
	//captchaImage.DrawTextNoise(gocaptcha.CaptchaComplexLower)
	//captchaImage.DrawHollowLine()
	//captchaImage.Drawline(3);

	code := gocaptcha.RandText(4)
	captchaImage.DrawText(code)
	// 默认boder

	//captchaImage.DrawBorder(gocaptcha.ColorToRGB(0x17A7A7A))

	var b bytes.Buffer
	captchaImage.SaveImage(&b, gocaptcha.ImageFormatPng)

	base64Str := base64.StdEncoding.EncodeToString(b.Bytes())
	res := map[string]interface{}{
		"code": code,
		"data": `data:image/png;base64,` + base64Str,
	}
	result, err := json.Marshal(res)
	if err != nil {
		log.Fatal("json转换失败")
	}

	w.Write([]byte(result))
}
