# CLAUDE.md — Project Context for Claude Code

## Project Overview

chuk.chat is a bilingual (EN/DE) marketing website for Chuk Chat, a privacy-focused AI chat application that exclusively uses open-weight AI models. Built with Hugo, deployed via Docker + Nginx.

## Tech Stack

- **Static Site Generator:** Hugo (latest, via `hugomods/hugo:latest`)
- **Web Server:** Nginx (Alpine)
- **Deployment:** Docker multi-stage build (Hugo build → Nginx serve)
- **CSS:** Custom vanilla CSS (`assets/css/styles.css`), no frameworks
- **JS:** Inline in templates, no build tools or npm dependencies
- **Design:** "Warm Paper Theme" — off-white (#FDFBF7), goldenrod accent (#B8860B)

## Directory Structure

```
├── assets/css/styles.css      # Main stylesheet (all CSS)
├── content/                   # Markdown content (EN: .md, DE: .de.md)
│   ├── _index.md/de.md       # Homepage
│   ├── blog/                  # Blog posts
│   ├── pricing.md/de.md      # Pricing page
│   ├── downloads.md/de.md    # Platform downloads
│   ├── privacy.md/de.md      # Privacy policy
│   ├── terms.md/de.md        # Terms of service
│   ├── faq.md/de.md          # FAQ
│   ├── impressum.md/de.md    # Legal notice (Impressum)
│   ├── cancellation.md/de.md # Cancellation policy
│   ├── support.md/de.md      # Support contact
│   └── subscription.md/de.md # Subscription status (payment callbacks)
├── layouts/
│   ├── _default/baseof.html   # Base template
│   ├── index.html             # Homepage layout
│   ├── page/                  # Page-specific layouts (pricing, legal, etc.)
│   ├── blog/                  # Blog list + single layouts
│   ├── partials/              # head.html, nav.html, footer.html
│   └── 404.html
├── static/                    # Images, favicon, badges
├── hugo.toml                  # Hugo config
├── nginx.conf                 # Nginx config with redirects
├── Dockerfile                 # Multi-stage Hugo → Nginx build
└── docker-compose.yml
```

## Languages

- **English (en):** Primary, weight=1
- **German (de):** Secondary, weight=2
- URL structure: `/en/...` and `/de/...` (`defaultContentLanguageInSubdir: true`)
- Content pattern: `filename.md` (EN), `filename.de.md` (DE)
- **All content changes must be made in both language files.**

## Key Business Rules

- **Open-weight models only.** No proprietary/closed models (no GPT-4, Claude, Gemini). Current models are sourced from providers like DeepSeek, Qwen, Moonshot AI, MiniMax, Z-AI, Alibaba.
- **Kleinunternehmerregelung (§19 UStG):** No VAT is charged. Do not reference VAT/MwSt in pricing.
- **Pricing:** €20/month, includes €16 AI credits.
- **Privacy:** E2E encryption (AES-256-GCM), data stored in Supabase, payments via Stripe.
- **GDPR compliance:** Legal basis section (Art. 6), supervisory authority is ULD Schleswig-Holstein.
- **German law applies** (excluding CISG), ODR platform referenced in Terms.

## Legal Pages Structure

Legal pages use `layout: "legal"` and are rendered by `layouts/page/legal.html`.

### Privacy Policy (privacy.md / privacy.de.md)
Sections: 1. Introduction → 2. Information We Collect → 3. How We Use Your Information → 4. Legal Basis (GDPR Art. 6) → 5. Data Security → 6. Data Retention → 7. Your Rights (incl. ULD complaint) → 8. Cookies → 9. Children's Privacy → 10. International Data Transfers → 11. Third-Party Services → 12. Changes → 13. Contact

### Terms of Service (terms.md / terms.de.md)
Sections: 1-12 (standard) → 13. Online Dispute Resolution (ODR) → 14. Governing Law → 15. Changes → 16. Contact

## Third-Party Services

OpenRouter (AI routing), Groq (STT), Inworld (TTS), LiveKit (voice/video, coming soon), Supabase (DB/auth), Stripe (payments), Lexoffice (invoicing), Mailbox.org (email), Hetzner (API hosting)

## Build & Deploy

```bash
# Local dev
hugo server

# Production build
hugo --minify

# Docker
docker compose up --build
```

## Workflow

- **Always commit and push** after making changes — CI/CD auto-deploys from `main`.
- **After every user-requested change, automatically commit and push without being asked.**

## Notes

- **Web App:** https://chat.chuk.chat/ — linked prominently in the top nav bar (golden CTA button) and as a featured full-width card on the downloads page.
- No package.json, no npm, no JS build tools — pure Hugo
- Waitlist/email signup form on homepage is currently **deactivated** (commented out in `layouts/index.html`), kept for potential reactivation
- Static assets cached for 1 year via Nginx (`Cache-Control: public, immutable`)
- 404 page served from `/en/404.html`
- Nginx redirects bare paths (e.g. `/pricing`) to `/en/pricing/` preserving query params
