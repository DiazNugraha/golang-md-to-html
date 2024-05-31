package main

import (
	"flag"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func main() {
	// parse args
	flag.Parse()
	arg := flag.Arg(0)
	fileName := flag.Arg(1)
	if arg == "" {
		panic("arg is empty")
	}

	// read file
	dat, err := os.ReadFile(arg)
	if err != nil {
		panic(err)
	}

	// convert markdown to html
	mardownText := mdToHtml(dat)

	// write to file
	os.WriteFile(fileName, mardownText, 0644)
}

func mdToHtml(md []byte) []byte {
	// create extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock

	// create markdown parser
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
