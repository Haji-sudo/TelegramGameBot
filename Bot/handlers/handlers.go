package handlers

import (
	c "dogegambling/config"
	"fmt"
	"strings"

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
	c.Bot.Handle(&BtnDeposit, func(ctx b.Context) error {
		UserID := ctx.Chat().ID
		if !UserExist(UserID) {
			return ctx.Send("Please /start Again")
		}
		user := GetUserFromDB(UserID)
		if user.DepositAddress == "" {
			user.CreateDepositAddress()
		}

		ctx.Send(DEPOSIT(CopyedString(user.DepositAddress)), b.ModeMarkdown)

		return nil
	})
	c.Bot.Handle(&BtnFAQ, func(ctx b.Context) error {
		UserID := ctx.Chat().ID
		if !UserExist(UserID) {
			return ctx.Send("Please /start Again")
		}

		ctx.Send(FAQ(), b.ModeMarkdown)

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

				}
				if user.Location == "main" {
					HandelMain(ctx, user)

				} else if user.Location == "games" {
					HandelGameBoard(ctx, user)

				} else if strings.Contains(user.Location, "dice") {
					HandelDice(ctx, user)

				}
				return nil
			} else {
				ctx.Send("Don't Spam Bitch")
			}
		} else {
			ctx.Send("Please /start Again")
		}

		return nil
	})

}

func HandelMain(ctx b.Context, user UserRedis) {
	input := ctx.Text()
	UserID := ctx.Chat().ID
	if input == BtnGames.Text {
		user.ChangeLocation("games")
		ctx.Send("GameBoard", GameMenu)

	} else if input == BtnReferrals.Text {
		link := fmt.Sprintf("t.me/%v?start=%v", c.BotUsername, UserID)
		ctx.Send(link)
	} else if input == BtnAccount.Text {
		userdata := GetUserFromDB(UserID)
		link := "tg://user?id=" + strconv.FormatInt(UserID, 10)
		ctx.Send(ACCOUNT(ctx.Chat().FirstName, link, userdata.Balance, userdata.Referrals, userdata.Warn, CopyedString(userdata.Wallet)), b.ModeMarkdown)
	}
}

func CopyedString(str string) string {
	return fmt.Sprintf("`%v`", str)
}
