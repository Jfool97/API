/*
@Author Shubhodeep Mukherjee
@Organization Google Developers Group VIT Vellore
	Isn't Go sexy?
	I know right!!
	Just like Shubhodeep
	I mean, have you seen the guy? xP
	#GDGSwag
*/

package api

import (
	"github.com/patrickmn/go-cache"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf"
	"go-MyVIT/api/login"
	"go-MyVIT/api/scrape"
	"time"
)

var cac *cache.Cache = cache.New(2*time.Minute, 30*time.Second)

//Executable script to Login
func LogIn(regno, password, baseuri string) *login.Response {
	bow := surf.NewBrowser()
	return login.NewLogin(bow, regno, password, baseuri, cac)
}

//Executable script to show timetable
func TimeTable(regno, password, baseuri string) *scrape.Timetable {
	bow := surf.NewBrowser()
	return scrape.ShowTimetable(bow, regno, password, baseuri, cac)
}

//Executable script to show Faculty Advisor details
func Advisor(regno, password, baseuri string) *scrape.Advisor {
	bow := surf.NewBrowser()
	return scrape.FacultyAdvisor(bow, regno, password, baseuri, cac)
}

//Executable script to show Attendance
func Attendance(regno, password, baseuri string) *scrape.Attendance {
	bow := surf.NewBrowser()
	return scrape.ShowAttendance(bow, regno, password, baseuri, cac)
}
func Schedule(regno, password, baseuri string) *scrape.ExamSchedule {
	bow := surf.NewBrowser()
	return scrape.ExmSchedule(bow, regno, password, baseuri, cac)
}

func AcademicHistory(regno, password, baseuri string) *scrape.AcademicStruct {
	bow := surf.NewBrowser()
	return scrape.Academics(bow, regno, password, baseuri, cac)
}

func CourseCoursesPage(regno, password, baseuri string) *scrape.CourseStruct {
	bow := surf.NewBrowser()
	return scrape.Courses(bow, regno, password, baseuri, cac)
}

func CourseSlotsPage(regno, password, baseuri, coursekey string) *scrape.SlotsStruct {
	bow := surf.NewBrowser()
	return scrape.Slots(bow, regno, password, baseuri, coursekey, cac)
}

func CourseFacPage(regno, password, baseuri, coursekey, slt string) *scrape.FacStruct {
	bow := surf.NewBrowser()
	return scrape.Facs(bow, regno, password, baseuri, coursekey, slt, cac)
}

func CourseDataPage(regno, password, baseuri, coursekey, slt, fac string) *scrape.CourseDataStruct {
	bow := surf.NewBrowser()
	return scrape.CourseData(bow, regno, password, baseuri, coursekey, slt, fac, cac)
}
