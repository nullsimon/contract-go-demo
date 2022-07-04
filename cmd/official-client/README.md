# 学习web3的第三天

前两天主要学习了solidity的语言语法，以及合约的知识，并且部署到一个测试环境。

今天呢，我们来操作线上的环境:),哈哈哈，不是真正的写，仅仅读取一些线上的交易消息，感受下ether的律动。


# 预备
## 1 申请一个infura的账号
直接官网申请即可，他是一个web2的服务，提供的api服务，把ether的信息转换成一个api和消息投递出来。你也可以本地启动一个ether的node，来接入node也可以。

[infura](https://docs.infura.io/infura/getting-started)

主要作用：提供一个proxy，把ether的信息转换成一个api和消息投递出来。

这里需要注册一个账户，并且拿到一个project-id,有足够的的免费额度，供你使用

## 2 打开etherscan，来进行一个信息的比对和查询，确保后续测试正常

[etherscan](https://etherscan.io/)

## 3 打开vscode，准备写代码

do by yourself

# 读ether和监听ether的交易消息
## 1 先来读取一个ether的最高交易区块

```go
// 获取最高交易区块，使用infura的api
client, err := rpc.Dial("https://mainnet.infura.io/v3/<YOUR-PROJECT-ID>")
if err != nil {
	log.Fatalf("Could not connect to Infura: %v", err)
}

var lastBlock Block
err = client.Call(&lastBlock, "eth_getBlockByNumber", "latest", true)
if err != nil {
	fmt.Println("Cannot get the latest block:", err)
	return
}
// 获取最高区块的高度
fmt.Printf("Latest block: %v\n", lastBlock.Number)
```

## 2 监听NewBlock事件
```go
// 获取最高交易区块，使用infura的websocket
ws, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/<YOUR-PROJECT-ID>")

if err != nil {
	log.Fatal(err)
}

// 设置监听事件
headers := make(chan *types.Header)
sub, err := ws.SubscribeNewHead(context.Background(), headers)
if err != nil {
	log.Fatal(err)
}

// 监听事件
for {
	select {
	case err := <-sub.Err():
		log.Fatal(err)
	case header := <-headers:
		fmt.Println(header.Hash().Hex()) 

		block, err := ws.BlockByHash(context.Background(), header.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(block.Hash().Hex())        
		fmt.Println(block.Number().Uint64())  
		fmt.Println(block.Time())           
		fmt.Println(block.Nonce())            
		fmt.Println(len(block.Transactions()))
	}
}
```

## 3 other things

# 结语

现在有一个go的代码环境，可以来进行web3的交互，交易等，也是操作过线上环境了。来吧，感受下ether的律动。那一笔笔的交易消息，就是这个时代的心跳。

# 参考文章
[web3-from-zero](https://kay-is.github.io/web3-from-zero/)

[infura](https://docs.infura.io/infura/getting-started)

[etherscan](https://etherscan.io/)

[goethereumbook](https://goethereumbook.org/)