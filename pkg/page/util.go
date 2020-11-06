package page

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/yuin/goldmark"
	"golang.org/x/net/html"
)

func MarkdownFileToNodes(path string) ([]Node, error) {
	// read data from file
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, errors.New("Couldn't read markdown file in the given path.")
	}
	// read markdown, and convert to html strings
	htmlText, err := MarkdownToHTML(string(data))
	if err != nil {
		return nil, errors.New("Couldn't parse markdown string to html.")
	}
	// create []Node from html string
	nodes, err := NodesFromHTML(htmlText)

	if err != nil {
		return nil, errors.New("Couldn't parse generated HTML from markdown.")
	}

	return nodes, nil

}

func MarkdownToHTML(markdownText string) (string, error) {
	var buf bytes.Buffer
	// "# Heading \n\n paragraph \n *italic paragraph* normal paragraph"
	source := []byte(markdownText)
	if err := goldmark.Convert(source, &buf); err != nil {
		return "", errors.New("Couldn't convert markdown to html. :(")
	}

	return buf.String(), nil
}

func NodesFromHTML(htmlText string) ([]Node, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlText))

	if err == nil {
		return traverseNodes(doc.Find("body").Contents()), nil
	}

	return nil, err
}

// traverse DOM for creating new nodes
func traverseNodes(selections *goquery.Selection) []Node {
	nodes := []Node{}

	var tag string
	var attrs map[string]string
	var element NodeElement

	selections.Each(func(_ int, child *goquery.Selection) {
		for _, node := range child.Nodes {
			switch node.Type {
			case html.TextNode:
				// ignore single occurences of '\n'
				if node.Data == "\n" {
					continue
				}
				nodes = append(nodes, node.Data) // append text
			case html.ElementNode:
				// attributes
				attrs = map[string]string{}
				for _, attr := range node.Attr {
					attrs[attr.Key] = attr.Val
				}
				// new node element
				if len(node.Namespace) > 0 {
					tag = fmt.Sprintf("%s.%s", node.Namespace, node.Data)
				} else {
					tag = node.Data
				}
				element = NodeElement{
					Tag:      tag,
					Attrs:    attrs,
					Children: traverseNodes(child.Contents()),
				}

				nodes = append(nodes, element) // append element
			default:
				continue // skip other things
			}
		}
	})

	return nodes
}
