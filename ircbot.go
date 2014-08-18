package main
import (
		"github.com/thoj/go-ircevent"
		"fmt"
        //"os"
        s "strings"
)
var roomName = "#phantasaia"
var nick = "mystbot"


func main() {
	con := irc.IRC(nick, nick)
	con.Password = "<password>"
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
        if e.Nick == nick{
            con.Privmsg(roomName, "Hello! I am Mystbot, a very poorly written attempt at being a bot...")
        }
        
        // Stub out timed and linecount messages here
        
    })
    con.AddCallback("PRIVMSG", func (e *irc.Event) {
    	if s.HasPrefix(e.Message(), "!mystbot"){
    		var commandPieces []string = s.Split(e.Message(), " ")
            var command string = commandPieces[1]
    		if s.ToLower(command) == "echo"{
                con.Privmsg(roomName, s.Join(commandPieces[2:len(commandPieces)], " ") )
            }
            if s.ToLower(command) == "leave"{
                var whoInfo string = con.Who(e.Nick)
                fmt.Printf(whoInfo)
                con.Privmsg(roomName, "It was a pleasure hanging out with you today! Bye!")
                //con.Disconnect()
                //os.Exit(0)
            }
    	}
    })
    con.Loop()
}