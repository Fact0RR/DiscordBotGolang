package saper

import (
	"log"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var saperStarted bool = false


func Controller(s *discordgo.Session, m *discordgo.MessageCreate,saperChan chan string) {
	if m.Content == "go" && !saperStarted {
		go Start(s, m, saperChan,&saperStarted)
		saperStarted = true
	} else if m.Content == "go" && saperStarted {
		//log.Println("kill")
		saperChan <- "kill"
		go Start(s, m, saperChan,&saperStarted)
	} else if saperStarted {
		matched, err := regexp.MatchString(`^\d(\s|[a-z])\d$`, m.Content)
		if err != nil {
			log.Println(err)
		}
		if matched {
			saperChan <- m.Content
		}
	}
}