---
key: b9-s
title: Sharing test cases of internet protocol with Go and OCI Artifacts
id: b9-s
format: conference
talkType: short_session
level: intermediate
tags:
  - B9-S
speakers:
  - moto_ishizawa
videoId: null
presentation: null
draft: false
---
インターネットは様々なプロトコルとその実装によって支えられています。しかしながら、モダンなプロトコルの仕様は複雑化しており、プロトコルを安全に正しく実装することがますます難しくなっています。このトークでは Go と OCI Artifacts を活用して、プロトコルの仕様に関するテストケースを開発者間で共有し、テストを容易にする試みについてご紹介します。

---
このトークでは Go と OCI Artifacts を利用して開発した protospec というツールを使用して、プロトコルのテストケースをサーバーやクライアントの開発者間で共有し、テストを容易にするための仕組みについて説明します。

protospec は YAML ファイルに記述されたプロトコルのテストケースを読み込み、サーバーやクライアントに対してテストを実行できます。テストケースを記述した複数の YAML ファイルは OCI Artifacts を使用して1つコンテナイメージとして保存でき、そのイメージは一般的なコンテナイメージと同様に、コンテナレジストリを通じて他の開発者と共有可能になっています。protospec は全て Go で実装されており、テストの実行やテストケースの共有などの機能を Go でどのように実装したかについて説明する予定です。

## 得られる知識

このトークは、参加者が以下の知識を獲得できることを想定しています。

- プロトコルの実装におけるテストの課題
- プロトコルの仕様に対するテストの方法
- Go で OCI Artifact を操作するための実装方法

## 前提知識

このトークでは以下の前提知識を必要とする想定です。

- Go の構文についての簡単な理解
- インターネットプロトコルについての簡単な理解

## アジェンダ

このトークでは以下のようなトピックについてお話する予定です。

- プロトコル実装における課題
    - テストケースの実装が難しい
    - テストケースの共有がしにくい
- HTTP/2 のテストツールによる試み
    - h2spec の開発と利用例の紹介
    - h2spec における課題
- より汎用的なテストツールの開発
    - DSL でテストケースが実装できる protospec
    - OCI Artifact を利用したテストケースの共有
    - Go で OCI Artifact を操作するための実装例の紹介
- 今後の展望
    - プロトコル実装のプラグイン化

## 関連リンク

トークで紹介する protospec のコードは後日 GitHub で公開予定です。

- https://github.com/summerwind/h2spec
- https://github.com/opencontainers/artifacts
