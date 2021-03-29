---
key: a8-s
title: Go製のネットワーククライアントに対する継続的Fuzzing
id: a8-s
format: conference
talkType: short_session
level: intermediate
tags:
  - A8-S
speakers:
  - yuki_ito
videoId: null
presentation: null
draft: false
---
Fuzzing とは fuzz と呼ばれる予測不可能なデータをインプットとして与えることによりその不具合を発見するテスト手法です。このセッションでは Fuzzing の概要、Goで書かれたソフトウェアに対する Fuzzing の手法、クレジットカード決済のためのISO8583というプロトコルを話すGoで書かれたネットワーククライアントに対して実践している継続的Fuzzingの実例について紹介します。

---
1. Fuzzing とは
2. Go で書かれたソフトウェアに対する Fuzzing
3. 継続的に Fuzzing を行っている実例と課題
