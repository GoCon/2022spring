---
key: a8-s
title: database/sqlパッケージを理解する
id: sivchari
format: conference
talkType: short_session
level: Intermediate
tags: []
speakers:
- sivchari
partner: null
videoId: null
presentation: null
draft: false
---
私たちはdatabase/sqlと好きなdatabaseのdriverをimportすることで、とても簡単にGoとdatabaseの通信をすることができるようになります。
とても便利ですが、なぜdatabaseのdriverはblank importをするだけでDBの種類を識別できているのでしょうか。
また、Goのdatabase/sqlはgoroutineで複数接続しても安全に処理することができますが、どのように制御しているのでしょうか。
このトークを聞くことにより、普段何気なく行っていたdatabase/sqlの仕組みを理解した上でアプリケーションを実装することができるようになります。