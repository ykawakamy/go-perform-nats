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
| -t (topic)| Publish/Subscribeするトピック名（Subject名）<p>複数トピック指定する場合、下記のセクションを参照してください。</p> | test | 〇 | 〇 |
| -m (bytes)| Publishするメッセージサイズ<p>※パフォーマンス測定用のデータが含まれるため、24バイト以上で指定すること</p> | 150 | 〇 | - |
| -c (count)| プロセスあたりのクライアント（スレッド）数を指定します。 | 10 | - | 〇 |
| -v (count)| トピックの振分数を指定します。1以上をしてしてください。 | 1 | - | 〇 |

## 複数トピック指定する場合
※この機能は検証ツールとしての機能です。
### 範囲指定
`/topic/:(L)-(H)`のように指定すると、L<=x<=Hの範囲のサブトピックとして使用します。
例えば、`/topic/1-3`の場合、`/topic/1`,`/topic/2`,`/topic/3`を使用します。

### 複数指定
`/topic/:1,2,3`のようにカンマ区切りとした場合、`/topic/1`,`/topic/2`,`/topic/3`を使用します。

## トピックの振分について
Subscriberでは``-v`,`-t`の複数トピック指定を組み合わせることでクライアントあたりのトピックを制御できます。
トピックを指定した振分数で分割したものを１クライアントに割り当てます。
例えば、`-v 80 -t /test/:1-4000`を指定した場合、1クライアントあたり50(=4000/80) トピックを指定します。

