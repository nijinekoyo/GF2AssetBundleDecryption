/*
 * @Author: nijineko
 * @Date: 2023-09-30 22:03:17
 * @LastEditTime: 2023-10-07 21:04:05
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
	"strings"

	"github.com/nijinekoyo/GF2AssetBundleDecryption/Decryption"
)

func main() {
	Model := flag.String("model", "ab", "指定解密模式，可选值为ab, story")
	AssetBundlePath := flag.String("ab_path", "./AssetBundles_Windows", "指定AssetBundle文件夹路径")
	AssetBundleDecryptedPath := flag.String("ab_decrypted_path", "./ab_decrypted_output", "指定AssetBundle文件解密后文件夹路径")
	TablePath := flag.String("table_path", "./Table", "指定Table文件夹路径")
	TableDecryptedPath := flag.String("table_decrypted_path", "./table_decrypted_output", "指定Table文件解析后文件夹路径")
	MaxPoolNum := flag.Int("max_pool", 20, "指定最大并发数")
	flag.Parse()

	switch *Model {
	case "ab":
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

		os.Exit(0)
	case "story":
		err := os.MkdirAll(*TableDecryptedPath, 0666)
		if err != nil {
			panic(err)
		}

		StoryDataFiles := []string{
			"LangPackageTableCnBuiltinData.bytes",
			"LangPackageTableCnData.bytes",
		}

		// 解析剧情文件
		for _, FilName := range StoryDataFiles {
			// 读取文件
			FileData, err := os.ReadFile(filepath.Join(*TablePath, FilName))
			if err != nil {
				panic(err)
			}

			// 解析
			DecryptedFile, err := Decryption.LangPackageTable(FileData)
			if err != nil {
				panic(err)
			}

			// 写入文件
			DecryptedFilePath := filepath.Join(*TableDecryptedPath, strings.Replace(filepath.Base(FilName), filepath.Ext(FilName), ".json", 1))

			err = os.WriteFile(DecryptedFilePath, DecryptedFile, 0666)
			if err != nil {
				panic(err)
			}

			fmt.Println("解析成功，文件已保存至:", DecryptedFilePath)
		}

		os.Exit(0)
	default:
		fmt.Println("未知的解密模式")
	}
}
