package scrapper

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Vacancy struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

const filename = "vacancies.json"

func Save(data []Vacancy) error {
	dataBytes, err := json.Marshal(data)
	if err != nil || dataBytes == nil {
		return err
	}

	return os.WriteFile(filename, dataBytes, 0644)
}

func Read() ([]Vacancy, error) {
	dataBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var vacancies []Vacancy
	err = json.Unmarshal(dataBytes, &vacancies)
	if err != nil {
		return nil, err
	}

	return vacancies, nil
}

type Job struct {
	Title       string
	Company     string
	Location    string
	Salary      string
	Description string
	DetailsURL  string
	Category    string
}

func Scrap() []Job {
	c := colly.NewCollector(
		colly.AllowedDomains("djinni.co"),
	)

	var jobs []Job
	page := 1
	baseURL := "https://djinni.co/jobs/?primary_keyword=Golang&page="

	// Find and visit all job listings
	c.OnHTML(".job-list-item", func(e *colly.HTMLElement) {
		title := e.ChildText(".job-list-item__title a")
		company := e.ChildText(".job-list-item__pic + a")
		location := strings.TrimSpace(e.ChildText(".location-text"))
		salary := e.ChildText(".public-salary-item")
		detailsURL := e.ChildAttr(".job-list-item__title a", "href")
		category := e.ChildText(".job-list-item__info a[href*='primary_keyword']")

		job := Job{
			Title:      title,
			Company:    company,
			Location:   location,
			Salary:     salary,
			DetailsURL: "https://djinni.co" + detailsURL,
			Category:   category,
		}
		jobs = append(jobs, job)
	})

	// Check if there is a next page
	c.OnHTML("ul.pagination", func(e *colly.HTMLElement) {
		if e.DOM.Find("li.page-item a.page-link").Length() > 0 {
			page++
			nextPageURL := fmt.Sprintf("%s%d", baseURL, page)
			fmt.Println("Visiting next page:", nextPageURL)
			c.Visit(nextPageURL)
		}
	})

	// Set error handler
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// Start scraping
	c.Visit(baseURL + "1")

	c.Wait()

	// Clone collector for fetching job descriptions
	jobDetailsCollector := c.Clone()

	// Fetch job descriptions
	for i := range jobs {
		job := &jobs[i]

		jobDetailsCollector.OnHTML(".job-post-description", func(e *colly.HTMLElement) {
			// Get the full description
			fullDescription := e.Text
			job.Description = strings.TrimSpace(fullDescription)
		})

		jobDetailsCollector.Visit(job.DetailsURL)
		jobDetailsCollector.Wait()
	}

	return jobs
}
