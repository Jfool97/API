package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/login"
	"strings"
	"sync"
)

type Advisor struct {
	Status string `json:"status"`
	Details map[string]string `json:"faculty_det"`
}


func FacultyAdvisor(bow *browser.Browser,regno, password string) *Advisor{
	response := login.NewLogin(bow,regno,password)
	status := "Success"
	dets := make(map[string]string)
	if response.Status == 0 {
		status = "Failure"
	} else {
		var wg sync.WaitGroup
		bow.Open("https://academics.vit.ac.in/student/faculty_advisor_view.asp")
		//Reload
		bow.Open("https://academics.vit.ac.in/student/faculty_advisor_view.asp")
		table := bow.Find("table").Eq(1)
		rows := table.Find("tr").Length()
		table.Find("tr").Each(func(i int, s *goquery.Selection){
			if i>0 && i<rows-1 {
				wg.Add(1)
				go func(){
					defer wg.Done()
					td := s.Find("td")
					dets[strings.TrimSpace(td.Eq(0).Text())] = strings.TrimSpace(td.Eq(1).Text())
					
				}()
			}
		})
		wg.Wait()
	}
	return &Advisor{
		Status: status,
		Details: dets,
	}	
}
