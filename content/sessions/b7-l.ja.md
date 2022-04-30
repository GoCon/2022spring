---
key: b7-l
title: Go で RDB に SQL でアクセスするためのライブラリ Kra の紹介
id: ryushi
format: conference
talkType: long_session
level: Intermediate
tags:
- B7-L
speakers:
- ryushi
partner: null
videoId: null
presentation: https://drive.google.com/file/d/1lsF7q74o0Akewk-5rUb9Jt9nA2MqpDDw/view
draft: false
---
Go で PostgreSQL を使ったアプリケーションを実装する際に、単一のトランザクション内で通常の SQL を発行すると共に CopyFrom を使ったバルクインサートができるDBアクセスライブラリKraを紹介します。