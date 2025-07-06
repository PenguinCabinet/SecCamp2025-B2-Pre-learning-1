# セキュリティキャンプ 2025 B2 事前学習1
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



### 2-3. 完了条件
- [x] [読了後に作成したコード(2-1. TypeScript)](./2.TypeScript/2-1.TypeScript/)
- [ ] 読了後に作成したコード(2-2. Fastify)
