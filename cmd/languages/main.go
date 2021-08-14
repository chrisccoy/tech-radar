/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"github.com/chrisccoy/tech-radar/pkg/types"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/xanzy/go-gitlab"
)

const  HighUse="High Use"
const  MediumUse="Medium Use"
const  LowUse="Low Use"
const HighUseColor="#93c47d"
const MediumUseColor="#93d2c2"
const LowUseColor="#fbdb84"


func main() {
//	getLanguagesByGroup()
	getLanguages()
}
type LanguageCoverage struct {
	ttlProject int
	ttlLangs int
	maxLang int
	lang map[string]int
}
func getLanguagesByGroup() {

	langCoverage := LanguageCoverage{lang: make(map[string]int)}
	git, err := gitlab.NewClient(os.Getenv("GITLAB_TOKEN"), gitlab.WithBaseURL(os.Getenv("GITLAB_URL")))
	if err != nil {
		log.Fatal(err)
	}
		grps, _, err:= git.Groups.GetGroup(1024)
		if err != nil {
			log.Fatal(err)
		}

		for _, s := range grps.Projects {
			languages, _, err := git.Projects.GetProjectLanguages(s.ID)
			if err != nil {
				log.Fatal(err)
			}
			for l, _ := range *languages {
				curr := langCoverage.lang[l]
				curr++
				langCoverage.ttlLangs++
				langCoverage.lang[l] = curr
				if curr > langCoverage.maxLang {
					langCoverage.maxLang=curr
				}

			}
		}
	radar:= buildRadarData(langCoverage)
	radarj, err :=json.MarshalIndent(radar, "", "    ")
	log.Printf("Json radar: %s", radarj)
	ioutil.WriteFile("techradar.json", radarj, 0644)
}

func getLanguages() {

	langCoverage := LanguageCoverage{lang: make(map[string]int)}
	git, err := gitlab.NewClient(os.Getenv("GITLAB_TOKEN"), gitlab.WithBaseURL(os.Getenv("GITLAB_URL")))
	if err != nil {
		log.Fatal(err)
	}

	opt := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 50,
			Page:    1,
		},
	}

	for {
		proj, resp, err := git.Projects.ListProjects(opt)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Current State: %v", resp)

		for _, s := range proj {
			languages, _, err := git.Projects.GetProjectLanguages(s.ID)
			if err != nil {
				log.Fatal(err)
			}
			if len(*languages) > 0 {
				langCoverage.ttlProject++
			}
			for l, _ := range *languages {
				curr := langCoverage.lang[l]
				curr++
				langCoverage.ttlLangs++
				langCoverage.lang[l] = curr
				if curr > langCoverage.maxLang {
					langCoverage.maxLang=curr
				}


			}
		}
		// Exit the loop when we've seen all pages.
		if resp.CurrentPage >= resp.TotalPages {
			break
		}

		// Update the page number to get the next page.
		opt.Page = resp.NextPage
	}
	log.Printf("Coverages: %v", langCoverage)
	radar:= buildRadarData(langCoverage)
	radarj, err :=json.MarshalIndent(radar, "", "    ")
	log.Printf("Json radar: %s", radarj)
	ioutil.WriteFile("techradar.json", radarj, 0644)
}

func buildRadarData(coverage LanguageCoverage)*types.TechRadar {
	radar := types.TechRadar{}
	radar.Quadrants=append(radar.Quadrants,types.Quadrant{ID: "Languages", Name: "Languages"})
	radar.Rings=append(radar.Rings, types.Ring{ID: HighUse, Name: HighUse, Color: HighUseColor}, types.Ring{ID: MediumUse, Name: MediumUse, Color: MediumUseColor}, types.Ring{ID: LowUse, Name: LowUse,Color: LowUseColor} )
	for k, i := range coverage.lang {
		entry:=types.Entry{Timeline: makeTimeLineEntry(k, float32(i)/float32(coverage.maxLang), float32(coverage.maxLang)/float32(coverage.ttlProject)),ID: k, Key: k, Title: k, URL:"#", Description: k, Quadrant: "Languages"}
		radar.Entries = append(radar.Entries, entry)

	}
	addFluff(&radar)
	return &radar

}

func addFluff(t *types.TechRadar) {
	t.Quadrants=append(t.Quadrants,types.Quadrant{ID: "Data Storage", Name: "Data Storage"},
									types.Quadrant{ID: "Methodology", Name: "Methodology"},
									types.Quadrant{ID: "Release Cadence", Name: "Release Cadence"})
	// Data Storage
	t.Entries = append(t.Entries, types.Entry{Timeline: []types.TimelineEntry{{RingID: LowUse, Date: time.Now()}}, ID: "MongoDB", Key: "MongoDB", Title: "MongoDB", URL:"#", Description: "MongoDB", Quadrant: "Data Storage"})
	t.Entries = append(t.Entries, types.Entry{Timeline: []types.TimelineEntry{{RingID: MediumUse, Date: time.Now()}}, ID: "Kafka", Key: "Kafka", Title: "Kafka", URL:"#", Description: "Kafka", Quadrant: "Data Storage"})
	t.Entries = append(t.Entries, types.Entry{Timeline: []types.TimelineEntry{{RingID: MediumUse, Date: time.Now()}}, ID: "Cassandra", Key: "Cassandra", Title: "Cassandra", URL:"#", Description: "Cassandra", Quadrant: "Data Storage"})
	t.Entries = append(t.Entries, types.Entry{Timeline: []types.TimelineEntry{{RingID: HighUse, Date: time.Now()}}, ID: "Postgres", Key: "Postgres", Title: "Postgres", URL:"#", Description: "Postgres", Quadrant: "Data Storage"})
	t.Entries = append(t.Entries, types.Entry{Timeline: []types.TimelineEntry{{RingID: LowUse, Date: time.Now()}}, ID: "Oracle", Key: "Oracle", Title: "Oracle", URL:"#", Description: "Oracle", Quadrant: "Data Storage"})
	// Methodology
	t.Entries = append(t.Entries, types.Entry{Timeline: []types.TimelineEntry{{RingID: HighUse, Date: time.Now()}}, ID: "Scrum", Key: "Scrum", Title: "Scrum", URL:"#", Description: "Scrum", Quadrant: "Methodology"})
	t.Entries = append(t.Entries, types.Entry{Timeline: []types.TimelineEntry{{RingID: MediumUse, Date: time.Now()}}, ID: "Waterfall", Key: "Waterfall", Title: "Waterfall", URL:"#", Description: "Waterfall", Quadrant: "Methodology"})
	t.Entries = append(t.Entries, types.Entry{Timeline: []types.TimelineEntry{{RingID: LowUse, Date: time.Now()}}, ID: "Wing-it", Key: "Wing-it", Title: "Wing-it", URL:"#", Description: "Wing-it", Quadrant: "Methodology"})
	// Release Cadence
	t.Entries = append(t.Entries, types.Entry{Timeline: []types.TimelineEntry{{RingID: MediumUse, Date: time.Now()}}, ID: "Bi-Weekly", Key: "Bi-Weekly", Title: "Bi-Weekly", URL:"#", Description: "Bi-Weekly", Quadrant: "Release Cadence"})
	t.Entries = append(t.Entries, types.Entry{Timeline: []types.TimelineEntry{{RingID: HighUse, Date: time.Now()}}, ID: "Quarterly", Key: "Quarterly", Title: "Quarterly", URL:"#", Description: "Quarterly", Quadrant: "Release Cadence"})
	t.Entries = append(t.Entries, types.Entry{Timeline: []types.TimelineEntry{{RingID: LowUse, Date: time.Now()}}, ID: "Constant", Key: "Constant", Title: "Constant", URL:"#", Description: "Constant", Quadrant: "Release Cadence"})
	t.Entries = append(t.Entries, types.Entry{Timeline: []types.TimelineEntry{{RingID: LowUse, Date: time.Now()}}, ID: "Annual", Key: "Annual", Title: "Annual", URL:"#", Description: "Annual", Quadrant: "Release Cadence"})
}


func makeTimeLineEntry(k string, actual float32, percentile float32) []types.TimelineEntry {
	tl:=[]types.TimelineEntry{}
	ring:= HighUse // default to high

	if actual/percentile < .5 && actual/percentile >= .10 {
		ring=MediumUse
	} else if actual/percentile < .10 {
		ring= LowUse
	}
	tl = append(tl, types.TimelineEntry{RingID: ring, Date: time.Now()})
	return tl
}
