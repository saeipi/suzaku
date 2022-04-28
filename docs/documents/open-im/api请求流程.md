#### 一、登录后的api请求
```
1、获取个人信息
 "/user/get_self_user_info"
2、获取自己发起的添加好友申请列表
"/friend/get_self_friend_apply_list"
3、获取好友列表
"/friend/get_friend_list"
4、获取自己发起的加入群申请
"/group/get_user_req_group_applicationList"
5、获取他人发起的添加好友申请列表
"/friend/get_friend_apply_list"
6、群主或管理员收到的加入群申请
"/group/get_recv_group_applicationList"
7、获取已加入群的列表
"/group/get_joined_group_list"
8、？？？
"/conversation/get_all_conversations"

9、获取好友在线状态
"/user/get_users_online_status"

10、聊天上传文件
"/third/minio_storage_credential"
```

#### 二、创建群后的api请求
```
1、创建群
"/group/create_group"
2、获取加入的群列表
"/group/get_joined_group_list"
3、获取当前群成员
 "/group/get_group_all_member_list"
```