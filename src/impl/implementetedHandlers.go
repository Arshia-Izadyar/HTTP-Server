package impl

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Arshia-Izadyar/HTTP-Server/src/http"
)

func EchoBody(r *http.HttpRequest) http.HttpResponse {
	tempStruct := struct {
		Echo string `json:"echo"`
	}{}
	json.Unmarshal([]byte(r.Body), &tempStruct)
	return http.Cr(200, map[string]string{
		"echo": tempStruct.Echo,
		"body": r.Body,
	})
}

func EchoParameter(r *http.HttpRequest) http.HttpResponse {
	param, err := r.UrlParams.Get("echo")
	if err != nil {
		return http.Cr(400, map[string]string{
			"status": "invalid body",
		})
	}
	return http.Cr(200, map[string]string{
		"echo": param,
	})
}

func ResponseHtml(r *http.HttpRequest) http.HttpResponse {
	f, err := os.Open("../impl/index.html")
	if err != nil {
		fmt.Println(err)
		return http.Cr(500, map[string]string{
			"status": "internal server error",
		})
	}
	html, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return http.Cr(500, map[string]string{
			"status": "internal server error",
		})
	}
	return http.Response(200, string(html), "text/html")
}

func ServeImage(r *http.HttpRequest) http.HttpResponse {
	f, err := os.Open("../static/bugs2.jpg")
	if err != nil {
		fmt.Println(err)
		return http.Cr(500, map[string]string{
			"status": "internal server error",
		})
	}
	html, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return http.Cr(500, map[string]string{
			"status": "internal server error",
		})
	}
	return http.Response(200, string(html), "image/jpg")
}

var fileExtentions map[string]string = map[string]string{
	"csv":  "text/csv",
	"pdf":  "application/pdf",
	"mp3":  "audio/mpeg",
	"png":  "image/png",
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"json": "application/json",
	"docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
}

func ServeFile(r *http.HttpRequest) http.HttpResponse {
	filename, err := r.UrlParams.Get("filename")
	if err != nil {
		fmt.Println(err)
		return http.Cr(500, map[string]string{
			"status": "internal server error",
		})
	}
	f, err := os.Open(fmt.Sprintf("../static/%s", filename))
	if err != nil {
		fmt.Println(err.Error())
		return http.Cr(404, map[string]string{
			"status": "file not found",
		})
	}
	splitedFileName := strings.Split(filename, ".")
	extention := splitedFileName[len(splitedFileName)-1]

	photo, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return http.Cr(500, map[string]string{
			"status": "internal server error",
		})
	}

	if _, ok := fileExtentions[extention]; !ok {
		extention = "octeted-stream"
	} else {
		extention = fileExtentions[extention]

	}
	return http.Response(200, string(photo), extention)
}
