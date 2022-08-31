package handlers

import (
	c "dogegambling/config"
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
			if user.Location != Main {
				user.ChangeLocation(Main)
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
	Admin.Handle("/add", func(ctx b.Context) error {
		userid, _ := strconv.Atoi(ctx.Args()[0])
		amount, _ := strconv.ParseFloat(ctx.Args()[1], 32)
		BetWin(int64(userid), float32(amount))
		ctx.Send("Sent")
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
				if user.Location != Main {
					if input == BtnHome.Text {
						user.ChangeLocation(Main)
						return ctx.Send("Home", MainMenu)

					}

				}
				if user.Location == Main {
					HandelMain(ctx, &user)
					return nil

				} else if user.Location == Games {
					HandelGameBoard(ctx, &user)
					return nil

				} else if strings.Contains(user.Location, "dice") {
					HandelDice(ctx, &user)
					return nil

				} else if strings.Contains(user.Location, "bowl") {
					HandelBowl(ctx, &user)
					return nil

				} else if strings.Contains(user.Location, "dart") {
					HandelDart(ctx, &user)
					return nil

				} else if strings.Contains(user.Location, "slot") {
					HandelSlot(ctx, &user)
					return nil

				} else if strings.Contains(user.Location, "basket") {
					HandelBasket(ctx, &user)
					return nil

				} else if strings.Contains(user.Location, "acc") {
					HandelAccount(ctx, &user)
					return nil
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
