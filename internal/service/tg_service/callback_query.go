package tg_service

import (
	"fmt"
	"myapp/internal/models"
	"myapp/pkg/files"
	my_regex "myapp/pkg/regex"
	"strconv"
	"strings"
	"time"
)

func (srv *TgService) HandleCallbackQuery(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	srv.l.Info(fmt.Sprintf("HandleCallbackQuery: fromId: %d, fromUsername: %s, cq.Data: %s", fromId, fromUsername, cq.Data))

	srv.Db.EditStep(fromId, fmt.Sprintf("кнопка: %s", cq.Data))

	srv.SendMsgToServer(fromId, "user", fmt.Sprintf("кнопка: %s", cq.Data))

	if cq.Data == "bad_answer_article" {
		srv.SendMessage(fromId, "❌ Ответ неверный, перечитай текст выше и попробуй еще раз")
		return nil
	}

	go func() {
		if cq.Data != "subscribe" {
			if fromId != 6151764130 {
				time.Sleep(time.Second)
				srv.EditMessageReplyMarkup(fromId, cq.Message.MessageId)
				srv.Db.UpdateLatsActiontime(fromId)
			}
		}
	}()

	// user, err := srv.Db.GetUserById(fromId)
	// if err != nil {
	// 	return fmt.Errorf("HandleCallbackQuery GetUserById err: %v", err)
	// }
	// if user.Id != 0 && user.Lives == 0 {
	// 	return nil
	// }

	if cq.Data == "delete_user_by_username_btn" {
		err := srv.CQ_delete_user_by_username_btn(m)
		if err != nil {
			srv.SendMessage(fromId, ERR_MSG)
			srv.SendMessage(fromId, err.Error())
		}
		return err
	}

	if cq.Data == "delete_user_by_id_btn" {
		err := srv.CQ_delete_user_by_id_btn(m)
		if err != nil {
			srv.SendMessage(fromId, ERR_MSG)
			srv.SendMessage(fromId, err.Error())
		}
		return err
	}

	if cq.Data == "start_game" {
		err := srv.CQ_start_game(m)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if cq.Data == "restart_game" {
		err := srv.CQ_restart_game(m)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if cq.Data == "subscribe" {
		err := srv.CQ_subscribe(m)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if cq.Data == "zabrat_instr" {
		err := srv.CQ_zabrat_instr(m)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if strings.HasPrefix(cq.Data, "show_q_") { // показать mil вопрос
		if strings.Contains(strings.ToLower(cq.Message.Text), "ответ неверный") || (cq.Message.Caption != nil &&  strings.Contains(strings.ToLower(*cq.Message.Caption), "ответ неверный")) {
			time.Sleep(time.Second)
			srv.DeleteMessage(fromId, cq.Message.MessageId)
			srv.DeleteMessage(fromId, cq.Message.MessageId-1)
		}

		qId := my_regex.GetStringInBetween(cq.Data, "show_q_", "_")
		qIdInt, _ := strconv.Atoi(qId)
		err := srv.ShowMilQ(fromId, qIdInt)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if strings.HasPrefix(cq.Data, "_lose_q_") { // показать "Попробовать еще раз" на вопрос
		qId := my_regex.GetStringInBetween(cq.Data, "_lose_q_", "_")
		err := srv.ShowQLose(fromId, qId)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if strings.HasPrefix(cq.Data, "_win_q_") {
		qId := my_regex.GetStringInBetween(cq.Data, "_win_q_", "_")
		err := srv.ShowQWin(fromId, qId)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if strings.HasPrefix(cq.Data, "prodolzit_") { // prodolzit_14_
		prodolzit_id := my_regex.GetStringInBetween(cq.Data, "prodolzit_", "_")
		err := srv.Prodolzit(fromId, prodolzit_id)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if cq.Data == "mailing_copy_btn" {
		err := srv.CQ_mailing_copy_btn(m)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	srv.Db.UpdateLatsActiontime(fromId)
	return nil
}

func (srv *TgService) CQ_start_game(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	srv.l.Info(fmt.Sprintf("CQ_start_game: fromId: %d, fromUsername: %s", fromId, fromUsername))

	srv.SendAnimMessage("-1", fromId, animTimeout250)
	srv.SendBalance(fromId, "1000", animTimeout250)
	srv.SendAnimMessageHTML("2", fromId, animTimeoutTest)
	srv.SendAnimMessage("4", fromId, animTimeoutTest)
	srv.Db.EditStep(fromId, "5")
	srv.SendAnimMessage("5", fromId, animTimeoutTest)

	err := srv.ShowMilQ(fromId, 1)
	if err != nil {
		return fmt.Errorf("CQ_start_game ShowMilQ1 err: %v", err)
	}

	return nil
}

func (srv *TgService) CQ_mailing_copy_btn(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	srv.l.Info(fmt.Sprintf("CQ_start_game: fromId: %d, fromUsername: %s", fromId, fromUsername))

	srv.SendForceReply(fromId, MAILING_COPY_STEP)

	return nil
}

func (srv *TgService) CQ_restart_game(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	fromFirstName := cq.From.FirstName
	srv.l.Info(fmt.Sprintf("CQ_restart_game: fromId: %d, fromUsername: %s", fromId, fromUsername))

	user, err := srv.Db.GetUserById(fromId)
	if err != nil {
		return fmt.Errorf("CQ_restart_game GetUserById err: %v", err)
	}
	if user.CreatedAt != "" && srv.IsIgnoreUser(fromId) {
		return nil
	}

	err = srv.Db.AddNewUser(fromId, fromUsername, fromFirstName)
	if err != nil {
		return fmt.Errorf("CQ_restart_game AddNewUser err: %v", err)
	}
	srv.Db.EditBotState(fromId, "")
	srv.Db.EditLives(fromId, 3)
	srv.SendMessageAndDb(fromId, fmt.Sprintf("Привет, %s 👋", fromFirstName))

	srv.Db.EditStep(fromId, "1")
	srv.SendAnimMessageHTML("1", fromId, animTimeout3000)

	time.Sleep(time.Millisecond * time.Duration(animTimeoutTest))
	
	text := "Прямо сейчас начинай игру и забирай бонус 1000₽ за уверенный старт! 🚀"
	replyMarkup := `{"inline_keyboard" : [
		[{ "text": "Начать игру", "callback_data": "start_game" }]
	]}`
	srv.SendMessageWRM(fromId, text, replyMarkup)
	
	srv.SendMsgToServer(fromId, "bot", text)
	srv.Db.UpdateLatsActiontime(fromId)

	return nil
}

func (srv *TgService) CQ_subscribe(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	srv.l.Info(fmt.Sprintf("CQ_subscribe: fromId: %d, fromUsername: %s", fromId, fromUsername))

	user, _ := srv.Db.GetUserById(fromId)
	scheme, _ := srv.Db.GetsSchemeById(user.Ref)
	ChatToCheck := scheme.ChatCheckId
	// ChatToCheck := -1001954824103
	// if user.Ref == "ref15" {
	// 	ChatToCheck = -1001771020146
	// }
	// if user.Ref == "ref6" {
	// 	ChatToCheck = -1001980240287
	// }

	GetChatMemberResp, err := srv.GetChatMember(fromId, ChatToCheck)
	if err != nil {
		return fmt.Errorf("CQ_subscribe GetChatMember fromId: %d, ChatToCheck: %d, err: %v", fromId, ChatToCheck, err)
	}
	if GetChatMemberResp.Result.Status != "member" && GetChatMemberResp.Result.Status != "creator" {
		logMess := fmt.Sprintf("CQ_subscribe GetChatMember bad resp: %+v", GetChatMemberResp)
		srv.l.Error(logMess)
		mess := "❌ вы не подписаны на канал!"
		srv.SendMessageAndDb(fromId, mess)
		srv.Db.EditStep(fromId, mess)
		return nil
	}

	go func() {
		time.Sleep(time.Second)
		srv.EditMessageReplyMarkup(fromId, cq.Message.MessageId)
	}()

	messText := "Отлично! Осталось 2 последних условия 😎\nСмотри кружочек 👇🏻"
	reply_markup := `{
		"keyboard" : [[{ "text": "Написать Марку", "resize": true }, { "text": "Часто задаваемые вопросы", "resize": true }]],
		"resize_keyboard": true
	}`
	_, err = srv.SendMessageWRM(fromId, messText, reply_markup)
	if err != nil {
		srv.l.Error("Написать Марку err: ", err)
	}
	time.Sleep(time.Second)

	lichka := user.Lichka
	scheme, _ = srv.Db.GetsSchemeByLichka(lichka)
	schemeLink := scheme.Link

	reply_markup = fmt.Sprintf(`{"inline_keyboard" : [
		[{ "text": "Зарегистрироваться", "url": "%s" }]
	]}`, schemeLink)
	
	futureJson := map[string]string{
		"video_note":   fmt.Sprintf("@%s", "./files/krug_2.mp4"),
		"chat_id": strconv.Itoa(fromId),
		"reply_markup": reply_markup,
	}
	cf, body, err := files.CreateForm(futureJson)
	if err != nil {
		return fmt.Errorf("CQ_subscribe CreateForm err: %v", err)
	}
	_, err = srv.SendVideoNote(body, cf)
	if err != nil {
		return fmt.Errorf("CQ_subscribe SendVideoNote err: %v", err)
	}

	// textMess := fmt.Sprintf(
	// 	"Переходи и регистрируйся по ссылке:\n\n%s\n\nДалее присылай сюда почту, на которую регистрировался 👇🏻",
	// 	srv.ChInfoToLinkHTML("", "ССЫЛКА"),
	// )
	// srv.SendMessageHTML(fromId, textMess)

	srv.Db.EditBotState(fromId, "wait_email")

	srv.SendMsgToServer(fromId, "bot", "wait_email")
	srv.SendMsgToServer(fromId, "bot", user.Ref)
	srv.SendMsgToServer(fromId, "bot", schemeLink)

	return nil
}

func (srv *TgService) CQ_zabrat_instr(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	// fromFirstName := cq.From.FirstName
	srv.l.Info(fmt.Sprintf("CQ_zabrat_instr: fromId: %d, fromUsername: %s", fromId, fromUsername))

	srv.Send3Kruga(fromId)

	return nil
}

func (srv *TgService) CQ_delete_user_by_username_btn(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	srv.l.Info(fmt.Sprintf("CQ_delete_user_by_username_btn: fromId: %d, fromUsername: %s", fromId, fromUsername))

	srv.SendForceReply(fromId, DEL_USER_MSG)
	return nil
}

func (srv *TgService) CQ_delete_user_by_id_btn(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	srv.l.Info(fmt.Sprintf("CQ_delete_user_by_id_btn: fromId: %d, fromUsername: %s", fromId, fromUsername))

	srv.SendForceReply(fromId, DEL_USER_ID_MSG)
	return nil
}
