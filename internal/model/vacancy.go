package model

type Vacancy struct {
	Text          string   `json:"text"`
	Category      string   `json:"category"`
	RequiredStack []string `json:"required_stack,omitempty"`
	WelcomeStack  []string `json:"welcome_stack,omitempty"`
}

// {
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
