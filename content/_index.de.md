---
bodyClass: "home uc-page"
title: "Chuk Chat — Privater KI-Chat aus Deutschland | Sicher & Anonym"
description: "Chuk Chat ist ein privater KI-Chat aus Deutschland ohne Data Mining, ohne Tracking und mit Ende-zu-Ende-Verschlüsselung. Nutze eine sichere ChatGPT-Alternative mit Open-Weight-Modellen."
keywords: "privater ki chat, deutscher ki chat, deutscher ai chat, privat chatgpt, privat gemini, privat gemien, chatgpt alternative deutschland, verschlüsselter ai chat, datenschutz ki chat, chat gbt alternative"
aliases:
  - "/use-cases/"

hero_eyebrow: "Privater KI-Chat aus Deutschland"
hero_title: "Ein privater Chat.<br>Viele Welten."
hero_lead: "Gründer, kleine Betriebe, Entwickler, Studierende. Alle nutzen dieselbe App für ganz unterschiedliche Arbeit. Scroll nach unten. Jede Welt ist ein echter Anwendungsfall."
scroll_hint: "Scrollen zum Start"
input_placeholder: "Frag mich alles!"
disclaimer: "Du chattest mit einer KI/LLM — Fehler möglich, Wichtiges prüfen."
mode_label: "Fast"
worlds_label: "Zu einer Welt springen"
of_label: "von"
ui:
  download: "Herunterladen"
  open: "Öffnen"
  preview: "Vorschau"
  code: "Code"
  version: "Version"

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
      chat_title: "Gespräch mit dem Chef"
      ask: "Wie sage ich meinem Chef, dass ich ausgebrannt bin?"
      meta: "Thought for 4s"
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
      chat_title: "Spoke Landingpage"
      ask: "Bau eine Landingpage für meinen Fahrrad-Reparaturservice und veröffentliche sie."
      running: "Running create artifact"
      steps: ["artifact manager", "create artifact"]
      meta: "Worked for 12s"
      art_title: "Spoke Landingpage"
      art_sub: "HTML · v1"
      art_type: "HTML"
      done: "Fertig. Deine Seite ist online, jeder mit dem Link kann sie öffnen:"
      url: "https://artifacts.chuk.chat/k7f2-spoke"
      brand: "Spoke"
      nav: ["Preise", "Gebiete", "Buchen"]
      headline: "Fahrrad-Reparatur vor deiner Tür."
      sub: "Wir kommen zu dir. Festpreise, Termin am selben Tag."
      cta: "Reparatur buchen"
      features: ["Termin am selben Tag", "Festpreise", "Alle Marken"]

  - id: "business"
    short: "Kleine Betriebe"
    scene: "day"
    tone: "light"
    label: "Einen kleinen Betrieb führen"
    title: "Der Bürokram, erledigt zwischen zwei Aufträgen."
    text: "Ein Tischler schreibt die Rechnung als sauberes PDF auf dem Weg zum nächsten Kunden. Eine Café-Inhaberin beantwortet Mails und plant die Schichten. Termine landen direkt im Kalender. Ganz ohne IT-Abteilung."
    chips: ["PDF-Dokumente", "E-Mail-Entwürfe", "Kalender und Erinnerungen"]
    mock:
      chat_title: "Rechnung Frau Weber"
      ask: "Rechnung für Frau Weber: Küchenregal, 6 Stunden à 58 €, Material 140 €. Als PDF."
      running: "Compiling document"
      steps: ["typst compile"]
      meta: "Worked for 6s"
      art_title: "Rechnung Frau Weber"
      art_sub: "Typst · PDF · v1"
      art_type: "Typst"
      doc_from: "Tischlerei Brandt · Hafenstraße 12 · 24103 Kiel"
      doc_title: "Rechnung"
      doc_no: "Nr. 2026-031"
      doc_to: "Frau Weber"
      rows:
        - ["Küchenregal, Arbeitszeit 6 h × 58 €", "348,00 €"]
        - ["Material", "140,00 €"]
      total_label: "Gesamt"
      total: "488,00 €"
      ask2: "Schreib Frau Weber noch eine kurze Mail dazu."
      mail_subject: "Ihre Rechnung 2026-031"
      mail_to: "weber@example.de"
      mail_body: "Liebe Frau Weber, vielen Dank für Ihren Auftrag. Hier ist die Rechnung für das Küchenregal. Gesamt: 488,00 €, zahlbar innerhalb von 14 Tagen. Viele Grüße, Jan Brandt"
      mail_button: "In Mail-App öffnen"

  - id: "connectors"
    short: "Connectors"
    scene: "network"
    tone: "light"
    label: "Alle Tools verbinden"
    title: "Alle deine Tools. Ein Chat, der sie steuert."
    text: "Wissensarbeiter verbinden Notion, Linear, Todoist, Dropbox und 50 weitere Dienste. Du meldest dich einmal im Browser an. Dann fragst du in normalen Worten, und Chuk Chat klickt sich für dich durch alle."
    chips: ["50+ Connectors (MCP)", "Anmeldung per OAuth", "Desktop und Android"]
    mock:
      chat_title: "Diese Woche"
      ask: "Was ist diese Woche fällig? Schau in Linear, Todoist und Notion und mach mir einen Plan."
      running: "Running todoist find-tasks"
      steps: ["linear list issues", "todoist find-tasks"]
      search: "Launch-Plan"
      meta: "Worked for 9s"
      lead: "Deine Woche, der Reihe nach:"
      answer:
        - "Mo: Login-Bug fixen (Linear, hohe Priorität)"
        - "Di: Q3-Bericht verschicken (Todoist)"
        - "Do: Launch-Plan prüfen (Notion)"
      orbit: ["github", "stripe", "dropbox", "figma", "calcom", "airtable", "zapier", "asana", "sentry", "canva", "supabase", "fastmail", "box", "vercel"]

  - id: "models"
    short: "Modelle"
    scene: "prism"
    tone: "light"
    label: "Jedes Modell, ein Budget"
    title: "Schluss mit fünf KI-Abos."
    text: "Manche Modelle schreiben besser. Manche denken tiefer. Manche sind schnell und günstig. Power-User wählen für jede Nachricht ein anderes Modell und zahlen aus einem Budget: 20 € im Monat, 16 € davon als KI-Guthaben. Kein zweites Konto, keine zweite Rechnung."
    chips: ["Frontier-Open-Weight-Modelle", "Wechsel pro Nachricht", "16 € KI-Guthaben inklusive"]
    prompt: "Schreib mein Anschreiben mit Kimi K3."
    mock:
      chat_title: "Anschreiben"
      modes: ["Fast", "Thinking"]
      reasoning: "Reasoning"
      reasoning_value: "Low"
      more: "More models"
      choose: "Choose model"
      models:
        - {logo: "deepseek.svg", name: "DeepSeek V4 Pro 0813"}
        - {logo: "moonshot.svg", name: "Kimi K3"}
        - {logo: "zai.svg", name: "GLM 5.3"}
        - {logo: "qwen.svg", name: "Qwen3.8 27B"}
        - {logo: "minimax.svg", name: "MiniMax M3"}
        - {logo: "mistral.svg", name: "Mistral Small 4"}
        - {logo: "openai.svg", name: "gpt-oss-120b"}
      sent:
        - {text: "Schreib mein Anschreiben für die Stelle als Büroleitung.", model: 1, meta: "Thought for 3s", answer: "Sehr geehrte Frau Berger, hiermit bewerbe ich mich als Büroleitung…"}
        - {text: "Prüf die Gehaltsrechnung: 4.250 € im Monat, 13 Gehälter.", model: 0, meta: "Thought for 6s", answer: "Stimmt: 4.250 € × 13 = 55.250 € im Jahr."}
        - {text: "Übersetz das Anschreiben ins Englische.", model: 3, meta: "Thought for 2s", answer: "Dear Ms Berger, I am applying for the office manager position…"}

  - id: "engineers"
    short: "Entwickler"
    scene: "terminal"
    tone: "dark"
    label: "Viele Aufgaben gleichzeitig"
    title: "Fünf Chats laufen. Keiner wartet auf dich."
    text: "Entwickler starten in einem Chat ein Refactoring, im nächsten eine Recherche und im dritten eine Fehlersuche. Antworten streamen weiter, wenn du den Chat wechselst. Code landet in bearbeitbaren Panels, und auf dem Desktop führt eine Sandbox-Shell Befehle für dich aus."
    chips: ["Paralleles Streaming", "Code-Artefakte", "Sandbox-Shell (Desktop)"]
    prompt: "Refactor den Auth-Service und lass den Formatter laufen."
    mock:
      nav_new: "Neuer Chat"
      nav_media: "Medien"
      nav_search: "Suche"
      user: "Lena Hartmann"
      credit: "24,60 €"
      groups:
        - label: "Angeheftet"
          chats:
            - {title: "Lasttest für die neue Sync-API", time: "Mo., 21. Sept.", live: true}
            - {title: "Code-Styleguide fürs Team", time: "Mi., 12. Aug."}
        - label: "Heute"
          chats:
            - {title: "Auth-Service refactoren", time: "14:32", live: true, active: true, done: 5}
            - {title: "Warum ist der CI-Build langsam?", time: "14:28", live: true, done: 4}
            - {title: "Migration auf Dart 3.13", time: "14:21", live: true}
            - {title: "Stacktrace erklären", time: "14:09", live: true}
            - {title: "Regex für deutsche IBANs", time: "13:40"}
        - label: "Diese Woche"
          count: 23
          chats:
            - {title: "Wackeliger Login-Test in der CI", time: "Sa., 26. Sept."}
            - {title: "Postgres-Index für die Volltextsuche", time: "Sa., 26. Sept."}
            - {title: "Multi-Stage-Dockerfile für die API", time: "Fr., 25. Sept."}
            - {title: "Rate-Limiter mit Redis", time: "Do., 24. Sept."}
            - {title: "Kotlin-Coroutines vs. Dart-Isolates", time: "Mi., 23. Sept."}
            - {title: "Release Notes für 2.3", time: "Di., 22. Sept."}
        - label: "Diesen Monat"
          count: 41
          chats:
            - {title: "Sentry für Flutter einrichten", time: "Sa., 19. Sept."}
            - {title: "GraphQL oder REST fürs Admin-Panel", time: "Do., 17. Sept."}
            - {title: "Speicherleck im Bild-Cache", time: "Mo., 14. Sept."}
            - {title: "Nginx-Konfiguration für WebSockets", time: "Fr., 11. Sept."}
      ask: "Refactor den Auth-Service und lass den Formatter laufen."
      running: "Running bash"
      search: "dart token refresh pattern"
      bash: "dart format lib/"
      meta: "Worked for 21s"
      answer: "Ich habe den Token-Refresh in eine eigene Methode gezogen und einen Retry ergänzt. Der Formatter hat 3 Dateien geändert."
      art_title: "auth_service.dart"
      art_sub: "Code · v2"

  - id: "prototypers"
    short: "Prototyper"
    scene: "blueprint"
    tone: "dark"
    label: "Der Prototyper"
    title: "Von der Skizze zum klickbaren Prototyp."
    text: "Prototyper skizzieren einen Ablauf, machen daraus eine klickbare HTML-Seite und generieren die Bilder dazu. Jede Idee kostet Minuten statt Tage. Die guten bleiben. Der Rest fliegt raus."
    chips: ["Excalidraw-Skizzen", "HTML-Vorschau", "Bildgenerierung"]
    mock:
      chat_title: "Signup-Ablauf"
      ask: "Skizzier einen Signup-Ablauf mit E-Mail-Code."
      steps: ["artifact manager"]
      meta: "Worked for 8s"
      art1_title: "Signup-Ablauf"
      art1_sub: "Excalidraw sketch · v1"
      art1_type: "Excalidraw"
      nodes: ["E-Mail eingeben", "Code gesendet", "Code eingeben", "Willkommen"]
      ask2: "Mach daraus eine klickbare Seite."
      meta2: "Worked for 11s"
      art2_title: "Signup-Seite"
      art2_sub: "HTML · v1"
      art2_type: "HTML"
      proto_title: "Konto erstellen"
      proto_input: "du@beispiel.de"
      proto_btn: "Code senden"
      ask3: "Füg ein Hero-Bild hinzu."
      meta3: "Worked for 5s"
      image_model: "Z-Image Turbo"

  - id: "research"
    short: "Recherche"
    scene: "library"
    tone: "light"
    label: "Recherche und Dokumente"
    title: "Sprich mit deinen Dokumenten. Antworten mit Quellen."
    text: "Studierende und Forschende hängen PDFs, Notizen und Paper an und stellen Fragen dazu. Chuk Chat sucht im Web, nennt seine Quellen, zeichnet Diagramme und setzt das Ergebnis als sauberes PDF. Alte Chats bleiben durchsuchbar."
    chips: ["PDFs und Dateianhänge", "Websuche mit Quellen", "Diagramme und PDF-Export"]
    mock:
      chat_title: "Befragungsmethoden"
      files: ["masterarbeit_entwurf.pdf", "mueller_2024.pdf"]
      ask: "Vergleich die Methoden beider Paper. Zeig die Rücklaufquoten als Diagramm."
      running: "Searching the web"
      search: "Rücklaufquote nach Befragungsmethode"
      sources:
        - {host: "nature.com", letter: "n", color: "#1F1F1F"}
        - {host: "arxiv.org", letter: "a", color: "#B31B1B"}
        - {host: "destatis.de", letter: "D", color: "#0B5CA8"}
        - {host: "pewresearch.org", letter: "P", color: "#2E6E8E"}
      source_count: "8 sources"
      meta: "Worked for 14s"
      lead: "Beide Paper messen die Rücklaufquote, aber Müller vergleicht vier Methoden. Persönlich wirkt am besten:"
      chart_title: "Rücklaufquote nach Befragungsmethode (%)"
      bars:
        - {label: "Online", value: 38}
        - {label: "Telefon", value: 22}
        - {label: "Post", value: 14}
        - {label: "Persönlich", value: 61}

  - id: "onthego"
    short: "Unterwegs"
    scene: "street"
    tone: "dark"
    label: "Unterwegs"
    title: "Home-Taste halten. Fragen. Erledigt."
    text: "Auf Android kann Chuk Chat dein Assistent sein. Halte die Home-Geste gedrückt: Er hört zu, liest mit, was auf dem Bildschirm ist, antwortet laut und handelt. Er stellt Wecker, findet Orte und startet die Navigation."
    chips: ["Android-Assistent", "Sprache rein, Sprache raus", "Orte, Routen, Wecker"]
    mock:
      clock: "18:42"
      brand: "CHUK CHAT"
      listening: "Ich höre zu"
      thinking: "Denke nach …"
      acting: "Führe aus …"
      heard: "Such eine Apotheke, die jetzt offen hat, und bring mich hin."
      tool: "Standort"
      tool2: "Orte: Apotheke"
      tool3: "Navigation: Apotheke am Markt"
      places_title: "Apotheke"
      places:
        - {name: "Apotheke am Markt", address: "Marktplatz 4", meta: "Apotheke · geöffnet bis 20:00", rating: "4.7", reviews: "128"}
        - {name: "Linden-Apotheke", address: "Lindenstraße 21", meta: "Apotheke · geöffnet bis 19:00", rating: "4.5", reviews: "86"}
        - {name: "Hafen-Apotheke", address: "Kaistraße 8", meta: "Apotheke · geöffnet bis 18:30", rating: "4.4", reviews: "51"}
      answer: "Drei Apotheken haben offen. Am nächsten ist die **Apotheke am Markt**, 400 m entfernt. Die Navigation läuft."
      action: "Navigation gestartet"
      action_detail: "Apotheke am Markt"
      pause: "Pausieren"
      date: "Fr., 25. Sep."
      temp: "14 °C"
      apps: ["Kalender", "Uhr", "Fotos", "Wetter", "Notizen", "Dateien", "Einstellungen", "Chuk Chat"]

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

hero_cta_download: "Download für Mac / Windows / Linux"
hero_cta_github: "Auf GitHub ansehen"
open_source_line: "100% Open Source · monatliches Abo · kein Token-Stress"

comparison_title: "Chuk Chat im Vergleich"
comparison:
  - feature: "Privatsphäre zuerst"
    chatgpt: false
    claude: false
    okara: true
    lumo: true
    chuk: true
  - feature: "Sitz in Deutschland"
    chatgpt: false
    claude: false
    okara: false
    lumo: false
    chuk: true
  - feature: "Sitz in der EU"
    chatgpt: false
    claude: false
    okara: false
    lumo: true
    chuk: true
  - feature: "Open Source"
    chatgpt: false
    claude: false
    okara: false
    lumo: true
    chuk: true
  - feature: "Frontier-Class-Modelle"
    chatgpt: true
    claude: true
    okara: false
    lumo: false
    chuk: true
  - feature: "Nur Open-Weight-Modelle"
    chatgpt: false
    claude: false
    okara: true
    lumo: true
    chuk: true
  - feature: "Ende-zu-Ende-verschlüsselt"
    chatgpt: false
    claude: false
    okara: true
    lumo: true
    chuk: true
  - feature: "Trainiert nie mit deinen Daten"
    chatgpt: false
    claude: false
    okara: true
    lumo: true
    chuk: true
  - feature: "Kein User-Tracking"
    chatgpt: false
    claude: false
    okara: false
    lumo: true
    chuk: true
  - feature: "DSGVO by Design"
    chatgpt: false
    claude: false
    okara: true
    lumo: true
    chuk: true
  - feature: "Plattformen"
    chatgpt_platforms: ["iOS", "Android", "Mac", "Windows", "Web"]
    claude_platforms: ["iOS", "Android", "Mac", "Windows", "Web"]
    okara_platforms: ["Web"]
    lumo_platforms: ["iOS", "Android", "Web"]
    chuk_platforms: ["Android", "Mac", "Windows", "Linux", "Web"]
---
