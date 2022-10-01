package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {
	godotenv.Load(".env")

	SLACKER_BOT := os.Getenv("SLACKER_BOT")
	// SLACKER_APP := os.Getenv("SLACKER_APP")
	fmt.Printf("its my key %s", SLACKER_BOT)
	bot := slacker.NewClient("xoxb-4055114362742-4091177104134-fEFUk1so0QZrCFREzJoLL51w", "xapp-1-A042SQ5SBNH-4091173300310-0b5700de4270974482b5afca60bed23e6a22d26bf86c66e97d0f49a6fa3a74a8")

	bot.Command("my job is <year>", &slacker.CommandDefinition{
		Description: "i am a lonely age calcultor",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			IntConv, err := strconv.Atoi(request.Param("year"))
			fmt.Print(err)
			age := 2022 - IntConv
			var r string
			switch {
			case age < 18:
				r = fmt.Sprintf("u are loli u pussy fucking %d y/o", age)
			default:
				r = fmt.Sprintf("u are fucking old u pussy fucking %d y/o", age)

			}
			response.Reply(r)

		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		fmt.Print(err)
	}

}
