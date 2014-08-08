package main
import (
		"github.com/thoj/go-ircevent"
		"fmt"
)
var roomName = "#my-bot-test"
import s "strings"

func main() {
	con := irc.IRC("mystbot","mystbot")
	con.Password = "oauth:fzfjkum1omh806t4c0f7q7t5a7y6qum"
	err := con.Connect("irc.twitch.tv:6667")
	if err != nil {
		fmt.Println("Failed Connecting")
		return
	}
	con.AddCallback("001", func (e *irc.Event) {
        con.Join(roomName)
    })
	con.AddCallback("JOIN", func (e *irc.Event) {
        con.Privmsg(roomName, "Hello! I am Mystbot, a very poorly written attempt at being a bot...")
    })
    con.AddCallback("PRIVMSG", func (e *irc.Event) {
    	if s.Contains(e.Message(), "!mystbot")
    		commandPieces = s.Split(e.Message(), " ")
    		
    	}
        con.Privmsg(roomName, e.Message())
    })
    con.Loop()
}