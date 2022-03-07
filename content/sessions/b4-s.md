---
key: b4-s
title: lock free な doubly linked list を実装していたらいつのまにか concurrent skip list map を実装していたでござる
id: kazuhisa_takei
format: conference
talkType: short_session
level: Advanced
tags: []
speakers:
- kazuhisa_takei
partner: null
videoId: null
presentation: null
draft: false
---
埋め込み型のlinux kernel のようなdoubly linked list を実装しだしたら、lock free にしたくなり、そのまま sync.Map に勝つべく, hash map を実装していたらしらないうちに ほぼskip list なhash map を実装するまでの顛末