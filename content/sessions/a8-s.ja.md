---
key: a8-s
title: database/sqlパッケージを理解する
id: sivchari
format: conference
talkType: short_session
level: Intermediate
tags:
- A8-S
speakers:
- sivchari
partner: null
videoId: GZCr-Dbi0YM?start=17023
presentation: https://docs.google.com/presentation/d/11jdcnZUm8pYfv6RijlvWEz8ZgK2tMfloIJyLAo-cMnI/edit#slide=id.p
draft: false
---
私たちはdatabase/sqlと好きなdatabaseのdriverをimportすることで、とても簡単にGoとdatabaseの通信をすることができるようになります。
とても便利ですが、なぜdatabaseのdriverはblank importをするだけでDBの種類を識別できているのでしょうか。
また、Goのdatabase/sqlはgoroutineで複数接続しても安全に処理することができますが、どのように制御しているのでしょうか。
このトークを聞くことにより、普段何気なく行っていたdatabase/sqlの仕組みを理解した上でアプリケーションを実装することができるようになります。