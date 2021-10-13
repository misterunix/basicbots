# **basicbots**

## What is it?
**basicbots** is a programming game where you write a program in *BASIC* and have it compete against other robots.

## UPDATE: 
- Endgame fixed.
- Teams is working!
  - demo and docs coming soon!

## Releases
- [v0.0.2b](https://github.com/misterunix/basicbots/releases/tag/v0.0.2b)
- [v0.0.2a](https://github.com/misterunix/basicbots/releases/tag/v0.0.2a)
- [v0.0.1c](https://github.com/misterunix/basicbots/releases/tag/v0.0.1c-beta)
- [v0.0.1b](https://github.com/misterunix/basicbots/releases/tag/v0.0.1b-alpha)
- [v0.0.1a](https://github.com/misterunix/basicbots/releases/tag/v0.0.1a-test-release)

## Player Documentation 
- [BASICBOT](BASICBOTS.md) - Not complete.
- [BASIC Commands](BASIC.md) - Quick notes, not complete.

## Example bots (most complex so far)
- Free for all
  - [blaster.bas](https://github.com/misterunix/basicbots/blob/main/robots/blaster.bas)
- Teams
  - [shooter.bas](https://github.com/misterunix/basicbots/blob/main/testbots/shooter.bas)
---

<p align="center">
<img width="100%" height="100%" src="images/current.gif">
</p>

## Example match results

     ./basicbots -m 311 robots/blaster.bas robots/corner-runner.bas robots/rook.bas robots/nexus.bas 
     blaster.bas          w:00294 t:00001 l:00016 p:00883
     corner-runner.bas    w:00005 t:00002 l:00304 p:00017
     rook.bas             w:00009 t:00000 l:00302 p:00027
     nexus.bas            w:00001 t:00001 l:00309 p:00004


blaster.bas is out preforming my expections.

## Teams with simple robots.

<p align="center">
  <img width="100%" height="100%" src="images/example6.gif">
</p>

## Match with teams
```
./basicbots -t -m 111 testbots/shooter.bas testbots/teamtest.bas testbots/shooter.bas testbots/teamtest.bas
shooter.bas          w:00052 t:00002 l:00057 p:00156
teamtest.bas         w:00052 t:00001 l:00057 p:00131
shooter.bas          w:00057 t:00001 l:00052 p:00172
teamtest.bas         w:00057 t:00002 l:00052 p:00143
Team1                w:00104 t:00003 l:00114 p:00287
Team2                w:00114 t:00003 l:00104 p:00315
```


## Play Testers
- Even if you don't know *BASIC* it's easy to learn. There are resources on the internet about programming in *BASIC*, look for the really old ones. I and as time goes on, others will be more than happy to help out.

## Contributors
- Documentation
- Programming
  - Bug fixes
  - Features
  - Enhancements
  - Commenting
 
## Who am I?
I am an old-school type of guy. I enjoy games from a time when graphics wasn't the main focus of a game and when you died, you had to start over. 

I can be reached via email *misterunix@gmail.com* or on Twitter as *misterunix*. 
