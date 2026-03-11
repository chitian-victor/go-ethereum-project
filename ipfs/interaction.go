package ipfs

import (
	"fmt"
	"io"
	"log"

	ipfsapi "github.com/ipfs/go-ipfs-api"
)

const testFileCID = "QmYL5ihKy6QHANJngMGUn7JUe8iJ7KytHx2aG3ShR6tDf5"

func Cat() {
	// 1. 连接到本地的 IPFS Daemon (5001 是 API 端口)
	sh := ipfsapi.NewShell("localhost:5001")

	// 2. 使用 Cat 方法获取数据流 (io.ReadCloser)
	reader, err := sh.Cat(testFileCID)
	if err != nil {
		log.Fatalf("从 IPFS 获取数据失败: %v", err)
	}
	defer reader.Close()

	// ---- 选项 A：直接把内容读进内存 (适合 JSON / 短文本) ----
	bytes, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("读取数据流失败: %v", err)
	}
	fmt.Printf("成功获取到内容:\n%s\n", string(bytes))
	// ---- 选项 B：如果是大文件(如视频)，直接以流的形式写入本地硬盘 ----
	/*
		outFile, _ := os.Create("downloaded_video.mp4")
		defer outFile.Close()
		io.Copy(outFile, reader) // 流式拷贝，不撑爆后端内存
	*/
}
