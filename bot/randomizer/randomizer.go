package randomizer

import (
	crypto "crypto/rand"
	"math/big"

	"github.com/bwmarrin/discordgo"
)


func Get_Random_Agent(s *discordgo.Session,m *discordgo.MessageCreate,agents []string) {
	if m.Content == "go"{
		r := NewCryptoRand(len(agents))
		s.ChannelMessageSend(m.ChannelID, agents[r])
	}
}

func NewCryptoRand(i int) int64 {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(int64(i)))
	if err != nil {
			panic(err)
	}
	return safeNum.Int64()
}