package parser

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/utils"
)

// IWatchParser apple watchに関するページのインタフェース
type IWatchParser interface {
	ParseWatchPage() (*model.Watch, error)
}

// WatchParser apple watchに関するページのパーサー
type WatchParser struct {
	Title     string
	AmountStr string
	DetailURL string
}

func loadWatchInformationFromDetailURL(watch *model.Watch, doc *goquery.Document) {
	detail := doc.Find(".as-productinfosection-mainpanel").First()
	detailRegExp, _ := regexp.Compile(`(\n|\s)`)
	detail.Find("div .para-list > p").Each(func(_ int, s *goquery.Selection) {
		text := detailRegExp.ReplaceAllLiteralString(s.Text(), "")
		if strings.Index(text, "月発売") > -1 {
			// 発売年月
			year, _ := strconv.Atoi(text[:4])
			month, _ := strconv.Atoi(text[strings.Index(text, "年"):strings.Index(text, "月")])
			watch.ReleaseDate = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		} else if strings.Index(text, "GB") > -1 {
			strage := strings.Replace(text, "容量", "", 1)
			watch.Strage = strage[:len(strage)-1]
		}
	})
}

// ParseWatchPage watchに関するページをパースして、watchに関する情報のオブジェクトを返却
func (parser *WatchParser) ParseWatchPage() (*model.Watch, error) {
	var watch model.Watch
	// 最初に詳細情報が取ってこれるかを確認
	doc, err := utils.GetGoQueryObject(parser.DetailURL)
	if err != nil {
		log.Errorf("cannot open detail product page. url: %s", parser.DetailURL)
		return nil, err
	}
	// 不要な部分を削除
	nameRegExp, _ := regexp.Compile(`\s*(（.+）|\[.+\])`)
	name := nameRegExp.ReplaceAllLiteralString(parser.Title, "")
	// Cellularモデルかどうか
	if strings.Index(parser.Title, "Cellular") > -1 {
		watch.IsCellular = true
	}
	// 名前・色
	colorRegExp, _ := regexp.Compile(`\d+mm`)
	strs := strings.Split(name, "- ")
	watch.Name = strs[0]
	watch.Color = colorRegExp.ReplaceAllLiteralString(strs[1], "")
	// 金額
	amountRegExp, _ := regexp.Compile(`(,|円（税別）|\s)`)
	watch.Amount, _ = strconv.Atoi(amountRegExp.ReplaceAllLiteralString(parser.AmountStr, ""))

	// その他の情報
	loadWatchInformationFromDetailURL(&watch, doc)
	return &watch, nil
}
