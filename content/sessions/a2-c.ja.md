---
key: a2-c
title: 創業以来のPHPシステムが生み出した混沌をGoへの移行で乗り越えた話
id: yappli
format: conference
talkType: sponsor_session
level: All
tags:
- A2-C
speakers:
- yosuke_moriya
partner: yappli
videoId: GZCr-Dbi0YM?start=1988
presentation: null
draft: false
---
Yappliは2013年創業の、ノーコードでアプリを作成できるプラットフォームです。
当初サーバーサイドはフレームワークを使わない生のPHPで実装されており、重複コードが多い・テストコードがない・依存が複雑といった様々な問題を抱えおり、2018年よりGoへの移行を始めました。
しかし長年のPHP運用により、DBには大量のJSONカラム・数値型または空文字列を取り得るフィールド・フォーマットの違う時刻など数々の混沌が蔓延っていました。
本発表ではこれらを乗り越えたノウハウを紹介しつつ、またPHPからGoへの移行で組織的に感じられた恩恵についても共有します。