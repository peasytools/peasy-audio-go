# peasy-audio-go

[![Go Reference](https://pkg.go.dev/badge/github.com/peasytools/peasy-audio-go.svg)](https://pkg.go.dev/github.com/peasytools/peasy-audio-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/peasytools/peasy-audio-go)](https://goreportcard.com/report/github.com/peasytools/peasy-audio-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go client for the [PeasyAudio](https://peasyaudio.com) API — audio trim, merge, convert, and normalize. Zero dependencies beyond the Go standard library.

Built from [PeasyAudio](https://peasyaudio.com), a comprehensive audio processing toolkit offering free online tools for trimming, merging, converting, and normalizing audio files across all major formats with detailed format guides and glossary.

> **Try the interactive tools at [peasyaudio.com](https://peasyaudio.com)** — [Audio Tools](https://peasyaudio.com/), [Audio Glossary](https://peasyaudio.com/glossary/), [Audio Guides](https://peasyaudio.com/guides/)

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

## Learn More

- **Tools**: [Audio Convert](https://peasyaudio.com/tools/audio-convert/) · [Audio Trim](https://peasyaudio.com/tools/audio-trim/) · [Audio Merge](https://peasyaudio.com/tools/audio-merge/) · [All Tools](https://peasyaudio.com/)
- **Guides**: [Audio Encoding Guide](https://peasyaudio.com/guides/audio-encoding/) · [All Guides](https://peasyaudio.com/guides/)
- **Glossary**: [MP3](https://peasyaudio.com/glossary/mp3/) · [WAV](https://peasyaudio.com/glossary/wav/) · [All Terms](https://peasyaudio.com/glossary/)
- **Formats**: [MP3](https://peasyaudio.com/formats/mp3/) · [WAV](https://peasyaudio.com/formats/wav/) · [All Formats](https://peasyaudio.com/formats/)
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
