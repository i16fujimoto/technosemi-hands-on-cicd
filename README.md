# technosemi-hands-on-cicd
2023年春 実践教育プログラム Web &lt;破> CI/CD入門 ハンズオンリポジトリ

## 開発環境の構築

本リポジトリの推奨される Go のバージョンは以下の通りです。<br>
インストールについてはこちらの[公式サイト](https://go.dev/doc/manage-install)を参照してください。

```shell
v1.21.7
```

## Github Actions
./github/workflows/ に実行するActionsの設定ファイルが格納されています。
- go_test.yml: Goのコードをビルドし，テストを実行する
- review-dog.yml: Goのコードを静的解析し，結果をレビュードッグでPR上にコメントで表示する
