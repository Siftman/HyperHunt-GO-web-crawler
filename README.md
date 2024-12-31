# 🕷️ HyperHunt-GO Web Crawler

A blazing fast, concurrent web crawler built in Go that helps you extract product information from e-commerce websites. This bad boy can handle sitemaps, crawl pages, and extract structured product data like a champ.

## 🚀 Features

- Sitemap detection and parsing
- Smart URL filtering and pattern matching
- Product schema extraction (supports JSON-LD)
- Fallback to OpenGraph meta tags
- CSV export functionality
- Concurrent crawling with rate limiting
- Price normalization (handles IRR currency)

## 🛠️ Tech Stack

- [Colly](https://github.com/gocolly/colly) - The beast powering our crawling
- [httpx](https://github.com/projectdiscovery/httpx) - For robust HTTP interactions
- [goflags](https://github.com/projectdiscovery/goflags) - CLI flag parsing
- [gologger](https://github.com/projectdiscovery/gologger) - Logging made sexy

## 🏃‍♂️ Quick Start

1. Clone this repo:
```bash
git clone https://github.com/yourusername/HyperHunt-GO-web-crawler.git
cd HyperHunt-GO-web-crawler
```

2. Install dependencies:
```bash
go mod download
```

3. Run it:
```bash
go run main.go
```

## 🎯 How It Works

1. First, it checks for a sitemap at common locations (`/sitemap.xml` or `/sitemap_index.xml`)
2. If found, it parses the sitemap to extract all product URLs
3. For each URL, it:
   - Attempts to extract product data from JSON-LD schema
   - Falls back to OpenGraph meta tags if needed
   - Normalizes prices and data formats
4. Exports results to CSV files:
   - `raw_links.csv`: All discovered URLs
   - `proper_urls.csv`: Filtered URLs matching product patterns

## 📦 Project Structure

```
.
├── main.go              # Entry point
├── pkg/
│   ├── crawler/        # Core crawling logic
│   ├── fileops/        # File operations (CSV handling)
│   ├── models/         # Data models
│   └── utils/          # Helper functions
```

## 🔧 Configuration

The crawler is configured to work with specific e-commerce sites out of the box. You can modify the base URL in `main.go`:

```go
baseURL := "https://your-target-site.com/"
```

## 🤝 Contributing

PRs are welcome! Just:

1. Fork it
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a PR