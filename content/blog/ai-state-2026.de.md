---
title: "Wir müssen über den aktuellen Zustand von KI reden"
date: 2026-01-16
description: "Es ist Januar 2026. Aus irgendeinem Grund werden alle KI-Chats immer schlechter. Ein frustrierter Blick auf Multi-Milliarden-Dollar-Unternehmen, die die Basics nicht hinbekommen."
author: "Chuk"
---

Es ist Januar 2026.
Aus irgendeinem Grund werden alle KI-Chats immer schlechter.

Übrigens: Behaltet immer im Hinterkopf, dass das Multi-Milliarden-Dollar-Unternehmen sind. Die machen das seit Jahren. Die haben die besten Programmierer, die man für Geld kaufen kann. Und trotzdem.

## Grok

Wenn ich Grok frage "I need a Z Image Turbo prompt to make an image like that"
gibt es mir das...

![Grok generiert Bilder statt einen Prompt](/blog/ai-state-2026/grok-1.png)

Für mich sieht das nicht wie ein Image-Prompt aus, oder?
Übrigens kein System-Prompt-Modding oder sowas, einfach nur grok.com.

Natürlich macht es den Job, wenn man ihm dann sagt, es soll einem geben, worum man gefragt hat. Übrigens nutzt es immer die Suche, auch wenn es keinen Sinn macht.

![Grok gibt endlich den Prompt, erklärt aber erstmal was Z Image Turbo ist](/blog/ai-state-2026/grok-2.png)

Lass mich weitermachen - ja, es hat jetzt meine Frage beantwortet, aber ich weiß bereits was Z Image Turbo ist und hatte nicht danach gefragt, es mir zu erklären. Es gibt bestimmt Leute, die das interessiert, und würde mich das interessieren, dann hätte ich danach gefragt, oder?

## Perplexity

Kommen wir zu Perplexity. Deren einziger Job ist Suche. Schauen wir mal, wie verschiedene Modelle eine einfache Frage zum neuesten Google Pixel handhaben.

**Standard-Modell:**

![Perplexity Standard-Modell stoppt bei Pixel 9](/blog/ai-state-2026/perplexity-default.png)

Hm, soweit ich weiß war meine Frage offen und du musstest herausfinden, was das neueste Google Pixel ist - und das ist die 10er Serie.
Und ja, das ist wahrscheinlich nicht in den Trainingsdaten des Modells, aber dein einziger Job ist Suche. Solltest du nicht zuerst herausfinden, was das neueste Google Pixel ist?

**Kimi K2:**

![Kimi K2 sagt Pixel 10 noch nicht erschienen](/blog/ai-state-2026/perplexity-kimi.png)

Es sagt wörtlich "Pixel 10: (not released yet)" - was einfach falsch ist. Beide Modelle haben das wahrscheinlich nicht in ihren Trainingsdaten, aber Perplexitys Job ist es, das zum Laufen zu bringen, oder? Also sagt dem Modell vielleicht, dass es gar nichts weiß, oder dass es Sachen, bei denen es nicht sicher sein kann oder die aktualisiert wurden, immer nachsuchen muss.

**Gemini 3 Flash:**

![Gemini 3 Flash scheitert auch](/blog/ai-state-2026/perplexity-gemini-flash.png)

Na, nicht wirklich. Aber das führende Modell Gemini 3 Pro, oder?

**Gemini 3 Pro (mit Reasoning):**

![Gemini 3 Pro mit Reasoning scheitert auch](/blog/ai-state-2026/perplexity-gemini-pro.png)

Ich denke, das ist die AGI, die wir alle brauchen!

Gleiches Ergebnis bei Sonnet 4.5 und GPT 5.2 übrigens - ich werde nicht jeden einzelnen screenshotten, aber ihr könnt es selbst ausprobieren.

**Grok:**

![Grok weicht der Frage aus](/blog/ai-state-2026/perplexity-grok.png)

"Pixel 10 and beyond - battery sizes vary by model; I can pull exact figures if you want the latest list" - Also weicht sogar Grok einfach der Frage aus, anstatt tatsächlich zu suchen.

**Deep Research:**

![Deep Research funktioniert endlich](/blog/ai-state-2026/perplexity-deep-research.png)

Endlich eine vollständige Tabelle mit Pixel 10. Aber ich musste Deep Research für eine einfache Frage nutzen, die jede Suche hätte beantworten sollen.

Ok, und ja - wenn ich explizit "to the latest" sage, macht es den Job:

![Mit explizitem to the latest funktioniert es](/blog/ai-state-2026/perplexity-to-latest.png)

Aber schau dir den vorherigen Screenshot an - ich habe bereits "starting from 7" gesagt, was impliziert, dass ich alles bis zum aktuellen Modell will. Jeder Mensch würde das verstehen.

**Gemini 3 Pro mit explizitem "to the latest" Prompt:**

![Gemini scheitert trotzdem mit expliziten Anweisungen](/blog/ai-state-2026/perplexity-gemini-to-latest.png)

Geht immer noch nur bis Pixel 9. Selbst mit expliziten Anweisungen.

All diese Modelle, getestet in Perplexity - einer Plattform, deren ganzer Zweck Suche ist - und die meisten von ihnen denken nicht daran, zuerst die aktuellen Informationen zu suchen.

Und ja, wenn du den Prompt ein bisschen änderst, funktioniert es manchmal bei manchen Modellen. Aber ist es mein Job, einen grundlegenden System-Prompt zu machen?

## Voice Mode

Reden wir über Spracheingabe. Ja, sie ist über die Jahre besser geworden. Sogar viel besser. Aber sie vermasselt immer noch so viel. Sie versteht dich immer noch ständig falsch. Egal ob Grok, ChatGPT oder Perplexity - alle haben dieses Problem.

Und OpenAI? Wusstet ihr, dass deren Voice-Mode-Modell immer noch GPT-4o mini ist, veröffentlicht am 18. Juli 2024? Wir haben jetzt GPT 5.2, das sie als so großartiges Modell bewerben. Warum ist Voice Mode dann immer noch auf 4o mini?

Und dann gibt es Claude. Deren Voice Mode ist ein kompletter Witz. Haben die überhaupt einen? Sie behaupten, das beste Coding-Modell zu haben - und ich werde dem nicht widersprechen, wahrscheinlich stimmt das - aber sie können keinen Voice Mode bauen? Ernsthaft?

## Claude Website

Reden wir über Anthropics Claude-Website, speziell in Firefox.

Erstens dauert es ewig zu laden. Wenn du dann endlich eine Nachricht sendest, kommt sie manchmal einfach nicht an. Die Nachricht verschwindet. Weg. Du musst sie nochmal schreiben.

Und die Artifacts? Die öffnen sich die Hälfte der Zeit nicht richtig. Du klickst drauf und nichts passiert, oder sie werden kaputt gerendert.

Das ist ein Unternehmen, das behauptet, eines der besten KI-Modelle der Welt zu haben, und sie können keine Website bauen, die in Firefox richtig funktioniert?

## Okara AI - "Privacy" Chat?

Jetzt ist das eine andere Kategorie. Okara AI ist kein Multi-Milliarden-Dollar-Unternehmen wie die anderen oben. Die anderen behaupten nicht mal, datenschutzorientiert zu sein. Aber Okara schon. Das ist deren ganzer Verkaufspunkt.

Also meine Frage: Warum sendest du meinen Standort?

![Okara Chain of Thought zeigt Standort](/blog/ai-state-2026/okara-1.png)

![Okara zeigt exakte Koordinaten](/blog/ai-state-2026/okara-2.png)

![Okara fragt nach deinem Standort](/blog/ai-state-2026/okara-3.png)

Nicht mein Standort übrigens.

Schau dir die Chain of Thought an: "The user is in Tespe, Germany (based on the location data provided)". Es zeigt buchstäblich meine Koordinaten: lat: 53.3991, lon: 10.4128, city: Tespe, country: DE.

Und dann fragt es "How's Tespe treating you today?"

Ein "Privacy" KI-Chat, der meine GPS-Koordinaten im System-Prompt hat. Richtig.

Und hier ist das Ding: Keine ihrer Datenschutz-Behauptungen ist verifizierbar. Sie können behaupten, was sie wollen. Sie haben einen GitHub-Account, aber da ist nichts drauf. Keine Repositories. Nichts Open Source. Wie soll ich ihnen also vertrauen?

Sie vermarkten sich auch als Nutzer von "Open Source Modellen" - aber es ist eigentlich Open Weights, was nicht dasselbe ist. Klar, man könnte argumentieren, das sei Haarspalterei, aber wenn du ein "Privacy"-Unternehmen bist, zählen Details.

Oh, und du kannst dich mit Google einloggen. Auf einer datenschutzorientierten Plattform. Und deren Passwort-System? 5 Stellen. Du kannst es ändern. Großartig.

Aber es sind nicht nur die Datenschutz-Probleme. Deren gesamte Plattform ist schlecht.

Keine Handy-App. In 2026. Ernsthaft?

Und die Art, wie sie Bildgenerierung handhaben, ist komplett kaputt. Du kannst sagen, dass du ein Text-Modell oder ein Bild-Modell willst, aber du kannst nicht natürlich im selben Chat wechseln wie bei jedem anderen KI-Chat. Es versucht zu erkennen, ob dein Prompt für ein Bild ist oder nicht, und liegt dabei ständig falsch. Du kannst nicht einfach ein einzelnes Wort tippen und erwarten, dass es aus dem Kontext versteht - es wird willkürlich entscheiden, ein Bild des Wortes "Punkt" zu generieren, anstatt zu verstehen, was du eigentlich meinst.

## Keine Cross-Platform Apps

Hier ist, was mich wirklich ärgert: Warum kann keines dieser Unternehmen Cross-Platform-Apps für Windows, macOS und Linux machen?

Alle Apps, die sie machen, sind in Electron oder sowas geschrieben. Es kann doch nicht so schwer sein, für alle Plattformen zu kompilieren, oder? Das Framework unterstützt es buchstäblich.

Und hier ist die echte Ironie: Diese Unternehmen haben die fortschrittlichsten KI-Modelle der Welt. Sie behaupten, KI werde Programmierer ersetzen. Aber sie können keine Desktop-App bauen, die auf drei Betriebssystemen läuft?

Ernsthaft?

## Fazit

Ich könnte tiefer gehen. Ich könnte mehr Plattformen testen, mehr kaputte Features zeigen, mehr Fehler dokumentieren. Aber ich denke, ihr versteht den Punkt.

Ich frage mich jeden einzelnen Tag: Nutzen die Leute, die diese Produkte bauen, sie eigentlich selbst? Weil es sich wirklich nicht so anfühlt.

Oh, und wenn du jetzt denkst "aber das ist, was die Mehrheit will, wie die KI antwortet" - die KI-Unternehmen bekommen Feedback und so weiter - na dann sind diese Leute Zeitverschwender und dumm. Man kann im Internet den ganzen Tag diskutieren und nie zu einem Ergebnis kommen.

Und ja, bei all dem könnte man sagen "es ist meine Schuld und ich hätte detaillierter sein müssen" oder so. Aber jeder Mensch hätte sofort verstanden, was ich meine, und meine Aussagen beinhalten auch das, was ich meine.

Was die Leute verstehen müssen: Ich will keinen Aufsatz darüber, ich will eine Antwort, ein Ergebnis. Ich kann verstehen, dass es eine Art menschliche Konversation geben soll - aber das interessiert mich nicht.

Ok klar, manche Nutzer mögen vielleicht die Art, wie die KI antwortet. Aber macht vielleicht einfach einen Modus, wo man kurze Antworten einschalten kann oder so. Ist das zu viel verlangt von einem Multi-Milliarden-Dollar-Unternehmen?

Deshalb baue ich meinen eigenen KI-Chat. Nicht weil ich denke, ich bin besser als diese Unternehmen. Sondern weil ich frustriert bin, und ich etwas will, das tatsächlich auf Nutzer-Feedback hört und sich verbessert. Wir werden sehen, wie das läuft.
