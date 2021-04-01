---
key: a11-s
title: FlutterとGoを組み合わせたモバイルアプリケーション開発
id: a11-s
format: conference
talkType: short_session
level: beginner
tags:
  - A11-S
speakers:
  - eternal_field
videoId: null
presentation: null
draft: false
---
FlutterはGoogleが開発したクロスプラットフォーム向けモバイルアプリケーションフレームワークです。

しかし、現実的にはiOS/Androidのためにそれぞれコードを書く場合は避けられません。その量が肥大化するにつれてFlutterの魅力は半減してしまいます。

そこで今回はできる限りiOS/Android用にコードを書くことを避けるために、Flutterからgo mobileで作成したiOS/Android用に作成したSDKを利用する方法とそのパターンをご紹介します。

---
# 目標
このトークの目標は「クロスプラットフォームなモバイルアプリ開発でgoの活用できること」を伝えることです。

# 課題
FlutterはGoogleが開発したクロスプラットフォーム向けではあるものの、
過去の資産を利用するためにAndroid/iOS用にそれぞれ別ライブラリや記述することは避けられません。
結果として意図せぬ依存を生むことになってしまい、バグを含む原因となることがあります。
また、Android/iOS用の記述が増加するにつれて、クラスプラットフォームのFlutterの魅力は半減してしまいます。

# 提案
goにはgomobileという協力なツールがあり、Android/iOSから呼び出し可能なSDKを作成することができます。
これを利用することで、従来ではAndroid/iOS用に書かざるを得なかった内容をgoに寄せることが可能となるだけなく、
goの資産をモバイルアプリで活用することが可能になります。
そこで、実際にどのように書けば良いのかなどをサンプルコードを交えつつご紹介します。











