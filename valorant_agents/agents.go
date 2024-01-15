package valorantagents

import "github.com/gocolly/colly"

func Get_Agents() []string {
	c := colly.NewCollector()
	slice := make([]string, 0, 50)

	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, el *colly.HTMLElement) {
			if i > 0 {
				slice = append(slice, el.ChildText("td:nth-child(2)"))
			}
		})
	})

	c.Visit("https://valorant.fandom.com/ru/wiki/%D0%90%D0%B3%D0%B5%D0%BD%D1%82%D1%8B")
	
	return slice
}