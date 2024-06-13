package analyzer

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/IraIvanishak/stack_app/internal/model"
	"github.com/IraIvanishak/stack_app/pkg/gpt"
)

type VacancyGPT struct {
	RequiredStack gpt.GPTModelProperty `json:"required_stack"`
	WelcomeStack  gpt.GPTModelProperty `json:"welcome_stack"`
}

var client = gpt.NewGPT(os.Getenv("OPENAI_API_KEY"), gpt.GPT3Dot5Turbo)

func Analyze(data []model.Vacancy) []model.Vacancy {
	vacancyInfo := gpt.GPTFunction{
		Name:        "vacancyInfo",
		Description: "Get the vacancy information from the body of the input text",
		Parameters: gpt.GPTFunctionParameter{
			Type: "object",
			Properties: VacancyGPT{
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
	wg.Add(len(data))

	vacancies := make([]model.Vacancy, len(data))

	t1 := time.Now()
	for i, text := range data {
		go func(text string) {
			defer wg.Done()

			vacancy := model.Vacancy{}
			cat := data[i].Category

			err := client.NewChatCompletion(text, &vacancy, vacancyInfo)
			if err != nil {
				fmt.Printf("error: %v", err)
			}

			vacancy.Text = text
			vacancy.Category = cat
			vacancies[i] = vacancy
		}(text.Text)
	}

	wg.Wait()
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	return vacancies
}
