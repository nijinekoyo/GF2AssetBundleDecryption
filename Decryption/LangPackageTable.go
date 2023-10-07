/*
 * @Author: nijineko
 * @Date: 2023-10-07 20:23:44
 * @LastEditTime: 2023-10-07 21:00:16
 * @LastEditors: nijineko
 * @Description: 剧情解包工具
 * @FilePath: \GF2AssetBundleDecryption\Decryption\LangPackageTable.go
 */
package Decryption

import (
	"bytes"
	"encoding/binary"
	"encoding/json"

	"github.com/nijinekoyo/GF2AssetBundleDecryption/Decryption/Proto"
	"google.golang.org/protobuf/proto"
)

/**
 * @description: 解析剧情文件
 * @param {[]byte} FileData 文件数据
 * @return {[]byte} 解析后的数据
 */
func LangPackageTable(FileData []byte) ([]byte, error) {
	// 读取前4个字节作为一个32位无符号整数
	var skip uint32
	err := binary.Read(bytes.NewReader(FileData[:4]), binary.LittleEndian, &skip)
	if err != nil {
		return nil, err
	}

	// 截取数据，跳过前部分
	FileData = FileData[skip+4:]

	// 解析Proto
	TextMapTable := &Proto.TextMapTable{}
	err = proto.Unmarshal(FileData, TextMapTable)
	if err != nil {
		return nil, err
	}

	// 转化为json
	TextMapTableJson, err := json.Marshal(TextMapTable)
	if err != nil {
		return nil, err
	}

	// 格式化json，便于阅读
	var IndentTextMapTableJson bytes.Buffer
	json.Indent(&IndentTextMapTableJson, TextMapTableJson, "", "    ")

	// 转化为[]byte，覆写原数据
	TextMapTableJson = IndentTextMapTableJson.Bytes()

	return TextMapTableJson, nil
}
