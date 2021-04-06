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
Protocol Buffers API v2が2020年にリリースされました。v2では、コード生成のライブラリなど開発者のインターフェースに関わる変更に加えて、Reflectionのサポートが提供されるようになりました。このReflection機能は、GoのビルトインのReflectionのように、Protocol Buffersのメッセージを動的に操作・参照することを可能にしてくれます。このトークでは、v2がリリースされた背景や変更点にも触れつつ、このReflection機能に焦点をあてて、サンプルコードを用いたりソースを追いかけながら使い方を紹介していきたいと思います。

---
# 得られること
- Protocol Buffers API v2がリリースされた背景と大きな変更点
- protoreflectの機能を使うことで何ができるか、実際にどのように使えばいいのか

# 主な対象者
- Protocol Buffers(gRPC含む)を使っている方
- Protocol Buffers API v2でなにができるようになったのか、とりあえず知っておきたい方
- Protocol Buffers(gRPC含む)に対して、動的なメッセージ操作・参照をしてみたい方
- protoreflectを多少触ってみたけど、まだよくわかっていない方

# 主なアジェンダ
- Why was the API v2 released?
- What was the remarkable changes on v2?
- Dive into protoreflect
  
メインのアジェンダでは、実際のユースケースを想定したコードを用いて、protoreflectの具体的な使い方を説明します。その中で、内部実装やライブラリの全体像についても触れたいと思います。

また、Protocol Buffersをそもそも触ったことがない方にも向けて、最低限ですが、そもそもProtocol Buffersがどんなものであるかも説明したいと思います。
