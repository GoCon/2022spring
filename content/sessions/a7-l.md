---
key: a7-l
title: 高速で統一的な自動生成ツールをprotocプラグインとして実装した話
id: hikyaru_suzuki
format: conference
talkType: long_session
level: Intermediate
tags: []
speakers:
- hikyaru_suzuki
partner: null
videoId: null
presentation: null
draft: false
---
Go言語には総称型が実装されていないためコードを自動生成して賄うことが多いです。

ここで自動生成のソースをGo言語自体とした場合、よくある手法としてreflectパッケージによる生成が行われますが、ソースが多いと実効速度がネックとなってしまいます。 また、StructTagを活用したオプション設定は便利ですが、文字列による設定なのでタイプミスも発生します。

そこでProtocolBuffersをソースとする自動生成ツールをprotocプラグインとして実装することで、オプションを型安全にしつつ実行速度を大幅に向上させることに成功しました。 今回はこのprotocプラグインを紹介します。