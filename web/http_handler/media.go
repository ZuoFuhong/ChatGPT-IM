package http_handler

import (
	"bytes"
	"fmt"
	"go-IM/consts"
	"go-IM/pkg/defs"
	"go-IM/pkg/util"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

type media struct{}

var Media = new(media)

// AudioTranscriptions 语音转文本
func (*media) AudioTranscriptions(w http.ResponseWriter, r *http.Request) {
	model := r.FormValue("model")
	file, _, err := r.FormFile("file")
	if err != nil {
		defs.Error(w, defs.ParameterError, err.Error())
		return
	}
	defer file.Close()
	reqBody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)
	_ = writer.WriteField("model", model)
	file2, _ := writer.CreateFormFile("file", fmt.Sprint(time.Now().UnixMilli(), ".mp3"))
	if _, err := io.Copy(file2, file); err != nil {
		defs.Error(w, defs.ParameterError, err.Error())
		return
	}
	if err := writer.Close(); err != nil {
		defs.Error(w, defs.ParameterError, err.Error())
		return
	}
	headers := map[string]string{
		"Authorization": "Bearer " + consts.APIKey,
		"Content-Type":  writer.FormDataContentType(),
	}
	rspBytes, err := util.DefaultClient.DoReq("POST", "https://api.openai.com/v1/audio/transcriptions", reqBody, headers)
	if err != nil {
		defs.Error(w, defs.Media, "语义转文本失败")
		return
	}
	_, _ = w.Write(rspBytes)
}
