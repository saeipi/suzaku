package client

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"suzaku/internal/msg_gateway/protocol"
	"suzaku/internal/msg_gateway/ws_server"
	"suzaku/pkg/constant"
	"suzaku/pkg/proto/pb_ws"
	"sync"
	"time"
)

type Client struct {
	conn *websocket.Conn
	// 用户ID
	userID string
	// 平台ID
	platformID int32
	// 上线时间戳（毫秒）
	onlineAt int64
	// Buffered channel of outbound messages.
	send chan []byte
	// 关闭通知
	close    chan []byte
	closed   bool
	nickname string
	sync.Mutex
}

func NewClient(userID string) (client *Client) {
	var (
		u    url.URL
		q    url.Values
		ts   int64
		conn *websocket.Conn
		resp *http.Response
		buf  []byte
		err  error
	)
	ts = time.Now().Unix()
	u = url.URL{Scheme: "ws", Host: "localhost:17778", Path: "/"}
	q = u.Query()
	q.Set("user_id", userID)
	q.Set("platform_id", "1")
	u.RawQuery = q.Encode()

	client = &Client{
		conn:       nil,
		userID:     userID,
		platformID: 1,
		onlineAt:   ts,
		send:       make(chan []byte, 100),
		close:      make(chan []byte),
		closed:     false,
	}
	conn, resp, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if resp != nil {
		defer resp.Body.Close()
		buf, err = ioutil.ReadAll(resp.Body)
		fmt.Println(string(buf))
	}
	if err != nil {
		return
	}
	client.conn = conn
	go client.write()
	go client.read()
	return
}

func (c *Client) closeConn() {
	c.Lock()
	if c.closed {
		c.Unlock()
		return
	}
	c.closed = true
	c.Unlock()

	c.conn.Close()
	close(c.send)
	close(c.close)
}

func (c *Client) read() {
	defer func() {
		c.closeConn()
	}()

	var (
		msgType int
		bufMsg  []byte
		err     error
	)
	for {
		if msgType, bufMsg, err = c.conn.ReadMessage(); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				//TODO:需要添加日志
				log.Printf("error: %v", err)
			}
			break
		}
		if msgType == websocket.PingMessage {
			continue
		}
		if len(bufMsg) == 0 {
			continue
		}
		go c.messageHandler(bufMsg)
	}
}

func (c *Client) write() {
	pingTicker := time.NewTicker(ws_server.WsPingPeriod)
	defer func() {
		pingTicker.Stop()
		c.closeConn()
	}()

	var (
		err     error
		message []byte
		ok      bool
	)
	c.conn.SetReadLimit(ws_server.WsMaxMessageSize)
	for {
		select {
		case message, ok = <-c.send:
			if ok == false {
				// chan 关闭
				return
			}
			if err = c.conn.SetWriteDeadline(time.Now().Add(ws_server.WsWriteWait)); err != nil {
				c.conn.WriteMessage(websocket.CloseMessage, ws_server.WsMsgBufClose)
				return
			}
			if err = c.conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
				return
			}
		case <-pingTicker.C:
			c.conn.SetWriteDeadline(time.Now().Add(ws_server.WsWriteWait))
			if err = c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case <-c.close:
			return
		}
	}
}

func (c *Client) messageHandler(message []byte) {
	var (
		buffer  *bytes.Buffer
		decoder *gob.Decoder
		req     protocol.MessageReq
		body    pb_ws.MsgData
		err     error
	)

	req = protocol.MessageReq{}
	buffer = bytes.NewBuffer(message)
	decoder = gob.NewDecoder(buffer)
	err = decoder.Decode(&req)
	if err != nil {
		fmt.Println("解析消息错误")
		return
	}
	if req.ReqIdentifier == 0 {
		return
	}
	if req.Data != nil {
		err = proto.Unmarshal(req.Data, &body)
		if err != nil {
			fmt.Println("解析消息本体错误")
			return
		}
	}
	fmt.Println("收到消息:", c.userID, req.ReqIdentifier, req.OperationID, req.SendID, req.Token)
	c.SendUser(req.SendID)
}

func (c *Client) SendUser(recvId string) (err error) {
	var (
		ts         int64
		contentMap map[string]interface{}
		bodyBytes  []byte
		req        protocol.MessageReq
		reqBytes   []byte
		msgData    pb_ws.MsgData
	)

	ts = time.Now().Unix()
	contentMap = map[string]interface{}{}
	contentMap["content"] = "文本聊天消息"

	msgData = pb_ws.MsgData{
		SendId:           c.userID, // 发送者ID
		RecvId:           recvId,   // 接收者ID
		GroupId:          "",
		ClientMsgId:      strconv.Itoa(int(ts)),
		ServerMsgId:      "",
		SenderPlatformId: 1,
		SenderNickname:   c.nickname,
		SenderFaceUrl:    "https://github.com/saeipi/suzaku/blob/main/assets/images/suzaku.jpg",
		SessionType:      1, // 单聊为1，群聊为2
		MsgFrom:          1,
		ContentType:      101,                  // 消息类型，101表示文本，102表示图片
		Content:          c.toByte(contentMap), // 内部是json 对象
		Seq:              1,
		SendTime:         ts,
		CreateTime:       ts,
		Status:           0,
		Options:          nil,
		OfflinePushInfo:  nil, // |否| 离线推送的具体内容，如果不填写，使用服务器默认推送标题
	}
	bodyBytes, _ = proto.Marshal(&msgData)
	req = protocol.MessageReq{
		ReqIdentifier: constant.WSSendMsg,
		Token:         strconv.Itoa(int(ts)) + ":" + c.userID,
		SendID:        c.userID,
		OperationID:   strconv.Itoa(int(ts)) + ":" + c.userID,
		MsgIncr:       strconv.Itoa(int(ts)) + ":" + c.userID,
		Data:          bodyBytes,
	}
	reqBytes = c.toByte(req)
	time.Sleep(time.Second * 2)
	c.Send(reqBytes)
	return
}

func (c *Client) toByte(data interface{}) (buf []byte) {
	var (
		dataBytes bytes.Buffer
		encoder   *gob.Encoder
		err       error
	)
	encoder = gob.NewEncoder(&dataBytes)
	err = encoder.Encode(data)
	if err != nil {
		return
	}
	return dataBytes.Bytes()
}

func (c *Client) Send(message []byte) {
	if c.closed {
		return
	}
	c.send <- message
}

func (c *Client) SendMessage(message []byte) (err error) {
	if c.closed {
		return
	}
	if err = c.conn.SetWriteDeadline(time.Now().Add(ws_server.WsWriteWait)); err != nil {
		c.conn.WriteMessage(websocket.CloseMessage, ws_server.WsMsgBufClose)
		return
	}
	if err = c.conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
		return
	}
	return
}

func (c *Client) Close() {
	if c.closed {
		return
	}
	c.close <- ws_server.WsMsgBufClose
}
