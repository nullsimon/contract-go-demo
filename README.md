# 学习web3的第二天

今天就学习下写以太坊合约，并且部署到以太坊网络上，再创建一个读合约的go版本的api server。

使用到的技术如下

* go
* ethereum
* solidity
* ganache-cli


# 预备：软件安装

## 1 安装golang

```shell
brew install golang
```

## 2 安装ethereum

follow official site [geth](https://geth.ethereum.org/docs/install-and-build/installing-geth#macos-via-homebrew)

主要作用，提供一些命令行，如 apigen ，生成合约的go代码

## 3 安装solidity

follow official site [soliditylang](https://docs.soliditylang.org/en/v0.8.2/installing-solidity.html)

主要作用，提供solidity 的命令行，如 solc

## 4 安装ganache-cli
```shell
brew install ganache-cli
```

主要作用，提供一个本地的节点，用来测试合约的功能

# 写合约主要源代码-solidity 代码

## 1 写合约的solidity代码
```javascript
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

contract MySmartContract {
    function Hello() public view returns (string memory) {
        return "Hello World";
    }
    function Greet(string memory str) public view returns (string memory) {
        return str;
    }
}
```

解释：首先是一个无参数的方法，Hello，一个是有参数的方法，Greet。

## 2 根据合约生成go代码

```shell
solc --optimize --abi ./contracts/MySmartContract.sol -o build
solc --optimize --bin ./contracts/MySmartContract.sol -o build
abigen --bin=./build/MySmartContract.bin --abi=./build/MySmartContract.abi --pkg=api --out=./api/MySmartContract.go
```

解释：solc 命令，将solidity代码转换成二进制文件，并且生成abi文件，apigen 生成go相关的代码，并且生成到api目录下。供后续deploy和client使用

## 3 写go的部署合约代码

```go
// 此处连接的是你本地的节点，也就是ganache-cli提供的，需要注意换成你自己的
client, err := ethclient.Dial("http://127.0.0.1:8545")
...
// 此处的私钥，换成ganache-cli 提供的，并且去掉开头的 0x 
privateKey, err := crypto.HexToECDSA("be1b85896f93f5d18fe2cf28b81daecbf790e33bdd96fb52a056d669b0c93cde")
...
// 部署代码到本地的以太坊网络上
address, tx, instance, err := api.DeployApi(auth, client)
```

## 4 写go的读合约代码

```go
// 这里的address是你部署的合约地址，也就是部署完成后，会提供你一个合约地址
conn, err := api.NewApi(common.HexToAddress("351dd6679c502b41c221ab749666d4ca6c6b8f5d"), client)

// 调用合约的方法使用
e.GET("/greet/:message", func(c echo.Context) error {
    message := c.Param("message")
    reply, err := conn.Greet(&bind.CallOpts{}, message)
    if err != nil {
        return err
    }
    return c.JSON(http.StatusOK, reply)
})
```

# 部署合约，与之进行交互

## 1 开启一个ganache-cli节点

一行命令即可

```bash
ganache-cli 
```

会给你自动生成一些账户，私钥公钥。

## 2 部署合约

成功之后，会给你一个合约地址，可以用来调用合约的方法，也就是client所需要用到的

0x351DD6679C502xxxxx21aB749666d4Ca6C6b8f5D

## 3 调用合约的方法

```shell
curl http://localhost:1323/greet/hello
curl http://localhost:1323/hello
```
成功！！！

# 参考文章
[Creating a simple Ethereum Smart Contract in Golang](https://towardsdev.com/creating-a-simple-ethereum-smart-contract-in-golang-138b9439f64e)
