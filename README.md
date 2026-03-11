#

## solidity 原生编译器
1. 安装方法 for mac
> [参考](https://docs.soliditylang.org/en/latest/installing-solidity.html#macos-packages)
```shell
brew update
brew upgrade 
brew tap ethereum/ethereum
brew install solidity
```
2. 升成 abi 和 bin 文件
> 如果用的 foundry 框架，可以直接从 out/contract_name.json里面提取出 abi 也可以
```shell
# 需要指定--include-path，代表该合约引用了哪些包，比如openzeppelin，同时要配合 --base-path 使用，代表当前执行路径
solc --abi --bin --include-path lib  ./src/PumpkinFaucet.sol --base-path . -o ./solc_out --overwrite
```

## go-ethereum
> 用于和以太坊交互的 golang 仓库

1. 安装 geth 自带的所有 cmd，也可以单独安装指定 cmd 
``` shell
git clone https://github.com/ethereum/go-ethereum.git
cd go-ethereum
make geth
```
2. 单独安装 abigen
> [指令说明](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings)
```shell
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```
3. 利用 abigen 解析 abi 和 bin 文件生成对应 go 文件
```shell
$ abigen --bin relative_path.bin --abi relative_path.abi --pkg go_pkg_name --type contract_name --out relative_path/contract_alias.go
```
示例：
```shell
abigen --abi abi/PumpkinToken.abi --pkg contract --type PumpkinToken --out contract/pumpkin_token.go
```


## IPFS
1. 安装 Kubo
> [参考](https://docs.ipfs.tech/install/command-line/#install-ipfs-using-package-managers)
```shell
brew install ipfs
```
2. 初始化 IPFS 节点
> 默认生成到～/.ipfs 目录下
```shell
ipfs init
```
3. 启动 IPFS 节点[轻易别启动，是将本机作为存储节点链接全球的]
```shell
ipfs daemon
# 这个可以离线启动，不对外链接开放
ipfs daemon --offline
```
4. 上传文件得到 CID、利用 CID 查看文件、下载文件
```shell
ipfs add {file_path}
ipfs cat {CID}
ipfs get {CID}
```





## ganache[deprecated]
> 一个基于 node.js用于测试的区块链网快速部署工具
1. npm 安装 & 运行
```shell
npm install -g ganache-cli
# 运行客户端
ganache-cli
```
2.推荐采用固定的助记词可以生成相同的地址
```shell
ganache-cli -m "you are really the very best person in the world"
```