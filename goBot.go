package gobot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + GoDotEnvVariable("BOTTOKEN"))

	if err != nil {
		fmt.Println("failed run a bot,", err)
		return
	}

	// botGO := handler.NewBotGo()
	// discord.AddHandler(botGO.CmdHandle)

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	stopBot := make(chan os.Signal, 1)

	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-stopBot

	discord.Close()
}
