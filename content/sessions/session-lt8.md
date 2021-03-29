---
key: lt8
title: CloudFirestoreとGo
id: lt8
format: conference
talkType: lt_session
level: intermediate
tags:
  - LT8
speakers:
  - mtskhs
videoId: null
presentation: null
draft: false
---
FirebaseOpenSourceにも掲載されているFirestoreの運用ツール[fsrpl](https://firebaseopensource.com/projects/matsu0228/fsrpl/)を開発した発表者が、FirestoreをGoで扱うための知見や、ツールfsrplについて発表する。Firestoreは近年アプリ開発においてますます重要となっているが、Goにおける知見やツールはまだ少ない。

本発表では知見として、FirestoreをGoで扱うための基礎やツールfsrplを紹介する。

---
## Purpose

- CloudFirestore(以降、Firestoreと呼ぶ)をGoで扱うための知見共有
  - 公式ライブラリ Firebase Admin Go SDK の使い方
  - NoSQLであるFirestoreをGoらしく扱うための方法
  - RDBとの違いを踏まえて開発する方法

- Firestoreの運用開発における課題と、開発したCLIツール[fsrpl](https://github.com/matsu0228/fsrpl)の紹介
- Firestoreだけに限らず、Goにおけるデータ変換テクニックや、NoSQLをGoで扱う方法といった一般的な知見の共有

## Why Firestore

- Firestoreは、ドキュメント指向のNoSQLデータベースである。特徴としてWebsocket通信によるリアルタイムでのデータ連携が可能であり、2019年初めにβ版から正式版となり東京リージョンが増えたことでアプリ開発における重要性が高まっている。
  - 一方で、まだ新しい技術であるため、商用利用している人もまだ少なく、Goで扱うための知見やツールについても、まだまとまったものはないように思う。
  - 発表者は、業務にてFirestoreをメインDBとして開発した商用アプリ[SpoLive](https://spolive.ntt.com/)において、主にそのバックエンド(Firestoreへのデータ連携)を担当している。この開発における知見を共有する。
