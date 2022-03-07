---
key: b3-l
title: ゼロから作る Protocol Buffer のパーサーとレキサー
id: yoheimuta
format: conference
talkType: long_session
level: All
tags: []
speakers:
- yoheimuta
partner: null
videoId: null
presentation: null
draft: false
---
インタプリタやコンパイラの基礎になる字句解析器（lexer）と構文解析器（parser）の実装はgoyaccなどのジェネレーターを使うか手書きするかの基本二択になります。goyaccに関する実践的な情報は多いですが学習カーブが伴います。私はProtocolBufferスキーマ定義ファイルのパーサーとレキサーをGoの標準パッケージだけで実装しました。ASTを標準出力するだけでなくVisitorパターンを導入すると使い勝手が増します。これらの知識は普段使う静的解析ツールのカスタマイズにも役立ちます。本セッションでは実際に動く最小構成の実装からはじめてGoでのプログラミング手法をご紹介します。