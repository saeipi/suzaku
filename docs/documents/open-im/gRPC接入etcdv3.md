OpenIM服务发现和负载均衡golang插件：gRPC接入etcdv3

## etcd简介

etcd是CoreOS团队于2013年6月发起的开源项目，它的目标是构建一个高可用的分布式键值(key-value)数据库。etcd内部采用 `raft`  协议作为一致性算法，etcd基于Go语言实现。

etcd作为服务发现系统，有以下的特点：

- 简单：安装配置简单，而且提供了HTTP API进行交互，使用也很简单
- 安全：支持SSL证书验证
- 快速：根据官方提供的benchmark数据，单实例支持每秒2k+读操作
- 可靠：采用raft算法，实现分布式系统数据的可用性和一致性

etcd项目地址：[](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fcoreos%2Fetcd%2F)[https://github.com/coreos/etcd/](https://github.com/coreos/etcd/)

**etcd典型应用场景-服务发现**

etcd比较多的应用场景是用于服务发现，服务发现(Service Discovery)要解决的是分布式系统中最常见的问题之一，即在同一个分布式集群中的进程或服务如何才能找到对方并建立连接。

从本质上说，服务发现就是要了解集群中是否有进程在监听upd或者tcp端口，并且通过名字就可以进行查找和链接。

要解决服务发现的问题，需要下面三大支柱，缺一不可。

- 一个强一致性、高可用的服务存储目录。

基于Ralf算法的etcd天生就是这样一个强一致性、高可用的服务存储目录。

- 一种注册服务和健康服务健康状况的机制。

用户可以在etcd中注册服务，并且对注册的服务配置key TTL，定时保持服务的心跳以达到监控健康状态的效果。

- 一种查找和连接服务的机制。

通过在etcd指定的主题下注册的服务业能在对应的主题下查找到。为了确保连接，我们可以在每个服务机器上都部署一个proxy模式的etcd，这样就可以确保访问etcd集群的服务都能够互相连接。  
![1.png](http://forum.rentsoft.cn/storage/attachments/2021/08/26/isoh7gHppGwVwmdw2dUe3EKTeIGboMJtGnewFf4U_thumb.png)

## gRPC简介

gRPC是谷歌开源的一款跨平台、高性能的RPC框架。gRPC是一个现代的开源高性能RPC框架，可以在任何环境下运行。在实际开发过程中，主要使用它来进行后端微服务的开发。

在gRPC中，客户端应用程序可以像本地对象那样直接调用另一台计算机上的服务器应用程序上的方法，从而更容易创建分布式应用程序和服务。与许多RPC系统一样，gRPC基于定义服务的思想，可以通过设置参数和返回类型来远程调用方法。在服务端，实现这个接口并运行gRPC服务器来处理客户端调用。客户端提供的方法（客户端与服务端的方法相同）。

如图所示，gRPC客户端和服务端可以在各种环境中运行和相互通信，并且可以用gRPC支持的任何语言编写。因此，可以用Go语言创建一个gRPC服务器，同时供PHP客户端和Android客户端等多个客户端调用，从而突破开发语言的限制。  
![2.jpg](http://forum.rentsoft.cn/storage/attachments/2021/08/26/7cyEfsOmm21Nt5Gw3rM0Il8HAymeRR4SjVXXz599_thumb.jpg)

## 服务注册：register.go

\`\`

```
/etcdAddr separated by commas
func RegisterEtcd(schema, etcdAddr, myHost string, myPort int, serviceName string, ttl int) error {
   cli, err := clientv3.New(clientv3.Config{
      Endpoints: strings.Split(etcdAddr, ","),
   })
   fmt.Println("RegisterEtcd")
   if err != nil {
      //    return fmt.Errorf("grpclb: create clientv3 client failed: %v", err)
      return fmt.Errorf("create etcd clientv3 client failed, errmsg:%v, etcd addr:%s", err, etcdAddr)
   }

   //lease
   ctx, cancel := context.WithCancel(context.Background())
   resp, err := cli.Grant(ctx, int64(ttl))
   if err != nil {
      return fmt.Errorf("grant failed")
   }

   //  schema:///serviceName/ip:port ->ip:port
   serviceValue := net.JoinHostPort(myHost, strconv.Itoa(myPort))
   serviceKey := GetPrefix(schema, serviceName) + serviceValue

   //set key->value
   if _, err := cli.Put(ctx, serviceKey, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
      return fmt.Errorf("put failed, errmsg:%v， key:%s, value:%s", err, serviceKey, serviceValue)
   }

   //keepalive
   kresp, err := cli.KeepAlive(ctx, resp.ID)
   if err != nil {
      return fmt.Errorf("keepalive faild, errmsg:%v, lease id:%d", err, resp.ID)
   }

   go func() {
   FLOOP:
      for {
         select {
         case _, ok := <-kresp:
            if ok == true {
            } else {
               break FLOOP
            }
         }
      }
   }()

   rEtcd = &RegEtcd{ctx: ctx,
      cli:    cli,
      cancel: cancel,
      key:    serviceKey}

   return nil
}
```

grpc模块在启动时调用RegisterEtcd注册，并定时lease

## 命名解析实现及服务发现：resolver.go

\`\`

```
type Resolver struct {
   cc                 resolver.ClientConn
   serviceName        string
   grpcClientConn     *grpc.ClientConn
   cli                *clientv3.Client
   schema             string
   etcdAddr           string
   watchStartRevision int64
}

var (
   nameResolver        = make(map[string]*Resolver)
   rwNameResolverMutex sync.RWMutex
)

func NewResolver(schema, etcdAddr, serviceName string) (*Resolver, error) {
   etcdCli, err := clientv3.New(clientv3.Config{
      Endpoints: strings.Split(etcdAddr, ","),
   })
   if err != nil {
      return nil, err
   }

   var r Resolver
   r.serviceName = serviceName
   r.cli = etcdCli
   r.schema = schema
   r.etcdAddr = etcdAddr
   resolver.Register(&r)

   conn, err := grpc.Dial(
      GetPrefix(schema, serviceName),
      grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
      grpc.WithInsecure(),
      grpc.WithTimeout(time.Duration(5)*time.Second),
   )
   if err == nil {
      r.grpcClientConn = conn
   }
   return &r, err
}

func (r1 *Resolver) ResolveNow(rn resolver.ResolveNowOptions) {
}

func (r1 *Resolver) Close() {
}

func GetConn(schema, etcdaddr, serviceName string) *grpc.ClientConn {
   rwNameResolverMutex.RLock()
   r, ok := nameResolver[schema+serviceName]
   rwNameResolverMutex.RUnlock()
   if ok {
      return r.grpcClientConn
   }

   rwNameResolverMutex.Lock()
   r, ok = nameResolver[schema+serviceName]
   if ok {
      rwNameResolverMutex.Unlock()
      return r.grpcClientConn
   }

   r, err := NewResolver(schema, etcdaddr, serviceName)
   if err != nil {
      rwNameResolverMutex.Unlock()
      return nil
   }

   nameResolver[schema+serviceName] = r
   rwNameResolverMutex.Unlock()
   return r.grpcClientConn
}

func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
   if r.cli == nil {
      return nil, fmt.Errorf("etcd clientv3 client failed, etcd:%s", target)
   }
   r.cc = cc

   ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
   //     "%s:///%s"
   prefix := GetPrefix(r.schema, r.serviceName)
   // get key first
   resp, err := r.cli.Get(ctx, prefix, clientv3.WithPrefix())
   if err == nil {
      var addrList []resolver.Address
      for i := range resp.Kvs {
         fmt.Println("init addr: ", string(resp.Kvs[i].Value))
         addrList = append(addrList, resolver.Address{Addr: string(resp.Kvs[i].Value)})
      }
      r.cc.UpdateState(resolver.State{Addresses: addrList})
      r.watchStartRevision = resp.Header.Revision + 1
      go r.watch(prefix, addrList)
   } else {
      return nil, fmt.Errorf("etcd get failed, prefix: %s", prefix)
   }

   return r, nil
}

func (r *Resolver) Scheme() string {
   return r.schema
}

func exists(addrList []resolver.Address, addr string) bool {
   for _, v := range addrList {
      if v.Addr == addr {
         return true
      }
   }
   return false
}

func remove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
   for i := range s {
      if s[i].Addr == addr {
         s[i] = s[len(s)-1]
         return s[:len(s)-1], true
      }
   }
   return nil, false
}

func (r *Resolver) watch(prefix string, addrList []resolver.Address) {
   rch := r.cli.Watch(context.Background(), prefix, clientv3.WithPrefix(), clientv3.WithPrefix())
   for n := range rch {
      flag := 0
      for _, ev := range n.Events {
         switch ev.Type {
         case mvccpb.PUT:
            if !exists(addrList, string(ev.Kv.Value)) {
               flag = 1
               addrList = append(addrList, resolver.Address{Addr: string(ev.Kv.Value)})
               fmt.Println("after add, new list: ", addrList)
            }
         case mvccpb.DELETE:
            fmt.Println("remove addr key: ", string(ev.Kv.Key), "value:", string(ev.Kv.Value))
            i := strings.LastIndexAny(string(ev.Kv.Key), "/")
            if i < 0 {
               return
            }
            t := string(ev.Kv.Key)[i+1:]
            fmt.Println("remove addr key: ", string(ev.Kv.Key), "value:", string(ev.Kv.Value), "addr:", t)
            if s, ok := remove(addrList, t); ok {
               flag = 1
               addrList = s
               fmt.Println("after remove, new list: ", addrList)
            }
         }
      }

      if flag == 1 {
         r.cc.UpdateState(resolver.State{Addresses: addrList})
         fmt.Println("update: ", addrList)
      }
   }
}
```

客户端先通过GetConn获取conn，然后再调用grpc服务，调用后不用关闭conn

## 服务端示例代码：server.go

\`\`

```
getcdv3.RegisterEtcd ("sk", etcdAddr, "127.0.0.1", port, "myrpc1", 10)
s := grpc.NewServer()
helloworld.RegisterHelloServer(s, &server{})
s.Serve(listener)
```

## 客户端示例代码：client.go

\`\`

```
p := getcdv3.GetConn("sk", etcdAddr, "myrpc1")
client := helloworld.NewHelloClient(p)
resp1, err := client.SayHello(context.Background(), &helloworld.HelloReq{Req: "world"})
```

总结：OpenIM集成此插件，实现了轻量级的服务发现机制，打造了基于集群的IM服务，各模块很方便平行扩展，方便运维。在使用grpc、etcd过程中，特别注意版本兼容问题，具体可以参考OpenIM的go.mod文件

![](http://forum.rentsoft.cn/storage/attachments/2021/08/26/isoh7gHppGwVwmdw2dUe3EKTeIGboMJtGnewFf4U.png?imageMogr2/format/webp/quality/80/interlace/1/ignore-error/1)

![](http://forum.rentsoft.cn/storage/attachments/2021/08/26/7cyEfsOmm21Nt5Gw3rM0Il8HAymeRR4SjVXXz599.jpg?imageMogr2/format/webp/quality/80/interlace/1/ignore-error/1)
