package main

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"golang.org/x/net/publicsuffix"
)

type Ncore struct {
	Username string
	Password string

	client *http.Client
}

func NewNcoreClient(username, password string) *Ncore {
	client := &Ncore{Username: username, Password: password}
	client.Connect()

	return client
}

func (n *Ncore) Connect() {
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, _ := cookiejar.New(&options)
	n.client = &http.Client{Jar: jar}
	n.Login()
}

func (n *Ncore) Login() {
	postData := url.Values{
		"set_lang":        {"hu"},
		"submitted":       {"1"},
		"nev":             {n.Username},
		"pass":            {n.Password},
		"ne_leptessen_ki": {"1"},
	}

	n.client.PostForm(n.getTargetURI("/login.php"), postData)
}

func (n *Ncore) Search(term string, categories string, limit int) []*Torrent {
	var categoryArray []string

	postData := url.Values{
		"mire": {term},
	}

	if len(categories) > 0 {
		categoryArray = strings.Split(categories, ",")
		postData.Add("tipus", "kivalasztottak_kozott")
	}

	for _, category := range categoryArray {
		postData.Add("kivalasztott_tipus[]", category)
	}

	resp, _ := n.client.PostForm(n.getTargetURI("/torrents.php"), postData)
	torrents := parseTorrentsFromSearch(resp, limit)

	return torrents
}

func (n *Ncore) Download(torrentId int64) {
	resp, _ := n.client.Get(
		n.getTargetURI(
			fmt.Sprintf("/torrents.php?action=download&id=%d", torrentId),
		),
	)
	defer resp.Body.Close()

	_, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	checkErr(err, "Download error!")

	file, err := os.Create(params["filename"])
	checkErr(err, "File open error!")

	content, _ := ioutil.ReadAll(resp.Body)

	file.Write(content)
	file.Close()

	fmt.Printf("Torrent file save as '%s'.", params["filename"])
}

// Private methods

func (n *Ncore) getTargetURI(path string) string {
	return fmt.Sprintf("https://ncore.cc%s", path)
}

func parseTorrentsFromSearch(resp *http.Response, limit int) []*Torrent {
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err, "Failed to parse HTML")

	torrents := make([]*Torrent, 0)

	lines := doc.Find(".box_torrent")
	for line := range lines.Nodes {
		torrent := lines.Eq(line)
		torrentLink := torrent.Find(".torrent_txt a")
		title, exists := torrentLink.Attr("title")

		if !exists {
			continue
		}

		if limit < 1 {
			break
		}
		limit--

		href, _ := torrentLink.Attr("href")
		sepIndex := strings.Index(href, "id=") + 3
		torrentId, _ := strconv.ParseInt(href[sepIndex:], 10, 32)

		seed, _ := strconv.ParseInt(torrent.Find(".box_s2").Text(), 10, 32)
		leech, _ := strconv.ParseInt(torrent.Find(".box_l2").Text(), 10, 32)

		category, _ := torrent.Find(".box_alap_img img").Attr("alt")

		torrents = append(
			torrents,
			&Torrent{
				Id:         torrentId,
				Name:       title,
				UploadedAt: torrent.Find(".box_feltoltve2").Text(),
				Download:   torrent.Find(".box_d2").Text(),
				Size:       torrent.Find(".box_meret2").Text(),
				Seed:       seed,
				Leech:      leech,
				Type:       category,
			},
		)
	}

	return torrents
}
