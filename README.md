# GF2AssetBundleDecryption
《少女前线2：追放》资产解密工具

## Use
1. AssetBundle解密模式  
直接将程序放置于`GF2Exilium\GF2 Game\GF2_Exilium_Data\LocalCache\Data`文件夹下，双击运行即可，解密后文件位于`ab_decrypted_output`文件夹

或使用以下命令启动
```
.\GF2AssetBundleDecryption.exe -ab_path './AssetBundles_Windows'
```
OR
```
go run . -ab_path './AssetBundles_Windows'
```

2. 剧情文本解析模式  
将程序放置在`GF2Exilium\GF2 Game\GF2_Exilium_Data\LocalCache\Data\Table`文件夹下，执行以下命令即可，解析后的json文件位于`table_decrypted_output`文件夹
```
.\GF2AssetBundleDecryption.exe -model story -table_path './Table'
```
OR
```
go run . -model story -table_path './Table'
```

## Parameter
```
-model string
    指定解密模式，可选值为ab, story (default "ab")
-max_pool int
    指定最大并发数 (default 20)
-ab_decrypted_path string
    指定AssetBundle文件解密后文件夹路径 (default "./ab_decrypted_output")
-ab_path string
    指定AssetBundle文件夹路径 (default "./AssetBundles_Windows")
-table_decrypted_path string
    指定Table文件解析后文件夹路径 (default "./table_decrypted_output")
-table_path string
    指定Table文件夹路径 (default "./Table")
```

# Question
1. Q: 如何解密音频文件？  
A: 请参考这个Issues [关于其他资源文件 #1](https://github.com/nijinekoyo/GF2AssetBundleDecryption/issues/1)

## Thanks
1. [GF2Decrypter](https://github.com/66hh/GF2Decrypter) 提供了解密数据的思路