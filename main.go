package main

import (
	"log"
	"os"

	"github.com/IraIvanishak/stack_app/pkg/gpt"
)

type Student struct {
	Name   gpt.GPTModelProperty `json:"name"`
	Major  gpt.GPTModelProperty `json:"major"`
	School gpt.GPTModelProperty `json:"school"`
	Grades gpt.GPTModelProperty `json:"grades"`
	Club   gpt.GPTModelProperty `json:"club"`
}

func main() {
	api_key := os.Getenv("OPENAI_API_KEY")
	client := gpt.NewGPT(api_key)

	studentFunction := gpt.GPTFunction{
		Name:        "extract_student_info",
		Description: "Get the student information from the body of the input text",
		Parameters: gpt.GPTFunctionParameter{
			Type: "object",
			Properties: Student{
				Name: gpt.GPTModelProperty{
					Type:        gpt.GPTPropertyTypeString,
					Description: "The name of the student",
				},
				Major: gpt.GPTModelProperty{
					Type:        gpt.GPTPropertyTypeString,
					Description: "The major of the student",
				},
				School: gpt.GPTModelProperty{
					Type:        gpt.GPTPropertyTypeString,
					Description: "The university name",
				},
				Grades: gpt.GPTModelProperty{
					Type:        gpt.GPTPropertyTypeInteger,
					Description: "GPA of the student",
				},
				Club: gpt.GPTModelProperty{
					Type:        gpt.GPTPropertyTypeString,
					Description: "The club the student is part of",
				},
			},
		},
	}

	// FROM https://www.datacamp.com/tutorial/open-ai-function-calling-tutorial
	studentDescription := "David Nguyen is a sophomore majoring in computer science at Stanford University. He is Asian American and has a 3.8 GPA. David is known for his programming skills and is an active member of the university's Robotics Club. He hopes to pursue a career in artificial intelligence after graduating."

	response, err := client.NewChatCompletion(studentDescription, gpt.GPT4, studentFunction)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response.Choices[0].Message.FunctionCall.Arguments)
}
