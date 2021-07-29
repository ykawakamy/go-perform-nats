# go-perform-nats
[nats.io](https://docs.nats.io/)のパフォーマンス測定用です。

# publisher
## 使い方
```
go run ./publisher 
# 引数の説明表示
go run ./publisher -h
```

# subscriber
## 使い方
```
go run ./subscriber
# 引数の説明表示
go run ./subscriber -h
```

# 引数の説明
| 引数 | 説明 | デフォルト値| PUB | SUB |
|------|------|------|------|------|
| -d (ns)| 計測時間をナノ秒で指定します | 5秒(5,000,000,000ns)| 〇 | 〇 |
| -i (count)| 計測回数を指定します。 | 20回 | 〇 | 〇 |
| -s (string)| 接続先URL | localhost:4222 | 〇 | 〇 |
| -t (topic)| Publish/SubscribeするTopic名（Subject名）<p>複数Topic指定する場合、下記の</p> | test | 〇 | 〇 |
| -m (bytes)| Publishするメッセージサイズ<p>※パフォーマンス測定用のデータが含まれるため、24バイト以上で指定すること</p> | 150 | 〇 | - |
| -c (count)| スレッド数を指定します。 | 10 | - | 〇 |
