---
key: b5-l
title: Mastering Tags of Structs - Goの構造体とタグを極める
id: b5-l
format: conference
talkType: long_session
level: advanced
tags:
  - B5-L
speakers:
  - yoshiki_shibukawa
videoId: null
presentation: null
draft: false
---
Struct and its tag are important building blocks to create large applications. They reduce boring codes and make application more declarative and help to write enterprise application, but creating library to process tag requires much amount of code with reflection / typing.

Goで大規模アプリケーションを作るためのビルディングブロックは、構造体とタグです。手続き的なコードを減らし、宣言的に解決する手法はGoをエンタープライズ開発で活用するためにますます必要となると思いますが、自作するにはタグを処理するコーディングは多数の型を網羅するコーディングが必要で長くなります。

This presentation introduce tag processing patterns that uses reflect package and understanding about tag, and how to handle tag to the people who are responsible for Go application architecture.

本プレゼンでは、Goプロジェクトを牽引する「アーキテクト」クラスの開発者を対象に、reflectパッケージを使ったタグ処理のパターンを紹介するとともに、タグ処理関数を実装するための考え方、ロジックなどを紹介します。

説明は日本語、資料は日英表記で発表します。

I make presentation in Japanese (little English), and its slide in Japanese and English.

---

Goではリフレクションよりも静的なコード生成を重視する文化ですが、
reflectパッケージが多様される数少ない活躍の場が構造体のタグの処理です。
Goを業務で使ったことがあれば、encoding/jsonを使ったマッピングや
RDBのレスポンスのマッピングなどでその恩恵に授かったことがある人が
ほとんどだと思います。

In Go's culture, reflection is less important than static code generation.
but processing tag is unchallenged position for reflect package. Almost all Gophers may benefit from JSON-struct mapping of encoding/json, row-struct mapping of ORM.

実装の効率を高めるには、汎用的なユーティリティを作成して少ない行数で
ロジックを実現できる、繰り返し実装を減らしてミスを減らすなど、
コーディングのテクニックで改善できるものもありますが、
Goらしい方法としてタグを使った宣言的記法を使う方法もあります。
高度なタグ処理関数を実装することは一定以上の品質のプログラムを一定の速度で作り続けるには極めて有効な方法です。

There are several ways to improve programming performance. Implementing utility functions and reducing code duplication are the most basic way of them. Declarative code using tag is a Go-ish option. To create library to process tag is very efficient way to keep development speed.

ざっとタグの活用例をリストアップしたのが次の項目です。

There are many application fo tag:

* DBのマッピング
* envconfigで環境変数のマッピング
* JSON/YAMLなどのマッピング
* バリデーション
* 比較など

* DB row mapping
* Environment variable mapping
* Data file mapping (like JSON/YAML)
* Validation
* Comparison etc

セッションの中では、タグ処理のロジックを3種類に類型化し、その4通りの組み合わせを紹介します。

I introduce three building blocks to implement tag processor and four combination ot them.

これらのパターンの実装にあたっては多くの人が使い慣れないreflectパッケージを使います。
構造体の探索の仕方のコツや、構造体のフィールドからの値の読み込み
フィールドへの値の設定について、さまざまなバリエーションを網羅する書き方を紹介します。
し、実用的なサンプルとして、*http.Requestの各要素（JSONの
ボディだけではなく、ヘッダーやクッキーなども）を構造体にマッピング
するコードを紹介します。

To implement tag processor, we use reflect package that are not familiar with regular application development. I introduce how to traverse struct and extracting/assigning value from/to struct's field. Then I describe
practical sample that maps *http.Request to struct (not only JSON body, but also header, cookie, etc).

タグの活用の可能性はまだ開拓されきっていないと思いますが、
その割に現時点でタグ処理のコードを書くための情報が十分にあるとは
言えないと思います。本発表を通じて、仕事の流れを作るようなアーキテクトクラスの開発者、
少ないコーディング量で使えるようなライブラリを実装したいと
考えているライブラリ作者の人たちに、Goらしさを持った
宣言的なコーディングテクニックを提供します。

I think tag's possibility is not completely unexplored, but the description about tag processing is not enough. I provides Go-ish declarative coding technique to architect or library programmer.