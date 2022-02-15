# Go Conference 2022 Spring

![GitHub Actions](https://github.com/GoCon/2022spring/actions/workflows/ci.yml/badge.svg?branch=main)

## 必要なもの

* Hugo

## テーマの取得

```sh
$ git submodule init
$ git submodule update
```

## ローカルでの動作確認

```sh
$ hugo server
```

## データ素材のダウンロード

下記のコマンドで[マスタデータのスプレッドシート](https://docs.google.com/spreadsheets/d/1cxTGkIUIQ7UCfFmXJXOl0shwGLKcJxo03gOwVZ7l1A8/edit#gid=1791874699)の内容を`./raw_data`フォルダにダウンロードします

```
make download
```

## CSVデータ素材のyamlへの変換

`scripts`配下のコードで`raw_data`配下のcsvファイルをyamlに変換し、`data`ディレクトリに配置します。

makeファイルに記載したコマンドで一括変換します。

```
make generate-yml
```