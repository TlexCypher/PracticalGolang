package main

import (
	"os"
	"strings"

	"github.com/alecthomas/kingpin/v2"
)

var (
	app      = kingpin.New("chat", "A command-line chat application.")
	debug    = app.Flag("debug", "Enable debug mode.").Bool()
	serverIP = app.Flag("server", "Server address").Default("127.0.0.1").IP()

	register     = app.Command("register", "Register a new user.")
	registerNick = register.Arg("nick", "Nickname for user.").Required().String()
	registerName = register.Arg("name", "Name of user.").Required().String()

	post        = app.Command("post", "Post a message to channel.")
	postImage   = post.Flag("image", "Image to post").File()
	postChannel = post.Arg("channel", "Channel to post to.").Required().String()
	postText    = post.Arg("text", "Text to post.").Required().Strings()
)

func Kingpin() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case register.FullCommand():
		println(*registerNick)

	case post.FullCommand():
		if *postImage != nil {
		}
		text := strings.Join(*postText, " ")
		println("Post:", text)
	}
}
