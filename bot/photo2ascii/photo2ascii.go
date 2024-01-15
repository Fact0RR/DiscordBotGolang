package photo2ascii

import (
	"bytes"
	"image"
	"io"
	"log"
	"net/http"
	

	"github.com/bwmarrin/discordgo"
	"github.com/qeesung/image2ascii/convert"
)

func ConvertPhoto2Ascii(s *discordgo.Session, m *discordgo.MessageCreate) {
	//log.Println(m.Attachments[0].ContentType)
	if len(m.Attachments)<1{
		s.ChannelMessageSend(m.ChannelID, "Прикрепите картинку в формате .png или .jpg")
		return
	}
	if (m.Attachments[0].ContentType != "image/png") && (m.Attachments[0].ContentType != "image/jpeg"){
		s.ChannelMessageSend(m.ChannelID, "Неверный формат (ваш -"+m.Attachments[0].ContentType+" нужен - .png или .jpg)")
		return
	}

	response, err := http.Get(m.Attachments[0].URL)
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, "Ошибка подключения через URL(проверьте логи)")
		return
	}
	defer response.Body.Close()

	b_im, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, "Ошибка преобразования в массив байтов(проверьте логи)")
		return
	}
	final_img, _, err := image.Decode(bytes.NewReader(b_im))
    if err != nil {
        log.Fatalln(err)
    }
	convertOptions := convert.DefaultOptions
	convertOptions.FitScreen = false
	convertOptions.StretchedScreen = false
	convertOptions.Colored = false
	convertOptions.FixedWidth = 69
	convertOptions.FixedHeight = 28

	// Create the image converter
	converter := convert.NewImageConverter()
	//log.Print(converter.Image2ASCIIString(final_img, &convertOptions))
	//s.ChannelMessageSend(m.ChannelID, "`"+converter.Image2ASCIIString(final_img, &convertOptions))
	result := converter.Image2ASCIIString(final_img, &convertOptions)
	s.ChannelMessageSend(m.ChannelID,"`"+result+"`")
	
}