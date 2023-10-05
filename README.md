# GF2AssetBundleDecryption
少女前线2：追放 AssetBundle 资产解密工具

## Use
直接将程序放置于`GF2Exilium\GF2 Game\GF2_Exilium_Data\LocalCache\Data`文件夹下，双击运行即可，解密后文件位于`decrypted_output`文件夹

或使用以下命令启动
```
go run . -ab_path './AssetBundles_Windows'
```

## Parameter
```
-ab_path string
    指定AssetBundle文件夹路径 (default "./AssetBundles_Windows")
-decrypted_path string
    指定解密后文件夹路径 (default "./decrypted_output")
-max_pool int
    指定最大并发数 (default 20)
```

# Question
1. Q: 如何解密音频文件？  
A: 请参考这个Issues [关于其他资源文件 #1](https://github.com/nijinekoyo/GF2AssetBundleDecryption/issues/1)

## Thanks
1. [GF2Decrypter](https://github.com/66hh/GF2Decrypter) 提供了解密数据的思路