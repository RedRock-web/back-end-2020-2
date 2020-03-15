package jwts

import (
	"back-end-2020-1/config"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//Jwt struct 表示一个 json web token
type Jwt struct {
	form      config.LoginForm
	Header    Header
	Payload   Payload
	Signature Signature
}

//NewJwt 返回一个 Jwt struct
func NewJwt() Jwt {
	return Jwt{}
}

//Create 返回 一个 特定的 Jwt
func (j *Jwt) Create(form config.LoginForm, key string) (string, error) {
	j.form = form
	j.Signature.key = key

	//获取 header and payload
	hAndp, err := j.headerAndPayload2str()
	if err != nil {
		errors.New("Header 和 Payload 拼接失败！")
		return "", err
	}

	//获取 signature
	signature, err := j.Signature.New(*j)
	if err != nil {
		errors.New("获取 Signature 失败！")
		return "", err
	}

	return hAndp + "." + signature, nil
}

func (j *Jwt) Check(token string, key string) (config.LoginForm, error) {
	//首先把 token 和划分为 3 部分
	arr := strings.Split(token, ".")
	if len(arr) != 3 {
		return config.LoginForm{}, errors.New("token error!")
	}

	//对 Header 解密
	_, err := base64.StdEncoding.DecodeString(arr[0])
	if err != nil {
		return config.LoginForm{}, errors.New("token error!")
	}

	//对 payload 解密
	pay, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		return config.LoginForm{}, errors.New("token error!")
	}

	//对 signature 解密
	_, err = base64.StdEncoding.DecodeString(arr[2])
	if err != nil {
		return config.LoginForm{}, errors.New("token error!")
	}

	hAndP := arr[0] + "." + arr[1]
	ss := base64.StdEncoding.EncodeToString(HmacSha256(hAndP, key))
	if res := strings.Compare(arr[2], ss); res != 0 {
		return config.LoginForm{}, errors.New("token error!")
	}

	var payload Payload
	json.Unmarshal(pay, &payload)

	return config.LoginForm{payload.Username, payload.Password}, nil
}

//Header 表示 Jwt 的 header
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

//New 返回一个 Alg 为 HS256, Typ 为 JWT 的 Header 对象
func (h *Header) New() Header {
	return Header{
		Alg: "HS256",
		Typ: "JWT",
	}
}

//Payload 表示 Jwt 的 payload
type Payload struct {
	Iss      string `json:"iss"`
	Exp      string `json:"exp"`
	Iat      string `json:"iat"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//New 返回一个特定的 Payload
func (p *Payload) New(form config.LoginForm) Payload {
	return Payload{
		Iss:      "redrock",
		Exp:      strconv.FormatInt(time.Now().Add(3*time.Hour).Unix(), 10),
		Iat:      strconv.FormatInt(time.Now().Unix(), 10),
		Username: form.Username,
		Password: form.Password,
	}
}

//Signature 表示 Jwt 的 signature
type Signature struct {
	key string
}

//New 返回 一个 signature
func (s *Signature) New(j Jwt) (string, error) {
	str, err := j.headerAndPayload2str()
	if err != nil {
		errors.New("Header 和 Payload 拼接失败！")
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(HmacSha256(str, s.key))
	return signature, nil
}

//headerAndPayload2str() 将 Json 格式 的 Header 和 Payload 转换为 String 后并拼接
func (j *Jwt) headerAndPayload2str() (string, error) {
	h, err := json.Marshal(j.Header.New())
	if err != nil {
		errors.New("解析 Header 出错！")
		return "", err
	}
	p, err := json.Marshal(j.Payload.New(j.form))
	if err != nil {
		errors.New("解析 Payload 出错！")
		return "", err
	}
	headerBase64 := base64.StdEncoding.EncodeToString(h)
	payloadBase64 := base64.StdEncoding.EncodeToString(p)

	return strings.Join([]string{headerBase64, payloadBase64}, "."), nil
}

func HmacSha256(str string, key string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(key))
	fmt.Println(mac.Sum(nil))
	return mac.Sum(nil)
}
