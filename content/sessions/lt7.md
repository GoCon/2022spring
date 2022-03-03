---
key: lt7
title: Go runtime の歩き方
id: convto
format: conference
talkType: lt_session
level: Intermediate
tags: []
speakers:
- convto
partner: null
videoId: null
presentation: null
draft: false
---
ある程度 Go での開発経験をつむと、 goroutine や channel の動作原理を知るために Go の runtime パッケージのソースを読みたいことがあると思います。

しかし Go の runtime は plan9 ベースのアセンブリでの実装があったり、他ではあまり見かけない compiler directive を使用していたり、 所見だと built in 関数の名称がわからなかったり、前知識なしにコードが追いづらいです。

そこでこれらの解説や objdump で built in 関数を追う手法を紹介し、 runtime パッケージを読むための前提知識をまとめます。