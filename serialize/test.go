package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"io"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	a := `{"data":{"urls":[{"long_url":"http://o9d.cn","IP":"127.0.0.1"},{"long_url":"http://huiyimei.com/name=阿贾克斯","IP":"192.168.1.1"}],"timestamp":1467012356},"meta":"ab"}`
	GenerateSign([]byte(a), 1467012356, "AABBABAB")
	b := `{"data":{"urls":[{"long_url":"http://o9d.cn","IP":"127.0.0.1"},{"long_url":"http://huiyimei.com/name=阿贾克斯","IP":"192.168.1.1"}],"timestamp":1467012356,"sign":"127AFCDCF37DAA9650A78D57D4EC0119"},"meta":"ab"}`
	is := ValidSign([]byte(b), "AABBABAB")
	fmt.Println(is)
}

//参数key全部按键值排序     ToUpper(md5(sha1(SecretKey1Value1Key2Value2SecretTime)))
func GenerateSign(requestData []byte, requestTime int64, secretKey string) string {
	var rdata map[string]interface{}
	json.Unmarshal([]byte(requestData), &rdata)
	str := Serialize(rdata)
	fmt.Println("rrrrr", str)
	// serial := secretKey + str.(string) + secretKey + strconv.FormatInt(int64(requestTime), 10)
	var serial bytes.Buffer
	serial.WriteString(secretKey)
	serial.WriteString(str.(string))
	serial.WriteString(secretKey)
	serial.WriteString(strconv.FormatInt(int64(requestTime), 10))
	fmt.Println("11-------------------------------", serial.String())
	urlencodeSerial := url.QueryEscape(serial.String())
	fmt.Println("22-------------------------------", urlencodeSerial)
	urlencodeBase64Serial := base64.StdEncoding.EncodeToString([]byte(urlencodeSerial))
	fmt.Println("33-------------------------------", urlencodeBase64Serial)
	sign, _ := Sha1(urlencodeBase64Serial)
	fmt.Println("44-------------------------------", sign)
	sign, _ = MD5(sign)

	fmt.Println("rrrrrrr", strings.ToUpper(sign))

	return strings.ToUpper(sign)
}

func Serialize(data interface{}) interface{} {
	fmt.Println("----", data, "---", reflect.TypeOf(data).Kind())
	// var str string
	var buffer bytes.Buffer
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(data)
		for i := 0; i < s.Len(); i++ {
			serial := Serialize(s.Index(i).Interface())
			if reflect.TypeOf(serial).Kind() == reflect.Float64 {
				serial = strconv.Itoa(int(serial.(float64)))
			}
			// str = str + strconv.Itoa(i) + serial.(string)
			buffer.WriteString(strconv.Itoa(i))
			buffer.WriteString(serial.(string))
		}
		return buffer.String()
	case reflect.Map:
		s := reflect.ValueOf(data)

		keys := s.MapKeys()
		//ksort
		sorted_keys := make([]string, 0)
		for _, key := range keys {
			sorted_keys = append(sorted_keys, key.Interface().(string))
		}
		sort.Strings(sorted_keys)
		for _, key := range sorted_keys {
			serial := Serialize(s.MapIndex(reflect.ValueOf(key)).Interface())
			if reflect.TypeOf(serial).Kind() == reflect.Float64 {
				serial = strconv.Itoa(int(serial.(float64)))
			}
			// str = str + key + serial.(string)
			buffer.WriteString(key)
			buffer.WriteString(serial.(string))
		}
		//     for _, key := range keys {
		//         serial := Serialize(s.MapIndex(reflect.ValueOf(key.String())).Interface(), true)
		//         if reflect.TypeOf(serial).Kind() == reflect.Float64 {
		//             serial = strconv.Itoa(int(serial.(float64)))
		//         }
		//         str = str + key.String() + serial.(string)
		//     }
		// }
		return buffer.String()
	}

	fmt.Println("pppppp")
	// fmt.Println(data)
	// switch data.(type) {
	// case []interface{}:
	// case map[string]interface{}:
	// 	for k, v := range data {
	//
	// 	}
	// }
	return data
}

//MD5 md5加密
func MD5(src string) (string, error) {
	return _hash(src, md5.New())
}

//Sha1 sha1加密
func Sha1(src string) (string, error) {
	return _hash(src, sha1.New())
}

func _hash(src string, h hash.Hash) (string, error) {
	if _, err := io.WriteString(h, src); err != nil {
		return "", err
	}
	sig := hex.EncodeToString(h.Sum(nil))
	return sig, nil
}

//ValidSign 签名验证
func ValidSign(requestData []byte, secretKey string) error {
	//取出sign Timestamp
	var rdata map[string]interface{}
	json.Unmarshal(requestData, &rdata)
	data, _ := rdata["data"].(map[string]interface{})
	sign := data["sign"].(string)
	timestamp := int64(data["timestamp"].(float64))

	//去除sign
	_, ok := data["sign"]
	if ok {
		delete(data, "sign")
	}

	jsonData, err := json.Marshal(rdata)
	if err != nil {
		return err
	}

	//生成签名
	fmt.Println("====", timestamp)
	signed := GenerateSign(jsonData, timestamp, secretKey)
	fmt.Println(sign, signed)

	//对比sign
	if sign != signed {
		return errors.New("签名错误")
	}

	//时间是否合理
	if diff := time.Now().Unix() - timestamp; diff > 600 {
		return errors.New("时间不对")
	}

	return nil
}
