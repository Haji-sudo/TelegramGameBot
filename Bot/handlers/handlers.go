package handlers

import (
	c "dogegambling/config"
	"fmt"

	"strconv"

	b "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func (h Handler) Init() {
	h.UserInit()
	MenuInint()
	Admin := c.Bot.Group()
	Admin.Use(middleware.Whitelist(c.Admins...))
	c.Bot.Handle("/start", func(ctx b.Context) error {
		UserID := ctx.Chat().ID
		if !UserExist(UserID) {
			if !UserExistInDB(UserID) {
				CreateUserInAllDB(UserID)
				if payload := ctx.Message().Payload; payload != "" {
					inviterID, err := strconv.ParseInt(payload, 10, 64)
					if err == nil {
						if UserID != inviterID { //Check Not Your Self
							if UserExistInDB(inviterID) {
								userdb := GetUserFromDB(inviterID)
								userdb.AddReferral()
								inviter := b.ChatID(inviterID)
								c.Bot.Send(inviter, "you got a new referral")
							}
						}
					}
				}
			}
			CreateInRedis(UserID)
		}

		user := GetUser(UserID)
		if user.NotSpam() && !user.IsLock {
			user.UpdateTime()
			user.lock()
			defer user.unlock()
			if user.Location != "main" {
				user.ChangeLocation("main")
			}
			link := "tg://user?id=" + strconv.FormatInt(ctx.Chat().ID, 10)
			return ctx.Send(START(ctx.Chat().FirstName, link), MainMenu, b.ModeMarkdown)
		}

		return nil
	})

	c.Bot.Handle(b.OnText, func(ctx b.Context) error {
		input := ctx.Text()
		UserID := ctx.Chat().ID

		if UserExist(UserID) {
			user := GetUser(UserID)

			if !user.IsLock && user.NotSpam() {
				user.UpdateTime()
				user.lock()
				defer user.unlock()
				if user.Location != "main" {
					if input == BtnHome.Text {
						user.ChangeLocation("main")
						return ctx.Send("Home", MainMenu)

					}

				} else if user.Location == "main" {

					if input == BtnGames.Text {
						user.ChangeLocation("game")
						return ctx.Send("GameBoard", GameMenu)

					} else if input == BtnReferrals.Text {
						link := fmt.Sprintf("t.me/%v?start=%v", c.BotUsername, UserID)
						return ctx.Send(link)
					} else if input == BtnAccount.Text {
						userdata := GetUserFromDB(UserID)
						link := "tg://user?id=" + strconv.FormatInt(UserID, 10)
						return ctx.Send(ACCOUNT(ctx.Chat().FirstName, link, userdata.Balance, userdata.Referrals, userdata.Warn, userdata.Wallet), b.ModeMarkdown)
					}

				} else if user.Location == "game" {
					if input == BtnDice.Text {
						user.ChangeLocation("dice")
						return ctx.Send("hi", DiceMenu)

					}
				}

			} else {
				ctx.Send("Don't Spam Bitch")
			}
		} else {
			ctx.Send("Please /start Again")
		}

		return nil
	})

}
