package slackdown

import (
	"bytes"
	"io"

	bf "gopkg.in/russross/blackfriday.v2"
)

// Renderer is the rendering interface for slack output.
type Renderer struct {
	w             bytes.Buffer
	lastOutputLen int
}

var itemLevel = 0

var (
	strongTag        = []byte("*")
	strikethroughTag = []byte("~")
	itemTag          = []byte("-")
	codeTag          = []byte("`")
	codeBlockTag     = []byte("```")
)

var (
	nlBytes    = []byte{'\n'}
	spaceBytes = []byte{' '}
)

var escapes = [256][]byte{
	'&': []byte(`&amp;`),
	'<': []byte(`&lt;`),
	'>': []byte(`&gt;`),
}

func (r *Renderer) esc(w io.Writer, text []byte) {
	var start, end int
	for end < len(text) {
		if escSeq := escapes[text[end]]; escSeq != nil {
			w.Write(text[start:end])
			w.Write(escSeq)
			start = end + 1
		}
		end++
	}

	if start < len(text) && end <= len(text) {
		w.Write(text[start:end])
	}
}

func (r *Renderer) out(w io.Writer, text []byte) {
	w.Write(text)
	r.lastOutputLen = len(text)
}

func (r *Renderer) cr(w io.Writer) {
	if r.lastOutputLen > 0 {
		r.out(w, nlBytes)
	}
}

// RenderNode parses a single node of a syntax tree.
func (r *Renderer) RenderNode(w io.Writer, node *bf.Node, entering bool) bf.WalkStatus {

	switch node.Type {
	case bf.Text:
		r.esc(w, node.Literal)
	case bf.Softbreak:
		break
	case bf.Hardbreak:
		break
	case bf.BlockQuote:
		break
	case bf.CodeBlock:
		r.out(w, codeBlockTag)
		r.esc(w, node.Literal)
		r.out(w, codeBlockTag)
		r.cr(w)
		r.cr(w)
		break
	case bf.Code:
		r.out(w, codeTag)
		r.esc(w, node.Literal)
		r.out(w, codeTag)
		break
	case bf.Emph:
		break
	case bf.Heading:
		if entering {
			r.out(w, strongTag)
		} else {
			r.out(w, strongTag)
			r.cr(w)
		}
	case bf.Image:
		break
	case bf.Item:
		if entering {
			r.out(w, spaceBytes)
			r.out(w, itemTag)
			r.out(w, spaceBytes)
		} else {
			r.cr(w)
		}
		break
	case bf.Link:
		break
	case bf.HorizontalRule:
		break
	case bf.List:
		if entering {
			itemLevel++
		} else {
			itemLevel--
			if itemLevel == 0 {
				r.cr(w)
				r.cr(w)
			}
		}
		break
	case bf.Document:
		break
	case bf.Paragraph:
		break
	case bf.Strong:
		break
	case bf.Del:
		r.out(w, strikethroughTag)
		break
	case bf.Table:
		break
	case bf.TableCell:
		break
	case bf.TableHead:
		break
	case bf.TableBody:
		break
	case bf.TableRow:
		break
	default:
		panic("Unknown node type " + node.Type.String())
	}
	return bf.GoToNext
}

// Render prints out the whole document from the ast.
func (r *Renderer) Render(ast *bf.Node) []byte {
	ast.Walk(func(node *bf.Node, entering bool) bf.WalkStatus {
		return r.RenderNode(&r.w, node, entering)
	})

	return r.w.Bytes()
}

// RenderHeader writes document header (unused).
func (r *Renderer) RenderHeader(w io.Writer, ast *bf.Node) {
}

// RenderFooter writes document footer (unused).
func (r *Renderer) RenderFooter(w io.Writer, ast *bf.Node) {
}

// Run prints out the confluence document.
func Run(input []byte, opts ...bf.Option) []byte {
	r := &Renderer{}
	optList := []bf.Option{bf.WithRenderer(r), bf.WithExtensions(bf.CommonExtensions)}
	optList = append(optList, opts...)
	parser := bf.New(optList...)
	ast := parser.Parse([]byte(input))
	return r.Render(ast)
}
