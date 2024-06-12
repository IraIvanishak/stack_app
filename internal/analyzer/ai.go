package analyzer

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/IraIvanishak/stack_app/data"
	"github.com/IraIvanishak/stack_app/internal/model"
	"github.com/IraIvanishak/stack_app/internal/utils"
	"github.com/IraIvanishak/stack_app/pkg/gpt"
)

type VacancyGPT struct {
	Categories    gpt.GPTModelProperty `json:"categories"`
	RequiredStack gpt.GPTModelProperty `json:"required_stack"`
	WelcomeStack  gpt.GPTModelProperty `json:"welcome_stack"`
}

var client = gpt.NewGPT(os.Getenv("OPENAI_API_KEY"), gpt.GPT3Dot5Turbo)

func Analyze() {
	vacancyInfo := gpt.GPTFunction{
		Name:        "vacancyInfo",
		Description: "Get the vacancy information from the body of the input text",
		Parameters: gpt.GPTFunctionParameter{
			Type: "object",
			Properties: VacancyGPT{
				Categories: gpt.GPTModelProperty{
					Type:        gpt.GPTPropertyTypeArray,
					Description: "Main tech stack from the title of the text. Don't include level, domain, etc.",
					Items: &gpt.GPTModelProperty{
						Type:        gpt.GPTPropertyTypeString,
						Description: "Tech unit from the title of the text",
					},
				},
				RequiredStack: gpt.GPTModelProperty{
					Type:        gpt.GPTPropertyTypeArray,
					Description: "Required stack of the vacancy from the body of the text",
					Items: &gpt.GPTModelProperty{
						Type:        gpt.GPTPropertyTypeString,
						Description: "Technology unit in standart naming",
					},
				},
				WelcomeStack: gpt.GPTModelProperty{
					Type:        gpt.GPTPropertyTypeArray,
					Description: "Welcome stack of the vacancy from the body of the text",
					Items: &gpt.GPTModelProperty{
						Type:        gpt.GPTPropertyTypeString,
						Description: "Technology unit in standart naming",
					},
				},
			},
		},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(data.TestData))

	vacancies := make([]model.Vacancy, len(data.TestData))

	t1 := time.Now()
	for i, text := range data.TestData {
		go func(text string) {
			defer wg.Done()

			vacancy := model.Vacancy{}

			err := client.NewChatCompletion(text, &vacancy, vacancyInfo)
			if err != nil {
				fmt.Printf("error: %v", err)
			}

			vacancies[i] = vacancy
		}(text)
	}

	wg.Wait()
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))

	utils.PrintVacanciesTable(vacancies)
}
