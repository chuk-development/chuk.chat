// shotconv serves the Android store screenshots of the chuk_chat repo as
// resized JPEGs. nginx's image_filter can resize but cannot change the
// format, and these screenshots are large PNGs (300 KB+ at 480 px wide).
// As JPEG they are a fraction of that.
//
// It listens on 127.0.0.1 only. nginx proxies /app-shots/<width>/<path>
// for phone screenshots here and caches the result (see nginx.conf).
package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/image/draw"
)

const upstream = "https://raw.githubusercontent.com/chuk-development/chuk_chat/master/"

// Same whitelist as nginx: only these widths and only the store screenshots.
var route = regexp.MustCompile(`^/(480|960)/(fastlane/metadata/android/(?:en-US|de-DE)/images/phoneScreenshots/[a-z0-9_]+\.png)$`)

var client = &http.Client{Timeout: 30 * time.Second}

func handle(w http.ResponseWriter, r *http.Request) {
	m := route.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	width, _ := strconv.Atoi(m[1])

	resp, err := client.Get(upstream + m[2])
	if err != nil {
		http.Error(w, "upstream error", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		http.NotFound(w, r)
		return
	}
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("upstream status %d", resp.StatusCode), http.StatusBadGateway)
		return
	}

	src, _, err := image.Decode(io.LimitReader(resp.Body, 40<<20))
	if err != nil {
		http.Error(w, "decode error", http.StatusBadGateway)
		return
	}

	b := src.Bounds()
	if width > b.Dx() {
		width = b.Dx()
	}
	height := b.Dy() * width / b.Dx()
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, b, draw.Src, nil)

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, dst, &jpeg.Options{Quality: 84}); err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	w.Write(buf.Bytes())
}

func main() {
	addr := os.Getenv("SHOTCONV_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8082"
	}
	log.Printf("shotconv listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(handle)))
}
