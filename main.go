package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	cobra "github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var kbotVersion string = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "kbot",
	Version: kbotVersion,
	Short:   "kbot is a very simple telebot",
}

var startkbotCmd = &cobra.Command{
	Use:   "start",
	Short: "This is start command",
	Long:  "Use 'start' to start kbot",
	Run: func(cmd *cobra.Command, args []string) {

		pref := telebot.Settings{
			URL:         "",
			Token:       os.Getenv("TELE_TOKEN"),
			Updates:     0,
			Poller:      &telebot.LongPoller{Timeout: 5 * time.Second},
			Synchronous: false,
			Verbose:     false,
			ParseMode:   "",
			OnError: func(error, telebot.Context) {
			},
			Client:  &http.Client{},
			Offline: false,
		}

		kbot, err := telebot.NewBot(pref)
		if err != nil {
			log.Fatal("Unknown error! Please check system variable TELE_TOKEN value")
			return
		}

		fmt.Printf("Telegram bot 'kbot' started!\n")

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			var reply error
			// m.Send("I am alive!")
			msg := m.Text()
			log.Println("Someone entered: " + msg)
			switch msg {
			case "/hello":
				reply = m.Send("Hello")
			case "/help":
				reply = m.Send("This is simple bot on Go.\nOnly /hello and /version are available for now")
			case "/version":
				reply = m.Send("Version:" + kbotVersion)
			default:
				reply = m.Send("Do not know what to answer. Please try /help for help")
			}

			return reply
		})

		kbot.Start()
	},
}

var helpkbootCmd = &cobra.Command{
	Use:   "help",
	Short: "This is help command",
	Long:  "Use 'help' to print the available help of the kbot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Nowhere to find any help. Sorry...\n")
	},
}

var versionkbotCmd = &cobra.Command{
	Use:   "version",
	Short: "This is version command",
	Long:  "Use 'version' to print the version of the kbot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("The current version is %s\n", kbotVersion)
	},
}

func main() {

	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(startkbotCmd)
	rootCmd.SetHelpCommand(helpkbootCmd)
	rootCmd.AddCommand(versionkbotCmd)
}
