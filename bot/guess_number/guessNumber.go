package guessnumber

import (
	crypto "crypto/rand"
	"math/big"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var started bool = false 
var gn int

func Start(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "go"{
		s.ChannelMessageSend(m.ChannelID, "я загадал число от 0 до 9(все включительно), попробуй угадать!")
		started = true
		gn = int(NewCryptoRand(10))
	}else if started {
		sn := strconv.Itoa(gn)
		if sn == m.Content {
			s.ChannelMessageSend(m.ChannelID, "ты угадал, "+sn+" это загаданное число!!!")
		} else{
			s.ChannelMessageSend(m.ChannelID, "ты не угадал, "+sn+" было загаданным числом(")
		}
		started = false
	}
}

func NewCryptoRand(i int) int64 {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(int64(i)))
	if err != nil {
			panic(err)
	}
	return safeNum.Int64()
}