package tg_service

import (
	"fmt"
	"strconv"
	"time"
)

func (srv *TgService) ShowMilQ(chatId, qNum int) error {
	time.Sleep(time.Millisecond * time.Duration(animTimeoutTest))

	textMap := map[int]string{
		1: "Первый вопрос 👆\n\nВыбери правильный ответ 👇",
		2: "Второй вопрос 👆\n\nВыбери правильный ответ 👇",
		3: "Третий вопрос 👆\n\nВыбери правильный ответ 👇",
	}
	fileNameMap := map[int]string{
		1:  "./files/mil_q1.jpg",
		2:  "./files/mil_q2.jpg",
		3:  "./files/mil_q9.jpg",
	}
	replyMarkupMap := map[int]string{
		1: `{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "_lose_q_1_" }, { "text": "B", "callback_data": "_win_q_1_" }, { "text": "C", "callback_data": "_lose_q_1_" }, { "text": "D", "callback_data": "_lose_q_1_" }]
		]}`,
		2: `{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "_lose_q_2_" }, { "text": "B", "callback_data": "_lose_q_2_" }, { "text": "C", "callback_data": "_lose_q_2_" }, { "text": "D", "callback_data": "_win_q_2_" }]
		]}`,
		3: `{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "_lose_q_3_" }, { "text": "B", "callback_data": "_lose_q_3_" }, { "text": "C", "callback_data": "_win_q_3_" }, { "text": "D", "callback_data": "_lose_q_3_" }]
		]}`,
	}

	text := textMap[qNum]
	replyMarkup := replyMarkupMap[qNum]
	fileNameInServer := fileNameMap[qNum]
	_, err := srv.SendPhotoWCaptionWRM(chatId, text, fileNameInServer, replyMarkup)
	if err != nil {
		return fmt.Errorf("ShowMilQ SendPhotoWCaptionWRM err: %v", err)
	}
	// srv.Db.EditStep(chatId, text)
	// srv.SendMsgToServer(chatId, "bot", text)
	return nil
}

func (srv *TgService) Prodolzit(chatId int, prodolzit_id string) error {
	time.Sleep(time.Second * 2)
	if prodolzit_id != "0" {
		srv.SendMessage(chatId, "Правильный ответ ✅")
	}
	prodolzitIdInt, _ := strconv.Atoi(prodolzit_id)

	if prodolzit_id == "0" {
		srv.SendAnimArticleHTMLV3("1.1", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 4) // https://t.me/c/1998413789/4
		srv.SendAnimArticleHTMLV3("1.2", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 9)

		text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "bad_answer_article" }, { "text": "B", "callback_data": "prodolzit_%d_" }, { "text": "C", "callback_data": "bad_answer_article" }, { "text": "D", "callback_data": "bad_answer_article" }]
		]}`, prodolzitIdInt+1)
		fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		return nil
	}
	if prodolzit_id == "1" {
		srv.SendAnimArticleHTMLV3("1.3", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 11)
		srv.SendAnimArticleHTMLV3("1.4", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 13)


		text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "bad_answer_article" }, { "text": "B", "callback_data": "bad_answer_article" }, { "text": "C", "callback_data": "prodolzit_%d_" }, { "text": "D", "callback_data": "bad_answer_article" }]
		]}`, prodolzitIdInt+1)
		fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		return nil
	}
	if prodolzit_id == "2" {
		srv.SendAnimArticleHTMLV3("1.5", chatId, 2000)
		
		text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "prodolzit_%d_" }, { "text": "B", "callback_data": "bad_answer_article" }, { "text": "C", "callback_data": "bad_answer_article" }, { "text": "D", "callback_data": "bad_answer_article" }]
		]}`, prodolzitIdInt+1)
		fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		return nil
	}
	if prodolzit_id == "3" {
		srv.CopyMessage(chatId, -1001998413789, 15)
		srv.SendAnimArticleHTMLV3("1.6", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 17)
		
		text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "bad_answer_article" }, { "text": "B", "callback_data": "prodolzit_%d_" }, { "text": "C", "callback_data": "bad_answer_article" }, { "text": "D", "callback_data": "bad_answer_article" }]
		]}`, prodolzitIdInt+1)
		fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		return nil
	}
	if prodolzit_id == "4" {
		srv.SendAnimArticleHTMLV3("1.7", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 19)
		
		text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "prodolzit_%d_" }, { "text": "B", "callback_data": "bad_answer_article" }, { "text": "C", "callback_data": "bad_answer_article" }, { "text": "D", "callback_data": "bad_answer_article" }]
		]}`, prodolzitIdInt+1)
		fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		return nil
	}
	if prodolzit_id == "5" {
		srv.SendAnimArticleHTMLV3("1.8", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 21)
		
		text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "bad_answer_article" }, { "text": "B", "callback_data": "prodolzit_%d_" }, { "text": "C", "callback_data": "bad_answer_article" }, { "text": "D", "callback_data": "bad_answer_article" }]
		]}`, prodolzitIdInt+1)
		fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		return nil
	}
	if prodolzit_id == "6" {
		srv.SendAnimArticleHTMLV3("1.9", chatId, 2000)

		text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "bad_answer_article" }, { "text": "B", "callback_data": "bad_answer_article" }, { "text": "C", "callback_data": "prodolzit_%d_" }, { "text": "D", "callback_data": "bad_answer_article" }]
		]}`, prodolzitIdInt+1)
		fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		return nil
	}
	if prodolzit_id == "7" {
		srv.SendAnimArticleHTMLV3("2.1", chatId, 2000)
		srv.SendAnimArticleHTMLV3("2.2", chatId, 2000)
		srv.SendAnimArticleHTMLV3("2.3", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 29)
		
		text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "bad_answer_article" }, { "text": "B", "callback_data": "prodolzit_%d_" }, { "text": "C", "callback_data": "bad_answer_article" }, { "text": "D", "callback_data": "bad_answer_article" }]
		]}`, prodolzitIdInt+1)
		fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		return nil
	}
	if prodolzit_id == "8" {
		srv.SendAnimArticleHTMLV3("2.4", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 31)
		
		text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "bad_answer_article" }, { "text": "B", "callback_data": "prodolzit_%d_" }, { "text": "C", "callback_data": "bad_answer_article" }, { "text": "D", "callback_data": "bad_answer_article" }]
		]}`, prodolzitIdInt+1)
		fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		return nil
	}
	if prodolzit_id == "9" {
		srv.SendAnimArticleHTMLV3("2.5", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 33)
		srv.SendAnimArticleHTMLV3("2.6", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 35)
		
		text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "bad_answer_article" }, { "text": "B", "callback_data": "prodolzit_%d_" }, { "text": "C", "callback_data": "bad_answer_article" }, { "text": "D", "callback_data": "bad_answer_article" }]
		]}`, prodolzitIdInt+1)
		fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		return nil
	}
	if prodolzit_id == "10" {
		srv.SendAnimArticleHTMLV3("2.7", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 37)
		srv.SendAnimArticleHTMLV3("2.8", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 39)
		
		text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "prodolzit_%d_" }, { "text": "B", "callback_data": "bad_answer_article" }, { "text": "C", "callback_data": "bad_answer_article" }, { "text": "D", "callback_data": "bad_answer_article" }]
		]}`, prodolzitIdInt+1)
		fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		return nil
	}
	if prodolzit_id == "11" {
		srv.SendAnimArticleHTMLV3("2.9", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 41)
		srv.SendAnimArticleHTMLV3("2.10", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 43)
		srv.SendAnimArticleHTMLV3("2.11", chatId, 2000)

		text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "prodolzit_%d_" }, { "text": "B", "callback_data": "bad_answer_article" }, { "text": "C", "callback_data": "bad_answer_article" }, { "text": "D", "callback_data": "bad_answer_article" }]
		]}`, prodolzitIdInt+1)
		fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		return nil
	}
	if prodolzit_id == "12" {
		srv.SendAnimArticleHTMLV3("3.3", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 51)
		
		text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "bad_answer_article" }, { "text": "B", "callback_data": "prodolzit_%d_" }, { "text": "C", "callback_data": "bad_answer_article" }, { "text": "D", "callback_data": "bad_answer_article" }]
		]}`, prodolzitIdInt+1)
		fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		return nil
	}
	if prodolzit_id == "13" {
		srv.SendAnimArticleHTMLV3("3.4", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 53)
		srv.SendAnimArticleHTMLV3("3.5", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 55)
		
		text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "bad_answer_article" }, { "text": "B", "callback_data": "bad_answer_article" }, { "text": "C", "callback_data": "prodolzit_%d_" }, { "text": "D", "callback_data": "bad_answer_article" }]
		]}`, prodolzitIdInt+1)
		fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		return nil
	}
	if prodolzit_id == "14" {
		srv.SendAnimArticleHTMLV3("3.6", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 57)
		srv.SendAnimArticleHTMLV3("3.7", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 59)
		srv.SendAnimArticleHTMLV3("3.8", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 65)
		srv.SendAnimArticleHTMLV3("3.9", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 66)
		
		// text := fmt.Sprintf("Вопрос %d из 15\n\nВыбери правильный ответ 👇", prodolzitIdInt+1)
		// reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
		// 	[ { "text": "A", "callback_data": "prodolzit_%d_" }, { "text": "B", "callback_data": "bad_answer_article_" } ]
		// ]}`, prodolzitIdInt+1)
		// fileName := fmt.Sprintf("./files/article_q_%d.jpg", prodolzitIdInt+1)
		// srv.SendPhotoWCaptionWRM(chatId, text, fileName, reply_markup)
		srv.Db.EditBotState(chatId, "read_article_after_TrurOrFalse_win")
		return nil
	}
	if prodolzit_id == "15" {
		// srv.SendAnimArticleHTMLV3("3.10", chatId, 2000)
		// messText := "Все условия выполнены! Поздравляю! 🎉\n\nЯ подключил к твоему аккаунту необходимые настройки, благодаря которым ты уже сегодня сможешь вытащить солидную прибыль.\n\nНиже отправляю тебе инструкцию, повторив которую ты уже сегодня заработаешь от 500.000₽👇\n\nВсё работает на 1.000%! Попробуй и убедись🤝"
		// reply_markup := `{"inline_keyboard" : [ [{ "text": "Забрать инструкцию", "callback_data": "zabrat_instr" }]]}`
		// srv.SendMessageWRM(chatId, messText, reply_markup)

		// user, _ := srv.Db.GetUserById(chatId)
		// lichka := user.Lichka
		// // lichkaUrl := fmt.Sprintf("https://t.me/%s", srv.DelAt(lichka))
		// scheme, _ := srv.Db.GetsSchemeByLichka(lichka)
	
		// base64Str := srv.CreateBase64UserData(chatId, user.Username, user.Firstname)
		// siteUrl := fmt.Sprintf("%s&data=%s", scheme.Link, base64Str)
	
		// mesgText := srv.GetActualSchema(chatId, siteUrl)
	
		// _, err := srv.SendMessageHTML(chatId, mesgText)
		// if err != nil {
		// 	srv.l.Error(fmt.Errorf("CQ_zabrat_instr SendMessageWRM err: %v", err))
		// }
		// srv.SendMsgToServer(chatId, "bot", mesgText)

		return nil
	}

	return nil
}

func (srv *TgService) ShowQLose(chatId int, q_num string) error {
	time.Sleep(time.Millisecond * time.Duration(animTimeoutTest))

	text := "Ответ неверный ❌\nК сожалению, ты ошибся, но шанс еще есть!\n\nЖми на кнопку 👇"
	reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
		[{ "text": "Попробовать еще раз", "callback_data": "show_q_%s_" }]
	]}`, q_num)
	srv.SendMessageWRM(chatId, text, reply_markup)

	// srv.SendMsgToServer(chatId, "bot", text)
	return nil
}

func (srv *TgService) ShowQWin(chatId int, q_num string) error {
	time.Sleep(time.Millisecond * time.Duration(animTimeoutTest))
	
	textMap := map[string]string{
		"1":  "Отлично, ты дал верный ответ ✅",
		"2":  "Снова в цель! ✅\nЕще один вопросик для победы 😏",
		"3": "Ответ верный ✅✅✅\nПоздравляю с победой 🎉",
	}

	if q_num == "1" {
		srv.SendMessageAndDb(chatId, textMap[q_num])
		time.Sleep(time.Millisecond * 2000)
		srv.ShowMilQ(chatId, 2)
		return nil
	}
	if q_num == "2" {
		srv.SendMessageAndDb(chatId, textMap[q_num])
		time.Sleep(time.Millisecond * 2000)
		srv.ShowMilQ(chatId, 3)
		return nil
	}
	if q_num == "3" {
		srv.SendMessageAndDb(chatId, textMap[q_num])
		time.Sleep(time.Millisecond * 2000)
		// srv.Db.EditStep(chatId, "6")
		// srv.SendAnimMessage("6", chatId, animTimeoutTest)
		// time.Sleep(time.Second)

		user, _ := srv.Db.GetUserById(chatId)
		scheme, _ := srv.Db.GetsSchemeById(user.Ref)
		chLink := scheme.ChatCheckLink
		// chLink := "https://t.me/+aC2nuKzjkTwxNzZh"
		// if user.Ref == "ref15" {
		// 	chLink = "https://t.me/+yGMKYazOioVkMTAx"
		// }
		// if user.Ref == "ref6" {
		// 	chLink = "https://t.me/+36S69s4Z9GY4NWUx"
		// }

		messText := fmt.Sprintf("Чтобы разблокировать награду и забрать её, тебе осталось выполнить 3 простейших условия:\n\n1. Подпишись на мой канал👇\n%s\n\nКак только подписался - жми кнопку ниже ⏬", chLink)
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "Подписался☑️", "callback_data": "subscribe" }]
		]}`
		srv.SendMessageWRM(chatId, messText, reply_markup)

		// srv.SendMsgToServer(chatId, "bot", "Ссылка")
		return nil
	}
	if q_num == "4" {
		srv.SendMessageAndDb(chatId, textMap[q_num])
		time.Sleep(time.Second * 2)
		srv.ShowMilQ(chatId, 5)
		return nil
	}
	if q_num == "5" {
		srv.SendMessageAndDb(chatId, textMap[q_num])
		time.Sleep(time.Second * 2)
		srv.ShowMilQ(chatId, 6)
		return nil
	}
	if q_num == "6" {
		srv.SendMessageAndDb(chatId, textMap[q_num])
		time.Sleep(time.Second * 2)

		srv.Db.EditStep(chatId, "8")
		srv.SendAnimMessageHTML("8", chatId, 2000)
		srv.Db.EditBotState(chatId, "read_article_after_KNB_win")
		return nil
	}
	if q_num == "7" {
		srv.SendMessageAndDb(chatId, textMap[q_num])
		time.Sleep(time.Second * 2)
		srv.ShowMilQ(chatId, 8)
		return nil
	}
	if q_num == "8" {
		srv.SendMessageAndDb(chatId, textMap[q_num])
		time.Sleep(time.Second * 2)
		srv.ShowMilQ(chatId, 9)
		return nil
	}
	if q_num == "9" {
		srv.SendMessageAndDb(chatId, textMap[q_num])
		time.Sleep(time.Second * 2)

		srv.Db.EditStep(chatId, "10")
		srv.SendAnimMessageHTML("10", chatId, 2000)
		srv.Db.EditBotState(chatId, "read_article_after_OIR_win")
		return nil
	}
	if q_num == "10" {
		srv.SendMessageAndDb(chatId, textMap[q_num])
		time.Sleep(time.Second * 2)
		srv.ShowMilQ(chatId, 11)
		return nil
	}
	if q_num == "11" {
		srv.SendMessageAndDb(chatId, textMap[q_num])
		time.Sleep(time.Second * 2)
		srv.ShowMilQ(chatId, 12)
		return nil
	}
	if q_num == "12" {
		srv.SendMessageAndDb(chatId, textMap[q_num])
		time.Sleep(time.Second * 2)

		srv.Db.EditStep(chatId, "12")
		srv.SendAnimMessageHTML("12", chatId, 2000)
		srv.Db.EditBotState(chatId, "read_article_after_TrurOrFalse_win")
		return nil
	}
	return nil
}


func (srv *TgService) ShowQLosePhoto(chatId int, q_num string) error {
	time.Sleep(time.Millisecond * time.Duration(animTimeoutTest))

	user, err := srv.Db.GetUserById(chatId)
	if err != nil {
		return fmt.Errorf("ShowQLosePhoto GetUserById err: %v", err)
	}
	if user.Lives == 0 {
		return fmt.Errorf("0 жизней")
	}
	newLivesCnt := user.Lives - 1
	srv.Db.EditLives(chatId, newLivesCnt)

	pushTextMap := map[int]string{
		1: "❤️❤️🖤\nОтвет неверный ❌\n\nУ тебя сгорела одна жизнь 😔",
		2: "❤️🖤🖤\nОтвет неверный ❌\n\nУ тебя сгорела вторая жизнь, и это очень печально 😒",
		3: "🖤🖤🖤\nОтвет неверный ❌\n\nУ тебя сгорели все жизни 🥶\n\nНо у тебя еще есть шанс восстановить их.",
	}
	messIndex := 3 - newLivesCnt
	text := pushTextMap[messIndex]
	fileNameInServer := fmt.Sprintf("./files/push_%d.jpg", messIndex)
	if newLivesCnt != 0 {
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[{ "text": "Попробовать еще раз", "callback_data": "show_q_%s_" }]
		]}`, q_num)
		_, err = srv.SendPhotoWCaptionWRM(chatId, text, fileNameInServer, reply_markup)
		if err != nil {
			return fmt.Errorf("ShowQLosePhoto SendPhotoWCaptionWRM err: %v", err)
		}
	} else {
		_, err = srv.SendPhotoWCaption(chatId, text, fileNameInServer)
		if err != nil {
			return fmt.Errorf("ShowQLosePhoto SendPhotoWCaptionWRM err: %v", err)
		}
	}
	srv.SendMsgToServer(chatId, "bot", text)

	if newLivesCnt == 0 && user.IsLastPush == 0 {
		huersStr, _ := srv.GetUserLeftTime(chatId)
		text = fmt.Sprintf("❗️У тебя есть %s на то, чтобы начать игру заново♻️\n\nЕсли ты не успеешь запустить игру за это время, то доступ к боту будет закрыт навсегда. Перезапуск бота не поможет, он просто перестанет работать для тебя ⛔️", huersStr)
		replyMarkup := `{"inline_keyboard" : [
			[{ "text": "ЗАБРАТЬ 100.000₽", "callback_data": "restart_game" }]
		]}`
		srv.SendMessageWRM(chatId, text, replyMarkup)
		srv.SendMsgToServer(chatId, "bot", text)
	}

	return nil
}
