package handlers

import (
	gateway "dogegambling/Gateway"
	c "dogegambling/config"
	"encoding/json"
	"strings"

	"strconv"

	b "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func (h Handler) Init() {
	h.UserInit()
	MenuInint()
	Admin := h.Bot.Group()
	Admin.Use(middleware.Whitelist(c.Admins...))
	h.Bot.Handle("/start", func(ctx b.Context) error {
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
								h.Bot.Send(inviter, "you got a new referral")
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
	h.Bot.Handle(&BtnDeposit, func(ctx b.Context) error {
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
	h.Bot.Handle(&BtnFAQ, func(ctx b.Context) error {
		UserID := ctx.Chat().ID
		if !UserExist(UserID) {
			return ctx.Send("Please /start Again")
		}

		ctx.Send(FAQ(), b.ModeMarkdown)

		return nil
	})
	h.Bot.Handle(b.OnText, func(ctx b.Context) error {
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
				} else if strings.Contains(user.Location, "withdraw") {
					HandelWithdraw(ctx, &user, h.Bot)
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

	Admin.Handle("/add", func(ctx b.Context) error {
		userid, _ := strconv.Atoi(ctx.Args()[0])
		amount, _ := strconv.ParseFloat(ctx.Args()[1], 32)
		BetWin(int64(userid), float32(amount))
		ctx.Send("Sent")
		return nil
	})
	Admin.Handle(&BtnConfirmWithdraw, func(ctx b.Context) error {
		user := GetUser(ctx.Sender().ID)
		if !user.IsLock && user.NotSpam() {
			user.UpdateTime()
			user.lock()
			defer user.unlock()

			ctx.Edit("🔁**Processing** **...**", b.ModeMarkdown)
			pid, _ := strconv.Atoi(ctx.Data())
			ctx.Edit("🔁**Processing** **...**\n Receiving Information ...", b.ModeMarkdown)
			w := GetPaymentByID(pid)
			if w.Status != "Pending" {
				ctx.Edit(ConfirmWithdrawTextChannel(w.UserRefer, w.Amount, w.TxID), b.ModeMarkdown)
				return nil
			}
			ctx.Edit("🔁**Processing** **...**\n Submit a withdrawal request ...", b.ModeMarkdown)
			res, err := gateway.Withdraw(float64(w.Amount), GetWalletAddress(w.UserRefer))
			if err != nil {
				return nil
			} else if res.Status != "success" {
				text, _ := json.Marshal(res)
				ctx.Send(string(text))
				return nil
			}
			ctx.Edit("🔁**Processing** **...**\n Record the result ...", b.ModeMarkdown)
			ConfirmWithdraw(pid, res.Data.Txid)
			ctx.Edit(ConfirmWithdrawTextChannel(w.UserRefer, w.Amount, res.Data.Txid), b.ModeMarkdown)
			SendToUser(h.Bot, w.UserRefer, ResponseConfirmWithdraw(res.Data.Txid))
			SendToChannel(h.Bot, c.TransactionChannelID, ResponseConfirmWithdraw(res.Data.Txid))
		}

		return nil
	})
	Admin.Handle(&BtnRejectWithdraw, func(ctx b.Context) error {
		user := GetUser(ctx.Sender().ID)
		if !user.IsLock && user.NotSpam() {
			user.UpdateTime()
			user.lock()
			defer user.unlock()

			ctx.Edit("🔁**Processing** **...**", b.ModeMarkdown)
			pid, _ := strconv.Atoi(ctx.Data())
			ctx.Edit("🔁**Processing** **...**\n Receiving Information ...", b.ModeMarkdown)
			w := GetPaymentByID(pid)
			if w.Status != "Pending" {
				ctx.Edit(ConfirmWithdrawTextChannel(w.UserRefer, w.Amount, w.TxID), b.ModeMarkdown)
				return nil
			}
			RejectWithdraw(pid)
			SendToUser(h.Bot, w.UserRefer, ResponseRejectWithdraw(pid))
			ctx.Edit(RejectWithdrawTextChannel(w.UserRefer, w.Amount, pid), b.ModeMarkdown)
		}
		return nil
	})
}
