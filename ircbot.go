package main
import (
		"github.com/thoj/go-ircevent"
		"fmt"
        "os"
        s "strings"
)
var roomName = "#phantasaia"


func main() {
	con := irc.IRC("mystbot","mystbot")
	con.Password = "oauth:fzfjkum1omh806t4c0f7q7t5a7y6qum"
	//err := con.Connect("irc.twitch.tv:6667")
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
        // Stub out timed and linecount messages here

        // stub out oping self

        //stub out oping users
        // irc.Mode(nickname, "+o")
    })
    con.AddCallback("PRIVMSG", func (e *irc.Event) {
    	if s.HasPrefix(e.Message(), "!mystbot"){
    		var commandPieces []string = s.Split(e.Message(), " ")
            var command string = commandPieces[1]
    		if s.ToLower(command) == "echo"{
                con.Privmsg(roomName, s.Join(commandPieces[2:len(commandPieces)], " ") )
            }
            if s.ToLower(command) == "leave"{
                con.Privmsg(roomName, "It was a pleasure hanging out with you today! Bye!")
                con.Quit()
                os.Exit(0)
            }
    	}
    })
    con.Loop()
}