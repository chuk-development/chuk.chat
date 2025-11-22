# Deployment Guide

This guide will help you deploy the Chuk Chat website with proper nginx configuration to support clean URLs (without `.html` extensions).

## The Problem

Your nginx logs show 404 errors when accessing `/subscription` because nginx is looking for a file without the `.html` extension.

## The Solution

I've created the necessary configuration files to handle clean URLs properly.

## Option 1: Docker Deployment (Recommended)

### Quick Start

1. **Build and run with Docker Compose:**
   ```bash
   docker-compose up -d --build
   ```

2. **Or build and run with Docker directly:**
   ```bash
   docker build -t chuk-chat .
   docker run -d -p 80:80 --name chuk-chat chuk-chat
   ```

3. **Verify it's working:**
   ```bash
   curl http://localhost/subscription?success=true
   curl http://localhost/dl?code=test123
   ```

### Update deployment:
```bash
docker-compose down
docker-compose up -d --build
```

## Option 2: Update Existing Nginx Configuration

If you're running nginx directly (not in Docker), update your nginx configuration:

1. **Find your nginx config:**
   ```bash
   # Common locations:
   # /etc/nginx/sites-available/default
   # /etc/nginx/conf.d/default.conf
   # /etc/nginx/nginx.conf
   ```

2. **Backup your current config:**
   ```bash
   sudo cp /etc/nginx/sites-available/default /etc/nginx/sites-available/default.backup
   ```

3. **Copy the new config:**
   ```bash
   sudo cp nginx.conf /etc/nginx/sites-available/default
   ```

4. **Test the configuration:**
   ```bash
   sudo nginx -t
   ```

5. **Reload nginx:**
   ```bash
   sudo systemctl reload nginx
   ```

## Option 3: Cloudflare/CDN Configuration

If you're using Cloudflare or another CDN:

1. Add a **Transform Rule** or **Page Rule**:
   - If URI Path matches `^/subscription$` → Rewrite to `/subscription.html`
   - If URI Path matches `^/dl$` → Rewrite to `/dl.html`
   - If URI Path matches `^/pricing$` → Rewrite to `/pricing.html`
   - If URI Path matches `^/downloads$` → Rewrite to `/downloads.html`
   - If URI Path matches `^/support$` → Rewrite to `/support.html`

## Testing Your URLs

After deployment, test these URLs:

- ✅ `https://chuk.chat/subscription?success=true`
- ✅ `https://chuk.chat/subscription?success=false`
- ✅ `https://chuk.chat/dl?code=4a4a274b-fca6-455f-aeb7-3575e5d6468d`
- ✅ `https://chuk.chat/pricing`
- ✅ `https://chuk.chat/downloads`
- ✅ `https://chuk.chat/support`

## Files Created

- `nginx.conf` - Nginx configuration with clean URL support
- `Dockerfile` - Docker container configuration
- `docker-compose.yml` - Docker Compose orchestration
- `subscription.html` - Payment success/failure page
- `dl.html` - Signup page with code processing

## Configuration Features

The nginx configuration includes:

- ✅ Clean URLs (no `.html` extension needed)
- ✅ Gzip compression for better performance
- ✅ Security headers (X-Frame-Options, X-XSS-Protection, etc.)
- ✅ Static asset caching (1 year for images, fonts, etc.)
- ✅ Fallback to index.html for 404s

## Troubleshooting

### Still getting 404 errors?

1. **Check if files exist:**
   ```bash
   ls -la /usr/share/nginx/html/
   ```

2. **Check nginx error logs:**
   ```bash
   docker logs chuk-chat  # For Docker
   sudo tail -f /var/log/nginx/error.log  # For system nginx
   ```

3. **Verify nginx config is loaded:**
   ```bash
   docker exec chuk-chat nginx -T  # For Docker
   sudo nginx -T  # For system nginx
   ```

### Need to update just the HTML files?

If using Docker with volumes (docker-compose.yml), just edit the HTML files locally - changes will reflect immediately.

If using Docker without volumes, rebuild:
```bash
docker-compose up -d --build
```

## Security Note

Remember to configure HTTPS/SSL in production! The current config only handles HTTP (port 80).

For HTTPS, you can use:
- Cloudflare (free SSL)
- Let's Encrypt with certbot
- Your hosting provider's SSL solution
