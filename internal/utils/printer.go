package utils

import (
	"fmt"
	"strings"

	"github.com/IraIvanishak/stack_app/internal/model"
)

func PrintVacanciesTable(vacancies []model.Vacancy) {
	if len(vacancies) == 0 {
		fmt.Println("No vacancies to display.")
		return
	}

	maxCategoryLen, maxRequiredStackLen, maxWelcomeStackLen := 10, 14, 13 // базові ширини на основі заголовків
	for _, v := range vacancies {
		for _, category := range v.Categories {
			if len(category) > maxCategoryLen {
				maxCategoryLen = len(category)
			}
		}
		for _, tech := range v.RequiredStack {
			if len(tech) > maxRequiredStackLen {
				maxRequiredStackLen = len(tech)
			}
		}
		for _, tech := range v.WelcomeStack {
			if len(tech) > maxWelcomeStackLen {
				maxWelcomeStackLen = len(tech)
			}
		}
	}

	printDivider := func() {
		fmt.Printf("+-%s-+-%s-+-%s-+\n", strings.Repeat("-", maxCategoryLen), strings.Repeat("-", maxRequiredStackLen), strings.Repeat("-", maxWelcomeStackLen))
	}

	printDivider()
	for _, v := range vacancies {
		for i := 0; i < max(len(v.Categories), len(v.RequiredStack), len(v.WelcomeStack)); i++ {
			category := ""
			requiredStack := ""
			welcomeStack := ""
			if i < len(v.Categories) {
				category = v.Categories[i]
			}
			if i < len(v.RequiredStack) {
				requiredStack = v.RequiredStack[i]
			}
			if i < len(v.WelcomeStack) {
				welcomeStack = v.WelcomeStack[i]
			}
			fmt.Printf("| %-*s | %-*s | %-*s |\n", maxCategoryLen, category, maxRequiredStackLen, requiredStack, maxWelcomeStackLen, welcomeStack)
		}
		printDivider()
	}
}

func max(values ...int) int {
	maxVal := values[0]
	for _, v := range values[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
