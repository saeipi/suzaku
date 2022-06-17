package config

var const_cfg = "server_ip: 0.0.0.0\n\ntsl:\n  cert_file: configs/tls/ca.pem\n  key_file: configs/tls/ca.key\n\netcd:\n  address: [ etcd:2379 ]\n  schema: suzaku\n  read_timeout: 5000\n  write_timeout: 5000\n  dial_timeout: 5000\n\nrpc_keepalive:\n  idle_timeout: 60000 #60s\n  force_close_wait: 20000 #20s\n  keep_alive_interval: 60000 #60s\n  keep_alive_timeout: 20000 #20s\n  max_life_time: 7200000 #2h\n\nrpc_port: #rpc服务端口 默认即可\n  user_port: [ 10100 ]\n  friend_port: [ 10200 ]\n  offline_message_port: [ 10300 ]\n  online_relay_port: [ 10400 ]\n  group_port: [ 10500 ]\n  auth_port: [ 10600 ]\n  push_port: [ 10700 ]\n  statistics_port: [ 10800 ]\n  message_cms_port: [ 10900 ]\n  admin_cms_port: [ 11000 ]\n  c2c:\n    callback_before_send_msg:\n      switch: false\n      timeout_strategy: 1\n    callback_after_send_msg:\n      switch: false\n  state:\n    state_change:\n      switch: false\n\nrpc_register_name: #rpc注册服务名，默认即可\n  user_name: User\n  friend_name: Friend\n  offline_message_name: OfflineMessage\n  push_name: Push\n  online_message_relay_name: OnlineMessageRelay\n  group_name: Group\n  auth_name: Auth\n  cache_name: Cache\n  statistics_name: Statistics\n  message_cms_name: MessageCMS\n  admin_cms_name: AdminCMS\n\nmysql:\n  address: [ mysql:3306 ]\n  username: root\n  password:\n  db: sdb\n  max_open_conn: 20\n  max_idle_conn: 10\n  conn_lifetime: 120000\n  charset: utf8\n\nmongo:\n  address: [ mongo:27017 ]\n  username: admin\n  password: 123456\n  db: suzaku\n  direct: false\n  timeout: 5000\n  max_pool_size: 20\n  retain_chat_records: 3650\n\nredis:\n  address: [ redis:6379 ]\n  db: 0\n  password: \"\"\n  prefix: \"SZK:\"\n\nkafka:\n  ws2mschat:\n    addr: [ kafka:9092 ]\n    topic: \"ws2ms_chat\"\n  ms2pschat:\n    addr: [ kafka:9092 ]\n    topic: \"ms2ps_chat\"\n  consumer_group_id:\n    msgToMongo: mongo\n    msgToMySql: mysql\n    msgToPush: push\n\njwt_auth:\n  auth_method: cookie\n  is_dev: true\n\ncasbin:\n  model-path: configs/casbin/rbac_model.conf\n\nenvironment:\n  run_model: dev\n\nmonlog:\n  batch_size: 100\n  commit_timeout: 1000\n  mongo:\n    address: [ mongo:27017 ]\n    username: admin\n    password: 123456\n    db: suzaku\n    direct: false\n    timeout: 5000\n    max_pool_size: 20\n    retain_chat_records: 3650\n\n# endpoints 内部组件间访问的端点host名称，访问时，可以内部直接访问 host:port 来访问\nendpoints:\n  api: suzaku_api\napi:\n  port: [ 10000 ]\n\nwebsocket:\n  port: [ 17778 ] # ws服务端口，默认即可，要开放此端口或做nginx转发\n  write_wait: 10000\n  pong_wait: 60000\n  max_message_size: 4096\n  read_buffer_size: 1024\n  write_buffer_size: 1024\n\ncredential: #腾讯cos，发送图片、视频、文件时需要，请自行申请后替换，必须修改\n  tencent:\n    app_id: 1302656840\n    region: ap-chengdu\n    bucket: echat-1302656840\n    secret_id: AKIDGNYVChzIQinu7QEgtNp0hnNgqcV8vZTC\n    secret_key: kz15vW83qM6dBUWIq681eBZA0c0vlIbe\n  minio: #MinIO 发送图片、视频、文件时需要，请自行申请后替换，必须修改。 客户端初始化时相应改动\n    bucket: suzaku\n    location: us-east-1\n    endpoint: http://127.0.0.1:9000\n    access_key: 17098899839\n    secret_key: 360001969\n\nlog:\n  storage_location: ../logs/\n  rotation_time: 24\n  rotation_count: 3 #日志数量\n  #日志级别 6表示全都打印，测试阶段建议设置为6\n  level: 6\n  es_address: [ 127.0.0.1:9201 ]\n  es_username: \"\"\n  es_password: \"\"\n  es_switch: false\n\nabnormal:\n  file: /var/log/suzaku/abnormal.log\n  level: 0\n\nzap:\n  encoder: console\n  directory: suzaku\n  show_line: true\n  encode_level: CapitalColor\n  stacktrace_key: stacktrace\n  log_stdout: true\n  segment:\n    maxsize: 500\n    maxage: 500\n    maxbackups: 500\n    localtime: true\n    compress: true\n\nsecret: saeipi\n\ncallback:\n  # callback url 需要自行更换callback url\n  callback_url : \"http://127.0.0.1:8080/callback\"\n  # 开启关闭操作前后回调的配置\n  callback_before_send_single_msg:\n    enable: false # 回调是否启用\n    callback_time_out: 2 # 回调超时时间\n    callback_failed_continue: true # 回调超时是否继续执行代码\n  callback_after_send_single_msg:\n    enable: false\n    callback_time_out: 2\n  callbackBeforeSendGroupMsg:\n    enable: false\n    callback_time_out: 2\n    callback_failed_continue: true\n  callback_after_send_group_msg:\n    enable: false\n    callback_time_out: 2\n  callback_word_filter:\n    enable: false\n    callback_time_out: 2\n    callback_failed_continue: true\n\n#ios系统推送声音以及标记计数\nios_push:\n  push_sound: \"xxx\"\n  badge_count: true"