package models

import "github.com/youngeek-0410/mottake/server/db"

func Init() {
	db.DB.AutoMigrate(&Customer{}, &Shop{}, &Genre{}, &FavoriteGenre{}, &RelatedGenre{}, &Receipt{}, &Purchase{}, &Menu{})

}

func InitGenres() {
	var genres = []Genre{
		{ID: 1, Name: "日本料理"},
		{ID: 2, Name: "寿司"},
		{ID: 3, Name: "魚介・海鮮料理"},
		{ID: 4, Name: "天ぷら・揚げ物"},
		{ID: 5, Name: "そば・うどん・麺類"},
		{ID: 6, Name: "うなぎ・あなご"},
		{ID: 7, Name: "焼鳥・串焼・鳥料理"},
		{ID: 8, Name: "すき焼き・しゃぶしゃぶ"},
		{ID: 9, Name: "おでん"},
		{ID: 10, Name: "お好み焼き・たこ焼き"},
		{ID: 11, Name: "郷土料理"},
		{ID: 12, Name: "丼もの"},
		{ID: 13, Name: "ステーキ・ハンバーグ"},
		{ID: 14, Name: "パスタ・ピザ"},
		{ID: 15, Name: "ハンバーガー"},
		{ID: 16, Name: "洋食"},
		{ID: 17, Name: "フレンチ"},
		{ID: 18, Name: "イタリアン"},
		{ID: 19, Name: "中華料理"},
		{ID: 20, Name: "餃子・肉まん"},
		{ID: 21, Name: "韓国料理"},
		{ID: 22, Name: "台湾料理"},
		{ID: 23, Name: "インド料理"},
		{ID: 24, Name: "メキシコ料理"},
		{ID: 25, Name: "カレー"},
		{ID: 26, Name: "鍋"},
		{ID: 27, Name: "居酒屋"},
		{ID: 29, Name: "バー"},
		{ID: 30, Name: "ファミレス"},
		{ID: 31, Name: "弁当・おにぎり"},
		{ID: 32, Name: "定食"},
		{ID: 33, Name: "レストラン"},
		{ID: 34, Name: "カフェ"},
		{ID: 35, Name: "喫茶店"},
		{ID: 36, Name: "パン"},
		{ID: 37, Name: "サンドイッチ"},
		{ID: 38, Name: "洋菓子"},
		{ID: 39, Name: "和菓子"},
		{ID: 40, Name: "酒類"},
	}
	db.DB.Save(&genres)

}
