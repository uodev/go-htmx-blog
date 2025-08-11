# 🦦 Go & Template & Htmx & Alpine

`html/template` ile **güvenli, modüler ve modern** server-side render (SSR) uygulamaları geliştirmek için örnek bir proje.  
Adım adım ilerleyen branch yapısıyla **temel Go template yapısından** başlayıp, **HTMX** ve **Alpine.js** ile interaktif hale gelen final sürüme kadar gidebilirsiniz.

---

## 🚀 Özellikler

- **Güvenli SSR** → `html/template` ile otomatik XSS koruması
- **Layout / Partial sistemi** → Tekrarlayan HTML’den kurtulma
- **TailwindCSS ile şık UI** → Hafif ve kolay özelleştirilebilir tasarım
- **HTMX entegrasyonu** → Sayfanın belirli alanlarını dinamik olarak güncelleme
- **Alpine.js entegrasyonu** → Küçük ama etkili interaktivite çözümleri

---

## 📌 Branch'lar

| Branch | Açıklama |
|--------|----------|
| **main** | Final sürüm — Go + html/template + HTMX + Alpine |
| **go-alpine** | Go + html/template + HTMX + Alpine (HTMX + Alpine demosu) |
| **go-htmx** | Go + html/template + HTMX (partial render örnekleri) |
| **go-html-template** | Go + html/template (layout + partial sistemi, temel SSR) |

> **İpucu:** Branch’leri sırayla inceleyerek gelişim sürecini görebilirsiniz.

---

## 📂 Proje Yapısı

```plaintext
/cmd/server/main.go        → Server giriş noktası
/internal/http/router.go   → Route tanımları
/views                     → Template dosyaları
  /layouts/base.html       → Ana layout
  /partials/nav.html       → Navigasyon
  /pages                   → Sayfa şablonları
/assets/tailwind.css       → Tailwind kaynak dosyası
/public/app.css            → Derlenmiş CSS çıktısı
