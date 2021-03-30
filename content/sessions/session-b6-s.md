---
key: b6-s
title: Dive into Go Protocol Buffers API v2 with the new reflection features
id: b6-s
format: conference
talkType: short_session
level: all
tags:
  - B6-S
speakers:
  - ryoya_sekino
videoId: null
presentation: null
draft: false
---
Protocol Buffers API v2が2020年にリリースされました。v2では、生成されたコードの分割やツールなど開発者のインターフェースに関わる変更に加えて、リフレションのサポートが提供されるようになりました。
新たに追加されたリフレション機能では、Goのビルトインのreflectを使うように、Protocol Bufferのメッセージへの操作・参照を可能にしてくれました。
このトークでは、v2がリリースされた背景やその他の変更点も触れつつ、このリフレション機能に焦点をあてて、サンプルコードを持ちたりソースを追いかけながら紹介していきたいと思います。

---
# Goal of the talk
- Let the audience know what was changed in Go protocol buffers API v2 with that purpose.
- Let the audience have the actual image of what came to be available by the reflection feature of API v2, which is the flagship feature of this update, so that they can use the feature in their development

# Current plan of the agenda
**1.  Introduction**  : About myself and the overview of the talk

**2. What's Protocol Buffers and gRPC** : Very rough summary for the audience who is not very familiar with Protocol Buffers and gRPC. I guess I cannot provide the detailed information due to the limit of the time.

**3. Why was the v2 released?** : Simply explain what was the problem in v1.

**4. What was the remarkable changes on v2?** : I will summarize the important changes on the API, focuses on the reflection.

**5. Dive into reflection**: This would be the most important agenda in my talk. I will explain what is actually available using the sample application or code and picking up a few actual source code of the "protobuf-go" library.

**6. Closing**: Just share the summary with the one slide.
