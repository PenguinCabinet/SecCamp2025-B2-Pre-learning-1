# セキュリティキャンプ 2025 B2 事前学習1
## 0. 動作確認済みの実行環境
|項目|バージョン|
|---|---|
|Windows|11|
|nodejs|20.10.0|
|Golang|1.22.4|
|Docker|20.10.16|
## 1. 環境の準備
詳しいソースコードは[1.environment_construction](1.environment_construction)をご覧ください。
### 1-1. Golang
VS Codeで本リポジトリを開き、Dev Containerの「コンテナを再度開く」からGolangを起動してください。   

VS Codeのterminalから、
```
go run 1.environment_construction/HelloWorld.go
```
と実行すれば、完了です。     
### 1-2. TypeScript
VS Codeで本リポジトリを開き、Dev Containerの「コンテナを再度開く」からTypeScriptを起動してください。   

VS Codeのterminalから、
```
tsc 1.environment_construction/HelloWorld.ts && node 1.environment_construction/HelloWorld.js
```
と実行すれば、完了です。     
### 1-3. 完了条件
- [x] [Golangで`Hello World.`を標準出力で表示](./1.environment_construction/HelloWorld.go)
- [x] [TypeScriptで`Hello World.`を標準出力で表示](./1.environment_construction/HelloWorld.ts)

## 2. TypeScript
### 2-1. TypeScript
[https://typescriptbook.jp/](https://typescriptbook.jp/)を読了後、作成したコードは[2.TypeScript/2-1.TypeScript](./2.TypeScript/2-1.TypeScript/)にあります。

[https://typescriptbook.jp/tutorials](https://typescriptbook.jp/tutorials)の各章と本リポジトリのパスの対応は下記の通りです。     
(「Vercelにデプロイしてみよう」の章だけパスではなくデプロイ先のURLになっています)

- [簡単な関数を作ってみよう](./2.TypeScript/2-1.TypeScript/simple-function/)
- [Reactでいいねボタンを作ろう](./2.TypeScript/2-1.TypeScript/like-button/)
- [Next.jsで猫画像ジェネレーターを作ろう](./2.TypeScript/2-1.TypeScript/random-cat/)
- [Vercelにデプロイしてみよう](https://sec-camp2025-b2-pre-learning-1-rand.vercel.app)
- [Jestでテストを書こう](./2.TypeScript/2-1.TypeScript/jest-tutorial/)
- [Reactコンポーネントのテストを書こう](./2.TypeScript/2-1.TypeScript/component-test-tutorial/)
- [Prettierでコード整形を自動化しよう](./2.TypeScript/2-1.TypeScript/prettier-tutorial/)
- [ESLintでTypeScriptのコーディング規約チェックを自動化しよう(JavaScript)](./2.TypeScript/2-1.TypeScript/eslint-tutorial/)
- [ESLintでTypeScriptのコーディング規約チェックを自動化しよう(TypeScript)](./2.TypeScript/2-1.TypeScript/eslint-typescript-tutorial/)

### 2-2. Fastifyの準備

　[https://fastify.dev/docs/latest/Guides/Getting-Started/](https://fastify.dev/docs/latest/Guides/Getting-Started/)を読了後、作成したコードは[2.TypeScript/2-2.Fastify](./2.TypeScript/2-2.Fastify/)にあります。

　各章と実行コマンドの対応は下記の通りです。  
#### Your first server
```
node src/main.js
```
#### Your first plugin
```
node src/main-plugin.js
```
#### MongoDB plugin
```
docker compose up -d
```
#### Validate your data
```
node src/main-validation.js
```
　下記のリクエストを投げて、検証してみましょう。
```
curl -X POST -H "Content-Type: application/json" -d '{"test":"test"}' localhost:3000
```
　下記のように、必須のプロパティがないため、はじかれます。
```
{"statusCode":400,"code":"FST_ERR_VALIDATION","error":"Bad Request","message":"body must have required property 'someKey'"}
```
　次に、必須のプロパティを加えたリクエストを投げます。
```
curl -X POST -H "Content-Type: application/json" -d '{"someKey":"test","someOtherKey":"test"}' localhost:3000
```
　今度は、下記の通り、someOtherKeyが数字でなければならないと、はじかれました。
```
{"statusCode":400,"code":"FST_ERR_VALIDATION","error":"Bad Request","message":"body/someOtherKey must be number"}
```
　そこで、someOtherKeyの値を数字にして、リクエストを送信します。
```
curl -X POST -H "Content-Type: application/json" -d '{"someKey":"test","someOtherKey":0}' localhost:3000     
```
　今度は、うまくいきました。
```
{"someKey":"test","someOtherKey":0}
```
#### Serialize your data
```
node src/main-serialization.js 
```
#### Parsing request payloads
```
node src/main-parsing.js 
```


### 2-3. 完了条件
- [x] [読了後に作成したコード(2-1. TypeScript)](./2.TypeScript/2-1.TypeScript/)
- [x] [読了後に作成したコード(2-2. Fastify)](./2.TypeScript/2-2.Fastify/)


## 3. Golang
### 3-1. Golang
#### Go Tour
　[https://go.dev/tour/welcome/1](https://go.dev/tour/welcome/1)を読了後、作成したコードは[3.Golang/Tour](./3.Golang/Tour/)にあります。      
　フォルダ構成は[go-tour-jpのリポジトリ](https://github.com/atotto/go-tour-jp/tree/jp/content)に準拠しました

　実行方法の例は下記の通りです。
```
cd 3.Golang/Tour/welcome   
go run hello.go
```

#### Go Tutorials
　[https://go.dev/doc/tutorial/](https://go.dev/doc/tutorial/)を読了後、作成したコードは[3.Golang/Tutorials](./3.Golang/Tutorials/)にあります。      
##### [Accessing a relational database](https://go.dev/doc/tutorial/database-access)
下記のコマンドで実行できます。
```
cd 3.Golang/Tutorials/data-access
docker compose up
```
##### [Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)
下記のコマンドで実行できます。
```
cd 3.Golang/web-service-gin
go run .
```
##### [Getting started with generics](https://go.dev/doc/tutorial/generics)
下記のコマンドで実行できます。
```
cd 3.Golang/Tutorials/generics
go run .
```
##### [Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)
下記のコマンドで実行できます。
```
go test -fuzz=Fuzz -fuzztime 30s
```



### 3-2. 完了条件
- [ ] [読了後に作成したコード(3-1. Golang)](./3.Golang/)
