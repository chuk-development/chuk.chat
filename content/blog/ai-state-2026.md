---
title: "We Need to Talk About the Current State of AI"
date: 2026-01-16
description: "It's January 2026. For some reason, all AI chats are getting worse. A frustrated look at multi-billion dollar companies that can't get the basics right."
author: "Chuk"
---

It's January 2026.
For some reason, all AI chats are getting worse.

Btw, always keep in mind that these are multi-billion dollar companies. They've been doing this for years. They have the best programmers money can buy. And yet.

## Grok

When I ask Grok "I need a Z Image Turbo prompt to make an image like that"
it gives me this...

![Grok generates images instead of a prompt](/blog/ai-state-2026/grok-1.avif)

For me that does not look like an image prompt, right?
Btw no system prompt modding or anything, just stock grok.com.

Of course if you then tell it to give you what you asked for, it does the job. By the way, it always uses search even when it doesn't make sense.

![Grok finally gives the prompt but explains what Z Image Turbo is](/blog/ai-state-2026/grok-2.avif)

Let me continue - yes it answered my question now, but I already know what Z Image Turbo is and I didn't ask for an explanation. There are certainly people who would be interested in that, and if I was interested, I would have asked, right?

## Perplexity

Let's move on to Perplexity. Their whole job is search. Let's see how different models handle a simple question about the latest Google Pixel.

**Default Model:**

![Perplexity default model stops at Pixel 9](/blog/ai-state-2026/perplexity-default.avif)

Hm, so as far as I know my question was open and you had to find out what the latest Google Pixel is - and it is the 10 series.
And yes, that is probably not in the model's training data, but your only job is search. Should you not first find out what the latest Google Pixel is?

**Kimi K2:**

![Kimi K2 says Pixel 10 not released yet](/blog/ai-state-2026/perplexity-kimi.avif)

It literally says "Pixel 10: (not released yet)" - which is just false. Both models probably do not have that in their training data, but Perplexity's job is to make it work, right? So maybe tell the model that it doesn't know shit at all, or that it always has to search for things it can't be sure about or that might be updated.

**Gemini 3 Flash:**

![Gemini 3 Flash also fails](/blog/ai-state-2026/perplexity-gemini-flash.avif)

Well, not really. But the leading model Gemini 3 Pro, right?

**Gemini 3 Pro (with reasoning):**

![Gemini 3 Pro with reasoning also fails](/blog/ai-state-2026/perplexity-gemini-pro.avif)

I think that is the AGI we all need!

Same result for Sonnet 4.5 and GPT 5.2 by the way - I'm not going to screenshot every single one, but you can try it yourself.

**Grok:**

![Grok dodges the question](/blog/ai-state-2026/perplexity-grok.avif)

"Pixel 10 and beyond - battery sizes vary by model; I can pull exact figures if you want the latest list" - So even Grok just dodges the question instead of actually searching.

**Deep Research:**

![Deep Research still only goes to Pixel 9a](/blog/ai-state-2026/perplexity-deep-research.avif)

Even Deep Research only goes up to Pixel 9a. Not Pixel 10. For a simple search question.

Ok, and yes - if I tell it "to the latest" explicitly, it finally does the job:

![With explicit to the latest it works](/blog/ai-state-2026/perplexity-to-latest.avif)

Now it shows Pixel 10. But I already said "starting from 7" which implies I want everything up to the current model. Every human would understand that. Why do I have to spell it out?

**Gemini 3 Pro with explicit "to the latest" prompt:**

![Gemini still fails with explicit instructions](/blog/ai-state-2026/perplexity-gemini-to-latest.avif)

Still only goes up to Pixel 9. Even with explicit instructions.

All of these models, tested in Perplexity - a platform whose entire purpose is search - and most of them don't think to actually search for the current information first.

And yes, if you change the prompt a little bit it does work sometimes on some models. But is it my job to do a basic system prompt?

## Voice Mode

Let's talk about voice input. Yes, it has gotten better over the years. Much better actually. But it still messes up so much. It still misunderstands you constantly. Doesn't matter if it's Grok, ChatGPT, or Perplexity - they all have this problem.

And OpenAI? Did you know that their voice mode model still is GPT-4o mini for free users and GPT-4o for paid users? Both released in 2024. We have GPT 5.2 now, which they claim is such a great model. So why is voice mode still on 4o? Not that it makes a big difference.

And then there's Claude. Their voice mode is a complete joke. Actually, do they even have one? They claim to have the best coding model - and I'm not going to argue with that, they probably do - but they can't build a voice mode? Seriously?

## Claude Website

Let's talk about Anthropic's Claude website, specifically on Firefox.

First of all, it takes forever to load. Then when you finally send a message, sometimes it just doesn't go through. The message disappears. Gone. You have to write it again.

And the Artifacts? They don't open properly half the time. You click on them and nothing happens, or they render broken.

This is a company that claims to have one of the best AI models in the world, and they can't make a website that works properly in Firefox?

## Okara AI - "Privacy" Chat?

Now, this is a different category. Okara AI is not a multi-billion dollar company like the others above. The others don't even claim to be privacy-focused. But Okara does. That's their whole selling point.

So my question is: why do you send my location?

![Okara Chain of Thought shows location](/blog/ai-state-2026/okara-1.avif)

![Okara shows exact coordinates](/blog/ai-state-2026/okara-2.avif)

![Okara asks about your location](/blog/ai-state-2026/okara-3.avif)

Not my location btw.

Look at the Chain of Thought: "The user is in Tespe, Germany (based on the location data provided)". It literally shows my coordinates: lat: 53.3991, lon: 10.4128, city: Tespe, country: DE.

And then it asks "How's Tespe treating you today?"

A "Privacy" AI chat that has my GPS coordinates in the system prompt. Right.

And here's the thing: none of their privacy claims are verifiable. They can claim whatever they want. They have a GitHub account, but there's nothing on it. No repositories. Nothing open source. So how am I supposed to trust them?

They also market themselves as using "open source models" - but it's actually open weights, which is not the same thing. Sure, you could argue that's nitpicking, but if you're a "privacy" company, details matter.

Oh, and you can log in with Google. On a privacy-focused platform. And their password system? 5 digits. You can change it. Great.

But it's not just the privacy issues. Their entire platform is bad.

No mobile app. In 2026. Seriously?

And the way they handle image generation is completely broken. You can tell it you want a text model or an image model, but you can't switch naturally in the same chat like every other AI chat. It tries to detect if your prompt is for an image or not, and it gets it wrong constantly. You can't just type a single word and have it understand from context - it will randomly decide to generate an image of the word "point" instead of understanding what you actually mean.

## No Cross-Platform Apps

Here's what really gets me: Why can none of these companies make cross-platform apps for Windows, macOS, and Linux?

All the apps they do make are written in Electron or something similar. It cannot be that hard to compile for all platforms, right? The framework literally supports it.

And here's the real irony: These companies have the most advanced AI models in the world. They claim AI will replace programmers. But they can't build a desktop app that runs on three operating systems?

Seriously?

## Update — February 11, 2026

It keeps going. Of course it does.

## DeepL

DeepL used to be simple. You go to the website, paste your text, get a translation. Done. That was it. That's all you needed.

Now look at this:

![DeepL's new AI platform homepage](/blog/ai-state-2026/deepl.avif)

"AI solutions that help you get work done." That's their headline now. Not "the world's best translator" or anything like that. AI solutions. They have tabs for "Translate speech", "Use API", "Automate tasks". Something called "DeepL Agent" that "helps you automate business tasks using simple language instructions."

But here's what actually gets me: You can only type 500 characters on the homepage. 500. That's like three sentences. Want more? "Open Translator". Want even more? "Start free trial". And if you want to translate more than 5,000 characters, you need to create an account.

You're a translation company. Your business customers pay you a lot of money. Why are you squeezing free users for 5,000 characters? Just let people translate their text. You can argue their product is done, there's nothing more to improve — fine. But instead of keeping it simple and fast, they turned it into this bloated "AI platform" that takes forever to load.

And yes, you can still get to the old translator page if you click through. But the main page — the one you land on when you go to deepl.com — used to be the translator. You'd go there, paste your text, translate. Now the main page is this AI marketing thing and the actual translator is buried behind a click.

I just want to go to DeepL, paste a text, and get a translation. That's it. Stop trying to be another AI company. The world doesn't need another "AI solutions" platform. It needs a fast, simple translator that just works.

## Perplexity — Again

Remember the Perplexity section from January? It gets better.

So Perplexity lets you choose different AI models. Cool feature. You'd expect that when you select a model, it actually uses that model, right?

![Perplexity says Gemini 3 Pro was unavailable](/blog/ai-state-2026/perplexity-model-unavailable.avif)

"Prepared using because Gemini 3 Pro was inapplicable or unavailable"

Read that again. It doesn't even tell you which model it actually used instead. "Prepared using because" — there's literally a word missing in the sentence. Grammar aside, it just silently switches to the default model and tells you after the fact. No warning beforehand, no "hey, this model is currently unavailable, want to pick another one?" — nothing. You write your prompt, you wait, and then at the very bottom it says "oh by the way, we didn't use the model you selected."

Why even offer model selection if you're just going to ignore it?

---

## Conclusion

I could go deeper. I could test more platforms, show more broken features, document more failures. But I think you get the point.

I ask myself this question every single day: Do the people who build these products actually use them? Because it really doesn't feel like it.

Oh, and if you now think "but that's what the majority wants, how the AI responds" - the AI companies get feedback and all that - well then those people are time wasters and dumb. You can argue on the internet all day and never reach a conclusion.

And yes, with all of this you could say "it's my fault and I should have been more detailed" or whatever. But every human would have immediately understood what I mean, and my statements already contain what I mean.

What people need to understand: I don't want an essay about it, I want an answer, a result. I can understand that there's supposed to be some kind of human conversation - but I don't care about that.

Ok sure, some users may like the way the AI responds. But maybe just make a mode where you can turn on short answers or something. Is that too much to ask from a multi-billion dollar company?

This is why I'm building my own AI chat. Not because I think I'm better than these companies. But because I'm frustrated, and I want something that actually listens to user feedback and improves. We'll see how that goes.
