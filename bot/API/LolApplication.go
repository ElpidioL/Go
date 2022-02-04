package API

import (
	PSB "bots/GOing/PostgreDB"
	"bots/GOing/modules"
	"bots/GOing/options"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

type MatchL struct {
	GameId       int    `json:"gameId"`
	GameMode     string `json:"gameMode"`
	Participants []struct {
		SpellId      int    `json:"spell1Id"`
		Spell2ID     int    `json:"spell2Id"`
		ChampionId   int    `json:"championId"`
		SummonerName string `json:"summonerName"`
		SummonerId   string `json:"summonerId"`
	}
}

type UserL struct {
	Id        string `json:"id"`
	AccountId string `json:"accountId"`
	Puuid     string `json:"puuid"`
	Name      string `json:"name"`
}

func NotifyLol(Discord *discordgo.Session) {

	//here i'll be handling the Lol API for now, since i need to update with a different rate than discord handler
	//this will be here.
	for true {
		allUsers := PSB.GetAllUsers()
		for index, _ := range allUsers {
			time.Sleep(500 * time.Millisecond)

			matchInfo := GetMatchLol(allUsers[index].Id)
			if matchInfo.GameId > 0 {
				alert := PSB.GetMatchDB(matchInfo.GameId, matchInfo.Participants[0].SummonerName)

				if alert {
					PSB.MatchRegister(matchInfo.GameId, allUsers[index].Name, GetChampName(matchInfo.Participants[0].ChampionId))
					discords := PSB.GetAllDiscords(allUsers[index].Puuid)
					message := fmt.Sprintf("O Crime foi iniciado, %s começou a gameplay criminosa jogando de %s em uma partida %s se preparem para o choro", matchInfo.Participants[0].SummonerName, GetChampName(matchInfo.Participants[0].ChampionId), matchInfo.GameMode)
					for index, _ := range discords {
						modules.SendMessage(Discord, discords[index].Discords_text, message, false)
					}
					//functions.PlayHorn(Discord, options.Guild, modules.FindVoiceChannel(Discord, options.Guild, options.Player))
					// i was playing a horn on a previous version, but since i can register a lot of players now, there is no way to keep track of the player to disturb
				}

			}
		}
		//sleep so i don't waste much process power for nothing (probably there is a better way to do that...)
		time.Sleep(60 * time.Second)
	}
}

func GetInfoApi(url string, bodyChan chan string) {
	resp, err := http.Get(url)
	if err != nil {
		//return "", "", ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var respBody string
	for _, value := range body { //here i convert runes to string that way because i don't know any better way.
		respBody += string(value) //dunno why its coming in runes in first place
	}
	//end

	bodyChan <- respBody
}

func GetUserLol(userNameLol string, guildId, discordChannel string) string {
	bodyChan := make(chan string)
	url := "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-name/" + userNameLol + "?api_key=" + options.LolKey
	go GetInfoApi(url, bodyChan)
	respBody := <-bodyChan

	user := UserL{}
	err := json.Unmarshal([]byte(respBody), &user)
	if err != nil {
		fmt.Println(err)
	}
	if len(user.Puuid) == 0 {
		return fmt.Sprintf("Player %s does not exist in BR server", userNameLol)
	}

	msg := PSB.UserRegister(user.Name, user.Id, user.Puuid, user.AccountId, guildId, discordChannel)

	return msg
}

func GetMatchLol(summonerID string) MatchL { //ChanSummonerName chan string, ChanGameMode chan string, ChanChampion chan string) {
	bodyChan := make(chan string)
	url := "https://br1.api.riotgames.com/lol/spectator/v4/active-games/by-summoner/" + summonerID + "?api_key=" + options.LolKey
	go GetInfoApi(url, bodyChan)
	respBody := <-bodyChan
	match := MatchL{}
	if len(respBody) <= 1 {
		return match

	} else {
		err := json.Unmarshal([]byte(respBody), &match)
		if err != nil {
			fmt.Println(err)
		}
		for index, _ := range match.Participants {
			if match.Participants[index].SummonerId == summonerID {
				match.Participants[0] = match.Participants[index]
				return match
			}
		}
	}
	return match
}

func GetChampName(id int) string {

	switch id {
	case 266:
		return "Aatrox"
	case 412:
		return "Thresh"
	case 23:
		return "Tryndamere"
	case 79:
		return "Gragas"
	case 69:
		return "Cassiopeia"
	case 136:
		return "Aurelion Sol"
	case 13:
		return "Ryze"
	case 78:
		return "Poppy"
	case 14:
		return "Sion"
	case 1:
		return "Annie"
	case 202:
		return "Jhin"
	case 43:
		return "Karma"
	case 111:
		return "Nautilus"
	case 240:
		return "Kled"
	case 99:
		return "Lux"
	case 103:
		return "Ahri"
	case 2:
		return "Olaf"
	case 112:
		return "Viktor"
	case 34:
		return "Anivia"
	case 27:
		return "Singed"
	case 86:
		return "Garen"
	case 127:
		return "Lissandra"
	case 57:
		return "Maokai"
	case 25:
		return "Morgana"
	case 28:
		return "Evelynn"
	case 105:
		return "Fizz"
	case 74:
		return "Heimerdinger"
	case 238:
		return "Zed"
	case 68:
		return "Rumble"
	case 82:
		return "Mordekaiser"
	case 37:
		return "Sona"
	case 96:
		return "Kog'Maw"
	case 55:
		return "Katarina"
	case 117:
		return "Lulu"
	case 22:
		return "Ashe"
	case 30:
		return "Karthus"
	case 12:
		return "Alistar"
	case 122:
		return "Darius"
	case 67:
		return "Vayne"
	case 110:
		return "Varus"
	case 77:
		return "Udyr"
	case 89:
		return "Leona"
	case 126:
		return "Jayce"
	case 134:
		return "Syndra"
	case 80:
		return "Pantheon"
	case 92:
		return "Riven"
	case 121:
		return "Kha'Zix"
	case 42:
		return "Corki"
	case 268:
		return "Azir"
	case 51:
		return "Caitlyn"
	case 76:
		return "Nidalee"
	case 85:
		return "Kennen"
	case 3:
		return "Galio"
	case 45:
		return "Veigar"
	case 432:
		return "Bard"
	case 150:
		return "Gnar"
	case 90:
		return "Malzahar"
	case 104:
		return "Graves"
	case 254:
		return "Vi"
	case 10:
		return "Kayle"
	case 39:
		return "Irelia"
	case 64:
		return "Lee Sin"
	case 420:
		return "Illaoi"
	case 60:
		return "Elise"
	case 106:
		return "Volibear"
	case 20:
		return "Nunu"
	case 4:
		return "Twisted Fate"
	case 24:
		return "Jax"
	case 102:
		return "Shyvana"
	case 429:
		return "Kalista"
	case 36:
		return "Dr. Mundo"
	case 427:
		return "Ivern"
	case 131:
		return "Diana"
	case 223:
		return "Tahm Kench"
	case 63:
		return "Brand"
	case 113:
		return "Sejuani"
	case 8:
		return "Vladimir"
	case 154:
		return "Zac"
	case 421:
		return "Rek'Sai"
	case 133:
		return "Quinn"
	case 84:
		return "Akali"
	case 163:
		return "Taliyah"
	case 18:
		return "Tristana"
	case 120:
		return "Hecarim"
	case 15:
		return "Sivir"
	case 236:
		return "Lucian"
	case 107:
		return "Rengar"
	case 19:
		return "Warwick"
	case 72:
		return "Skarner"
	case 54:
		return "Malphite"
	case 157:
		return "Yasuo"
	case 101:
		return "Xerath"
	case 17:
		return "Teemo"
	case 75:
		return "Nasus"
	case 58:
		return "Renekton"
	case 119:
		return "Draven"
	case 35:
		return "Shaco"
	case 50:
		return "Swain"
	case 91:
		return "Talon"
	case 40:
		return "Janna"
	case 115:
		return "Ziggs"
	case 245:
		return "Ekko"
	case 61:
		return "Orianna"
	case 114:
		return "Fiora"
	case 9:
		return "Fiddlesticks"
	case 31:
		return "Cho'Gath"
	case 33:
		return "Rammus"
	case 7:
		return "LeBlanc"
	case 16:
		return "Soraka"
	case 26:
		return "Zilean"
	case 56:
		return "Nocturne"
	case 222:
		return "Jinx"
	case 83:
		return "Yorick"
	case 6:
		return "Urgot"
	case 203:
		return "Kindred"
	case 21:
		return "Miss Fortune"
	case 62:
		return "Wukong"
	case 53:
		return "Blitzcrank"
	case 98:
		return "Shen"
	case 201:
		return "Braum"
	case 5:
		return "Xin Zhao"
	case 29:
		return "Twitch"
	case 11:
		return "Master Yi"
	case 44:
		return "Taric"
	case 32:
		return "Amumu"
	case 41:
		return "Gangplank"
	case 48:
		return "Trundle"
	case 38:
		return "Kassadin"
	case 161:
		return "Vel'Koz"
	case 143:
		return "Zyra"
	case 267:
		return "Nami"
	case 59:
		return "Jarvan IV"
	case 81:
		return "Ezreal"
	case 164:
		return "Camille"
	case 498:
		return "Xayah"
	case 497:
		return "Rakan"
	case 246:
		return "Qiyana"
	case 350:
		return "Yuumi"
	case 517:
		return "Sylas"
	case 555:
		return "Pyke"
	case 145:
		return "Kai'Sa"
	case 142:
		return "Zoe"
	case 516:
		return "Ornn"
	case 141:
		return "Kayn"
	case 711:
		return "Vex"
	case 166:
		return "Akshan"
	case 887:
		return "Gwen"
	case 526:
		return "Rell"
	case 360:
		return "Samira"
	case 777:
		return "Yone"
	case 876:
		return "Lillia"
	case 875:
		return "Sett"
	case 523:
		return "Aphelios"
	case 235:
		return "Senna"
	case 221:
		return "Zeri"

	}
	return "deu ruim no codigo do ctr c ctr v, me notifica no privado seu champ"
}
