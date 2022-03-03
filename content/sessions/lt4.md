---
key: lt4
title: 大規模ゲーム開発におけるContext活用パターン
id: 8kka
format: conference
talkType: lt_session
level: All
tags: []
speakers:
- 8kka
partner: null
videoId: null
presentation: null
draft: false
---
GoのContextではリクエストスコープの値を伝播する事ができます。
主な利用例として認証トークンを伝搬させる手法がありますが、ゲーム開発においては他にも様々なContextの利用方法があります。

このトークでは大規模ゲーム開発で利用しているContextの活用パターンについてお話しします。
Contextを活用することで、DBアクセス頻度を少なくしたり、レスポンスサイズを小さくする工夫ができたため、その実装手法を共有します。