# 🦦 Go & Template & Htmx & Alpine

Bu repo, **Go ile html/template, HTMX ve Alpine.js kullanarak adım adım blog uygulaması geliştirme** serisindeki örnek kodları içerir.  
Her branch, serinin farklı bir aşamasındaki kod hâlini barındırır.  

📚 Blog serisinin tüm yazılarına buradan ulaşabilirsiniz:  
[🔗 Blog Yazı Serisi Ana Sayfa](https://blog.uygarceylan.net/)

---

## 📌 Branch'lar ve Blog Yazıları

| Branch Adı             | Açıklama | Blog Yazısı |
|------------------------|----------|-------------|
| **main**               | Go + html/template + HTMX + Alpine (final sürüm) | [Bölüm 4: Üçünü Birleştirmek – Final Uygulama]() |
| **go-alpine**          | Go + html/template + HTMX + Alpine (HTMX + Alpine demosu) | [Bölüm 3: Alpine.js ile Dinamiklik Katmak]() |
| **go-htmx**            | Go + html/template + HTMX (partial render örnekleri) | [Bölüm 2: HTMX ile Progressive Enhancement]() |
| **go-html-template**   | Go + html/template (layout + partial sistemi, temel SSR) | [Bölüm 1: Go html/template ile SSR Mantığı](https://blog.uygarceylan.net/go-html-template-ile-temiz-ui-base-partial-funcmap) |

---

## 🚀 Projeyi Çalıştırma

### 1️⃣ Gerekli araçlar
- [Go 1.22+](https://go.dev/dl/)
- [Node.js 20+](https://nodejs.org/en/download/) (TailwindCSS derlemesi için)
- [npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)

---

### 2️⃣ Kurulum

```bash
# Repo'yu klonla
git clone https://github.com/uodev/go-htmx-blog.git
cd go-htmx-blog

# Branch seç (örnek: temel html/template versiyonu)
git checkout go-html-template

# Tailwind bağımlılıklarını yükle
npm install
