/*
 * @Author: nijineko
 * @Date: 2023-09-30 22:03:49
 * @LastEditTime: 2023-10-07 20:22:48
 * @LastEditors: nijineko
 * @Description: AssetBundle解密工具
 * @FilePath: \GF2AssetBundleDecryption\Decryption\AssetBundle.go
 */
package Decryption

/**
 * @description: AssetBundle异或计算
 * @param {[]byte} FileData 文件数据
 * @param {[]byte} Key 密钥
 * @return {[]byte} 解密后的数据
 */
func AssetBundleXOR(FileData []byte, Key []byte) []byte {
	var Size int
	if len(Key) < len(FileData) {
		Size = len(Key)
	} else {
		Size = len(FileData)
	}

	XORData := make([]byte, Size)
	for i := 0; i < Size; i++ {
		XORData[i] = FileData[i] ^ Key[i]
	}
	return XORData
}

/**
 * @description: 解密AssetBundle
 * @param {[]byte} FileData 文件数据
 * @return {[]byte} 解密后的数据
 */
func AssetBundle(FileData []byte) []byte {
	Key := []byte{0x55, 0x6E, 0x69, 0x74, 0x79, 0x46, 0x53, 0x00, 0x00, 0x00, 0x00, 0x07, 0x35, 0x2E, 0x78, 0x2E}

	// XOR计算文件密钥
	FileKey := AssetBundleXOR(FileData, Key)

	// 解密文件
	var Size int
	if 0x1000*8 < len(FileData) {
		Size = 0x1000 * 8
	} else {
		Size = len(FileData)
	}

	for i := 0; i < Size; i++ {
		// 解密文件，替换原数据
		FileData[i] = FileData[i] ^ FileKey[i%len(FileKey)]
	}

	return FileData
}
