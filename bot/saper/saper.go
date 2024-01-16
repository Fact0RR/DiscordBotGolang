package saper

import (
	crypto "crypto/rand"
	_"log"
	"math/big"
	"strconv"

	"github.com/bwmarrin/discordgo"
)



type UpField struct{
	cells [10][10]int
}

func (uf *UpField) UpFieldFeel(){
	for i:= 0; i < len(uf.cells); i++ {
		for j := 0; j < len(uf.cells[i]); j++ {
			uf.cells[i][j] = -1;
		}
	}
}

type DownField struct{
	cells [10][10]int
	bombs int
}

func (df *DownField) DownFieldFeel(){
	for i :=  0; i < len(df.cells); i++ {
		for j := 0; j < len(df.cells[i]); j++{
			
			r:= NewCryptoRand(10)
			//log.Println(r)
			if r == 9 {
				df.cells[i][j] = 9
				df.bombs++
			}
		}
	}

	for i := 0; i < len(df.cells); i++ {
		for j := 0; j < len(df.cells[i]); j++ {
			bombsAround := 0
			if (df.cells[i][j] != 9) {
				if (i + 1 < len(df.cells)) {
					if (df.cells[i + 1][j] == 9) {
						bombsAround++
					}
				}
				if (j + 1 < len(df.cells[i])) {
					if (df.cells[i][j + 1] == 9) {
						bombsAround++
					}
				}
				if (i - 1 >= 0) {
					if (df.cells[i - 1][j] == 9) {
						bombsAround++
					}
				}
				if (j - 1 >= 0) {
					if (df.cells[i][j - 1] == 9) {
						bombsAround++
					}
				}
				if (i + 1 < len(df.cells) && j + 1 < len(df.cells[i])) {
					if (df.cells[i + 1][j + 1] == 9) {
						bombsAround++
					}
				}
				if (i + 1 < len(df.cells) && j - 1 >= 0) {
					if (df.cells[i + 1][j - 1] == 9) {
						bombsAround++
					}
				}
				if (i - 1 >= 0 && j + 1 < len(df.cells[i])) {
					if (df.cells[i - 1][j + 1] == 9) {
						bombsAround++
					}
				}
				if (i - 1 >= 0 && j - 1 >= 0) {
					if (df.cells[i - 1][j - 1] == 9) {
						bombsAround++
					}
				}
				df.cells[i][j] = bombsAround
			}

		}
	}
}

func NewCryptoRand(i int) int64 {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(int64(i)))
	if err != nil {
			panic(err)
	}
	return safeNum.Int64()
}

func Start(s *discordgo.Session,m *discordgo.MessageCreate, c chan string,saperStarted *bool){
	
	var upField UpField
	var downField DownField

	upField.UpFieldFeel()
	downField.DownFieldFeel()

	s.ChannelMessageSend(m.ChannelID, "Бомб на поле: "+ strconv.Itoa(downField.bombs)+"!")
	s.ChannelMessageSend(m.ChannelID, "`"+ showUpField(&upField)+"`")

	for openCells(&upField) != downField.bombs {

		// synchronized (SanitarBot.saperThread){
		// 	SanitarBot.saperThread.wait();
		// }
		chat :=  <-c 
			//log.Println(chat)
			if (chat == "kill"){
				s.ChannelMessageSend(m.ChannelID, "kill")
				return
			}
		line,_ := strconv.Atoi(chat[0:1])
		column,_ := strconv.Atoi(chat[2:3])
		
		if (chat[1:2]!="f"){
				//int line = setLine();
				//int column = setColumn();
				
				
			if(upField.cells[line][column] == -2){
				continue;
			}
			//pw.print(line);
			//pw.print(" " + column + "\n");
			//move++;
			if (downField.cells[line][column] == 9) {
				break;
			} else if (downField.cells[line][column] == 0) {

				upField.cells[line][column] = downField.cells[line][column];
				clearTheField(&upField, &downField);
				//System.out.println(showUpField(upField));
				s.ChannelMessageSend(m.ChannelID, "`"+ showUpField(&upField)+"`")
			} else {
				upField.cells[line][column] = downField.cells[line][column];
				//System.out.println(showUpField(upField));
				s.ChannelMessageSend(m.ChannelID, "`"+ showUpField(&upField)+"`")
			}
		}else {
			//log.Println("Ждем ответ из чата")
			chat :=  <-c 
			//log.Println(chat)
			if (chat == "kill"){
				s.ChannelMessageSend(m.ChannelID, "kill")
				return
			}
			 
			line,_ := strconv.Atoi(chat[0:1])
			column,_ := strconv.Atoi(chat[2:3])
			if (upField.cells[line][column] == -1){
				upField.cells[line][column] = -2;
				s.ChannelMessageSend(m.ChannelID, "`" + showUpField(&upField) + "`")
			}else if(upField.cells[line][column] == -2){
				upField.cells[line][column] = -1;
				s.ChannelMessageSend(m.ChannelID, "`" + showUpField(&upField) + "`")
			}
		}
	}
	if (openCells(&upField) == downField.bombs) {
		//вы победили
		//System.out.println("Вы победили!!!");
		//pw.print("Всего ходов:"+move+"\n");
		//pw.print("Результат: победа\n");
		//SanitarBot.eventForSaper.getChannel().sendMessage("Вы победили!!!").queue();
		//System.out.println(showDownField(downField));
		//SanitarBot.eventForSaper.getChannel().sendMessage("`"+ showDownField(downField)+"`").queue();
		s.ChannelMessageSend(m.ChannelID,"Вы победили!!!")
		s.ChannelMessageSend(m.ChannelID,"`"+ showDownField(&downField)+"`")

	} else {
		//System.out.println("Вы проиграли");
		//pw.print("Всего ходов:"+move+"\n");
		//pw.print("Результат: проигрыш\n");
		//SanitarBot.eventForSaper.getChannel().sendMessage("Вы проиграли").queue();
		//System.out.println(showDownField(downField));
		//SanitarBot.eventForSaper.getChannel().sendMessage("`"+ showDownField(downField)+"`").queue();
		//log.Println(openCells(&upField),downField.bombs)
		s.ChannelMessageSend(m.ChannelID,"Вы проиграли")
		s.ChannelMessageSend(m.ChannelID,"`"+ showDownField(&downField)+"`")
	}
	//pw.close();
	//SanitarBot.eventForSaper.getChannel().sendMessage(showStatistic()).queue();
	*saperStarted = false
}


func showUpField(upField *UpField) string{
	mainString := "   ";
	for i:= 0; i < len(upField.cells); i++{

		mainString = mainString + " " + strconv.Itoa(i);
	}
	mainString = mainString + "\n   ";
	for i := 0; i < len(upField.cells); i++{

		mainString = mainString + " -";
	}
	mainString = mainString + "\n";
	for i:= 0; i < len(upField.cells); i++ {

		for j := 0; j < len(upField.cells[i]); j++{
			if j == 0 {
				if(upField.cells[i][j]==-1) {
					mainString = mainString + strconv.Itoa(i) + "|  #";
				}else if(upField.cells[i][j]==0) {
					mainString = mainString + strconv.Itoa(i) + "|   ";
				}else if(upField.cells[i][j]==-2) {
					mainString = mainString + strconv.Itoa(i) + "|  >";
				}else{
					mainString = mainString + strconv.Itoa(i) + "|  " + strconv.Itoa(upField.cells[i][j])+"";
				}
			}else if(j == len(upField.cells[i])-1){
				if(upField.cells[i][j]==-1) {
					mainString =mainString + " #  |" + strconv.Itoa(i);
				}else if(upField.cells[i][j]==0) {
					mainString =mainString + "    |" + strconv.Itoa(i);
				}else if(upField.cells[i][j]==-2) {
					mainString =mainString + " >  |" + strconv.Itoa(i);
				}else{
					mainString = mainString  + " " + strconv.Itoa(upField.cells[i][j])+"  |"+ strconv.Itoa(i);
				}
			}else{
				if(upField.cells[i][j]==-1) {
					mainString =mainString + " #";
				}else if(upField.cells[i][j]==0) {
					mainString =mainString + "  ";
				}else if(upField.cells[i][j]==-2) {
					mainString =mainString + " >";
				}else{
					mainString = mainString  + " " + strconv.Itoa(upField.cells[i][j]);
				}
			}
		}
		mainString = mainString + "\n";
	}
	mainString =mainString + "   ";

	for i:= 0; i < len(upField.cells); i++{

		mainString = mainString + " -";
	}
	mainString = mainString+ "\n";
	mainString =mainString + "   ";
	for i := 0; i < len(upField.cells); i++{
		mainString = mainString + " " + strconv.Itoa(i);
	}
	return mainString;
}

func clearTheField(upField *UpField,downField *DownField){
	for k := 0; k < len(upField.cells); k++ {
		for i := 0; i < len(upField.cells); i++ {
			for j := 0; j < len(upField.cells[i]); j++ {
				if (upField.cells[i][j] == 0) {
					if (i + 1 < len(upField.cells)) {
						upField.cells[i + 1][j] = downField.cells[i + 1][j];
					}
					if (j + 1 < len(upField.cells[i])) {
						upField.cells[i][j + 1] = downField.cells[i][j + 1];
					}
					if (i - 1 >= 0) {
						upField.cells[i - 1][j] = downField.cells[i - 1][j];
					}
					if (j - 1 >= 0) {
						upField.cells[i][j - 1] = downField.cells[i][j - 1];
					}
					if (i + 1 < len(upField.cells) && j + 1 < len(upField.cells[i])) {
						upField.cells[i + 1][j + 1] = downField.cells[i + 1][j + 1];
					}
					if (i + 1 < len(upField.cells) && j - 1 >= 0) {
						upField.cells[i + 1][j - 1] = downField.cells[i + 1][j - 1];
					}
					if (i - 1 >= 0 && j + 1 < len(upField.cells[i])) {
						upField.cells[i - 1][j + 1] = downField.cells[i - 1][j + 1];
						//if (upField.cells[i - 1][j + 1] == 0);
					}
					if (i - 1 >= 0 && j - 1 >= 0) {
						upField.cells[i - 1][j - 1] = downField.cells[i - 1][j - 1];

					}
				}

			}
		}
		for i:= len(upField.cells)-1; i >=0 ; i-- {
			for j := len(upField.cells[i])-1; j >=0 ; j-- {
				if (upField.cells[i][j] == 0) {
					if (i + 1 < len(upField.cells)) {
						upField.cells[i + 1][j] = downField.cells[i + 1][j];
					}
					if (j + 1 < len(upField.cells[i])) {
						upField.cells[i][j + 1] = downField.cells[i][j + 1];
					}
					if (i - 1 >= 0) {
						upField.cells[i - 1][j] = downField.cells[i - 1][j];
					}
					if (j - 1 >= 0) {
						upField.cells[i][j - 1] = downField.cells[i][j - 1];
					}
					if (i + 1 < len(upField.cells) && j + 1 < len(upField.cells[i])) {
						upField.cells[i + 1][j + 1] = downField.cells[i + 1][j + 1];
					}
					if (i + 1 < len(upField.cells) && j - 1 >= 0) {
						upField.cells[i + 1][j - 1] = downField.cells[i + 1][j - 1];
					}
					if (i - 1 >= 0 && j + 1 < len(upField.cells[i])) {
						upField.cells[i - 1][j + 1] = downField.cells[i - 1][j + 1];
						//if (upField.cells[i - 1][j + 1] == 0);
					}
					if (i - 1 >= 0 && j - 1 >= 0) {
						upField.cells[i - 1][j - 1] = downField.cells[i - 1][j - 1];

					}
				}

			}
		}
	}
}

func showDownField(upField *DownField)string{
	mainString := "   ";
        for i := 0; i < len(upField.cells); i++{

            mainString = mainString + " " + strconv.Itoa(i);
        }
        mainString = mainString + "\n   ";
        for i := 0; i < len(upField.cells); i++{

            mainString = mainString + " -";
        }
        mainString = mainString + "\n";
        for i := 0; i < len(upField.cells); i++ {

            for j := 0; j < len(upField.cells[i]); j++{
                if(j == 0){
                    if(upField.cells[i][j]== 9) {
                        mainString = mainString + strconv.Itoa(i) + "|  *";
                    }else if(upField.cells[i][j]==0) {
                        mainString = mainString + strconv.Itoa(i) + "|   ";
                    }else{
                        mainString = mainString + strconv.Itoa(i) + "|  " + strconv.Itoa(upField.cells[i][j])+"";
                    }
                }else if(j == len(upField.cells[i])-1){
                    if(upField.cells[i][j]==9) {
                        mainString =mainString + " *  |" + strconv.Itoa(i);
                    }else if(upField.cells[i][j]==0) {
                        mainString =mainString + "    |" + strconv.Itoa(i);
                    }else{
                        mainString = mainString  + " " + strconv.Itoa(upField.cells[i][j])+"  |"+ strconv.Itoa(i);
                    }
                }else{
                    if(upField.cells[i][j]==9) {
                        mainString =mainString + " *";
                    }else if(upField.cells[i][j]==0) {
                        mainString =mainString + "  ";
                    }else
                    {
                        mainString = mainString  + " " + strconv.Itoa(upField.cells[i][j]);
                    }
                }
            }
            mainString = mainString + "\n";
        }
        mainString =mainString + "   ";

        for i := 0; i < len(upField.cells); i++{

            mainString = mainString + " -";
        }
        mainString = mainString+ "\n";
        mainString =mainString + "   ";
        for i := 0; i < len(upField.cells); i++{

            mainString = mainString + " " + strconv.Itoa(i);
        }
        return mainString;
}

func openCells(upField *UpField) int{
	oc := 0;
	for i:=0;i<len(upField.cells);i++{
		for j:=0;j<len(upField.cells[i]);j++{
			if((upField.cells[i][j]==-1)||(upField.cells[i][j]==-2)){
				oc++;
			}
		}
	}

	return oc;
}