
#### 发送单聊群聊消息
##### 1、请求参数
| 参数名              | 类型     | 必选 | 说明                                                                                                    |
|------------------|--------|----|-------------------------------------------------------------------------------------------------------|
| operationID      | string | 是  | 操作ID，保持唯一，建议用当前时间微秒+随机数                                                                               |
| sendID           | string | 是  | 发送者ID                                                                                                 |
| recvID           | string | 否  | 接收者ID，单聊为用户ID，如果是群聊，则不填                                                                               |
| groupID          | string | 否  | 群聊ID，如果为单聊，则不填                                                                                        |
| senderNickname   | string | 否  | 发送者昵称                                                                                                 |
| senderFaceURL    | string | 否  | 发送者头像                                                                                                 |
| senderPlatformID | int    | 否  | 发送者平台号，模拟用户发送时填写， 1-&gt;IOS,2-&gt;Android,3-&gt;Windows,4-&gt;OSX,5-&gt;Web,5-&gt;MiniWeb,7-&gt;Linux |
| forceList        | json数组 | 否  | 当聊天类型为群聊时，使用@指定强推用户userID列表                                                                           |
| content          | json对象 | 是  | 消息的具体内容，内部是json 对象，其他消息的详细字段请参考消息类型格式描述文档                                                             |
| contentType      | int    | 是  | 消息类型，101表示文本，102表示图片..详细参考消息类型格式描述文档                                                                  |
| sessionType      | int    | 是  | 发送的消息是单聊还是群聊,单聊为1，群聊为2                                                                                |
| isOnlineOnly     | bool   | 否  | 改字段设置为true时候，发送的消息服务器不会存储，接收者在线才会收到并存储到本地，不在线该消息丢失，当消息的类型为113-&gt;typing时候，接收者本地也不会做存储                |
| offlinePushInfo  | json对象 | 否  | 离线推送的具体内容，如果不填写，使用服务器默认推送标题                                                                           |
| title            | string | 否  | 推送的标题                                                                                                 |
| desc             | string | 否  | 推送的具体描述                                                                                               |
| ex               | string | 否  | 扩展字段                                                                                                  |
| iOSPushSound     | string | 否  | IOS的推送声音                                                                                              |
| iOSBadgeCount    | bool   | 否  | IOS推送消息是否计入桌面图标未读数                                                                                    |

##### 2、返回参数
| 参数名         | 类型     | 说明                       |
|-------------|--------|--------------------------|
| errCode     | int    | 0成功，非0失败                 |
| errMsg      | string | 错误信息                     |
| sendTime    | int    | 消息发送的具体时间，具体为毫秒的时间戳      |
| serverMsgID | string | 服务器生成的消息的唯一ID            |
| clientMsgID | string | 客户端生成的消息唯一ID，默认情况使用这个为主键 |

#### 管理员发送通知类型消息
##### 1、请求参数
| 参数名                                           | 类型     | 必选 | 说明                                                                                                    |
|-----------------------------------------------|--------|----|-------------------------------------------------------------------------------------------------------|
| operationID                                   | string | 是  | 操作ID，保持唯一，建议用当前时间微秒+随机数，用于后台链路追踪问题使用                                                                  |
| sendID                                        | string | 是  | 管理员ID，为后台config文件中配置的管理员ID中一个，默认openIM123456                                                          |
| recvID                                        | string | 是  | 接收者userID                                                                                             |
| senderPlatformID                              | int    | 否  | 发送者平台号，模拟用户发送时填写， 1-&gt;IOS,2-&gt;Android,3-&gt;Windows,4-&gt;OSX,5-&gt;Web,5-&gt;MiniWeb,7-&gt;Linux |
| senderFaceURL                                 | string | 否  | 发送者头像，用于客户端通知会话产生                                                                                     |
| senderNickname                                | string | 是  | 发送者昵称，用于客户端通知会话产生                                                                                     |
| content                                       | object | 是  | 消息的具体内容，内部是json 对象                                                                                    |
| notificationName                              | string | 是  | 通知标题                                                                                                  |
| notificationFaceURL                           | string | 是  | 通知头像                                                                                                  |
| notificationType                              | int    | 是  | 通知类型，如：1代表入职通知，2代表离职通知                                                                                |
| text                                          | string | 是  | 通知正文e                                                                                                 |
| externalUrl                                   | string | 否  | 通知点击后需要跳转到的地址链接(不填则无需跳转)                                                                              |
| mixType                                       | int    | 是  | 通知混合类型0：纯文字通知1：文字+图片通知2：文字+视频通知3：文字+文件通知4:   文字+语音通知5:   文字+语音+图片通知                                   |
| pictureElem                                   | object | 否  | 图片元素对象                                                                                                |
| sourcePicture                                 | object | 否  | 原图                                                                                                    |
| bigPicture                                    | object | 否  | 大图                                                                                                    |
| snapshotPicture                               | object | 否  | 缩略图                                                                                                   |
| soundElem                                     | object | 否  | 声音元素对象                                                                                                |
| videoElem                                     | object | 否  | 视频元素对象                                                                                                |
| fileElem                                      | object | 否  | 文件元素对象                                                                                                |
| uuid                                          | string | 否  | 对象唯一ID用于缓存使用                                                                                          |
| type/videoType/                               | string | 否  | 图片类型/视频类型                                                                                             |
| size/dataSize/videoSize/snapshotSize/fileSize | int    | 否  | 多媒体文件大小，单位字节                                                                                          |
| width/snapshotWidth                           | int    | 否  | 图片/视频缩略图宽度                                                                                            |
| height/snapshotHeight                         | int    | 否  | 图片/视频缩略图高度                                                                                            |
| url/sourceUrl/videoUrl                        | string | 否  | 图片/文件/视频的URL                                                                                          |
| sourcePath/soundPath/videoPath/filePath       | string | 否  | 文件路径，可不填写                                                                                             |
| fileName                                      | string | 否  | 文件名字                                                                                                  |
| ex                                            | string | 否  | 扩展字段                                                                                                  |
| contentType                                   | int    | 是  | 消息类型固定为1400                                                                                           |
| sessionType                                   | int    | 是  | 通知会话类型固定为4                                                                                            |
| isOnlineOnly                                  | bool   | 否  | 改字段设置为true时候，发送的消息服务器不会存储，接收者在线才会收到，不在线该消息丢失。                                                         |
| offlinePushInfo                               | object | 否  | 离线推送的具体内容，如果不填写，使用服务器默认推送标题                                                                           |
| title                                         | string | 否  | 推送的标题                                                                                                 |
| desc                                          | string | 否  | 推送的具体描述                                                                                               |
| ex                                            | string | 否  | 扩展字段                                                                                                  |
| iOSPushSound                                  | string | 否  | IOS的推送声音                                                                                              |
| iOSBadgeCount                                 | bool   | 否  | IOS推送消息是否计入桌面图标未读数                                                                                    |

##### 2、返回参数
| 参数名         | 类型     | 说明                       |
|-------------|--------|--------------------------|
| errCode     | int    | 0成功，非0失败                 |
| errMsg      | string | 错误信息                     |
| sendTime    | int    | 消息发送的具体时间，具体为毫秒的时间戳      |
| serverMsgID | string | 服务器生成的消息的唯一ID            |
| clientMsgID | string | 客户端生成的消息唯一ID，默认情况使用这个为主键 |

#### 消息类型格式描述
##### ContentType消息类型说明
| ContentType值 | 类型说明      |
|--------------|-----------|
| 101          | 文本消息      |
| 102          | 图片消息      |
| 103          | 音频消息      |
| 104          | 视频消息      |
| 105          | 文件消息      |
| 106          | 群聊中的@类型消息 |
| 107          | 合并转发类型消息  |
| 108          | 名片消息      |
| 109          | 地理位置类型消息  |
| 110          | 自定义消息     |
| 111          | 撤回类型消息    |
| 112          | 已读回执类型消息  |
| 114          | 引用类型消息    |

##### Content具体内容
###### 文本消息
| 参数名  | 必选 | 类型     | 说明        |
|------|----|--------|-----------|
| text | 是  | string | 文本消息的具体内容 |
###### 自定义消息
| 参数名         | 类型          | 必选 | 说明                             |
|-------------|-------------|----|--------------------------------|
| data        | json string | 是  | 用户自定义的消息为json对象转换后的string      |
| description | json string | 否  | 扩展的描述信息为json对象转换后的string，可以不使用 |
| extension   | json string | 否  | 扩展字段，暂时不使用                     |

##### OpenIM字段说明
| 参数名          | 类型     | 最大字符串长度限制               | 说明                                                       | 取值范围  |
|--------------|--------|-------------------------|----------------------------------------------------------|-------|
| secret       | string | 32                      | OpenIM秘钥，服务端配置文件config.yaml的secret字段，注意安全保存              | 字符串即可 |
| platform     | int    | 用户登录或注册的平台类型            | iOS 1, Android 2, Windows 3, OSX 4, WEB 5, 小程序 6，linux 7 |
| userID       | string | 64                      | 用户 ID，必须保证IM内唯一                                          | 字符串即可 |
| nickname     | string | 255                     | 用户昵称或者群昵称                                                | 字符串即可 |
| faceURL      | string | 255                     | 用户头像或者群头像url，根据上下文理解                                     |
| gender       | int    | 用户性别                    | 1 表示男，2 表示女                                              |
| phoneNumber  | string | 32                      | 用户手机号码，包括地区，(如香港：+852-xxxxxxxx)，                         |
| birth        | uint32 | 用户生日，Unix时间戳（秒）         |
| email        | string | 64                      | 邮箱地址                                                     |
| ex           | string | 1024                    | 扩展字段，用户可自行扩展，建议封装成 JSON 字符串                              |
| operationID  | string | 操作ID，保持唯一，建议用当前时间微秒+随机数 |
| expiredTime  | int    | 过期时间，单位（秒）              |
| roleLevel    | int    | 群内成员类型                  | 1普通成员，2群主，3管理员                                           |
| groupType    | int    | 群类型                     | 目前统一填0                                                   |
| ownerUserID  | string | 64                      | 群主UserID                                                 |
| groupName    | string | 255                     | 群名称                                                      |
| notification | string | 255                     | 群公告                                                      |
| introduction | string | 255                     | 群介绍                                                      |
| memberList   | json数组 | 成员列表                    |
| reason       | string | 64                      | 原因，比如踢人等原因                                               |
| token        | string | 调用api时设置到请求header中      |
| userIDList   | json数组 | 用户的userID列表             |
