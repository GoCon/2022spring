---
key: b8-s
title: 「自動コード生成ツール」を20分で作れるようになろう
id: riita10069
format: conference
talkType: short_session
level: Beginner
tags: []
speakers:
- riita10069
partner: null
videoId: null
presentation: null
draft: false
---
皆さんは、普段からGo言語でコードを書いていると思います。
「いつも同じようなプログラムだから自動生成したいな」と感じることはありませんか？

例えば、
- 単にCRUDするだけのコード
- if err != nil 的なエラーハンドリング
- natsやKafkaからドメインイベントを取ってくるworker
- 副作用のないメソッドのユニットテスト

のようなものです。

スキーマ定義のみでCRUDするAPIを生成するジェネレータを作成したので、
コード生成をしたことない人向けに紹介したいと思います。

Keyword:
ast, roche, jeniffer, cobra, afero