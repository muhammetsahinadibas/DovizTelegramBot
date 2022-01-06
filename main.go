package main

// Developer: Muhammet Sahin Adibas
// Twitter: twitter.com/muhammetadibas
// Blog: muhammetsahinadibas.com.tr
// Github: github.com/muhammetsahinadibas

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var usd, eur, gbp, grGold string

func currency() (string, string, string, string) {
	url := "https://www.doviz.com/"

	res, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	if res.StatusCode != 200 {
		log.Panic("Status Code Error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Panic(err)
	}

	doc.Find(".market-data .item .value").Each(func(index int, item *goquery.Selection) {
		dataKey, _ := item.Attr("data-socket-key")

		switch dataKey {
		case "USD":
			usd = item.Text()
		case "EUR":
			eur = item.Text()
		case "GBP":
			gbp = item.Text()
		case "gram-altin":
			grGold = item.Text()
		}

	})

	return usd, eur, gbp, grGold
}

func sendMessage() {
	currency()

	BotToken := ""  // Telegram Bot Token (1234567890:ABC-DEF1234ghIkl-zyx57W2v1u123ew11)
	ChannelID := "" // Telegram Channel ID (-1234567890)

	messageText := "➔ Dolar: " + usd + " ₺" +
		"\n➔ Euro: " + eur + " ₺" +
		"\n➔ İngiliz Sterlini: " + gbp + " ₺" +
		"\n➔ Gram Altın: " + grGold + " ₺" +
		"\n📍 @kurfiyatlari"

	data := url.Values{
		"chat_id": {ChannelID},
		"text":    {messageText},
	}

	http.PostForm("https://api.telegram.org/bot"+BotToken+"/sendMessage", data)
}

func main() {

	for {
		hour, minute, _ := time.Now().Clock()
		clock := strconv.Itoa(hour) + ":" + strconv.Itoa(minute)

		for i := 0; i <= 24; i++ {
			newClock := strconv.Itoa(i) + ":" + "0"
			if clock == newClock {
				sendMessage()
				time.Sleep(3480 * time.Second)
			}
		}

		time.Sleep(30 * time.Second)
	}

}
