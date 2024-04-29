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

//go:embed cleanDB.ts
var CleanDB []byte

//go:embed drizzle.config.ts
var DrizzleConfig []byte

//go:embed migrate.ts
var Migrate []byte

//go:embed schema.ts
var Schema []byte

//go:embed db.ts
var DB []byte

