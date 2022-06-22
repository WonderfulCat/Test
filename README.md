# Test

## 一些注意点

/storeItem itemId itemNum index 命令未写入itemType

/clearup   按类型和数量排序无法实现,***所有使用的是itemId及itemNum进行的排序***.


## 目录结构: 
```
│  docker-compose.yml    #compose 
│  Dockerfile            #dockerfile **只有server端, client没有dockerfile**
├─config                 #itemdata配置文件
└─src                    
    ├─client             #用客户端(内有打包好的可执行文件)
    ├─test_common        #全局变量
    ├─test_constant      #全局常量
    ├─test_impl          #逻辑实现
    ├─test_interface     #逻辑接口
    ├─test_model         #通信数据定义
    ├─test_net           #网络部分
    │  ├─net_impl        #网络部分实现
    │  └─net_interface   #网络部分接口
    ├─test_pb            #item PB文件
    ├─test_service       #对外接口
    └─utils              #测试
```

## 外部通信
**如果使用第三方客户端需遵循数据包规则**

**数据包格式:**
|uint32|uint32|[]byte|
|-|-|-|
|包长|CODE|DATA| 

**CODE:**
|code|command|
|-|-|
|1000|/login|
|1001|/whichAlliance|
|1002|/createAlliance|
|1003|/joinAlliance|
|1004|/dismissAlliance|
|1005|/increaseCapacity|
|1006|/storeItem|
|1007|/destoryItem|
|1008|/clearUp|
|1009|/allianceList|
|1010|/getItemList|  
#1010为额外增加命令,返回所有仓库物品信息 

**DATA:**
```
类似如下:
   type LoginRequestInfo struct {
	        Name string `json:"name"`
	        Pswd string `json:"pswd"`
    }
使用json序列化为[]byte后存入数据包DATA中.
所有数据结构体定义在 : /src/test_model/test_model.go文件中. 
```

## 内部通讯
```         
创建连接-读取: accept -> 建立连接  -> 读取 -> 放入工作队列           
逻辑处理: 工作队列 -> 处理逻辑 -> 放入写入队列
消息返回:  写入队列 -> 发送消息

1. 每次connection都会创建2个协程分别用于读写操作.
2. 工作队列没有做协程池,只有一个channel做循环逻辑处理.
```

## 性能测试
```
整理仓库算法:
```
|0|1|2|3|...|
|-|-|-|-|-|
|id:5 / num:5|id:3 / num:2|id:4 / num:5|id:5 / num:4|...| 


```
遍历仓库将同item_id相同的num叠加:
{id:5 num:9} #index:0 和  index:3 同为id:5 num:累加
{id:3 num:2}
{id:4 num:5}
{...}

再以id大小排序,遍历map,以最大堆叠数依次放回仓库.
```
|0|1|2|3|...|
|-|-|-|-|-|
|id:3 / num:2|id:4 / num:5|id:5 / num:5|id:5 / num:4|...| 
 



只自测了一些功能,性能测试没做.

## 客户端
**打包**
```
/src/client/ 文件夹下己有打好的win,linux可执行文件.
也可进入到 /src/client/ 文件夹目录下自行打包.
```
**IP地址设置:**
```
默认连接地址为:0.0.0.0:8080,如需更新在启动时加入: -ip xxx.xxx.xxx.xxx:port 参数更改.
示例: ./client -ip 127.0.0.1:8888
```


## 部署
```
由于没有做外部配置文件,如有特别需求更改监听IP:PORT(默认全为 0.0.0.0:8080)
1.SERVER端可在/src/main.go 中进行修改.
2.CLIENT端可在启动时加入 -ip 参数进行修改.
```
```
映射端口可自行更改:
Dockerfile 暴露端口为: 8080
docker-compose 中对外映射端口同为 : 8080
```



