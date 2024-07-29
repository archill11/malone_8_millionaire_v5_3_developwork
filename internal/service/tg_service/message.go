package tg_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"myapp/internal/models"
	"myapp/pkg/files"
	my_regex "myapp/pkg/regex"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

func (srv *TgService) HandleMessage(m models.Update) error {
	msgText := m.Message.Text
	fromUsername := m.Message.From.UserName
	fromId := m.Message.From.Id
	srv.l.Info(fmt.Sprintf("HandleMessage: fromId-%d fromUsername-%s, msgText-%s", fromId, fromUsername, msgText))

	// srv.Db.EditStep(fromId, msgText)

	go func() {
		if msgText != "/admin" && msgText != "/start" && strings.HasPrefix(msgText, "add_am_") && msgText != "wait_email" {
			srv.l.Warn("syka 2")
			time.Sleep(time.Second*4)
			srv.EditMessageReplyMarkup(fromId, m.Message.MessageId)
			for i:=m.Message.MessageId; i >= m.Message.MessageId-25; i-- {
				user, _ := srv.Db.GetUserById(fromId)
				if i == user.NotDelMessId {
					break
				}
				srv.DeleteMessage(fromId, i)
				time.Sleep(time.Millisecond*300)
			}
			// srv.Db.UpdateLatsActiontime(fromId)
		}
	}()

	srv.SendMsgToServer(fromId, "user", msgText)

	if msgText == "/admin" {
		err := srv.M_admin(m)
		if err != nil {
			srv.SendMessage(fromId, ERR_MSG)
			srv.SendMessage(fromId, err.Error())
		}
		return err
	}

	user, err := srv.Db.GetUserById(fromId)
	if err != nil {
		return fmt.Errorf("HandleMessage GetUserById err: %v", err)
	}
	// if user.Id != 0 && user.Lives == 0 {
	// 	return nil
	// }

	if msgText == "/help" {
		srv.SendMessageAndDb(fromId, "@millioner_support\nвот контакт для связи")
		srv.Db.UpdateLatsActiontime(fromId)
		srv.Db.UpdateFeedbackTime(fromId)
		return nil
	}

	if msgText == "Написать Марку" {
		user, _ := srv.Db.GetUserById(fromId)
		lichka := user.Lichka
		lichkaUrl := fmt.Sprintf("https://t.me/%s", srv.DelAt(lichka))
		messText := "Если у тебя имеются какие-то вопросы - смело задавай мне их в лс по кнопке ниже 👇🏻"
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[{ "text": "Написать Марку", "url": "%s" }]
		]}`, lichkaUrl)
		srv.SendMessageWRM(fromId, messText, reply_markup)
		return nil
	}

	if msgText == "Часто задаваемые вопросы" {
		srv.CQ_frequently_questions_btn(m)
		return nil
	}



	if user.IsLastPush == 1 {
		srv.SendMessageAndDb(fromId, "бот вам больше не доступен")
		return nil
	}

	if strings.HasPrefix(msgText, "/start") { // https://t.me/tgbotusername?start=ref01 -> /start ref01
		err := srv.M_start(m)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		srv.Db.UpdateFeedbackTime(fromId)
		return err
	}

	if strings.HasPrefix(msgText, "add_am_") { // add_am_1.1_
		animMessId := my_regex.GetStringInBetween(msgText, "add_am_", "_")
		if animMessId == "" {
			return fmt.Errorf("некоректный Id статьи")
		}
		srv.Db.EditBotState(fromId, msgText)
		
		animMess, err := srv.Db.GetAminMessByTxtId(animMessId)
		if err != nil {
			return fmt.Errorf("M_state GetAminMessByTxtId err: %v", err)
		}
		if animMess.TxtMess != "" {
			srv.SendMessage(fromId, "прежняя версия👇")
			srv.SendMessageHTML(fromId, srv.ReplaceHtmlTag(animMess.TxtMess))
		} else {
			srv.SendMessage(fromId, "в базе еще нет этой статьи")
		}

		srv.SendMessage(fromId, fmt.Sprintf("Ожидание поста для статьи %v", animMessId))
		return nil
	}

	err = srv.M_state(m)
	if err != nil {
		srv.SendMessageAndDb(fromId, ERR_MSG)
		srv.SendMessageAndDb(fromId, err.Error())
	}
	srv.Db.UpdateLatsActiontime(fromId)
	srv.Db.UpdateFeedbackTime(fromId)
	return err
}

func (srv *TgService) M_start(m models.Update) error {
	fromId := m.Message.Chat.Id
	msgText := m.Message.Text
	fromFirstName := m.Message.From.FirstName
	fromUsername := m.Message.From.UserName
	srv.l.Info(fmt.Sprintf("M_start: fromId: %d, fromUsername: %s, msgText: %s", fromId, fromUsername, msgText))

	messId :=  m.Message.MessageId
	srv.Db.EditNotDelMessId(fromId, messId)

	refArr := strings.Split(msgText, " ")
	ref := ""
	if len(refArr) > 1 {
		ref = refArr[1]
	}

	// user, err := srv.Db.GetUserById(fromId)
	// if err != nil {
	// 	return fmt.Errorf("M_start GetUserById err: %v", err)
	// }
	// if user.CreatedAt != "" && srv.IsIgnoreUser(fromId) {
	// 	text := "К сожалению, время истекло и бот для вас больше недоступен.\nВы можете обратиться в поддержку через команду /help"
	// 	srv.SendMessageAndDb(fromId, text)
	// 	return nil
	// }

	err := srv.Db.AddNewUser(fromId, fromUsername, fromFirstName)
	if err != nil {
		return fmt.Errorf("M_start AddNewUser err: %v", err)
	}
	if ref == "" {
		ref = "developwork"
	}
	srv.Db.EditRef(fromId, ref)
	scheme, _ := srv.Db.GetsSchemeById(ref)
	lichka := scheme.Lichka
	srv.Db.EditLichka(fromId, lichka)
	if fromId == 1394096901 {
		srv.Db.EditAdmin(fromId, 1)
	}
	srv.Db.EditBotState(fromId, "")
	srv.Db.EditLives(fromId, 3)
	srv.Db.EditStep(fromId, "1")
	// srv.SendMessageAndDb(fromId, fmt.Sprintf("Привет, %s 👋", fromFirstName))
	// srv.SendAnimMessageHTML("1", fromId, animTimeout3000)

	// time.Sleep(time.Millisecond * time.Duration(animTimeoutTest))

	// text := "Прямо сейчас начинай игру и забирай бонус 1000₽ за уверенный старт! 🚀"
	// replyMarkup := `{"inline_keyboard" : [
	// 	[{ "text": "Начать игру", "callback_data": "start_game" }]
	// ]}`
	// srv.SendMessageWRM(fromId, text, replyMarkup)
	
	// srv.SendMsgToServer(fromId, "bot", text)

	futureJson := map[string]string{
		"video_note":   fmt.Sprintf("@%s", "./files/krug_1.mp4"),
		"chat_id": strconv.Itoa(fromId),
	}
	cf, body, err := files.CreateForm(futureJson)
	if err != nil {
		return fmt.Errorf("HandleVideoNote CreateFormV2 err: %v", err)
	}
	srv.SendVideoNote(body, cf)

	srv.Db.EditBotState(fromId, "read_article_after_KNB_win")
	time.Sleep(time.Second*20)
	srv.SendMessage(fromId, "Введи кодовое слово ниже 👇🏻")

	return nil
}

func (srv *TgService) M_state(m models.Update) error {
	fromId := m.Message.Chat.Id
	msgText := m.Message.Text
	fromUsername := m.Message.From.UserName
	fromFirstName := m.Message.From.FirstName
	srv.l.Info(fmt.Sprintf("M_state: fromId: %d, fromUsername: %s, msgText: %s", fromId, fromUsername, msgText))

	user, err := srv.Db.GetUserById(fromId)
	if err != nil {
		srv.l.Warn(fmt.Errorf("M_state GetUserById err: %v", err))
	}
	srv.Db.UpdateLatsActiontime(fromId)
	if user.BotState == "" {
		return nil
	}

	if strings.HasPrefix(user.BotState, "add_am_") {
		animMessId := my_regex.GetStringInBetween(user.BotState, "add_am_", "_")
		animMess, err := srv.Db.GetAminMessByTxtId(animMessId)
		if err != nil {
			return fmt.Errorf("M_state GetAminMessByTxtId err: %v", err)
		}
		// if animMess.TxtMess != "" {
		// 	srv.SendMessage(fromId, "прежняя версия👇")
		// 	srv.SendMessage(fromId, animMess.TxtMess)
		// }
		// srv.l.Info("m.Message.Entities:", m.Message.Entities)
		if fromId == 6151764130 {
			srv.SendMessage(fromId, fmt.Sprintf("m.Message.Entities: %+v", m.Message.Entities))
		}
		
		rawMess := msgText
		htmlMessRune := make([]rune, 0)

		for _, v := range rawMess {
			htmlMessRune = append(htmlMessRune, v)
		}

		ttt := make([]PushEntityFormat, 0)

		// for _, v := range m.Message.Entities {
		for i := len(m.Message.Entities)-1; i >= 0; i-- {
			v := m.Message.Entities[i]
			entityType := v.Type
			entityStart := v.Offset
			entityEnd := v.Offset + v.Length
			var entityStartSymb string
			var entityEndSymb string
			if entityType == "bold" {
				entityStartSymb = "|<"
				entityEndSymb = ">|"
			} else if entityType == "underline" {
				entityStartSymb = "_<"
				entityEndSymb = ">_"
			} else {
				continue
			}
			ttt = append(ttt, PushEntityFormat{EntityIndex: entityEnd, EntitySymb: []rune(entityEndSymb)})
			ttt = append(ttt, PushEntityFormat{EntityIndex: entityStart, EntitySymb: []rune(entityStartSymb)})
		}
		sort.Slice(ttt, func(i, j int) bool {
			return ttt[i].EntityIndex > ttt[j].EntityIndex
		})
		if fromId == 6151764130 {
			srv.SendMessage(fromId, fmt.Sprintf("ttt - %+v", ttt))
		}
		srv.l.Info("1")
		for _, v := range ttt {
			srv.l.Info("2")
			srv.l.Info(fmt.Sprintf("InsertSliceInSlice len(a): %d, index: %d, ", len(htmlMessRune), v.EntityIndex))
			htmlMessRune = InsertSliceInSlice(htmlMessRune, v.EntityIndex, []rune(v.EntitySymb))
		}
		srv.l.Info("3")

		msgText = string(htmlMessRune)
		_, err = srv.SendMessage(fromId, "новая версия статьи👇")
		if err != nil {
			srv.l.Error("SendMessage err", err)
		}
		_, err = srv.SendMessage(fromId, msgText)
		if err != nil {
			srv.l.Error("SendMessage err", err)
		}
		_, err = srv.SendMessage(fromId, "новая версия статьи format👇")
		if err != nil {
			srv.l.Error("SendMessage err", err)
		}
		_, err = srv.SendMessageHTML(fromId, srv.ReplaceHtmlTag(msgText))
		if err != nil {
			srv.l.Error("SendMessage err", err)
		}

		if animMess.TxtId != "" {
			err = srv.Db.EditAnimMessText(animMessId, msgText)
			if err != nil {
				return fmt.Errorf("M_state EditAnimMessText err: %v", err)
			}
			srv.SendMessage(fromId, "статья обновлена")
		} else {
			err = srv.Db.AddNewAminMess(animMessId, msgText)
			if err != nil {
				return fmt.Errorf("M_state AddNewAminMess err: %v", err)
			}
			srv.SendMessage(fromId, "статья добавлена")
		}
		srv.Db.EditBotState(fromId, "")

		srv.SendMessage(fromId, "новая версия статьи anim 👇")
		err = srv.SendAnimArticleHTMLV3(animMessId, fromId, 2000)
		if err != nil {
			srv.SendMessage(fromId, fmt.Sprintf("ERR: %v", err))
			// srv.SendMessage(fromId, "статья не обновлена")
		}

		return nil
	}

	if user.BotState == "read_article_after_KNB_win" { // Го, ко, коу, гоу, гэу
		if !strings.HasPrefix(strings.ToLower(msgText), "го") && !strings.HasPrefix(strings.ToLower(msgText), "ко") && !strings.HasPrefix(strings.ToLower(msgText), "гэ") && !strings.HasPrefix(strings.ToLower(msgText), "go") {
			srv.SendMessageAndDb(fromId, "❌ Вы неверно ввели кодовое слово, сверьтесь с лонгридом и попробуйте еще раз")
			return nil
		}

		srv.Db.EditBotState(fromId, "")
		// srv.SendAnimMessage("-1", fromId, animTimeout250)
		// srv.SendBalance(fromId, "30.000", animTimeout250)
		// srv.Db.EditStep(fromId, "9")
		// srv.SendAnimMessageHTML("9", fromId, animTimeoutTest)

		text := "Ну что, поехали, ответь правильно на 3 вопроса и уже сегодня сможешь заработать от 500.000₽ 😏"
		srv.SendMessage(fromId, text)
		err = srv.ShowMilQ(fromId, 1)
		if err != nil {
			return fmt.Errorf("M_state ShowMilQ err: %v", err)
		}

		// text := "Предлагаю тебе ответить на один вопрос 😏\nЗа него ты получишь +25.000₽ к банку💸"
		// replyMarkup :=`{"inline_keyboard" : [
		// 	[ { "text": "Давай попробуем", "callback_data": "show_q_3_" } ]
		// ]}`
		// srv.SendMessageWRM(fromId, text, replyMarkup)

		// srv.ShowMilQ(fromId, 2)
		// srv.Db.EditStep(fromId, "7")
		// srv.SendMsgToServer(fromId, "bot", text)
		return nil
	}

	if user.BotState == "wait_email" {
		msgTextEmail := msgText
		url := fmt.Sprintf("%s/api/v1/user?email=%s", srv.Cfg.ServerUrl, msgTextEmail)
		srv.l.Info("M_state wait_email иду к API", url)
		response, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("M_state wait_email Post err: %v", err)
		}
		srv.l.Info("M_state wait_email сходил к API")
		defer response.Body.Close()
	
		if response.StatusCode != http.StatusOK {
			bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
				return fmt.Errorf("M_state wait_email ReadAll err: %v", err)
			}
			return fmt.Errorf("M_state wait_email post %s bad response: [%d] %v", url, response.StatusCode, string(bodyBytes))
		}
	
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("M_state wait_email ReadAll err: %v", err)
		}
	
		resp := struct{
			Status string `json:"status"`
			Data   string `json:"data"`
		}{}
		json.Unmarshal(bodyBytes, &resp)
	
		if resp.Status == "success" {

			srv.Db.EditBotState(fromId, "")
			srv.Db.EditEmail(fromId, msgTextEmail)
			user, _ := srv.Db.GetUserById(fromId)
			scheme, _ := srv.Db.GetsSchemeById(user.Ref)
			lichkaId := scheme.LichkaId
			// lichkaId := 6405739421
			// if srv.DelAt(user.Lichka) == "mark_odlncov" {
			// 	lichkaId = 6865167980
			// }
			// lichka, tgId,  _ := srv.GetLichka()
			// srv.Db.EditLichka(fromId, lichka)
			// mess := fmt.Sprintf("Ваша личка %s", srv.AddAt(lichka))
			// srv.SendMessage(fromId, mess)

			url := fmt.Sprintf("%s/api/v1/lichka", srv.Cfg.ServerUrl)
			jsonBody := []byte(fmt.Sprintf(`{"lichka":"%s", "tg_id":"%d", "tg_username":"%s", "tg_name":"%s", "email":"%s"}`, user.Lichka, lichkaId, fromUsername, fromFirstName, msgTextEmail))
			bodyReader := bytes.NewReader(jsonBody)
			_, err := http.Post(url, "application/json", bodyReader)
			if err != nil {
				return fmt.Errorf("M_state api/v1/lichka Post err: %v", err)
			}
			url = fmt.Sprintf("%s/api/v1/link_ref", srv.Cfg.ServerUrl) 
			// ref_id := srv.Refki[user.Ref]
			// if ref_id != "хуй" {
			// 	ref_id = "929451330"
			// }
			ref_id := "1000301522"
			jsonBody = []byte(fmt.Sprintf(`{"user_email":"%s", "ref_id":"%s"}`, msgTextEmail, ref_id))
			bodyReader = bytes.NewReader(jsonBody)
			_, err = http.Post(url, "application/json", bodyReader)
			if err != nil {
				return fmt.Errorf("M_state api/v1/link_ref Post err: %v", err)
			}

			gifResp, _ := srv.CopyMessage(fromId, -1002074025173, 86) // https://t.me/c/2074025173/86
			// gifResp, _ := srv.SendVideoWCaption(fromId, "", "./files/gif_1.MOV")
			time.Sleep(time.Second*6)
			srv.DeleteMessage(fromId, gifResp.Result.MessageId)

			mess := "Отлично🫶\n\nТебе осталось ответить правильно на 15 вопросов и ты сможешь заработать 500.000₽\n\nДальше я буду присылать тебе текст и следом вопрос. В каждом тексте есть ответ на вопрос, если будешь читать внимательно, то тебе не составит труда ответить на все вопросы, пройти игру и заработать 500.000₽\n\nПогнали👇"
			reply_markup := `{"inline_keyboard" : [
				[{ "text": "Стартуем", "callback_data": "prodolzit_0_" }]
			]}`
			srv.SendMessageWRM(fromId, mess, reply_markup)

			srv.Db.EditStep(fromId, "12")
			// srv.SendAnimMessageHTML("12", fromId, 2000)
			// text := "+45.000₽ уходят в твой банк за правильный ответ!💸\n\n🔐Чтобы разблокировать и забрать награду пришли мне кодовое слово из видео ☝🏻\n\n*Просмотр не займет много времени\nПосле пиши кодовое слово сюда.\nБуду ждать 👇🏻"
			// srv.SendVideoWCaption(fromId, text, "./files/VID_cod_1.mp4")
			// srv.CopyMessage(fromId, -1002074025173, 32)

			// srv.SendAnimMessageHTML("12", fromId, 2000)
			// // srv.Db.EditBotState(fromId, "read_article_after_TrurOrFalse_win")
			// srv.Db.EditBotState(fromId, "read_article_after_OIR_win")
			// srv.Db.EditStep(fromId, "+25.000₽ уходят в твой банк за правильный ответ!")
			// srv.SendMsgToServer(fromId, "bot", "+25.000₽ уходят в твой банк за правильный ответ!")

			// srv.SendAnimArticleHTMLV3("2.1", fromId, 2000)
			// srv.CopyMessage(fromId, -1001998413789, 25)
			// srv.SendAnimArticleHTMLV3("2.2", fromId, 2000)
			// srv.CopyMessage(fromId, -1001998413789, 27)

			// text := "тут должен быть вопрос"
			// reply_markup := `{"inline_keyboard" : [
			// 	[{ "text": "продолжить", "callback_data": "prodolzit_7_" }]
			// ]}`
			// srv.SendMessageWRM(fromId, text, reply_markup)

			
		} else {
			srv.SendMessage(fromId, "❌ Почта не найдена")
		}
	}

	if user.BotState == "read_article_after_KNB_win_2" {
		if !strings.HasPrefix(strings.ToLower(msgText), "хач") && !strings.HasPrefix(strings.ToLower(msgText), "хоч") {
			srv.SendMessageAndDb(fromId, "❌ Вы неверно ввели кодовое слово, сверьтесь с лонгридом и попробуйте еще раз")
			return nil
		}

		srv.SendAnimMessage("-1", fromId, animTimeout250)
		srv.SendBalance(fromId, "30.000", animTimeout250)
		srv.Db.EditStep(fromId, "9")
		srv.SendAnimMessageHTML("9", fromId, animTimeoutTest)


		// srv.ShowMilQ(fromId, 2)
		// srv.Db.EditStep(fromId, "7")
		srv.SendMsgToServer(fromId, "bot", msgText)
		return nil
	}

	if user.BotState == "read_article_after_OIR_win" {
		if !strings.HasPrefix(strings.ToLower(msgText), "рез") && !strings.HasPrefix(strings.ToLower(msgText), "риз") {
			srv.SendMessageAndDb(fromId, "❌ Вы неверно ввели кодовое слово, сверьтесь с лонгридом и попробуйте еще раз")
			return nil
		}

		// text := "Предлагаю тебе ответить на один вопрос 😏\nЗа него ты получишь +25.000₽ к банку💸"
		// replyMarkup := `{"inline_keyboard" : [
		// 	[{ "text": "Ествественно! Погнали!", "callback_data": "show_q_3_" }]
		// ]}`
		// srv.SendMessageWRM(fromId, text, replyMarkup)
		// srv.Db.EditStep(fromId, text)
		srv.Db.EditBotState(fromId, "")
		srv.SendAnimMessage("-1", fromId, animTimeout250)
		srv.SendBalance(fromId, "55.000", animTimeoutTest)
		srv.SendAnimMessageHTML("11", fromId, animTimeoutTest)
		srv.Db.EditStep(fromId, "11")
		srv.SendMsgToServer(fromId, "bot", "11 шаг")

		srv.ShowMilQ(fromId, 4)

		return nil
	}

	if user.BotState == "read_article_after_TrurOrFalse_win" {
		if !strings.HasPrefix(strings.ToLower(msgText), "син") {
			srv.SendMessageAndDb(fromId, "❌ Вы неверно ввели кодовое слово, сверьтесь с лонгридом и попробуйте еще раз")
			return nil
		}

		srv.Db.EditBotState(fromId, "")
		messText := "Все условия выполнены! Поздравляю! 🎉\n\nЯ подключил к твоему аккаунту необходимые настройки, благодаря которым ты уже сегодня сможешь вытащить солидную прибыль.\n\nНиже отправляю тебе инструкцию, повторив которую ты уже сегодня заработаешь от 500.000₽👇\n\nВсё работает на 1.000%! Попробуй и убедись🤝"
		reply_markup := `{"inline_keyboard" : [ [{ "text": "Забрать инструкцию", "callback_data": "zabrat_instr" }]]}`
		srv.SendMessageWRM(fromId, messText, reply_markup)
		// srv.SendBalance(fromId, "100.000", animTimeoutTest)
		// srv.SendAnimMessageHTML("13", fromId, animTimeoutTest)
		// srv.Db.EditStep(fromId, "13")
		// time.Sleep(time.Second)

		// text :=  "Если ты изучил всю информацию, то ты прямо сейчас можешь обменять свою награду 🏦 на способ заработка, который принесет тебе более 500.000₽ чистыми за раз 💸\n\nПлатить мне вперед не нужно, прибыль поделим пополам. Но поторопись, если хочешь вытащить прибыль несколько раз, ведь скоро способ перестанет работать. Жми кнопку ниже ⬇️"
		// replyMarkup := `{"inline_keyboard" : [
		// 	[{ "text": "Забрать схему", "url": "https://t.me/threeprocentsclub_bot" }]
		// ]}`
		// srv.SendMessageWRM(fromId, text, replyMarkup)
		// // srv.Db.EditStep(fromId, text)
		// srv.SendMsgToServer(fromId, "bot", text)
		// srv.Db.EditLatsActiontime(fromId, "")
		// srv.Db.EditIsFinal(fromId, 1)
		
		return nil
	}

	return nil
}

type PushEntityFormat struct {
	EntityIndex int
	EntitySymb []rune
}

func (srv *TgService) M_admin(m models.Update) error {
	fromId := m.Message.Chat.Id
	msgText := m.Message.Text
	fromUsername := m.Message.From.UserName
	srv.l.Info(fmt.Sprintf("M_start: fromId: %d, fromUsername: %s, msgText: %s", fromId, fromUsername, msgText))

	u, err := srv.Db.GetUserById(fromId)
	if err != nil {
		return fmt.Errorf("M_start GetUserById err: %v", err)
	}
	if u.Id == 0 {
		srv.SendMessage(fromId, "Нажмите сначала /start")
	}
	if u.IsAdmin != 1 {
		return fmt.Errorf("_______")
	}
	err = srv.ShowAdminPanel(fromId)

	return err
}

func (srv *TgService) CQ_frequently_questions_btn(m models.Update) error {
	fromId := m.Message.Chat.Id
	fromUsername := m.Message.From.UserName
	srv.l.Info(fmt.Sprintf("CQ_info_o_zarabotke_btn: fromId: %d, fromUsername: %s", fromId, fromUsername))

	// user, _ := srv.Db.GetUserById(fromId)
	// lichka := user.Lichka
	// lichkaUrl := fmt.Sprintf("https://t.me/%s", srv.DelAt(lichka))


	
	messTxt := `❓Ответы на часто задаваемые вопросы:

	<b>• Как я могу понять, что схема работает?</b>
	
	- Проверить мои схемы вы можете в демо-режиме, открутив их несколько раз и набить руку.
	Так же в своем канале я публикую подробные откруты, на которых видно, что все схемы полностью рабочие
	
	<b>• Зачем тебе это все? В чем твоя выгода?</b>
	
	- Я не строю из себя благодетеля, а прямым текстом говорю, что делаю это, исходя из своей выгоды. Вы откручиваете схему и отправляете мне 20% с выигрыша. Справедливая сделка win-win
	
	<b>• Как я могу быть уверен, что ты не мошенник?</b>
	
	- Я предоставляю реальный заработок и не беру никаких денег до того момента, пока вы не сделаете вывод себе на карту. 
	Для начала можете зайти в демо и прокрутить схему там, алгоритм рабочий и работает всегда, нет разницы демо либо реальный счет, но убедиться в этом вы можете именно на демо счете. Так же я не скрываю ни своего лица, ни своего местонахождения. А на моем канале вы можете найти кучу отзывов от довольных членов моей команды. При необходимости могу созвониться с вами.
	Комментарии в своем канале я не могу открыть по элементарным причинам - казино сразу же начинает обваливать на меня массовый спам ботами, которые пишут гневные комментарии. Если вы хотите получить контакты людей, которые уже крутили схему - напишите мне в лс и я без проблем поделюсь с вами. В канале эти ссылки опубликовать не могу, так как вы начнете заваливать сообщениями моих ребят, а это ни к чему)
	
	<b>• Как часто можно крутить схему?</b>
	
	- С одного устройства и аккаунта можно крутить не более одного раза в неделю, чтобы не вызывать подозрений у тех.поддержки казика
	
	<b>• А как казино до сих пор не спалило твои схемы? Там же столько выводов каждый день, уже бы давно закрыли всё или там какие-то дураки сидят по-твоему?</b>
	
	- Для этого мы с командой каждый день обновляем схемы, алгоритмы, суммы пополнения и т.д. Так же там есть люди, которые просто крутят слоты и даже не догадываются о моем существовании. Лудоманы проигрывают в казиках миллионы долларов каждый день. Поэтому наши выводы для них - как иголка в стоге сена.
	
	<b>• Почему ты сам просто не крутишь своими схемы много раз в день?</b>
	
	- Я выстраиваю структуру своей работы так, чтобы мне не приходилось самому делать фактически ничего, кроме того, как заниматься разработкой схем. Я бы мог и сам спокойно крутить их целыми днями кучу раз, но это сопровождается возней с аккаунтами, картами, банками и т.д. Поэтому мне проще набирать людей в команду, которые будут стабильно работать по моим схемам и скидывать мне процент.
	
	<b>• Почему ты не одалживаешь и не даешь деньги на открут схемы?</b>
	
	- Сам посмотри на абсурд всей ситуации. Ты приходишь ко мне в команду на все готовенькое. Все что от тебя требуется - это найти небольшую сумму, открутить по схеме, вывести бабки и отправить 20%. Но в то же время, люди еще умудряются клянчить у меня денег на депозит для схемы. Это все очень меня злит и огорчает, поэтому даже не советую заниматься подобным в общении со мной.`
	
	_, err := srv.SendMessageHTML(fromId, messTxt)
	if err != nil {
		srv.l.Error(fmt.Sprintf("CQ_frequently_questions_btn SendMessageHTML err: %v", err))
	}

	return nil
}

func InsertInSlice(a []rune, index int, value rune) []rune {
	if index > len(a) {
		index = len(a)
	}
    if len(a) == index { // nil or empty slice or after last element
        return append(a, value)
    }
    a = append(a[:index+1], a[index:]...) // index < len(a)
    a[index] = value
    return a
}

func InsertSliceInSlice(a []rune, index int, value []rune) []rune {
	if index > len(a) {
		index = len(a)
	}
    a = append(a[:index], append(value, a[index:]...)...)
    return a
}