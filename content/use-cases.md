---
title: "Use Cases: Private AI for Every Kind of Work"
layout: "usecases"
type: "page"
translationKey: "use-cases"
bodyClass: "uc-page"
description: "Founders, small businesses, developers, students: nine real ways people use Chuk Chat, the private AI chat from Germany with end-to-end encryption and open-weight models."
keywords: "private ai chat use cases, ai for small business, ai for founders, encrypted ai assistant, open-weight ai models, mcp connectors, ai pdf invoice, android ai assistant"

hero_eyebrow: "Use cases"
hero_title: "One private chat.<br>Many worlds."
hero_lead: "Founders, small businesses, developers, students. They use the same app for very different work. Scroll down. Each world is a real use case."
scroll_hint: "Scroll to begin"
input_placeholder: "Ask me anything !"
mode_label: "Fast"
worlds_label: "Jump to a world"
of_label: "of"

worlds:
  - id: "private"
    short: "Private"
    scene: "night"
    tone: "dark"
    label: "The private questions"
    title: "The questions you would not ask anyone else."
    text: "Health, money, a hard talk with your boss. People ask Chuk Chat the things they would never type into a data-hungry AI. Your chats are encrypted on your device before we store them. We keep only ciphertext."
    chips: ["End-to-end encrypted storage", "Never used for training", "No tracking"]
    mock:
      ask: "How do I tell my boss that I am burned out?"
      meta: "Thought for 4s"
      lead: "Start with facts, not with blame."
      answer:
        - "Ask for 20 minutes at a calm moment."
        - "Say what changed: sleep, focus, sick days."
        - "Bring one clear request, like fewer projects for a month."
      stored: "What our server stores"
      cipher: "AES-256-GCM · the key stays on your device"

  - id: "founders"
    short: "Founders"
    scene: "dawn"
    tone: "light"
    label: "Building a product to sell"
    title: "From idea to launch page in one chat."
    text: "Founders research the market, write the pitch, draft the code and publish a landing page with a public link. Connect GitHub, Stripe or Vercel and work with your real project. Your idea stays your idea. We do not train on it."
    chips: ["Web research", "Artifacts with public links", "GitHub · Stripe · Vercel"]
    mock:
      ask: "Build a landing page for my bike repair service and publish it."
      meta: "Worked for 12s"
      tool: "create_artifact"
      done: "Done. Your page is live, and anyone with the link can open it."
      file: "spoke-landing.html"
      tab_preview: "Preview"
      tab_code: "Code"
      brand: "Spoke"
      nav: ["Prices", "Areas", "Book"]
      headline: "Bike repair at your door."
      sub: "We come to you. Fixed prices, same-day visits."
      cta: "Book a repair"
      features: ["Same-day visits", "Fixed prices", "All brands"]
      public: "Public link"
      url: "artifacts.chuk.chat/k7f2-spoke"
      copy: "Copy"
      connected: "Connected"
      connectors: ["github", "stripe", "vercel"]

  - id: "business"
    short: "Small business"
    scene: "day"
    tone: "light"
    label: "Running a small business"
    title: "The office work, done between two jobs."
    text: "A carpenter writes the invoice as a clean PDF on the way to the next customer. A café owner answers emails and plans the shifts. Appointments go straight into the calendar. No IT department needed."
    chips: ["PDF documents", "Email drafts", "Calendar and reminders"]
    mock:
      ask: "Write the invoice for Mrs Weber: kitchen shelf, 6 hours at €58, material €140. As PDF."
      meta: "Worked for 6s"
      tool: "typst_compile"
      file: "Invoice_2026-031.pdf"
      download: "Download"
      doc_title: "Invoice"
      doc_no: "No. 2026-031"
      doc_to: "Mrs Weber"
      rows:
        - ["Kitchen shelf, labour 6 h × €58", "€348.00"]
        - ["Material", "€140.00"]
      total_label: "Total"
      total: "€488.00"
      ask2: "Remind me on Tuesday at 9 to call her."
      event: "Call Mrs Weber"
      event_time: "Tuesday · 09:00"
      event_btn: "Save to calendar"

  - id: "connectors"
    short: "Connectors"
    scene: "network"
    tone: "light"
    label: "Connecting every tool"
    title: "All your tools. One chat to drive them."
    text: "Knowledge workers connect Notion, Linear, Todoist, Dropbox and 50 more services. You sign in once in the browser. Then you ask in plain words, and Chuk Chat does the clicking across all of them."
    chips: ["50+ connectors (MCP)", "Sign in with OAuth", "Desktop and Android"]
    mock:
      ask: "What is due this week? Check Linear, Todoist and Notion, then make me a plan."
      meta: "Worked for 9s"
      calls:
        - {logo: "linear", name: "Linear", result: "7 open issues"}
        - {logo: "todoist", name: "Todoist", result: "12 tasks"}
        - {logo: "notion", name: "Notion", result: "3 pages"}
      lead: "Your week, in order:"
      answer:
        - "Mon: fix the login bug (Linear, high priority)"
        - "Tue: send the Q3 report (Todoist)"
        - "Thu: review the launch plan (Notion)"
      orbit: ["github", "stripe", "dropbox", "figma", "calcom", "airtable", "zapier", "asana", "sentry", "canva", "supabase", "fastmail", "box", "vercel"]

  - id: "models"
    prompt: "Write my cover letter with Kimi K3."
    short: "Models"
    scene: "prism"
    tone: "light"
    label: "Every model, one budget"
    title: "Stop paying for five AI subscriptions."
    text: "Some models write better. Some think deeper. Some are fast and cheap. Power users pick a different model for each message and pay from one budget: €20 per month, with €16 in AI credits included. No second account, no second invoice."
    chips: ["Frontier open-weight models", "Switch for each message", "€16 AI credits included"]
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
        - {text: "Write my cover letter", model: 1}
        - {text: "Check the salary maths", model: 0}
        - {text: "Translate it into German", model: 3}
      credits_label: "AI credits this month"
      credits: ["€16.00", "€15.97", "€15.92", "€15.90"]

  - id: "engineers"
    prompt: "Refactor the auth service and run the formatter."
    short: "Engineers"
    scene: "terminal"
    tone: "dark"
    label: "Many tasks at once"
    title: "Five chats running. None of them waits for you."
    text: "Engineers start a refactor in one chat, a research question in the next and a bug hunt in a third. Answers keep streaming when you switch chats. Code lands in editable panels, and on desktop a sandboxed shell runs commands for you."
    chips: ["Parallel streaming", "Code artifacts", "Sandboxed shell (desktop)"]
    mock:
      chats:
        - {title: "Refactor the auth service", live: true}
        - {title: "Why is the CI build slow?", live: true, finishes: true}
        - {title: "Regex for German IBANs", live: false}
        - {title: "Migrate to Dart 3.13", live: true}
        - {title: "Explain this stack trace", live: false}
      file: "auth_service.dart"
      shell: "$ dart format lib/"
      shell_out: "Formatted 3 files (1 changed)"
      toast: "“Why is the CI build slow?” has an answer"

  - id: "prototypers"
    short: "Prototypers"
    scene: "blueprint"
    tone: "dark"
    label: "The prototyper"
    title: "Ten ideas tested before lunch."
    text: "Prototypers sketch a flow as a diagram, turn it into a clickable HTML page and generate the images for it. Each idea takes minutes, not days. Keep the good ones. Delete the rest."
    chips: ["Diagrams (Mermaid, Excalidraw)", "HTML prototypes", "Image generation"]
    mock:
      ask: "Sketch a signup flow with an email code."
      nodes: ["Enter email", "Code sent", "Enter code", "Welcome"]
      ask2: "Make it clickable. Add a hero image."
      diagram: "signup-flow.mmd"
      proto: "signup.html"
      proto_title: "Create your account"
      proto_input: "you@example.com"
      proto_btn: "Send code"
      image: "hero.png · generated"

  - id: "research"
    short: "Research"
    scene: "library"
    tone: "light"
    label: "Research and documents"
    title: "Talk to your documents. Get answers with sources."
    text: "Students and researchers attach PDFs, notes and papers and ask questions about them. Chuk Chat searches the web, names its sources, draws charts and writes the result as a typeset PDF. Old chats stay searchable."
    chips: ["PDF and file attachments", "Web search with sources", "Charts and PDF export"]
    mock:
      files: ["thesis_draft.pdf", "miller_2024.pdf"]
      ask: "Compare the methods in both papers. Show the response rates as a chart."
      meta: "Worked for 14s"
      sources: ["nature.com", "arxiv.org", "destatis.de", "+ 5"]
      chart_title: "Response rate by survey method"
      bars:
        - {label: "Online", value: 38}
        - {label: "Phone", value: 22}
        - {label: "Mail", value: 14}
        - {label: "In person", value: 61}
      export: "methods_comparison.pdf"
      export_note: "Typeset PDF · 4 pages"

  - id: "onthego"
    short: "On the go"
    scene: "street"
    tone: "dark"
    label: "On the go"
    title: "Hold the home button. Ask. Done."
    text: "On Android, Chuk Chat can be your assistant. Hold the home gesture, and it listens, reads what is on the screen, answers out loud and acts. It sets alarms, finds places and starts the navigation."
    chips: ["Android assistant", "Voice in, voice out", "Places, routes, alarms"]
    mock:
      listening: "Listening…"
      heard: "Find a pharmacy that is open now and take me there."
      place: "Pharmacy at the market"
      place_meta: "400 m · open until 20:00"
      action: "Navigation started"
      clock: "18:42"

finale_eyebrow: "What connects them"
finale_title: "Nine worlds. One rule:<br>your chats belong to you."
finale_text: "One thread runs through all of these. Your conversations are the most personal data you create. So they are encrypted on your device, answered by open-weight models and never used for training. No ads, no tracking, run from Germany."
facts:
  - {icon: "layers", title: "Open-weight models only", text: "Transparent models. You always know what runs your AI."}
  - {icon: "lock", title: "End-to-end encrypted", text: "Chats are encrypted on your device before we store them."}
  - {icon: "shield", title: "Made in Germany", text: "Operated from Germany under EU law. GDPR by design."}
  - {icon: "globe", title: "No tracking, no training", text: "No ads, no profiling. Your data does not train anything."}
cta_title: "Find your world."
cta_text: "€20 per month, €16 of it as AI credits. Web, Mac, Windows, Linux and Android."
cta_download: "Download Chuk Chat"
cta_web: "Open the Web App"
---
