package main

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var wsConn *websocket.Conn

func main() {
	websocketAddr := "riveng36-sht.gw.riven.panda.tv:8080"
	u := url.URL{Scheme: "wss", Host: websocketAddr, Path: "/"}

	origin := "https://www.panda.tv"
	wsHeaders := http.Header{
		"Origin": {origin},
	}

	wd := websocket.Dialer{
		HandshakeTimeout: time.Second * 5,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		Proxy: http.ProxyFromEnvironment,
	}

	var err error
	wsConn, _, err = wd.Dial(u.String(), wsHeaders)
	if err != nil {
		log.Fatal("-dial:", err)
	}

	defer func() {
		wsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "88"))
		wsConn.Close()
	}()

	// 获取房间信息
	param, err := getChatParam(20641)
	if err != nil {
		return
	}
	fmt.Println("++++++++++++: ", param)

	// handshake
	handshake(param)

	// keepAlive

	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}
}

var pandaStart = []byte{0x00, 0x06, 0x00, 0x02}
var pandaHeartbeat = []byte{0x00, 0x06, 0x00, 0x00}
var pandaResponse = []byte{0x00, 0x06, 0x00, 0x06} //连接弹幕服务器响应
var pandaReceiveMsg = []byte{0x00, 0x06, 0x00, 0x03}
var pandaHeartbeatResponse = []byte{0x00, 0x06, 0x00, 0x01}

type PandaChatParam struct {
	Rid          int64    `json:"rid"`
	Appid        string   `json:"appid"`
	Ts           int64    `json:"ts"`
	Sign         string   `json:"sign"`
	AuthType     string   `json:"authType"`
	ChatAddrList []string `json:"chat_addr_list"`
}

type pandaChatData struct {
	Data PandaChatParam `json:"data"`
}

func getChatParam(roomid int64) (*PandaChatParam, error) {
	u1 := fmt.Sprintf("http://riven.panda.tv/chatroom/getinfo?roomid=%d", roomid)
	var chatData pandaChatData
	err := GetJson(u1, &chatData)
	if err != nil {
		return nil, err
	}
	u2 := fmt.Sprintf("http://api.homer.panda.tv/chatroom/getinfo?rid=%d&roomid=%d&retry=0&sign=%s&ts=%d&_=%d",
		chatData.Data.Rid, roomid, chatData.Data.Sign, chatData.Data.Ts, time.Now().Unix()*1000)
	var chatData2 pandaChatData
	err = GetJson(u2, &chatData2)
	if err != nil {
		return nil, err
	}
	return &chatData2.Data, nil
}

func GetJson(url string, v interface{}) error {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("get error ", url, err)
		return err
	}

	if resp.StatusCode != http.StatusOK || resp.Body == nil {
		fmt.Println("status code error ", url, resp.StatusCode)
		return errors.New(fmt.Sprint("status code ", resp.StatusCode))
	}

	defer resp.Body.Close()
	d := json.NewDecoder(resp.Body)
	err = d.Decode(v)
	if err != nil {
		fmt.Println("json error ", url, err)
		return err
	}

	return nil
}

func handshake(param *PandaChatParam) error {
	data := fmt.Sprintf("u:%d@%s\nk:1\nt:300\nts:%d\nsign:%s\nauthtype:%s", param.Rid, param.Appid, param.Ts, param.Sign, param.AuthType)
	l := len(data)

	msg := make([]byte, 4+2+l)
	copy(msg, pandaStart)
	binary.BigEndian.PutUint16(msg[4:], uint16(l))
	// copy(msg[4:], []byte{byte(l >> 8), byte(l & 0xff)})
	copy(msg[4+2:], []byte(data))
	// err := wsConn.WriteMessage(websocket.BinaryMessage, msg)
	err := wsConn.WriteMessage(websocket.BinaryMessage, msg)
	if err != nil {
		return err
	}
	n, buff, err := wsConn.ReadMessage()
	if err != nil {
		return err
	}
	if n != 2 || !bytes.Equal(buff[:4], pandaResponse) {
		return errors.New("response error")
	}

	length := int((uint(buff[4]) << 8) + uint(buff[5]))
	if length > 255 {
		return errors.New("invalid response length flag")
	}
	n, buff2, err := wsConn.ReadMessage()
	if n == 2 || bytes.Equal(buff2[:4], pandaReceiveMsg) {
		log.Printf("开始接收弹幕消息: %v", pandaReceiveMsg)
	}

	return nil
}
