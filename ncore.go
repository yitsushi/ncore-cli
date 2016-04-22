package main

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
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

	resp, err := n.client.PostForm(n.getTargetURI("/torrents.php"), postData)
	checkErr(err, "Connection error!")
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err, "Failed to parse HTML")

	torrents := make([]*Torrent, 0)
	lines := doc.Find(".box_torrent")

	for line := range lines.Nodes {
		torrent := NewTorrentFromLine(lines.Eq(line))
		if torrent == nil {
			continue
		}
		torrents = append(torrents, torrent)
		if limit < 1 {
			break
		}
		limit--
	}

	// torrents := parseTorrentsFromSearch(resp, limit)

	return torrents
}

func (n *Ncore) Download(torrentId int64) {
	resp, err := n.client.Get(
		n.getTargetURI(
			fmt.Sprintf("/torrents.php?action=download&id=%d", torrentId),
		),
	)
	checkErr(err, "Connection Error!")
	defer resp.Body.Close()

	_, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	checkErr(err, "Download error!")

	file, err := os.Create(params["filename"])
	checkErr(err, "File open error!")

	content, err := ioutil.ReadAll(resp.Body)
	checkErr(err, "Parse error!")

	file.Write(content)
	file.Close()

	fmt.Printf("Torrent file save as '%s'.", params["filename"])
}

// Private methods

func (n *Ncore) getTargetURI(path string) string {
	return fmt.Sprintf("https://ncore.cc%s", path)
}
