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
	git, err := gitlab.NewClient("tf8dgQf_wkMysAfpcaTw", gitlab.WithBaseURL("https://git.ecd.axway.org"))
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
	git, err := gitlab.NewClient("tf8dgQf_wkMysAfpcaTw", gitlab.WithBaseURL("https://git.ecd.axway.org"))
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
			langCoverage.ttlProject++
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
		// Exit the loop when we've seen all pages.
		if resp.CurrentPage >= resp.TotalPages {
			break
		}

		// Update the page number to get the next page.
		opt.Page = resp.NextPage
	}
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
		entry:=types.Entry{Timeline: makeTimeLineEntry(k,i, coverage.maxLang),ID: k, Key: k, Title: k, URL:"#", Description: k, Quadrant: "Languages"}
		radar.Entries = append(radar.Entries, entry)

	}
	return &radar

}
func makeTimeLineEntry(k string, i int, max int) []types.TimelineEntry {
	tl:=[]types.TimelineEntry{}
	ring:= HighUse // default to high
	percentile:=float32(i/max)
	if percentile < .5 && percentile >= .25 {
		ring=MediumUse
	} else if percentile < .25 {
		ring= LowUse
	}
	tl = append(tl, types.TimelineEntry{RingID: ring, Date: time.Now()})
	return tl
}
