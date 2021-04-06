---
key: b4-s
title: 'cobra-cmder: Goの言語機能を活用したシンプルなCLIツール構成法'
id: b4-s
format: conference
talkType: short_session
level: beginner
tags:
  - B4-S
speakerscomder:
  - yaegashi_takeshi
videoId: null
presentation: null
draft: false
---
この講演では筆者が愛用するcobra-cmderというライブラリと共に構成するGoのCLIツールの開発手法について紹介します。cobra-cmderは既存の著名なライブラリであるspf13/cobraを利用していますが、リフレクションや構造体埋め込みといったGoの言語機能の活用によりCLIツールのシンプルな実装を実現していることが特徴です。グローバル変数やinit()を使わず、サブコマンド単位の共通のデータ変数・関数の管理やユニットテストが簡便であり、またサブコマンドの追加も既存コードの変更なしで可能という利点があります。筆者がこれまでに作成したCLIツールの事例とあわせて説明します。

講演で取り上げるライブラリ cobra-cmder および CLI ツール事例は次の場所にあります。
- <https://github.com/yaegashi/cobra-cmder>
- <https://github.com/yaegashi/contest.go>
- <https://github.com/yaegashi/azbill>
- <https://github.com/yaegashi/customazed>

---
次の記事をベースにGoの初心者レベルでもCLIが作れるようになる内容で講演します。

- [cobra-cmder で Go の CLI を簡単に作る](https://qiita.com/yaegashi/items/9acc1d22bcc247542a4e)
