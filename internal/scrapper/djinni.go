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
}

func Scrap() []Job {
	c := colly.NewCollector(
		colly.AllowedDomains("djinni.co"),
	)

	var jobs []Job

	// Find and visit all job listings
	c.OnHTML(".job-list-item", func(e *colly.HTMLElement) {
		title := e.ChildText(".job-list-item__title a")
		company := e.ChildText(".job-list-item__pic + a")
		location := strings.TrimSpace(e.ChildText(".location-text"))
		salary := e.ChildText(".public-salary-item")
		detailsURL := e.ChildAttr(".job-list-item__title a", "href")

		job := Job{
			Title:      title,
			Company:    company,
			Location:   location,
			Salary:     salary,
			DetailsURL: "https://djinni.co" + detailsURL,
		}
		jobs = append(jobs, job)
	})

	// Paginate to the next page
	c.OnHTML(".pagination li.next a", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(nextPage))
	})

	// Set error handler
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// Start scraping
	c.Visit("https://djinni.co/jobs/?primary_keyword=Golang")

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
