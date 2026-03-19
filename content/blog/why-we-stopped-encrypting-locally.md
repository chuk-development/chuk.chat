---
title: "Why We Stopped Encrypting Chat Data Locally"
date: 2026-03-19
description: "We encrypted everything — including the local cache on your device. The decryption key was stored right next to the data. That was a mistake."
author: "Chuk"
---

Chuk Chat encrypts every message with AES-256-GCM before it touches our Supabase backend. End-to-end, zero knowledge, the whole deal. We were proud of that.

So proud, in fact, that we also encrypted the local cache on your device. Every chat stored on disk was encrypted at rest, even though the decryption key lived right next to it in the app's secure storage.

That was a mistake.

## The problem

Loading the chat list means reading every cached conversation. With local encryption enabled, that meant *decrypting* every cached conversation. On a phone with 315 chats, the app had to spin up a background isolate, decrypt each one, and ferry the results back to the main thread.

The numbers were ugly:

| Metric | Before | After |
|--------|--------|-------|
| Chat list load | ~1.5s frame drop | Instant |
| Scroll jank | 5-7% of frames | ~0% |
| Background work | Continuous decrypt queue | None |
| Search | Required full preload | Reads directly |

A 1.5-second frame drop is not a subtle hitch. It is the app freezing while you stare at a blank screen. On mid-range Android hardware, it was worse.

## Why local encryption was security theater

Here is the core issue: the encryption key was stored on the same device as the encrypted data. This is like locking your front door and taping the key to the doorframe.

If an attacker has physical access to your unlocked phone, they can just open the app. If they have root access, they can read the app's secure storage and grab the key. If they have remote access, they can take screenshots or tap the Export button. In every realistic threat scenario, local encryption adds zero protection.

This is not a novel insight. Look at what the major encrypted messengers do:

- **Signal** -- plaintext SQLite database, local storage
- **WhatsApp** -- plaintext SQLite database, local storage
- **iMessage** -- plaintext SQLite database, local storage
- **Telegram** -- plaintext local storage
- **Slack, Discord** -- plaintext local storage

Every single one of them stores chat data unencrypted on the device. Encryption protects data *in transit* and *on the server*. The device's own security model -- Android's app sandbox, iOS's data protection -- handles local isolation. Other apps cannot read your chat database. That is the OS's job, not ours.

The only apps that encrypt local data behind a separate passphrase are password managers. And they work because the decryption key is *your master password*, which lives in your head, not on the device. We were not doing that. We were encrypting with a key the app could silently access at any time.

## What we changed

The fix was simple:

- **Local cache:** now stores plaintext JSON. No decryption step on load.
- **Server storage:** unchanged. Still AES-256-GCM, still end-to-end encrypted, still zero knowledge.
- **Transport:** unchanged. Still TLS + our own encryption layer.

```
Before:
  [Server] --AES-256-GCM--> [Transport] --TLS--> [Device: encrypted cache]
                                                        |
                                                  decrypt on every read
                                                  (isolate, 1.5s jank)

After:
  [Server] --AES-256-GCM--> [Transport] --TLS--> [Device: plaintext cache]
                                                        |
                                                  read directly
                                                  (instant)
```

The security boundary that matters -- protecting your data from anyone who accesses our servers -- is completely unchanged. The only thing we removed is a decryption step that burned CPU cycles to protect data from an attacker who, by definition, already had the key.

## The lesson

More encryption is not always more security. Encryption is a tool with a specific job: keeping data unreadable to anyone without the key. When the key is sitting next to the data, you are not adding security. You are adding latency.

We should have caught this earlier. We didn't, because "encrypt everything" *sounds* right. It took profiling frame drops on a Pixel 6a to make us actually question whether local encryption was doing anything useful.

It wasn't. Now it's gone, and the app is faster for it.
