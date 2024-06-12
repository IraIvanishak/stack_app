package model

type Vacancy struct {
	Categories    []string `json:"categories"`
	RequiredStack []string `json:"required_stack"`
	WelcomeStack  []string `json:"welcome_stack"`
}

// {
// 	"categories": [ //з тайтлу
// 	  "React",
// 	  "Node"
// 	],
// 	"required_stack": [
// 	  "JS",
// 	  "React",
// 	  "SCSS",
// 	  "Node.js",
// 	  "Next.js"
// 	]
// 	"welcome_stack": [
// 	  "TS",
// 	  "AWS"
// 	]
//   }
