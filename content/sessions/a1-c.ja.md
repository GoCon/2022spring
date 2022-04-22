---
key: a1-c
title: 人間の直感に対応させた複雑度<Cognitive Complexity>の計測ツールをgo/astで実装してみよう
id: bitkey
format: conference
talkType: sponsor_session
level: Beginner
tags:
- A1-C
speakers:
- norinoriki0
partner: bitkey
videoId: null
presentation: null
draft: false
---
Cognitive Complexity (認知的複雑度) というのは、プログラムの複雑さを測る指標の一つです。以前から使われてきた Cyclomatic Complexity (循環的複雑度) と比べて、人間の理解の難しさにフォーカスしたものであるという特徴があります。

今回の発表では、Go の標準ライブラリである、抽象構文木(AST)を扱うためのライブラリ(go/ast) を使って、どのように Cognitive Complexity の計測ツールを実装できるかについてお話します。静的解析ツールを作るときや go/ast を使うにあたって知っておいたほうが良い情報などについても触れる予定です。go/ast をまだ使ったことがない人が「ちょっと静的解析ツール作ってみような」と思えるようなお話ができればと思っています。