# peasy-audio-go

[![Go Reference](https://pkg.go.dev/badge/github.com/peasytools/peasy-audio-go.svg)](https://pkg.go.dev/github.com/peasytools/peasy-audio-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/peasytools/peasy-audio-go)](https://goreportcard.com/report/github.com/peasytools/peasy-audio-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go client for the [PeasyAudio](https://peasyaudio.com) API — analyze BPM, calculate bitrate, and convert audio formats. Built with `net/http`, `encoding/json`, and zero external dependencies.

Built from [PeasyAudio](https://peasyaudio.com), a comprehensive audio toolkit offering free online tools for analyzing tempo, calculating file sizes, comparing audio formats, and converting between MP3, WAV, FLAC, OGG, and AAC. The site includes in-depth guides on lossless vs. lossy audio encoding, format comparison charts, and a glossary covering concepts from bitrate and sample rate to audio codecs and clipping.

> **Try the interactive tools at [peasyaudio.com](https://peasyaudio.com)** — [Audio BPM Analyzer](https://peasyaudio.com/audio/audio-bpm/), [Audio Frequency Calculator](https://peasyaudio.com/audio/audio-freq/), [Audio File Size Calculator](https://peasyaudio.com/audio/audio-filesize/), and more.

<p align="center">
  <img src="demo.gif" alt="peasy-audio-go demo — audio BPM analysis and format conversion tools in Go terminal" width="800">
</p>

## Table of Contents

- [Install](#install)
- [Quick Start](#quick-start)
- [What You Can Do](#what-you-can-do)
  - [Audio Analysis Tools](#audio-analysis-tools)
  - [Browse Reference Content](#browse-reference-content)
  - [Search and Discovery](#search-and-discovery)
- [API Client](#api-client)
  - [Available Methods](#available-methods)
- [Learn More About Audio Tools](#learn-more-about-audio-tools)
- [Also Available](#also-available)
- [Peasy Developer Tools](#peasy-developer-tools)
- [License](#license)

## Install

```bash
go get github.com/peasytools/peasy-audio-go
```

Requires Go 1.21+.

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	peasyaudio "github.com/peasytools/peasy-audio-go"
)

func main() {
	client := peasyaudio.New()
	ctx := context.Background()

	// List available audio tools
	tools, err := client.ListTools(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tools.Results {
		fmt.Printf("%s: %s\n", t.Name, t.Description)
	}
}
```

## What You Can Do

### Audio Analysis Tools

Digital audio is represented as a series of samples captured at a fixed rate — CD-quality audio uses 44,100 samples per second (44.1 kHz) with 16-bit depth, producing 1,411 kbps of uncompressed data. Lossy codecs like MP3 and AAC reduce this dramatically (128-320 kbps) by discarding inaudible frequencies using psychoacoustic models, while lossless codecs like FLAC compress without any data loss. PeasyAudio provides calculators and analysis tools for understanding these encoding parameters.

| Tool | Slug | Description |
|------|------|-------------|
| BPM Analyzer | `audio-bpm` | Calculate beats per minute for tempo analysis |
| Frequency Calculator | `audio-freq` | Compute audio frequency values and wavelengths |
| File Size Calculator | `audio-filesize` | Estimate file sizes for different bitrate and duration combinations |

```go
// Retrieve audio analysis tools and explore their capabilities
client := peasyaudio.New()
ctx := context.Background()

// Get the BPM analyzer tool for tempo detection
tool, err := client.GetTool(ctx, "audio-bpm")
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Tool: %s\n", tool.Name)         // Audio BPM analyzer name
fmt.Printf("Description: %s\n", tool.Description) // How BPM detection works

// List all available audio tools with pagination
tools, err := client.ListTools(ctx, &peasyaudio.ListOptions{Page: 1, Limit: 20})
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Total audio tools available: %d\n", tools.Count)
```

Learn more: [Audio BPM Analyzer](https://peasyaudio.com/audio/audio-bpm/) · [Audio Format Comparison](https://peasyaudio.com/guides/audio-format-comparison/) · [Convert Between Audio Formats](https://peasyaudio.com/guides/convert-between-audio-formats/)

### Browse Reference Content

PeasyAudio includes a comprehensive glossary of audio engineering terminology and practical guides for common workflows. The glossary covers foundational concepts like bitrate (the number of bits processed per second, determining audio quality and file size), sample rate (how many times per second the audio signal is measured), WAV (Microsoft's uncompressed audio container), and FLAC (Free Lossless Audio Codec, the open-source standard for archival-quality audio).

| Term | Description |
|------|-------------|
| [Bitrate](https://peasyaudio.com/glossary/bitrate/) | Bits per second — determines audio quality and file size |
| [Sample Rate](https://peasyaudio.com/glossary/sample-rate/) | Samples per second — 44.1 kHz (CD), 48 kHz (video), 96 kHz (hi-res) |
| [WAV](https://peasyaudio.com/glossary/wav/) | Waveform Audio File Format — uncompressed PCM container |
| [FLAC](https://peasyaudio.com/glossary/flac/) | Free Lossless Audio Codec — open-source lossless compression |

```go
// Browse the audio glossary for encoding and format terminology
glossary, err := client.ListGlossary(ctx, &peasyaudio.ListOptions{
	Search: str("bitrate"), // Search for audio encoding concepts
})
if err != nil {
	log.Fatal(err)
}
for _, term := range glossary.Results {
	fmt.Printf("%s: %s\n", term.Term, term.Definition)
}

// Read a guide comparing lossless vs lossy audio formats
guide, err := client.GetGuide(ctx, "audio-format-comparison")
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Guide: %s (Level: %s)\n", guide.Title, guide.AudienceLevel)
```

Learn more: [Audio Glossary](https://peasyaudio.com/glossary/) · [Audio Format Comparison](https://peasyaudio.com/guides/audio-format-comparison/) · [Convert Between Audio Formats](https://peasyaudio.com/guides/convert-between-audio-formats/)

### Search and Discovery

The API supports full-text search across all content types — tools, glossary terms, guides, use cases, and format documentation. Search results are grouped by content type, making it easy to find the right tool or reference for any audio workflow.

```go
// Search across all audio content — tools, glossary, guides, and formats
results, err := client.Search(ctx, "convert flac", nil)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Found %d tools, %d glossary terms, %d guides\n",
	len(results.Results.Tools),
	len(results.Results.Glossary),
	len(results.Results.Guides),
)

// Discover format conversion paths — what can WAV convert to?
conversions, err := client.ListConversions(ctx, &peasyaudio.ListConversionsOptions{
	Source: str("wav"), // Find all formats WAV can be converted to
})
if err != nil {
	log.Fatal(err)
}
for _, c := range conversions.Results {
	fmt.Printf("%s → %s\n", c.SourceFormat, c.TargetFormat)
}
```

Learn more: [REST API Docs](https://peasyaudio.com/developers/) · [All Audio Tools](https://peasyaudio.com/)

## API Client

The client wraps the [PeasyAudio REST API](https://peasyaudio.com/developers/) with typed Go structs and zero external dependencies.

```go
client := peasyaudio.New()
// Or with a custom base URL:
// client := peasyaudio.New(peasyaudio.WithBaseURL("https://custom.example.com"))
ctx := context.Background()

// List tools with pagination
tools, _ := client.ListTools(ctx, &peasyaudio.ListOptions{Page: 1, Limit: 10})

// Get a specific tool by slug
tool, _ := client.GetTool(ctx, "audio-convert")
fmt.Println(tool.Name, tool.Description)

// Search across all content
results, _ := client.Search(ctx, "trim mp3", nil)
fmt.Printf("Found %d tools\n", len(results.Results.Tools))

// Browse the glossary
glossary, _ := client.ListGlossary(ctx, &peasyaudio.ListOptions{Search: str("bitrate")})
for _, term := range glossary.Results {
	fmt.Printf("%s: %s\n", term.Term, term.Definition)
}

// Discover guides
guides, _ := client.ListGuides(ctx, &peasyaudio.ListGuidesOptions{Category: str("encoding")})
for _, g := range guides.Results {
	fmt.Printf("%s (%s)\n", g.Title, g.AudienceLevel)
}

// List file format conversions
conversions, _ := client.ListConversions(ctx, &peasyaudio.ListConversionsOptions{Source: str("wav")})

// Get format details
format, _ := client.GetFormat(ctx, "mp3")
fmt.Printf("%s (%s): %s\n", format.Name, format.Extension, format.MimeType)
```

Helper for optional string parameters:

```go
func str(s string) *string { return &s }
```

### Available Methods

| Method | Description |
|--------|-------------|
| `ListTools(ctx, opts)` | List tools (paginated, filterable) |
| `GetTool(ctx, slug)` | Get tool by slug |
| `ListCategories(ctx, opts)` | List tool categories |
| `ListFormats(ctx, opts)` | List file formats |
| `GetFormat(ctx, slug)` | Get format by slug |
| `ListConversions(ctx, opts)` | List format conversions |
| `ListGlossary(ctx, opts)` | List glossary terms |
| `GetGlossaryTerm(ctx, slug)` | Get glossary term |
| `ListGuides(ctx, opts)` | List guides |
| `GetGuide(ctx, slug)` | Get guide by slug |
| `ListUseCases(ctx, opts)` | List use cases |
| `Search(ctx, query, limit)` | Search across all content |
| `ListSites(ctx)` | List Peasy sites |
| `OpenAPISpec(ctx)` | Get OpenAPI specification |

Full API documentation at [peasyaudio.com/developers/](https://peasyaudio.com/developers/).
OpenAPI 3.1.0 spec: [peasyaudio.com/api/openapi.json](https://peasyaudio.com/api/openapi.json).

## Learn More About Audio Tools

- **Tools**: [Audio BPM Analyzer](https://peasyaudio.com/audio/audio-bpm/) · [Audio Frequency Calculator](https://peasyaudio.com/audio/audio-freq/) · [Audio File Size Calculator](https://peasyaudio.com/audio/audio-filesize/) · [All Tools](https://peasyaudio.com/)
- **Guides**: [Audio Format Comparison](https://peasyaudio.com/guides/audio-format-comparison/) · [Convert Between Audio Formats](https://peasyaudio.com/guides/convert-between-audio-formats/) · [All Guides](https://peasyaudio.com/guides/)
- **Glossary**: [Bitrate](https://peasyaudio.com/glossary/bitrate/) · [Sample Rate](https://peasyaudio.com/glossary/sample-rate/) · [WAV](https://peasyaudio.com/glossary/wav/) · [FLAC](https://peasyaudio.com/glossary/flac/) · [All Terms](https://peasyaudio.com/glossary/)
- **Formats**: [All Formats](https://peasyaudio.com/formats/)
- **API**: [REST API Docs](https://peasyaudio.com/developers/) · [OpenAPI Spec](https://peasyaudio.com/api/openapi.json)

## Also Available

| Language | Package | Install |
|----------|---------|---------|
| **Python** | [peasy-audio](https://pypi.org/project/peasy-audio/) | `pip install "peasy-audio[all]"` |
| **TypeScript** | [peasy-audio](https://www.npmjs.com/package/peasy-audio) | `npm install peasy-audio` |
| **Rust** | [peasy-audio](https://crates.io/crates/peasy-audio) | `cargo add peasy-audio` |
| **Ruby** | [peasy-audio](https://rubygems.org/gems/peasy-audio) | `gem install peasy-audio` |

## Peasy Developer Tools

Part of the [Peasy Tools](https://peasytools.com) open-source developer ecosystem.

| Package | PyPI | npm | Go | Description |
|---------|------|-----|----|-------------|
| peasy-pdf | [PyPI](https://pypi.org/project/peasy-pdf/) | [npm](https://www.npmjs.com/package/peasy-pdf) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-pdf-go) | PDF merge, split, rotate, compress — [peasypdf.com](https://peasypdf.com) |
| peasy-image | [PyPI](https://pypi.org/project/peasy-image/) | [npm](https://www.npmjs.com/package/peasy-image) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-image-go) | Image resize, crop, convert, compress — [peasyimage.com](https://peasyimage.com) |
| **peasy-audio** | [PyPI](https://pypi.org/project/peasy-audio/) | [npm](https://www.npmjs.com/package/peasy-audio) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-audio-go) | **Audio trim, merge, convert, normalize — [peasyaudio.com](https://peasyaudio.com)** |
| peasy-video | [PyPI](https://pypi.org/project/peasy-video/) | [npm](https://www.npmjs.com/package/peasy-video) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-video-go) | Video trim, resize, thumbnails, GIF — [peasyvideo.com](https://peasyvideo.com) |
| peasy-css | [PyPI](https://pypi.org/project/peasy-css/) | [npm](https://www.npmjs.com/package/peasy-css) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-css-go) | CSS minify, format, analyze — [peasycss.com](https://peasycss.com) |
| peasy-compress | [PyPI](https://pypi.org/project/peasy-compress/) | [npm](https://www.npmjs.com/package/peasy-compress) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-compress-go) | ZIP, TAR, gzip compression — [peasytools.com](https://peasytools.com) |
| peasy-document | [PyPI](https://pypi.org/project/peasy-document/) | [npm](https://www.npmjs.com/package/peasy-document) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-document-go) | Markdown, HTML, CSV, JSON conversion — [peasyformats.com](https://peasyformats.com) |
| peasytext | [PyPI](https://pypi.org/project/peasytext/) | [npm](https://www.npmjs.com/package/peasytext) | [Go](https://pkg.go.dev/github.com/peasytools/peasytext-go) | Text case conversion, slugify, word count — [peasytext.com](https://peasytext.com) |

## License

MIT
