## 产品功能更新日志

**[1.0版本]包含功能如下**

<font face="微软雅黑" size=4 color=#32CD32>多端登录</font>

| **多端登录类型**                                             |
| ------------------------------------------------------------ |
| 只允许一端登录，Windows、Web、Android、iOS 彼此互踢          |
| 桌面PC 与 Web 端互踢、移动 Android 和 iOS 端互踢、桌面与移动端可同时登录 |
| 各端均可以同时登录在线                                       |



<font face="微软雅黑" size=4 color=#32CD32>消息类型</font>

| **消息类型** | **备注**                                                     |
| ------------ | ------------------------------------------------------------ |
| 文本消息     | 消息内容为普通文本                                           |
| 图片消息     | 消息内容为图片 URL 地址、尺寸、图片大小等信息，包括原图、大图、缩略图 |
| 语音消息     | 消息内容为语音文件的 URL 地址、时长、大小、格式等信息；      |
| 视频消息     | 消息内容为视频文件的 URL 地址、时长、大小、格式等信息；      |
| Tips 消息    | 包括群通知，例如例如有成员进出群组，群的描述信息被修改，群成员的资料发生变化等；也包括好友资料修改通知； |
| 文件消息     | 消息内容为文件的 URL 地址、大小、格式等信息，格式不限        |
| 地理位置消息 | 消息内容为地理位置标题、经度、纬度信息                       |
| 自定义消息   | 开发者任何自定义的消息类型                                   |



<font face="微软雅黑" size=4 color=#32CD32>消息功能</font>

| 功能类型 | 功能描述                                                     |
| :------- | :----------------------------------------------------------- |
| 离线消息 | 用户登录后退到后台，当有用户给其发消息时，即时通信 IM 支持离线推送 |
| 漫游消息 | 在新设备登录时，将服务器记录(云端)的历史消息存储进行同步，默认保存14天，开发者可以在配置文件中设置 |
| 多端同步 | 多终端消息同步，可同时收到消息                               |
| 历史消息 | 支持本地历史消息和云端历史消息                               |
| 消息撤回 | 撤回投递成功的消息，默认撤回 2 分钟内的消息，时间可以自行修改。撤回操作仅支持单聊和群聊消息 |
| 已读回执 | 查看单聊会话中对方的已读未读状态                             |
| @功能    | @ 的人在收到消息时，需要在 UI 上做特殊展示和提醒             |
| 正在输入 | 对方在输入框输入时，展示在本端                               |
| 消息删除 | 使用消息的 remove 方法可以在本地删除消息                     |

 

<font face="微软雅黑" size=4 color=#32CD32>用户资料托管</font>

| **功能**           | **功能描述**                                               |
| ------------------ | ---------------------------------------------------------- |
| 设置用户资料       | 用户设置自己的昵称、验证方式、头像、性别、年龄、位置等资料 |
| 获取用户资料       | 用户查看自己、好友及陌生人资料                             |
| 按字段获取用户资料 | 按照特定字段获取用户资料                                   |



<font face="微软雅黑" size=4 color=#32CD32>用户关系托管</font>

| **功能**         | **功能描述**                                                 |
| ---------------- | ------------------------------------------------------------ |
| 查找好友         | 可通过用户帐号 ID 查找好友                                   |
| 申请添加好友     | 要选择默认是否需要申请理由，目前是默认不需要                 |
| 添加好友         | 发送添加好友请求                                             |
| 删除好友         | 成为好友后可以删除好友                                       |
| 获取所有好友     | 获取所有好友的详细信息                                       |
| 同意/拒绝好友    | 收到请求加好友请求的系统通知后，可以通过或者拒绝，处理后申请人能收到通知 |
| 添加用户到黑名单 | 把用户拉黑，如果此前是好友关系不解除好友关系，通过字段来标识为黑名单 |
| 移除黑名单       | 把用户从黑名单中移除                                         |

 

<font face="微软雅黑" size=4 color=#32CD32>群组</font>

| **功能**       | **功能描述**                                                 |
| -------------- | ------------------------------------------------------------ |
| 角色           | 可通过用户帐号 ID 查找好友                                   |
| 群资料修改     | 群主可以修改群头像、群名、群公告                             |
| 申请加群       | 任何人可以提交加群申请                                       |
| 加群审批       | 群主审批或拒绝加群申请，申请人能收到审批结果                 |
| 邀请加群       | 群成员可以邀请加群，无需经过被邀请人和群主同意               |
| 群主退群       | 群主退群前，需先转让群主                                     |
| 移出成员       | 群主把成员从群中踢出                                         |
| 转让群主       | 群主转让给新群主                                             |
| 成员变更通知   | 所有成员都能收到通知，包括成员进群、退群、邀请进群、成员被踢等 |
| 群资料变更通知 | 所有成员都能收到通知，包括群头像、群名、群公告被修改         |



 <font face="微软雅黑" size=4 color=#32CD32>数据统计</font>

| **统计项** | **统计项说明**           |
| ---------- | ------------------------ |
| 新增用户数 | 每日注册 ID 的数量       |
| 活跃用户数 | 登录用户数               |
| 累计用户数 | 截止昨日所有注册用户数   |
| 单聊消息数 | 包括每日新增单聊消息数量 |

 
