package main

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func NewTorrentFromLine(torrent *goquery.Selection) *Torrent {
	torrentLink := torrent.Find(".torrent_txt a")
	title, exists := torrentLink.Attr("title")

	if !exists {
		return nil
	}

	href, _ := torrentLink.Attr("href")
	sepIndex := strings.Index(href, "id=") + 3
	torrentId, err := strconv.ParseInt(href[sepIndex:], 10, 32)
	checkErr(err, "nCore::TorrentId parse error!")

	seed, err := strconv.ParseInt(torrent.Find(".box_s2").Text(), 10, 32)
	checkErr(err, "nCore::Seed parse error!")
	leech, err := strconv.ParseInt(torrent.Find(".box_l2").Text(), 10, 32)
	checkErr(err, "nCore::Leech parse error!")

	category, _ := torrent.Find(".box_alap_img img").Attr("alt")

	return &Torrent{
		Id:         torrentId,
		Name:       title,
		UploadedAt: torrent.Find(".box_feltoltve2").Text(),
		Download:   torrent.Find(".box_d2").Text(),
		Size:       torrent.Find(".box_meret2").Text(),
		Seed:       seed,
		Leech:      leech,
		Type:       category,
	}
}
