package wapi

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Player struct {
	Name       string
	Number     string
	Position   string
	Height     string
	Weight     string
	Birthdate  string
	Age        string
	Experience string
	School     string
	Aquired    string
}

type Game struct {
	Opponent string
	Date     string
	Location string
	Status   string
	Notes    string
}

type TeamInfo struct {
	Record string
	Seed   string
}

func GetRoster() ([]Player, error) {
	response, err := http.Get("https://www.nba.com/team/1610612744")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	var players []Player

	doc.Find("div.TeamRoster_tableContainer__CUtM0").Each(func(i int, s *goquery.Selection) {
		s.Find("tbody").Each(func(j int, row *goquery.Selection) {
			row.Find("tr").Each(func(k int, t *goquery.Selection) {
				name := t.Find("td.primary.text").Text()
				number := t.Find("td:nth-of-type(2)").Text()
				position := t.Find("td:nth-of-type(3)").Text()
				height := t.Find("td:nth-of-type(4)").Text()
				weight := t.Find("td:nth-of-type(5)").Text()
				birthdate := t.Find("td:nth-of-type(6)").Text()
				age := t.Find("td:nth-of-type(7)").Text()
				experience := t.Find("td:nth-of-type(8)").Text()
				school := t.Find("td:nth-of-type(9)").Text()
				aquired := t.Find("td:nth-of-type(10)").Text()

				players = append(players, Player{Name: name, Number: number, Position: position, Height: height, Weight: weight,
					Birthdate: birthdate, Age: age, Experience: experience, School: school, Aquired: aquired})
			})
		})
	})
	return players, nil
}

func GetSchedule() ([]Game, error) {
	response, err := http.Get("https://www.nba.com/team/1610612744/schedule")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	var games []Game

	doc.Find("tbody.Crom_body__UYOcU").Each(func(i int, s *goquery.Selection) {
		s.Find("tr.TeamScheduleTable_game__STLzU").Each(func(j int, row *goquery.Selection) {
			date := row.Find("td.Crom_text__NpR1_:nth-of-type(1)").Text()
			opponent := row.Find("td.Crom_text__NpR1_:nth-of-type(2)").Text()
			location := row.Find("td.Crom_text__NpR1_:nth-of-type(5)").Text()
			status := row.Find("td.Crom_text__NpR1_:nth-of-type(3)").Text()
			notes := row.Find("td.Crom_text__NpR1_:nth-of-type(6)").Text()
			games = append(games, Game{Date: date, Opponent: opponent, Location: location, Status: status, Notes: notes})
		})
	})
	return games, nil
}

func GetTeamInfo() (TeamInfo, error) {
	response, err := http.Get("https://www.nba.com/stats/team/1610612744/traditional?SeasonType=Regular+Season")
	if err != nil {
		return TeamInfo{}, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return TeamInfo{}, err
	}

	var teamInfo TeamInfo
	doc.Find("div.TeamHeader_record__wzofp").Each(func(i int, s *goquery.Selection) {
		record := s.Find("span:nth-of-type(1)").Text()
		seed := s.Find("span:nth-of-type(2)").Text()

		teamInfo.Record = record

		parts := strings.Split(seed, " ")
		if len(parts) > 0 {
			seedNumber := strings.Trim(parts[2], "| th")
			teamInfo.Seed = seedNumber
		}
	})
	return teamInfo, nil
}
