---
key: a3-l
title: Dissecting Slices, Maps and Channels in Go
id: jespinog
format: conference
talkType: long_session
level: Intermediate
tags: []
speakers:
- jespinog
partner: null
videoId: null
presentation: null
draft: false
---
My talk is about how slices, maps and channels work in the Go runtime. The idea is to use unsafe to extract the memory state on runtime and analyze how that is changing over time when you operate with the slices, maps and channels. We will see how channel buffers or maps buckets work under the hood.