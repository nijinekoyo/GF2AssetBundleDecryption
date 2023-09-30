/*
 * @Author: nijineko
 * @Date: 2023-09-30 22:25:06
 * @LastEditTime: 2023-09-30 22:25:09
 * @LastEditors: nijineko
 * @Description: 工具封装
 * @FilePath: \GF2AssetBundleDecryption\tools.go
 */
package main

import (
	"os"
	"path/filepath"
)

/**
 * @description: 遍历出文件夹内所有文件路径
 * @param {string} DirPth
 * @return {[]string} 文件路径
 * @return {error} 错误
 */
func TraverseFolders(DirPth string) ([]string, error) {
	DirPth = filepath.Clean(DirPth)
	var dirs []string
	dir, err := os.ReadDir(DirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	var Files []string
	for _, fi := range dir {
		// 目录, 继续递归遍历
		if fi.IsDir() {
			dirs = append(dirs, filepath.Clean(DirPth+PthSep+fi.Name()))
			TraverseFolders(DirPth + PthSep + fi.Name())
		} else {
			Files = append(Files, filepath.Clean(DirPth+PthSep+fi.Name()))
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := TraverseFolders(table)
		for _, temp1 := range temp {
			Files = append(Files, filepath.Clean(temp1))
		}
	}

	return Files, nil
}
