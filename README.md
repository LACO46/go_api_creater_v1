# Golang API creater

- Golang APIを作成するためにDockerが必要です

## APIの作成方法について

### APIを作成するコマンドについて

以下のコマンドで、APIを作成することができる。

```
$ sh api_creater.sh ~/Desktop/
```

- 引数にAPIを作成するディレクトリを指定する

ディレクトリの指定が間違っていたり未入力だった場合以下ように表示される

```
Please set api dir
```

### API名の入力について

以下の表示が表示されたら、API名を入力する

```
Press add api name: 
```

- 未入力の場合、アプリケーションは終了する

### APIのIPアドレスについて

以下の表示が表示されたら、APIのIPアドレスを入力する

```
Press add api ip address: 
```

- 未入力の場合、アプリケーションは終了する
- 数値以外が入力された場合、アプリケーションは終了する

### DBのIPアドレスについて

以下の表示が表示されたら、DBのIPアドレスを入力する

```
Press add db ip address:
```

- 未入力の場合、アプリケーションは終了する
- 数値以外が入力された場合、アプリケーションは終了する


## APIのビルドについて

作成されたAPIのディレクトリに移動し、以下のコマンドで、APIをビルドすることができる

```
$ docker-compose up --build
```

## APIのマイグレーションについて

以下のコマンドで、APIのマイグレーションを行うことができる

```
$ docker-compose exec <API NAME>_api sh
$ cd setting && sql-migrate up
$ exit
```

マイグレーションが完了したら、以下のように表示される

```
Applied 1 migration
```

## APIの確認について

以下のコマンドで、APIの作成が確認できる

```
$ curl localhost:8080/
```

正常に、APIが起動できている場合、以下のように表示される

```
[
  {
    "Id": "b4c977cf-fb30-4322-a83c-dd7f341adf5d",
    "Name": "hello world"
  }
]
```

- Idはuuidのため、APIごとに変動する