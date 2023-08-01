![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

# 使用方法
## 配置
    // 必选配置，指定endpoint、用户名、密码
    var client = aliyun_igraph_go_sdk.NewClient("your_endpoint", "your_user_name", "your_password", "your_src")
    your_endpoint: igraph-cn-xxxx.igraph.aliyuncs.com
    your_user_name: 购买实例时设置的用户名
    your_pass_word: 购买实例时设置的密码
    your_src: 用来标记来源的标识

    var config = aliyun_igraph_go_sdk.ClientConfig{
				MaxConnsPerHost: 128,
				RequestTimeout:  100 * time.Millisecond,
			}
    RequestTimeout: 请求超时设置 默认是1s
    MaxConnsPerHost: 单机连接数上限 默认512

## 查询使用样例

    var client = aliyun_igraph_go_sdk.NewClient("http://igraph-cn-xxxx.igraph.aliyuncs", "username", "password", "src")
    client.InitConfig(config)
    m := make(map[string]string)
    readRequest := &aliyun_igraph_go_sdk.ReadRequest{QueryString: "GremlinQuery", QueryParams: m}
    resp, err := client.Read(*readRequest)

## 更新使用样例
    graphName := "graphName"
    instanceName := "igraph-cn-xxxx"
    labelName := "labelName"
    pkey := "pkfieldName"
    request := NewWriteRequest(WriteTypeAdd, instanceName, tableName, labelName, pkey, "", map[string]string{})
    request.AddContent("field1", "1")
    request.AddContent("field2", "1")
    request.AddContent("field3", "1")
    request.AddContent(pkey, "1")
    resp, err := client.Write(*request)