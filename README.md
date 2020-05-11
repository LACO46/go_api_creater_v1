# MeCab Result Record Api

## dockerコンテナのビルド方法

```
docker build -t mecab_result_record_api .
```

## 起動コマンド

```
docker run -it -p 8080:80 -v /<絶対PATH>/mecab_result_record_api:/home mecab_result_record_api
```