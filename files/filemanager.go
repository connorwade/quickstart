package files

import (
	_ "embed"
)

//go:embed app.css
var AppCSS []byte

//go:embed postcss.config.js
var PostCSSConfig []byte

//go:embed tailwind.config.js
var TailwindConfig []byte
