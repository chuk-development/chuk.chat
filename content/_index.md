---
bodyClass: "home uc-page"
title: "Chuk Chat — Private AI Chat from Germany | Secure & Anonymous"
description: "Chuk Chat is a private AI chat app from Germany with no data mining, no tracking, and end-to-end encryption. Use a secure ChatGPT alternative powered only by open-weight models."
keywords: "private ai chat, german ai chat, private chatgpt, private gemini, private claude alternative, private chatgpt alternative, encrypted ai chat, privacy ai assistant, no tracking ai chat, chat gbt alternative"
aliases:
  - "/en/use-cases/"

hero_eyebrow: "Private AI chat from Germany"
hero_title: "One private chat.<br>Many worlds."
hero_lead: "Founders, small businesses, developers, students. They use the same app for very different work. Scroll down. Each world is a real use case."
scroll_hint: "Scroll to begin"
input_placeholder: "Ask me anything !"
disclaimer: "You're chatting with an AI/LLM — it can be wrong. Check key info."
mode_label: "Fast"
worlds_label: "Jump to a world"
of_label: "of"
ui:
  download: "Download"
  open: "Open"
  preview: "Preview"
  code: "Code"
  version: "Version"

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
      chat_title: "Burnout talk with my boss"
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
      chat_title: "Spoke landing page"
      ask: "Build a landing page for my bike repair service and publish it."
      running: "Running create artifact"
      steps: ["artifact manager", "create artifact"]
      meta: "Worked for 12s"
      art_title: "Spoke landing page"
      art_sub: "HTML · v1"
      art_type: "HTML"
      done: "Done. Your page is live, and anyone with the link can open it:"
      url: "https://artifacts.chuk.chat/k7f2-spoke"
      brand: "Spoke"
      nav: ["Prices", "Areas", "Book"]
      headline: "Bike repair at your door."
      sub: "We come to you. Fixed prices, same-day visits."
      cta: "Book a repair"
      features: ["Same-day visits", "Fixed prices", "All brands"]

  - id: "business"
    short: "Small business"
    scene: "day"
    tone: "light"
    label: "Running a small business"
    title: "The office work, done between two jobs."
    text: "A carpenter writes the invoice as a clean PDF on the way to the next customer. A café owner answers emails and plans the shifts. Appointments go straight into the calendar. No IT department needed."
    chips: ["PDF documents", "Email drafts", "Calendar and reminders"]
    mock:
      chat_title: "Invoice Mrs Weber"
      ask: "Write the invoice for Mrs Weber: kitchen shelf, 6 hours at €58, material €140. As PDF."
      running: "Compiling document"
      steps: ["typst compile"]
      meta: "Worked for 6s"
      art_title: "Invoice Mrs Weber"
      art_sub: "Typst · PDF · v1"
      art_type: "Typst"
      doc_from: "Brandt Carpentry · Hafenstraße 12 · 24103 Kiel"
      doc_title: "Invoice"
      doc_no: "No. 2026-031"
      doc_to: "Mrs Weber"
      rows:
        - ["Kitchen shelf, labour 6 h × €58", "€348.00"]
        - ["Material", "€140.00"]
      total_label: "Total"
      total: "€488.00"
      ask2: "Now write Mrs Weber a short email about it."
      mail_subject: "Your invoice 2026-031"
      mail_to: "weber@example.de"
      mail_body: "Dear Mrs Weber, thank you for your order. Here is the invoice for the kitchen shelf. Total: €488.00, payable within 14 days. Kind regards, Jan Brandt"
      mail_button: "Open in Mail App"

  - id: "connectors"
    short: "Connectors"
    scene: "network"
    tone: "light"
    label: "Connecting every tool"
    title: "All your tools. One chat to drive them."
    text: "Knowledge workers connect Notion, Linear, Todoist, Dropbox and 50 more services. You sign in once in the browser. Then you ask in plain words, and Chuk Chat does the clicking across all of them."
    chips: ["50+ connectors (MCP)", "Sign in with OAuth", "Desktop and Android"]
    mock:
      chat_title: "This week"
      ask: "What is due this week? Check Linear, Todoist and Notion, then make me a plan."
      running: "Running todoist find-tasks"
      steps: ["linear list issues", "todoist find-tasks"]
      search: "launch plan"
      meta: "Worked for 9s"
      lead: "Your week, in order:"
      answer:
        - "Mon: fix the login bug (Linear, high priority)"
        - "Tue: send the Q3 report (Todoist)"
        - "Thu: review the launch plan (Notion)"
      orbit: ["github", "stripe", "dropbox", "figma", "calcom", "airtable", "zapier", "asana", "sentry", "canva", "supabase", "fastmail", "box", "vercel"]

  - id: "models"
    short: "Models"
    scene: "prism"
    tone: "light"
    label: "Every model, one budget"
    title: "Stop paying for five AI subscriptions."
    text: "Some models write better. Some think deeper. Some are fast and cheap. Power users pick a different model for each message and pay from one budget: €20 per month, with €16 in AI credits included. No second account, no second invoice."
    chips: ["Frontier open-weight models", "Switch for each message", "€16 AI credits included"]
    prompt: "Write my cover letter with Kimi K3."
    mock:
      chat_title: "Cover letter"
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
        - {text: "Write my cover letter for the office manager job.", model: 1, meta: "Thought for 3s", answer: "Dear Ms Berger, I am applying for the office manager position…"}
        - {text: "Check the salary maths: €4,250 per month, 13 salaries.", model: 0, meta: "Thought for 6s", answer: "Correct: €4,250 × 13 = €55,250 per year."}
        - {text: "Translate the letter into German.", model: 3, meta: "Thought for 2s", answer: "Sehr geehrte Frau Berger, hiermit bewerbe ich mich…"}

  - id: "engineers"
    short: "Engineers"
    scene: "terminal"
    tone: "dark"
    label: "Many tasks at once"
    title: "Five chats running. None of them waits for you."
    text: "Engineers start a refactor in one chat, a research question in the next and a bug hunt in a third. Answers keep streaming when you switch chats. Code lands in editable panels, and on desktop a sandboxed shell runs commands for you."
    chips: ["Parallel streaming", "Code artifacts", "Sandboxed shell (desktop)"]
    prompt: "Refactor the auth service and run the formatter."
    mock:
      nav_new: "New chat"
      nav_media: "Media"
      nav_search: "Search"
      user: "Lena Hartmann"
      credit: "€24.60"
      groups:
        - label: "Pinned"
          chats:
            - {title: "Load test the new sync API", time: "Mon, Sep 21", live: true}
            - {title: "Team code style guide", time: "Wed, Aug 12"}
        - label: "Today"
          chats:
            - {title: "Refactor the auth service", time: "14:32", live: true, active: true, done: 5}
            - {title: "Why is the CI build slow?", time: "14:28", live: true, done: 4}
            - {title: "Migrate to Dart 3.13", time: "14:21", live: true}
            - {title: "Explain this stack trace", time: "14:09", live: true}
            - {title: "Regex for German IBANs", time: "13:40"}
        - label: "This week"
          count: 23
          chats:
            - {title: "Flaky login test on CI", time: "Sat, Sep 26"}
            - {title: "Postgres index for full-text search", time: "Sat, Sep 26"}
            - {title: "Multi-stage Dockerfile for the API", time: "Fri, Sep 25"}
            - {title: "Rate limiter with Redis", time: "Thu, Sep 24"}
            - {title: "Kotlin coroutines vs. Dart isolates", time: "Wed, Sep 23"}
            - {title: "Release notes for 2.3", time: "Tue, Sep 22"}
        - label: "This month"
          count: 41
          chats:
            - {title: "Set up Sentry for Flutter", time: "Sat, Sep 19"}
            - {title: "GraphQL or REST for the admin panel", time: "Thu, Sep 17"}
            - {title: "Memory leak in the image cache", time: "Mon, Sep 14"}
            - {title: "Nginx config for WebSockets", time: "Fri, Sep 11"}
      ask: "Refactor the auth service and run the formatter."
      running: "Running bash"
      search: "dart token refresh pattern"
      bash: "dart format lib/"
      meta: "Worked for 21s"
      answer: "I moved the token refresh into its own method and added a retry. The formatter changed 3 files."
      art_title: "auth_service.dart"
      art_sub: "Code · v2"

  - id: "prototypers"
    short: "Prototypers"
    scene: "blueprint"
    tone: "dark"
    label: "The prototyper"
    title: "From a sketch to a clickable prototype."
    text: "Prototypers sketch a flow, turn it into a clickable HTML page and generate the images for it. Each idea takes minutes, not days. Keep the good ones. Delete the rest."
    chips: ["Excalidraw sketches", "HTML previews", "Image generation"]
    mock:
      chat_title: "Signup flow"
      ask: "Sketch a signup flow with an email code."
      steps: ["artifact manager"]
      meta: "Worked for 8s"
      art1_title: "Signup flow"
      art1_sub: "Excalidraw sketch · v1"
      art1_type: "Excalidraw"
      nodes: ["Enter email", "Code sent", "Enter code", "Welcome"]
      ask2: "Make it a clickable page."
      meta2: "Worked for 11s"
      art2_title: "Signup page"
      art2_sub: "HTML · v1"
      art2_type: "HTML"
      proto_title: "Create your account"
      proto_input: "you@example.com"
      proto_btn: "Send code"
      ask3: "Add a hero image."
      meta3: "Worked for 5s"
      image_model: "Z-Image Turbo"

  - id: "research"
    short: "Research"
    scene: "library"
    tone: "light"
    label: "Research and documents"
    title: "Talk to your documents. Get answers with sources."
    text: "Students and researchers attach PDFs, notes and papers and ask questions about them. Chuk Chat searches the web, names its sources, draws charts and writes the result as a typeset PDF. Old chats stay searchable."
    chips: ["PDF and file attachments", "Web search with sources", "Charts and PDF export"]
    mock:
      chat_title: "Survey methods"
      files: ["thesis_draft.pdf", "miller_2024.pdf"]
      ask: "Compare the methods in both papers. Show the response rates as a chart."
      running: "Searching the web"
      search: "survey response rates by method"
      sources:
        - {host: "nature.com", letter: "n", color: "#1F1F1F"}
        - {host: "arxiv.org", letter: "a", color: "#B31B1B"}
        - {host: "destatis.de", letter: "D", color: "#0B5CA8"}
        - {host: "pewresearch.org", letter: "P", color: "#2E6E8E"}
      source_count: "8 sources"
      meta: "Worked for 14s"
      lead: "Both papers measure response rates, but Miller compares four survey methods. In person works best:"
      chart_title: "Response rate by survey method (%)"
      bars:
        - {label: "Online", value: 38}
        - {label: "Phone", value: 22}
        - {label: "Mail", value: 14}
        - {label: "In person", value: 61}

  - id: "onthego"
    short: "On the go"
    scene: "street"
    tone: "dark"
    label: "On the go"
    title: "Hold the home button. Ask. Done."
    text: "On Android, Chuk Chat can be your assistant. Hold the home gesture, and it listens, reads what is on the screen, answers out loud and acts. It sets alarms, finds places and starts the navigation."
    chips: ["Android assistant", "Voice in, voice out", "Places, routes, alarms"]
    mock:
      clock: "18:42"
      brand: "CHUK CHAT"
      listening: "Listening"
      thinking: "Thinking …"
      acting: "Working …"
      heard: "Find a pharmacy that is open now and take me there."
      tool: "Location"
      tool2: "Places: Pharmacy"
      tool3: "Navigation: Pharmacy at the market"
      places_title: "Pharmacy"
      places:
        - {name: "Pharmacy at the market", address: "Marktplatz 4", meta: "Pharmacy · open until 20:00", rating: "4.7", reviews: "128"}
        - {name: "Linden Pharmacy", address: "Lindenstraße 21", meta: "Pharmacy · open until 19:00", rating: "4.5", reviews: "86"}
        - {name: "Harbour Pharmacy", address: "Kaistraße 8", meta: "Pharmacy · open until 18:30", rating: "4.4", reviews: "51"}
      answer: "Three pharmacies are open. The closest is **Pharmacy at the market**, 400 m away. Navigation is running."
      action: "Navigation started"
      action_detail: "Pharmacy at the market"
      pause: "Pause"
      date: "Fri, Sep 25"
      temp: "14°C"
      apps: ["Calendar", "Clock", "Photos", "Weather", "Notes", "Files", "Settings", "Chuk Chat"]

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

hero_cta_download: "Download for Mac / Windows / Linux"
hero_cta_github: "Check on GitHub"
open_source_line: "100% open source · monthly subscription · no token stress"

comparison_title: "How Chuk Chat Compares"
comparison:
  - feature: "Privacy First"
    chatgpt: false
    claude: false
    okara: true
    lumo: true
    chuk: true
  - feature: "Based in Germany"
    chatgpt: false
    claude: false
    okara: false
    lumo: false
    chuk: true
  - feature: "Based in EU"
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
  - feature: "Frontier-Class Models"
    chatgpt: true
    claude: true
    okara: false
    lumo: false
    chuk: true
  - feature: "Open-Weight Models Only"
    chatgpt: false
    claude: false
    okara: true
    lumo: true
    chuk: true
  - feature: "End-to-End Encrypted"
    chatgpt: false
    claude: false
    okara: true
    lumo: true
    chuk: true
  - feature: "Never Trains on Your Data"
    chatgpt: false
    claude: false
    okara: true
    lumo: true
    chuk: true
  - feature: "No User Tracking"
    chatgpt: false
    claude: false
    okara: false
    lumo: true
    chuk: true
  - feature: "GDPR by Design"
    chatgpt: false
    claude: false
    okara: true
    lumo: true
    chuk: true
  - feature: "Platforms"
    chatgpt_platforms: ["iOS", "Android", "Mac", "Windows", "Web"]
    claude_platforms: ["iOS", "Android", "Mac", "Windows", "Web"]
    okara_platforms: ["Web"]
    lumo_platforms: ["iOS", "Android", "Web"]
    chuk_platforms: ["Android", "Mac", "Windows", "Linux", "Web"]
---
