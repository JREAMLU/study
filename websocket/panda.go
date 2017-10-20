package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	purl "net/url"
	"strconv"
	"time"
)

func main() {
	// go ListenRoom("6666")
	ListenRoom("20641")
}

var (
	LiveCheckInterval = 10 * time.Minute
	KeepAliveInterval = time.Minute
)

var pandaStart = []byte{0x00, 0x06, 0x00, 0x02}
var pandaHeartbeat = []byte{0x00, 0x06, 0x00, 0x00}
var pandaResponse = []byte{0x00, 0x06, 0x00, 0x06} //连接弹幕服务器响应
var pandaReceiveMsg = []byte{0x00, 0x06, 0x00, 0x03}
var pandaHeartbeatResponse = []byte{0x00, 0x06, 0x00, 0x01}

var ChatAddrList = []string{}

const pandaIgnoreByteLength = 16 //弹幕消息体忽略的字节数

type Panda struct {
	room int64
	conn *net.TCPConn
	exit bool
	pool chan *string
}

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

func ListenRoom(roomID string) {
	p, _ := NewPanda(fmt.Sprintf("https://www.panda.tv/%v", roomID))
	ch, err := p.StartReceive()
	if err != nil {
		fmt.Println("++++++++++++: StartReceive 错误")
		return
	}
	for x := range ch {
		fmt.Println("++++++++++++: ", *x)
	}
}

func NewPanda(url string) (*Panda, error) {
	u, err := purl.Parse(url)
	if err != nil || len(u.Path) <= 1 {
		return nil, fmt.Errorf("invalid room url %s", url)
	}
	room := u.Path[1:]
	id, err := strconv.ParseInt(room, 10, 64)
	if err != nil {
		return nil, err
	}
	return &Panda{id, nil, false, make(chan *string, 1024)}, nil
}

func (p *Panda) init(flag int, chatAddrListIndex int) error {
	param, err := p.getChatParam()
	if err != nil {
		return err
	}

	for _, addrs := range param.ChatAddrList {
		ChatAddrList = append(ChatAddrList, addrs)
	}
	UniqueSliceString(&ChatAddrList)

	var adList string
	if chatAddrListIndex > (len(param.ChatAddrList) - 1) {
		chatAddrListIndex = len(param.ChatAddrList) - 1
	}
	switch flag {
	case 0:
		adList = param.ChatAddrList[chatAddrListIndex]
	case 1:
		adList = ChatAddrList[chatAddrListIndex]
	}

	log.Printf("ChatAddrList列表: %v, 当前连接地址: %v", ChatAddrList, adList)

	// flag = 0 默认取、flag = 1
	conn, err := net.Dial("tcp", adList)
	if err != nil {
		fmt.Println("dial", err)
		return err
	}
	p.conn = conn.(*net.TCPConn)
	p.handshake(param)
	return nil
}

func (p *Panda) StartReceive() (<-chan *string, error) {
	b, err := p.IsLive()
	if err != nil {
		return nil, err
	}

	go p.startReceive(b)
	return p.pool, nil
}

func (p *Panda) startReceive(live bool) {
	pc := 0
	defer func() {
		if pc != 1 {
			p.Close()
		}
		log.Printf("over receive room: %v", p.room)
		str := "EOF"
		p.pool <- &str
	}()

	var err error
	if !live {
		for !p.exit {
			time.Sleep(LiveCheckInterval)
			live, err = p.IsLive()
			if err != nil {
				log.Printf("over receive room 2: %v err: %v", p.room, err)
				pc = 1
				str := "EOF"
				p.pool <- &str
				return
			}
		}
	}

	tryTimes := 1
	flag := 0
	index := 0
Retry:
	err = p.init(flag, index)
	if err != nil {
		return
	}

	// flag == 0 尝试4次，flag == 1 重置 index
	index++
	if tryTimes == 4 {
		flag = 1
		index = 0
	}
	tryTimes++

	p.keepAlive()

	b := make([]byte, 512*1024)
	start := 0
	end := 0
	lastOffset := 0

	for !p.exit {
		start = 0
		end = 0
		n, err := p.conn.Read(b[lastOffset:])
		if err != nil {
			log.Printf("over receive room 3: %v err: %v", p.room, err)
			p.Close()
			// 计数
			if tryTimes <= 10 {
				time.Sleep(time.Second * 4)
				log.Printf("房间 %v 尝试重试，当前重试次数: %v", p.room, tryTimes)
				goto Retry
			}
			str := "EOF"
			p.pool <- &str
			return
		}
		end = lastOffset + n

		for {
			dealed := p.dealBuffer(b[start:end])
			start += dealed
			if dealed == 0 {
				copied := copy(b, b[start:end])
				lastOffset = copied
				break
			}
		}
	}
}

var typeStart = []byte(`{"type"`)

func (p *Panda) dealBuffer(buff []byte) int {
	l := len(buff)
	if l <= 4 {
		return 0 // no deal
	}

	if bytes.Equal(buff[:4], pandaReceiveMsg) {
		if l < 4+2 { // msg length + body not enough, wait next
			return 0
		}
		length := uint(buff[4]<<8) + uint(buff[5])
		pos := int(4 + 2 + length)
		if l < pos+4+pandaIgnoreByteLength {
			return 0
		}

		// msgLen := int((uint(buff[pos]) << 24) + (uint(buff[pos+1]) << 16) + (uint(buff[pos+2]) << 8) + uint(buff[pos+3]))
		msgLen := int(binary.BigEndian.Uint32(buff[pos:]))
		if l < pos+4+msgLen {
			return 0
		}
		pos += 4 + pandaIgnoreByteLength
		strBytes := buff[pos : pos+msgLen-pandaIgnoreByteLength]

		// 弹幕有时有bug，多条消息并在一起，需要拆开
		var n = 0
		for {
			n = bytes.LastIndex(strBytes, typeStart)
			if n == -1 {
				// for ugly string like
				// {"data":{"content":{"val":5819.173616,"c_lv":14,"c_lv_val":5180,"n_lv":15,"n_lv_val":6449},"to":{"toRoom":"66666"},"from":{}},"type":"212"}
				n = bytes.LastIndex(strBytes, []byte(`{"data"`))
				if n == -1 {
					fmt.Println("invalid string", string(strBytes))
					break
				}
			}
			str := string(strBytes[n:])
			p.pool <- &str
			if n == 0 {
				break
			}
			strBytes = strBytes[:n-pandaIgnoreByteLength]
		}

		return pos + msgLen - pandaIgnoreByteLength

	} else if bytes.Equal(buff[:4], pandaHeartbeatResponse) {
		// fmt.Println("heartbeat")
		return 4
	}
	fmt.Println(hex.EncodeToString(buff[:4]))
	return 4
}

func (p *Panda) keepAlive() {
	go func() {
		for {
			i, err := p.conn.Write(pandaHeartbeat)
			if err != nil {
				log.Printf("%v-房间 退出心跳 周期 %v, err: %v", p.room, KeepAliveInterval, err)
				return
			}
			log.Printf("%v-房间 发送心跳 周期 %v, i: %v, err: %v", p.room, KeepAliveInterval, i, err)
			time.Sleep(KeepAliveInterval)
		}
	}()
}

func (p *Panda) IsLive() (bool, error) {
	return true, nil
}

func (p *Panda) handshake(param *PandaChatParam) error {
	data := fmt.Sprintf("u:%d@%s\nk:1\nt:300\nts:%d\nsign:%s\nauthtype:%s", param.Rid, param.Appid, param.Ts, param.Sign, param.AuthType)
	l := len(data)

	msg := make([]byte, 4+2+l)
	copy(msg, pandaStart)
	binary.BigEndian.PutUint16(msg[4:], uint16(l))
	// copy(msg[4:], []byte{byte(l >> 8), byte(l & 0xff)})
	copy(msg[4+2:], []byte(data))
	_, err := p.conn.Write(msg)
	if err != nil {
		return err
	}
	buff := make([]byte, 6)
	n, err := p.conn.Read(buff)
	if err != nil {
		return err
	}
	if n != 6 || !bytes.Equal(buff[:4], pandaResponse) {
		return errors.New("response error")
	}

	length := int((uint(buff[4]) << 8) + uint(buff[5]))
	if length > 255 {
		return errors.New("invalid response length flag")
	}
	buff2 := make([]byte, length)
	p.conn.Read(buff2)
	return nil
}

func (p *Panda) getChatParam() (*PandaChatParam, error) {
	// u1 := fmt.Sprintf("http://www.panda.tv/ajax_chatinfo?roomid=%d&_=%d", p.room, time.Now().Unix()*1000)
	u1 := fmt.Sprintf("http://riven.panda.tv/chatroom/getinfo?roomid=%d", p.room)
	var chatData pandaChatData
	err := GetJson(u1, &chatData)
	if err != nil {
		return nil, err
	}
	u2 := fmt.Sprintf("http://api.homer.panda.tv/chatroom/getinfo?rid=%d&roomid=%d&retry=0&sign=%s&ts=%d&_=%d",
		chatData.Data.Rid, p.room, chatData.Data.Sign, chatData.Data.Ts, time.Now().Unix()*1000)
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

func (p *Panda) Close() {
	p.conn.Close()
}

func UniqueSliceString(slice *[]string) {
	found := make(map[string]bool)
	total := 0
	for i, val := range *slice {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*slice)[total] = (*slice)[i]
			total++
		}
	}

	*slice = (*slice)[:total]
}
