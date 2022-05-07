package client

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"suzaku/internal/msg_gateway/protocol"
	"suzaku/internal/msg_gateway/ws_server"
	"suzaku/pkg/constant"
	"suzaku/pkg/proto/pb_ws"
	"suzaku/pkg/utils"
	"sync"
	"time"
)

type Client struct {
	rwLock sync.RWMutex
	mgr    *Manager
	conn   *websocket.Conn
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
	curCount int
	endCount int
}

func NewClient(userID string, mgr *Manager) (client *Client) {
	var (
		u    url.URL
		q    url.Values
		ts   int64
		conn *websocket.Conn
		// resp *http.Response
		// buf []byte
		err error
	)
	ts = time.Now().Unix()
	// localhost:17778
	u = url.URL{Scheme: "ws", Host: "10.0.115.108:17778", Path: "/"}
	q = u.Query()
	q.Set("user_id", userID)
	q.Set("platform_id", "1")
	u.RawQuery = q.Encode()

	client = &Client{
		mgr:        mgr,
		conn:       nil,
		userID:     userID,
		platformID: 1,
		onlineAt:   ts,
		send:       make(chan []byte, 100),
		close:      make(chan []byte),
		closed:     false,
		nickname:   userID,
		curCount:   0,
		endCount:   10000, // rand.Intn(100-10) + 10
	}

	conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("创建连接失败")
		return
	}
	client.conn = conn
	go client.write()
	go client.read()
	return
}

func (c *Client) closeConn() {
	fmt.Println("客户端断开连接")
	c.rwLock.Lock()
	if c.closed {
		c.rwLock.Unlock()
		return
	}
	c.closed = true
	close(c.send)
	close(c.close)
	c.rwLock.Unlock()

	c.conn.Close()
	c.mgr.unregister <- c
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

var msgCount = 0

func (c *Client) messageHandler(message []byte) {
	if c.curCount > c.endCount {
		c.Close()
		return
	}
	c.curCount++
	var (
		req     protocol.MessageReq
		msgData pb_ws.MsgData
		err     error
	)

	req = protocol.MessageReq{}
	err = utils.BufferDecode(message, &req)
	if err != nil {
		fmt.Println("解析消息错误")
		return
	}
	if req.ReqIdentifier == 0 {
		return
	}
	if req.Data == nil {
		return
	}
	msgData = pb_ws.MsgData{}
	err = proto.Unmarshal(req.Data, &msgData)
	if err != nil {
		return
	}
	if msgData.MsgFrom == 0 {
		return
	}
	msgCount++
	fmt.Println("收到消息:", c.userID, req.SendID, string(msgData.Content), msgCount)
}

func (c *Client) SendGroup(groupId string) (err error) {
	var (
		ts        int64
		bodyBytes []byte
		req       protocol.MessageReq
		reqBytes  []byte
		msgData   pb_ws.MsgData
	)
	if c.conn == nil {
		return
	}
	if c.closed == true {
		return
	}
	ts = utils.GetCurrentTimestampByMill()
	msgData = pb_ws.MsgData{
		SendId:           c.userID, // 发送者ID
		RecvId:           "",       // 接收者ID
		GroupId:          groupId,
		ClientMsgId:      utils.GetMsgID(c.userID),
		ServerMsgId:      "",
		SenderPlatformId: 1,
		SenderNickname:   c.nickname,
		SenderAvatarUrl:  "https://github.com/saeipi/suzaku/blob/main/assets/images/suzaku.jpg",
		SessionType:      2, // 单聊为1，群聊为2
		MsgFrom:          int32(rand.Uint64()) + 1,
		ContentType:      101, // 消息类型，101表示文本，102表示图片
		Content:          nil, // 内部是json 对象
		Seq:              1,
		CreatedTs:        ts,
		Status:           0,
		Options:          nil,
		OfflinePushInfo:  nil, // |否| 离线推送的具体内容，如果不填写，使用服务器默认推送标题
	}
	msgData.Content = utils.Str2Bytes("群文本聊天消息" + c.userID)
	bodyBytes, err = proto.Marshal(&msgData)
	if err != nil {
		return
	}
	req = protocol.MessageReq{
		ReqIdentifier: constant.WSSendMsg,
		Token:         strconv.Itoa(int(ts)) + ":" + c.userID,
		SendID:        c.userID,
		OperationID:   c.userID,
		MsgIncr:       utils.GenMsgIncr(c.userID),
		Data:          bodyBytes,
	}
	reqBytes, err = utils.ObjEncode(req)
	if err != nil {
		return
	}
	c.Send(reqBytes)
	return
}

func (c *Client) Send(message []byte) {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if c.closed {
		return
	}
	c.send <- message
}

func (c *Client) Close() {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if c.closed {
		return
	}
	c.close <- ws_server.WsMsgBufClose
}
