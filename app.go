package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lifei6671/gocaptcha"
)

const (
	dx = 150
	dy = 50
)

func main() {

	err := gocaptcha.ReadFonts("fonts", ".ttf")
	if err != nil {
		fmt.Println(err)
		return
	}
	http.HandleFunc("/", Get)
	fmt.Println("服务已启动...3000")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {

	captchaImage, err := gocaptcha.NewCaptchaImage(dx, dy, gocaptcha.RandLightColor())

	captchaImage.DrawNoise(gocaptcha.CaptchaComplexLower)

	captchaImage.DrawTextNoise(gocaptcha.CaptchaComplexLower)

	code := gocaptcha.RandText(4)

	captchaImage.DrawText(code)
	//captchaImage.Drawline(3);
	captchaImage.DrawBorder(gocaptcha.ColorToRGB(0x17A7A7A))
	captchaImage.DrawSineLine()

	//captchaImage.DrawHollowLine()
	if err != nil {
		fmt.Println(err)
	}
	var b bytes.Buffer
	captchaImage.SaveImage(&b, gocaptcha.ImageFormatJpeg)

	base64Str := base64.StdEncoding.EncodeToString(b.Bytes())
	w.Header().Set("Content-Type", "application/json")
	res := map[string]interface{}{
		"code": code,
		"data": base64Str,
	}
	result, err := json.Marshal(res)
	if err != nil {
		log.Fatal("json转换失败")
	}
	w.WriteHeader(200)
	w.Write([]byte(result))
}
sudo docker login --username=412790861 ccr.ccs.tencentyun.com
sudo docker tag 539234e8e884 ccr.ccs.tencentyun.com/showers/demo:v1
sudo docker tag 539234e8e884 ccr.ccs.tencentyun.com/showers/docker:[tag]
sudo docker push ccr.ccs.tencentyun.com/showers/demo
/home/ubuntu/.docker
sudo docker login --username=ubuntu suyulin.com:5000/v2
DOCKER_OPTS="--registry-mirror=suyulin.com:5000"