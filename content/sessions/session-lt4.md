---
key: lt4
title: How do I made a powerful cache system using Go
id: lt4
format: conference
talkType: lt_session
level: all
tags:
  - LT4
speakers:
  - sylvain
videoId: null
presentation: null
draft: false
---
I’m Træfik user since v1.4 but there was no caching system. I scrolled over the internet to know if any solution exists but nothing appear then I decided to write my own Træfik cache system.

---
I discovered Go language but didn’t have any time to follow multiples tutorials to learn it. But one day I discovered Træfik reverse-proxy project when I wanted to switch my infrastructure into fully dockerized one. I’m Træfik user since v1.4 but after many months using it I encountered an issue : there were no caching system in this reverse-proxy. I scrolled over the internet to know if any solution exists but nothing appears.

Then I decided to write my own Træfik cache system, but the main question was “Which language?”
- PHP ? Nah.
- Nodejs ? What a joke !
- C++ ? I didn’t learn this language at school and it’s really insane to learn.

Then I was on Træfik github repository when I decided to write it in Go. Another good point: that’s compatible with docker integration.

So I started the project and called it Souin
Let's see together how I bring it up from code to deployment.than no contribution.

Experience:

* It’s very scary being a beginner, but we should never forget that we were all beginners once.
The need for having a fresh perspective is indispensable for the growth of the community. My journey began when I started challenging myself for every possible aspect of organising a meetup:- <br> 
<br>
   - My [thought process]: Never organized any tech-event, how will I do it?!
   - My [action]: I struggled for days, & got the opportunity of conducting a collaborative meet up with other well-established local-tech-communities and finally organized the very first WomenWhoGo, Delhi Chapter meet up on 25th March'18 and now this community is 250+ members strong. 
   - My [thought process]: I have no experience in designing creatives, how will I make it? What if folks will not like it? What should the colours of flyers/posters be?! <br>
   - My [action]: Explored online tools, conducted user interviews and re-iterated process of designing creatives/flyers/posters. Over a period of time, they were loved by many! <br>
   - My [thought process]: I’m not pro in Golang. I can't give talks. Also, I’ve stage fright. So, I just can’t be a speaker ever. <br>
   - My [action]: Asked college seniors/friends/posted on social media if they could give a talk during meet-ups. Sometimes received a positive response but most of the times struggled and I finally realized that I need to trash away my fears. Gave multiple talks, not only at local meetups but at the international platform as well for example at FOSDEM’19/20. <br>


In nutshell, I realised that we have to be the change we seek. Community gives you the abundance of opportunities which requires patience, practise and perseverance. <br>

Nothing makes me happier than if towards the end of my talk, we join hands and pledge to be open to everyone and make everyone feel as ONE and build/nurture our tech-communities! :)
