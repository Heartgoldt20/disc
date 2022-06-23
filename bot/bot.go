package bot

import (
	"fmt"

	"github.com/akhil/discord-ping/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

//bot start
func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is runnendeee broer!")
}

//dit is een fucn genaamd messagehandler waarom is deze functie er? nou om ervoor te zorgen dat de bot een mesage stuurt moet er een functie voor toegevoegd worden.
//en dat is deze messageHandler je ziet dat er een if statement is waarom? nou als messagecontent (wat iedmand stuurt) gelijk is aan !ping of !danny dan reageerd hij met ChannelMessageSend.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "!ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "ping mij niet neger")
	}

	if m.Content == "!danny" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "danny is gay lol")
	}

}
