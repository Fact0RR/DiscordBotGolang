package bot

import (
	guessnumber "bot/bot/guess_number"
	"bot/bot/photo2ascii"
	"bot/bot/randomizer"
	"bot/structs"

	"log"

	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)




func Run(conf structs.Conf,agents []string) {

	sess, err:= discordgo.New(conf.Type+" "+conf.Token)
        if err != nil{
                log.Fatal(err)
        }

        sess.AddHandler(func (s *discordgo.Session, m *discordgo.MessageCreate)  {

			//log.Println(m.ChannelID)
			if m.Author.ID != "930498017948209293"{
				if m.ChannelID == "995790861126344926" {
					randomizer.Get_Random_Agent(s,m,agents)
				}else if m.ChannelID == "1196503435676221571" {
					photo2ascii.ConvertPhoto2Ascii(s,m)
				}else if m.ChannelID == "994743377436348459"{
					guessnumber.Start(s,m)
				}
			}
        })

        err = sess.Open()
        if err!= nil{
                log.Fatal(err)
        }
        defer sess.Close()
		log.Println("ingnited")
        c:= make(chan os.Signal,1)
		signal.Notify(c,os.Interrupt)
		<-c

}