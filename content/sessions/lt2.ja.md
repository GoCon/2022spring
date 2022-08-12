---
key: lt2
title: Python製の姓名分割ライブラリをGoに移植した話
id: glassmonekey
format: conference
talkType: lt_session
level: All
tags:
- LT2
speakers:
- glassmonekey
partner: null
videoId: GZCr-Dbi0YM?start=28535
presentation: https://speakerdeck.com/nagano/pythonzhi-falsexing-ming-fen-ge-raiburariwogoniyi-zhi-sitahua
draft: false
---
一般的にわかち書きでは無い日本語で姓名から「姓＋名」の分割を行うことは困難です。
しかし、Python製の姓名分割ライブラリ(https://github.com/rskmoi/namedivider-python)を用いるとある程度精度良く分割は可能です。
そこでシングルバイナリで扱えるGoのメリットを活かして、Python製の姓名分割ライブラリをGoに移植した話をします。
その際移植で工夫した点や気をつけた点をお話します。