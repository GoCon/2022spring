---
key: lt1
title: Goの標準機能で簡易システムを低コストで作成するテクニック
id: yuji_shimada
format: conference
talkType: lt_session
level: Beginner
tags: []
speakers:
- yuji_shimada
partner: null
videoId: null
presentation: null
draft: false
---
Goによるゲーム開発のプロジェクトにて、Go標準の機能を複数活用し、実装コストの低いデバッグツール開発基盤を作成しました。
本セッションのデバッグツールとは、動作確認やデバッグのため、ユーザーのレベル操作や所持アイテムの増減等、データの状態を操作・参照するAPI群、及びwebツールを指します。
デバッグツールの拡充は、ゲームに限らず様々な分野で成果物の品質に直結しますが、ドメイン固有のロジック実装に集中するためにはより低コストで開発できることが望ましいです。
本セッションでは、低コストで開発を行えるデバッグツール基盤をどのように実現したのかを紹介します。