package crypto

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

// 将繁体中文转换为简体中文
func SimplifiedChinese(str string) string {
	input := []byte(str)
	output := make([]byte, len(str))
	tr := simplifiedchinese.GBK.NewDecoder()
	_, _, err := tr.Transform(output, input, true)
	if err != nil {
		panic(err)
	}
	return string(output)
}
func StringToBase64(str string) string {
	base64Str := base64.StdEncoding.EncodeToString([]byte(str))
	return base64Str
}

func StringToJson(data interface{}) string {
	d, _ := json.Marshal(data)
	return string(d)
}
func ToIndentJSON(obj any) (string, error) {
	bs, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = json.Indent(buf, bs, "", "\t")
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// 编码结构体为XML格式
func XML(data any) string {
	v, _ := xml.Marshal(data)
	return string(v)
}

// GbkToUtf8 transforms GBK to UTF8.
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	return io.ReadAll(reader)
}

// GbkToUtf8Str transforms GBK to UTF8 string.
func GbkToUtf8Str(s string) string {
	dst, _ := GbkToUtf8([]byte(s))
	return string(dst)
}

// Utf8ToGbk transforms utf8 to gbk.
func Utf8ToGbk(b []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(b), simplifiedchinese.GBK.NewEncoder())
	return io.ReadAll(reader)
}

// Utf8ToGbkStr transforms utf8 to gbk string.
func Utf8ToGbkStr(s string) string {
	dst, _ := Utf8ToGbk([]byte(s))
	return string(dst)
}
