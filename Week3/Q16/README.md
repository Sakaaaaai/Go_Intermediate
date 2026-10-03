# 問16｜JSONのパースと生成

## 目標
`encoding/json` パッケージを使って、GoのstructとJSONを相互変換（エンコード・デコード）する方法を学びます。

## 今回の演習で学ぶこと
この問題では、WebAPIとのデータやり取りで必須となるJSONシリアライズ・デシリアライズを学びます。

---

### なぜJSONが重要か

現代のWebアプリケーションでは、フロントエンド（ブラウザ）とバックエンド（サーバー）の間でデータをやり取りするために **JSON（JavaScript Object Notation）** が標準的に使われています。

```
ブラウザ:  JavaScript  ←→  JSON  ←→  Go サーバー
モバイル:  Swift/Kotlin ←→  JSON  ←→  Go サーバー
他サービス: Python/Ruby ←→  JSON  ←→  Go サーバー
```

JSONは人間にも読みやすく、様々な言語で対応しているため「共通語」のような存在です。

---

**1. 構造体タグ（Struct Tag）でJSONキーをマッピング**

Goの構造体フィールドは大文字始まりですが、JSONのキーは一般的に小文字スタートです。`json:"..."` タグで対応を指定します。

```go
type User struct {
    Name string `json:"name"` // JSON では "name" キーになる（小文字）
    Age  int    `json:"age"`  // JSON では "age" キーになる
}
```

タグを書かないと：
```go
type User struct {
    Name string  // JSON では "Name"（大文字始まり）になってしまう
}
```

> **構造体タグの書き方**: バッククォート `` ` `` で囲み、`json:"キー名"` と書きます。スペースが入ることに注意してください（スペースなしだと認識されない）。

---

**2. Goの構造体 → JSON（Marshal = シリアライズ）**

```go
user := User{Name: "Alice", Age: 30}
jsonBytes, err := json.Marshal(user)
if err != nil {
    log.Fatal(err)
}
// jsonBytes = []byte(`{"name":"Alice","age":30}`)
jsonStr := string(jsonBytes) // string に変換
fmt.Println(jsonStr)         // {"name":"Alice","age":30}
```

`json.Marshal` は構造体を `[]byte` に変換します。

---

**3. JSON → Goの構造体（Unmarshal = デシリアライズ）**

```go
jsonStr := `{"name":"Bob","age":25}`
var user User
err := json.Unmarshal([]byte(jsonStr), &user) // ポインタを渡す！
if err != nil {
    log.Fatal(err)
}
fmt.Println(user.Name) // "Bob"
fmt.Println(user.Age)  // 25
```

> **なぜポインタを渡すのか？** `json.Unmarshal` は変数にデータを書き込む必要があるため、`&user`（アドレス）を渡します。ポインタなしでは書き込めません。

---

**4. json.MarshalIndent で整形出力**

```go
jsonBytes, err := json.MarshalIndent(user, "", "  ") // インデント付き
fmt.Println(string(jsonBytes))
// {
//   "name": "Alice",
//   "age": 30
// }
```

開発時のデバッグや設定ファイルの生成に便利です。

---

**5. 特殊タグオプション**

```go
type User struct {
    Name     string `json:"name"`
    Password string `json:"-"`            // "-" = JSONには含めない（重要！）
    Age      int    `json:"age,omitempty"` // 0値のときはキーを省略
    Email    string `json:"email,omitempty"` // 空文字のときは省略
}
```

| タグオプション | 意味 |
|--------------|------|
| `json:"key"` | JSONキー名を指定 |
| `json:"-"` | このフィールドをJSONに含めない |
| `json:"key,omitempty"` | ゼロ値（0, "", nil, false）のとき省略 |
| `json:"key,string"` | 数値を文字列として扱う |

---

**6. json.NewEncoder / json.NewDecoder（ストリーム版）**

Q25で学ぶHTTPサーバーでは、ファイルや通信ストリームに直接JSON書き込む際に使います：

```go
// 通常版（[]byte 経由）
bytes, _ := json.Marshal(user)
w.Write(bytes)

// ストリーム版（直接書き込む・効率的）
json.NewEncoder(w).Encode(user)
```

---

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

1. `Name string` と `Age int` を持ち、JSONキーが `"name"`, `"age"` になる `User` 構造体を定義してください。
2. `EncodeUser(u User) string` 関数：`json.Marshal` でJSON文字列に変換して返してください。
3. `DecodeUser(jsonStr string) User` 関数：`json.Unmarshal` でJSON文字列をUserに変換して返してください。

### 実行方法
```bash
cd /project/workspace/Week3/Q16
go run main.go
```

### テストでの確認
```bash
cd /project/workspace/Week3/Q16
go test main.go main_test.go
```
## ヒント
- タグの書き方：`` `json:"key_name"` ``（バッククォートで囲む）
- Marshal の戻り値は `([]byte, error)` → `string(bytes)` で文字列化
- Unmarshal の第2引数はポインタ（`&user`）を渡す

## 達成基準
- [ ] 構造体タグでJSONキー名を指定できる
- [ ] `json.Marshal` でJSONを生成できる
- [ ] `json.Unmarshal` でJSONを構造体に変換できる




