---
key: a4-s
title: Going places...during an outage
id: a4-s
format: conference
talkType: short_session
level: all
tags:
  - A4-S
speakers:
  - sushant_bhadkamkar
videoId: null
presentation: null
draft: false
---
Building an effective fallback in a distributed system can be hard. Should our systems simply give up during an outage then? Or can we do better?
Join to learn some of the challenges of building a fallback and how we overcame them to reduce business impact, and not leave all our users stranded.

---
Lyft is ride hailing app that helps users find nearby drivers and quickly get to a place they want to go. If Lyft is down, users cannot be matched with drivers and the drivers are not able to make money on the platform. Thus, it is critical we do our best to keep our systems always available.
However, there are several reasons distributed fallback strategies fail in practice:

- they are hard to to simulate and test,
- the strategies themselves might fail or have latent bugs,
- they can make the outage worse.

In this talk we will discuss a best-effort fallback we built in Go that has helped us serve 8000 rides during a sustained cloud provider outage. We hope the audience will learn from our design and practical experience and start thinking about improving resiliency of their systems using similar techniques.
