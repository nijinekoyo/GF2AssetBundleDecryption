/*
 * @Author: nijineko
 * @Date: 2023-09-30 22:03:17
 * @LastEditTime: 2023-10-01 02:18:40
 * @LastEditors: nijineko
 * @Description: main.go
 * @FilePath: \GF2AssetBundleDecryption\main.go
 */
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nijinekoyo/GF2AssetBundleDecryption/Decryption"
)

func main() {
	AssetBundlePath := flag.String("ab_path", "./AssetBundles_Windows", "指定AssetBundle文件夹路径")
	AssetBundleDecryptedPath := flag.String("ab_decrypted_path", "./ab_decrypted_output", "指定AssetBundle文件解密后文件夹路径")
	MaxPoolNum := flag.Int("max_pool", 20, "指定最大并发数")
	flag.Parse()

	err := os.MkdirAll(*AssetBundleDecryptedPath, 0666)
	if err != nil {
		panic(err)
	}

	// 遍历AssetBundle文件夹，找出所有.bundle文件
	var AssetBundleFiles []string
	AllFiles, err := TraverseFolders(*AssetBundlePath)
	if err != nil {
		panic(err)
	}
	// 过滤出.bundle文件
	for _, file := range AllFiles {
		if filepath.Ext(file) == ".bundle" {
			AssetBundleFiles = append(AssetBundleFiles, file)
		}
	}

	Pool := NewPool(*MaxPoolNum)

	// 解密AssetBundle
	for _, FilePath := range AssetBundleFiles {
		Pool.Add(1)

		go func() {
			FileData, err := os.ReadFile(FilePath)
			if err != nil {
				panic(err)
			}

			// 解密
			DecryptedFile := Decryption.AssetBundle(FileData)

			// 写入文件
			DecryptedFilePath := filepath.Join(*AssetBundleDecryptedPath, filepath.Base(FilePath))

			err = os.WriteFile(DecryptedFilePath, DecryptedFile, 0666)
			if err != nil {
				panic(err)
			}

			fmt.Println("解密成功，文件已保存至:", DecryptedFilePath)
			Pool.Done()
		}()

		Pool.Wait()
	}
}
