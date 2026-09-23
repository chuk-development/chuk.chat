---
title: "Datenschutzerklärung"
layout: "legal"
type: "page"
translationKey: "privacy"
updated: "Zuletzt aktualisiert: 23. September 2026"
tldr: "Ihre Chats werden auf Ihrem Gerät Ende-zu-Ende-verschlüsselt, bevor sie gespeichert werden. Wir können Ihre Nachrichten nicht lesen. Wir sammeln nur die minimal notwendigen Daten zur Bereitstellung des Dienstes."
---

## 1. Einleitung

Diese Datenschutzerklärung erklärt, wie chuk.chat ("wir", "uns" oder "unser") Ihre Informationen sammelt, verwendet und schützt, wenn Sie unsere KI-Chat-Anwendung und Dienste nutzen.

Wir nehmen Ihre Privatsphäre ernst und haben unseren Dienst mit Datenschutz als Kernprinzip entwickelt.

## 2. Informationen, die wir sammeln

### 2.1 Kontoinformationen

Wenn Sie ein Konto erstellen, sammeln wir:

- E-Mail-Adresse (in Klartext in der Supabase-Datenbank gespeichert)
- Anzeigename (optional, in Klartext gespeichert)
- Passwort (als sicherer Hash gespeichert, niemals in Klartext)

**Wichtig:** Ihre E-Mail und Ihr Benutzername werden im Standarddatenbankformat (nicht verschlüsselt) gespeichert, da dies für die Authentifizierung und E-Mail-Kommunikation erforderlich ist. Nur Ihre Chat-Nachrichten werden verschlüsselt.

### 2.2 Chat-Daten

Alle Ihre Chat-Nachrichten und Konversationen werden auf Ihrem Gerät mit AES-256-GCM-Verschlüsselung Ende-zu-Ende-verschlüsselt, bevor sie an Supabase gesendet werden. Das bedeutet:

- Ihr Verschlüsselungsschlüssel wird von Ihrem Passwort abgeleitet und verlässt niemals Ihr Gerät
- Wir speichern nur verschlüsselte Daten und können Ihre Nachrichten nicht entschlüsseln
- Nur Sie können Ihren Chat-Verlauf lesen

### 2.3 Nutzungsinformationen

Wir sammeln die folgenden Nutzungsdaten für Abrechnungs- und Dienstbetriebszwecke:

- Abonnementstatus und Zahlungsinformationen (sicher über Stripe verarbeitet)
- **Token-Nutzungsdaten:** Für jede KI-Nachricht protokollieren wir:
  - Eingabe-Tokens (Prompt-Länge)
  - Ausgabe-Tokens (Antwortlänge)
  - Verwendeter Modellname (z.B. "deepseek-r1-0528", "qwen3-235b-a22b")
  - Anbietername (z.B. "DeepSeek", "Qwen")
  - Zeitstempel

Diese Daten werden in Supabase gespeichert, um Ihre Nutzung und Abrechnung zu berechnen. Wir speichern NICHT den tatsächlichen Nachrichteninhalt in diesen Protokollen - nur die Token-Anzahl und Metadaten.

**Betriebsmetriken (PostHog):** Um den Dienst stabil zu betreiben, senden wir pseudonyme technische Ereignisse an PostHog: verwendetes Modell und Anbieter, Token-Anzahl, Latenz, Kosten, HTTP-Status, Fehlertyp, API-Endpunkt und App-Version. Ihr Konto wird dabei durch einen verschlüsselten Hash (HMAC) ersetzt, den PostHog nicht auf Sie zurückführen kann. Diese Ereignisse enthalten nie Ihre Prompts, KI-Antworten, E-Mail-Adresse, Ihren Namen, Ihre IP-Adresse oder Zahlungskennungen.

**Wir sammeln NICHT:** Werbe-IDs, Geräte-Fingerprints oder den Inhalt Ihrer Nachrichten zu Analysezwecken. Diese Website enthält kein Tracking-Skript und die Apps kein Analyse-SDK.

### 2.4 Drittanbieter-KI-Dienste

Wenn Sie KI-Funktionen verwenden, werden Ihre Daten von spezialisierten Drittanbieterdiensten verarbeitet:

- **Textgenerierung (LLMs):** Nachrichten werden entweder an OpenRouter gesendet, das sie nur an Anbieter von open-weight KI-Modellen weiterleitet, oder direkt an einen dieser Inferenz-Anbieter: DeepInfra, Baseten, Fireworks AI, OrcaRouter oder RunAnywhere. Über OpenRouter können auch weitere Inferenz-Anbieter eine Anfrage verarbeiten. **Wir unterstützen nur open-weight Modelle** (z.B. Llama, Mistral, Qwen, DeepSeek). Geschlossene/proprietäre Modelle (Claude, GPT-4, Gemini) sind nicht verfügbar. **Modelltraining ist deaktiviert**.
- **Sprach-zu-Text:** Audio wird von Whisper auf der Groq-Infrastruktur zur Transkription verarbeitet.
- **Bildgenerierung:** Bild-Prompts und bei der Bildbearbeitung das Ausgangsbild werden von Replicate verarbeitet. Die erzeugten Bilder werden an Sie ausgeliefert.
- **Websuche:** Wenn die KI im Web sucht, wird die Suchanfrage an die Brave-Search-API gesendet. Ihre Kontodaten werden dabei nicht übermittelt.
- **Embeddings:** Text, der nach Bedeutung durchsuchbar sein soll, wird von DeepInfra oder Fireworks AI in Vektoren umgewandelt.
- **Text-zu-Sprache:** Text wird mit Inworld TTS in Sprache umgewandelt.
- **Sprach- und Videomodi (Demnächst):** Echtzeit-Sprach- und Videokommunikation wird über die LiveKit-Infrastruktur verarbeitet.

## 3. Wie wir Ihre Informationen verwenden

Wir verwenden gesammelte Informationen, um:

- Den Dienst bereitzustellen und aufrechtzuerhalten
- Die Nutzung zu berechnen und Abrechnung basierend auf Token-Verbrauch zu generieren
- Ihr Abonnement und Ihre Zahlungen zu verarbeiten
- Wichtige Dienstupdates und Sicherheitsbenachrichtigungen zu senden
- Kundensupport bereitzustellen
- Betrug oder Missbrauch zu erkennen und zu verhindern

**Wir verkaufen Ihre persönlichen Informationen nicht an Dritte.**

## 4. Rechtsgrundlage der Verarbeitung (DSGVO Art. 6)

Wir verarbeiten Ihre personenbezogenen Daten auf folgenden Rechtsgrundlagen:

- **Vertragserfüllung (Art. 6 Abs. 1 lit. b DSGVO):** Die Verarbeitung von Kontodaten, Chat-Daten, Zahlungsinformationen und Nutzungsdaten ist zur Erbringung unseres Dienstes und zur Erfüllung unserer vertraglichen Verpflichtungen erforderlich.
- **Berechtigtes Interesse (Art. 6 Abs. 1 lit. f DSGVO):** Wir verarbeiten Daten zur Betrugserkennung, für Sicherheitsmaßnahmen und zur Dienstverbesserung. Dazu gehören auch die pseudonymen Betriebsmetriken an PostHog. Unser berechtigtes Interesse ist die Aufrechterhaltung eines sicheren und funktionsfähigen Dienstes.
- **Rechtliche Verpflichtung (Art. 6 Abs. 1 lit. c DSGVO):** Wir können Daten verarbeiten, um gesetzliche Anforderungen wie Steuervorschriften und Anfragen von Strafverfolgungsbehörden zu erfüllen.
- **Einwilligung (Art. 6 Abs. 1 lit. a DSGVO):** Wo erforderlich, holen wir Ihre ausdrückliche Einwilligung vor der Verarbeitung ein. Sie können Ihre Einwilligung jederzeit widerrufen.

## 5. Datensicherheit

Wir implementieren branchenübliche Sicherheitsmaßnahmen zum Schutz Ihrer Daten:

- **Ende-zu-Ende-Verschlüsselung:** Alle Chat-Daten werden mit AES-256-GCM vor der Speicherung verschlüsselt
- **Sichere Authentifizierung:** Passwörter werden mit bcrypt mit starken Arbeitsfaktoren gehasht
- **HTTPS/TLS:** Alle Datenübertragungen werden mit TLS 1.3 verschlüsselt
- **Sichere Infrastruktur:** Gehostet auf Supabase mit SOC 2 Type II-Konformität

## 6. Datenspeicherung

- **Kontodaten:** Gespeichert, bis Sie Ihr Konto löschen
- **Chat-Daten:** Verschlüsselt gespeichert, bis Sie einzelne Chats manuell löschen
- **Token-Nutzungsprotokolle:** Für Abrechnungszwecke gespeichert und gelöscht, wenn Ihr Konto gelöscht wird
- **Zahlungsinformationen:** Sicher von Stripe gespeichert

## 7. Ihre Rechte

Sie haben das Recht auf:

- **Zugang:** Eine Kopie Ihrer personenbezogenen Daten anzufordern
- **Löschung:** Ihr Konto und alle zugehörigen Daten jederzeit zu löschen
- **Korrektur:** Ihre Kontoinformationen zu aktualisieren
- **Export:** Ihre Chat-Daten (entschlüsselt) aus der App herunterzuladen

Um diese Rechte auszuüben, kontaktieren Sie uns unter [privacy@chuk.chat](mailto:privacy@chuk.chat) oder verwenden Sie die In-App-Kontoeinstellungen.

- **Beschwerde:** Sie haben das Recht, eine Beschwerde bei einer Aufsichtsbehörde einzureichen. Die für uns zuständige Aufsichtsbehörde ist:
  Unabhängiges Landeszentrum für Datenschutz Schleswig-Holstein (ULD)
  Holstenstraße 98, 24103 Kiel
  https://www.datenschutzzentrum.de
  E-Mail: mail@datenschutzzentrum.de

## 8. Cookies und Tracking

Unsere Website verwendet minimale Cookies für wesentliche Funktionen:

- **Authentifizierungs-Cookies:** Um Sie eingeloggt zu halten
- **Präferenz-Cookies:** Um Ihr Theme und Ihre Einstellungen zu speichern

Wir verwenden keine Werbe-Cookies oder Tracking-Skripte von Drittanbietern.

## 9. Drittanbieterdienste

Wir verwenden die folgenden Drittanbieterdienste:

- **OpenRouter:** KI-Modell-Routing. [Datenschutz](https://openrouter.ai/privacy)
- **DeepInfra:** KI-Inferenz und Embeddings, direkt und über OpenRouter. [Datenschutz](https://deepinfra.com/privacy)
- **Baseten:** KI-Inferenz, direkt und über OpenRouter. [Datenschutz](https://www.baseten.co/privacy-policy/)
- **Fireworks AI:** KI-Inferenz und Embeddings. [Datenschutz](https://fireworks.ai/privacy-policy)
- **OrcaRouter:** KI-Inferenz. [Website](https://orcarouter.ai)
- **RunAnywhere:** KI-Inferenz. [Website](https://runanywhere.ai)
- **Replicate:** Bildgenerierung und -bearbeitung. [Datenschutz](https://replicate.com/privacy)
- **Brave Search:** Websuche für KI-Antworten. [Datenschutz](https://search.brave.com/help/privacy-policy)
- **PostHog:** Pseudonyme Betriebsmetriken, siehe Abschnitt 2.3. [Datenschutz](https://posthog.com/privacy)
- **Groq:** Sprach-zu-Text. [Datenschutz](https://groq.com/privacy-policy/)
- **Inworld:** Text-zu-Sprache. [Datenschutz](https://www.inworld.ai/privacy-policy)
- **Supabase:** Datenbank und Authentifizierung. [Datenschutz](https://supabase.com/privacy)
- **Stripe:** Zahlungsabwicklung. [Datenschutz](https://stripe.com/privacy)
- **Lexoffice:** Rechnungserstellung. [Datenschutz](https://www.lexoffice.de/datenschutz/)
- **Hetzner:** API-Server-Hosting. [Datenschutz](https://www.hetzner.com/legal/privacy-policy)

**Übermittlung in Drittländer:** Mehrere dieser Dienste (zum Beispiel OpenRouter, DeepInfra, Baseten, Fireworks AI, Replicate, Groq und PostHog) haben ihren Sitz in den USA. Die Übermittlung an diese Anbieter erfolgt auf Grundlage des EU-US Data Privacy Framework, sofern der Anbieter zertifiziert ist, und andernfalls auf Grundlage der Standardvertragsklauseln der EU-Kommission (Art. 46 Abs. 2 lit. c DSGVO).

## 10. Datenschutz von Kindern

Unser Dienst ist nicht für Kinder unter 16 Jahren bestimmt. Wir sammeln wissentlich keine personenbezogenen Daten von Kindern unter 16 Jahren. Wenn Sie glauben, dass wir Informationen von einem Kind unter 16 Jahren gesammelt haben, kontaktieren Sie uns bitte umgehend.

Gemäß Art. 8 DSGVO müssen Nutzer in der Europäischen Union mindestens 16 Jahre alt sein, um unseren Dienst zu nutzen.

## 11. Änderungen dieser Datenschutzerklärung

Wir können diese Datenschutzerklärung von Zeit zu Zeit aktualisieren. Wir werden Sie über wesentliche Änderungen per E-Mail oder über die App informieren. Die fortgesetzte Nutzung des Dienstes nach Änderungen gilt als Akzeptanz der aktualisierten Richtlinie.

## 12. Kontakt

E-Mail: [privacy@chuk.chat](mailto:privacy@chuk.chat)
Support: [support@chuk.chat](mailto:support@chuk.chat)
