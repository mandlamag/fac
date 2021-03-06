package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

var YAML = Register(MustNewLexer(
	&Config{
		Name:      "YAML",
		Aliases:   []string{"yaml"},
		Filenames: []string{"*.yaml", "*.yml"},
		MimeTypes: []string{"text/x-yaml"},
	},
	Rules{
		"root": {
			Include("whitespace"),
			{`#.*`, Comment, nil},
			{`![^\s]+`, CommentPreproc, nil},
			{`\s&[^\s]+`, CommentPreproc, nil},
			{`\s\*[^\s]+`, CommentPreproc, nil},
			{`^%include\s+[^\n\r]+`, CommentPreproc, nil},
			{`([>|])(\s+)((?:(?:.*?$)(?:[\n\r]*?\2)?)*)`, ByGroups(StringDoc, StringDoc, StringDoc), nil},
			Include("value"),
			{`[?:,\[\]]`, Punctuation, nil},
			{`.`, Text, nil},
		},
		"value": {
			{Words(``, `\b`, "true", "false", "null"), KeywordConstant, nil},
			{`"(?:\\.|[^"])*"`, StringDouble, nil},
			{`'(?:\\.|[^'])*'`, StringSingle, nil},
			{`\d\d\d\d-\d\d-\d\d([T ]\d\d:\d\d:\d\d(\.\d+)?(Z|\s+[-+]\d+)?)?`, LiteralDate, nil},
			{`\b[+\-]?(0x[\da-f]+|0o[0-7]+|(\d+\.?\d*|\.?\d+)(e[\+\-]?\d+)?|\.inf|\.nan)\b`, Number, nil},
		},
		"whitespace": {
			{`\s+`, Whitespace, nil},
		},
	},
))
