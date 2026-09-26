---
title: "Anwendungsfälle: Private KI für jede Art von Arbeit"
layout: "usecases"
type: "page"
translationKey: "use-cases"
bodyClass: "uc-page"
description: "Gründer, kleine Betriebe, Entwickler, Studierende: neun echte Anwendungsfälle für Chuk Chat, den privaten KI-Chat aus Deutschland mit Ende-zu-Ende-Verschlüsselung und Open-Weight-Modellen."
keywords: "ki chat anwendungsfälle, ki für kleine unternehmen, ki für gründer, verschlüsselter ki assistent, open-weight ki modelle, mcp connectors, rechnung als pdf mit ki, android ki assistent"

hero_eyebrow: "Anwendungsfälle"
hero_title: "Ein privater Chat.<br>Viele Welten."
hero_lead: "Gründer, kleine Betriebe, Entwickler, Studierende. Alle nutzen dieselbe App für ganz unterschiedliche Arbeit. Scroll nach unten. Jede Welt ist ein echter Anwendungsfall."
scroll_hint: "Scrollen zum Start"
input_placeholder: "Frag mich alles!"
mode_label: "Fast"
worlds_label: "Zu einer Welt springen"
of_label: "von"

worlds:
  - id: "private"
    short: "Privat"
    scene: "night"
    tone: "dark"
    label: "Die privaten Fragen"
    title: "Die Fragen, die du sonst niemandem stellst."
    text: "Gesundheit, Geld, ein schwieriges Gespräch mit dem Chef. Genau das fragen Leute Chuk Chat: Dinge, die sie nie in eine datenhungrige KI tippen würden. Deine Chats werden auf deinem Gerät verschlüsselt, bevor wir sie speichern. Bei uns liegt nur Chiffretext."
    chips: ["Ende-zu-Ende verschlüsselt gespeichert", "Nie fürs Training genutzt", "Kein Tracking"]
    mock:
      ask: "Wie sage ich meinem Chef, dass ich ausgebrannt bin?"
      meta: "4 s nachgedacht"
      lead: "Fang mit Fakten an, nicht mit Vorwürfen."
      answer:
        - "Bitte in einem ruhigen Moment um 20 Minuten."
        - "Sag, was sich verändert hat: Schlaf, Fokus, Krankentage."
        - "Bring eine klare Bitte mit, etwa weniger Projekte für einen Monat."
      stored: "Was unser Server speichert"
      cipher: "AES-256-GCM · der Schlüssel bleibt auf deinem Gerät"

  - id: "founders"
    short: "Gründer"
    scene: "dawn"
    tone: "light"
    label: "Ein Produkt bauen, das man verkauft"
    title: "Von der Idee zur Launch-Seite in einem Chat."
    text: "Gründer recherchieren den Markt, schreiben den Pitch, entwerfen den Code und veröffentlichen eine Landingpage mit öffentlichem Link. Verbinde GitHub, Stripe oder Vercel und arbeite mit deinem echten Projekt. Deine Idee bleibt deine Idee. Wir trainieren nicht damit."
    chips: ["Web-Recherche", "Artefakte mit öffentlichem Link", "GitHub · Stripe · Vercel"]
    mock:
      ask: "Bau eine Landingpage für meinen Fahrrad-Reparaturservice und veröffentliche sie."
      meta: "12 s gearbeitet"
      tool: "create_artifact"
      done: "Fertig. Deine Seite ist online, jeder mit dem Link kann sie öffnen."
      file: "spoke-landing.html"
      tab_preview: "Vorschau"
      tab_code: "Code"
      brand: "Spoke"
      nav: ["Preise", "Gebiete", "Buchen"]
      headline: "Fahrrad-Reparatur vor deiner Tür."
      sub: "Wir kommen zu dir. Festpreise, Termin am selben Tag."
      cta: "Reparatur buchen"
      features: ["Termin am selben Tag", "Festpreise", "Alle Marken"]
      public: "Öffentlicher Link"
      url: "artifacts.chuk.chat/k7f2-spoke"
      copy: "Kopieren"
      connected: "Verbunden"
      connectors: ["github", "stripe", "vercel"]

  - id: "business"
    short: "Kleine Betriebe"
    scene: "day"
    tone: "light"
    label: "Einen kleinen Betrieb führen"
    title: "Der Bürokram, erledigt zwischen zwei Aufträgen."
    text: "Ein Tischler schreibt die Rechnung als sauberes PDF auf dem Weg zum nächsten Kunden. Eine Café-Inhaberin beantwortet Mails und plant die Schichten. Termine landen direkt im Kalender. Ganz ohne IT-Abteilung."
    chips: ["PDF-Dokumente", "E-Mail-Entwürfe", "Kalender und Erinnerungen"]
    mock:
      ask: "Rechnung für Frau Weber: Küchenregal, 6 Stunden à 58 €, Material 140 €. Als PDF."
      meta: "6 s gearbeitet"
      tool: "typst_compile"
      file: "Rechnung_2026-031.pdf"
      download: "Herunterladen"
      doc_title: "Rechnung"
      doc_no: "Nr. 2026-031"
      doc_to: "Frau Weber"
      rows:
        - ["Küchenregal, Arbeitszeit 6 h × 58 €", "348,00 €"]
        - ["Material", "140,00 €"]
      total_label: "Gesamt"
      total: "488,00 €"
      ask2: "Erinner mich Dienstag um 9, sie anzurufen."
      event: "Frau Weber anrufen"
      event_time: "Dienstag · 09:00"
      event_btn: "Im Kalender speichern"

  - id: "connectors"
    short: "Connectors"
    scene: "network"
    tone: "light"
    label: "Alle Tools verbinden"
    title: "Alle deine Tools. Ein Chat, der sie steuert."
    text: "Wissensarbeiter verbinden Notion, Linear, Todoist, Dropbox und 50 weitere Dienste. Du meldest dich einmal im Browser an. Dann fragst du in normalen Worten, und Chuk Chat klickt sich für dich durch alle."
    chips: ["50+ Connectors (MCP)", "Anmeldung per OAuth", "Desktop und Android"]
    mock:
      ask: "Was ist diese Woche fällig? Schau in Linear, Todoist und Notion und mach mir einen Plan."
      meta: "9 s gearbeitet"
      calls:
        - {logo: "linear", name: "Linear", result: "7 offene Issues"}
        - {logo: "todoist", name: "Todoist", result: "12 Aufgaben"}
        - {logo: "notion", name: "Notion", result: "3 Seiten"}
      lead: "Deine Woche, der Reihe nach:"
      answer:
        - "Mo: Login-Bug fixen (Linear, hohe Priorität)"
        - "Di: Q3-Bericht verschicken (Todoist)"
        - "Do: Launch-Plan prüfen (Notion)"
      orbit: ["github", "stripe", "dropbox", "figma", "calcom", "airtable", "zapier", "asana", "sentry", "canva", "supabase", "fastmail", "box", "vercel"]

  - id: "models"
    prompt: "Schreib mein Anschreiben mit Kimi K3."
    short: "Modelle"
    scene: "prism"
    tone: "light"
    label: "Jedes Modell, ein Budget"
    title: "Schluss mit fünf KI-Abos."
    text: "Manche Modelle schreiben besser. Manche denken tiefer. Manche sind schnell und günstig. Power-User wählen für jede Nachricht ein anderes Modell und zahlen aus einem Budget: 20 € im Monat, 16 € davon als KI-Guthaben. Kein zweites Konto, keine zweite Rechnung."
    chips: ["Frontier-Open-Weight-Modelle", "Wechsel pro Nachricht", "16 € KI-Guthaben inklusive"]
    mock:
      modes: ["Fast", "Thinking"]
      models:
        - {logo: "deepseek.svg", name: "DeepSeek V4 Pro"}
        - {logo: "moonshot.svg", name: "Kimi K3"}
        - {logo: "zai.svg", name: "GLM 5.3"}
        - {logo: "qwen.svg", name: "Qwen3.8"}
        - {logo: "minimax.svg", name: "MiniMax M3"}
        - {logo: "mistral.svg", name: "Mistral Small 4"}
        - {logo: "openai.svg", name: "gpt-oss-120b"}
      sent:
        - {text: "Schreib mein Anschreiben", model: 1}
        - {text: "Prüf die Gehaltsrechnung", model: 0}
        - {text: "Übersetz es ins Englische", model: 3}
      credits_label: "KI-Guthaben diesen Monat"
      credits: ["16,00 €", "15,97 €", "15,92 €", "15,90 €"]

  - id: "engineers"
    prompt: "Refactor den Auth-Service und lass den Formatter laufen."
    short: "Entwickler"
    scene: "terminal"
    tone: "dark"
    label: "Viele Aufgaben gleichzeitig"
    title: "Fünf Chats laufen. Keiner wartet auf dich."
    text: "Entwickler starten in einem Chat ein Refactoring, im nächsten eine Recherche und im dritten eine Fehlersuche. Antworten streamen weiter, wenn du den Chat wechselst. Code landet in bearbeitbaren Panels, und auf dem Desktop führt eine Sandbox-Shell Befehle für dich aus."
    chips: ["Paralleles Streaming", "Code-Artefakte", "Sandbox-Shell (Desktop)"]
    mock:
      chats:
        - {title: "Auth-Service refactoren", live: true}
        - {title: "Warum ist der CI-Build langsam?", live: true, finishes: true}
        - {title: "Regex für deutsche IBANs", live: false}
        - {title: "Migration auf Dart 3.13", live: true}
        - {title: "Stacktrace erklären", live: false}
      file: "auth_service.dart"
      shell: "$ dart format lib/"
      shell_out: "Formatted 3 files (1 changed)"
      toast: "„Warum ist der CI-Build langsam?“ hat eine Antwort"

  - id: "prototypers"
    short: "Prototyper"
    scene: "blueprint"
    tone: "dark"
    label: "Der Prototyper"
    title: "Zehn Ideen getestet, noch vor dem Mittag."
    text: "Prototyper skizzieren einen Ablauf als Diagramm, machen daraus eine klickbare HTML-Seite und generieren die Bilder dazu. Jede Idee kostet Minuten statt Tage. Die guten bleiben. Der Rest fliegt raus."
    chips: ["Diagramme (Mermaid, Excalidraw)", "HTML-Prototypen", "Bildgenerierung"]
    mock:
      ask: "Skizzier einen Signup-Ablauf mit E-Mail-Code."
      nodes: ["E-Mail eingeben", "Code gesendet", "Code eingeben", "Willkommen"]
      ask2: "Mach es klickbar. Und ein Hero-Bild dazu."
      diagram: "signup-flow.mmd"
      proto: "signup.html"
      proto_title: "Konto erstellen"
      proto_input: "du@beispiel.de"
      proto_btn: "Code senden"
      image: "hero.png · generiert"

  - id: "research"
    short: "Recherche"
    scene: "library"
    tone: "light"
    label: "Recherche und Dokumente"
    title: "Sprich mit deinen Dokumenten. Antworten mit Quellen."
    text: "Studierende und Forschende hängen PDFs, Notizen und Paper an und stellen Fragen dazu. Chuk Chat sucht im Web, nennt seine Quellen, zeichnet Diagramme und setzt das Ergebnis als sauberes PDF. Alte Chats bleiben durchsuchbar."
    chips: ["PDFs und Dateianhänge", "Websuche mit Quellen", "Diagramme und PDF-Export"]
    mock:
      files: ["masterarbeit_entwurf.pdf", "mueller_2024.pdf"]
      ask: "Vergleich die Methoden beider Paper. Zeig die Rücklaufquoten als Diagramm."
      meta: "14 s gearbeitet"
      sources: ["nature.com", "arxiv.org", "destatis.de", "+ 5"]
      chart_title: "Rücklaufquote nach Befragungsmethode"
      bars:
        - {label: "Online", value: 38}
        - {label: "Telefon", value: 22}
        - {label: "Post", value: 14}
        - {label: "Persönlich", value: 61}
      export: "methodenvergleich.pdf"
      export_note: "Gesetztes PDF · 4 Seiten"

  - id: "onthego"
    short: "Unterwegs"
    scene: "street"
    tone: "dark"
    label: "Unterwegs"
    title: "Home-Taste halten. Fragen. Erledigt."
    text: "Auf Android kann Chuk Chat dein Assistent sein. Halte die Home-Geste gedrückt: Er hört zu, liest mit, was auf dem Bildschirm ist, antwortet laut und handelt. Er stellt Wecker, findet Orte und startet die Navigation."
    chips: ["Android-Assistent", "Sprache rein, Sprache raus", "Orte, Routen, Wecker"]
    mock:
      listening: "Ich höre zu…"
      heard: "Such eine Apotheke, die jetzt offen hat, und bring mich hin."
      place: "Apotheke am Markt"
      place_meta: "400 m · geöffnet bis 20:00"
      action: "Navigation gestartet"
      clock: "18:42"

finale_eyebrow: "Was sie verbindet"
finale_title: "Neun Welten. Eine Regel:<br>Deine Chats gehören dir."
finale_text: "Durch all diese Welten zieht sich ein Faden. Deine Gespräche sind die persönlichsten Daten, die du erzeugst. Deshalb werden sie auf deinem Gerät verschlüsselt, von Open-Weight-Modellen beantwortet und nie fürs Training genutzt. Keine Werbung, kein Tracking, betrieben aus Deutschland."
facts:
  - {icon: "layers", title: "Nur Open-Weight-Modelle", text: "Transparente Modelle. Du weißt immer, was deine KI antreibt."}
  - {icon: "lock", title: "Ende-zu-Ende verschlüsselt", text: "Chats werden auf deinem Gerät verschlüsselt, bevor wir sie speichern."}
  - {icon: "shield", title: "Made in Germany", text: "Betrieben aus Deutschland unter EU-Recht. DSGVO von Anfang an."}
  - {icon: "globe", title: "Kein Tracking, kein Training", text: "Keine Werbung, kein Profiling. Mit deinen Daten wird nichts trainiert."}
cta_title: "Finde deine Welt."
cta_text: "20 € im Monat, 16 € davon als KI-Guthaben. Web, Mac, Windows, Linux und Android."
cta_download: "Chuk Chat herunterladen"
cta_web: "Web App öffnen"
---
