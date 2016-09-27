package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	displayBanner()

	// CLI flags
	keyword := flag.String("s", "", "Search keyword")
	limit := flag.Int("l", 25, "Limit results; Max 25")
	categories := flag.String("c", "", categoryHelp())
	nocolor := flag.Bool("b", false, "Black & White; aka no color")
	torrentId := flag.Int64("d", 0, "Download torrent by ID")
	hitandrun := flag.Bool("r", false, "HitAndRun list")
	oneline := flag.Bool("1", false, "Oneline display mode")
	flag.Parse()

	setColorMode(*nocolor)

	username := os.Getenv("NCORE_LOGIN")
	password := os.Getenv("NCORE_PASSWORD")

	if len(username) == 0 || len(password) == 0 {
		fmt.Println("Set up your login credentials with environment variables!")
		fmt.Println("  NCORE_LOGIN=\"nCoreLoginUsername\"")
		fmt.Println("  NCORE_PASSWORD=\"nCoreLoginPassword\"")
		return
	}

	client := NewNcoreClient(username, password)

	if *torrentId > 0 {
		client.Download(*torrentId)
		return
	}

	if *hitandrun {
		torrents := client.HitAndRun()
		for _, torrent := range torrents {
			if *oneline {
				fmt.Println(torrent.ToString())
			} else {
				fmt.Println(torrent.ToStringMultiLine())
			}
		}
		return
	}

	if len(*keyword) > 0 {
		torrents := client.Search(*keyword, *categories, *limit)

		for _, torrent := range torrents {
			if *oneline {
				fmt.Println(torrent.ToString())
			} else {
				fmt.Println(torrent.ToStringMultiLine())
			}
		}

		return
	}

	flag.PrintDefaults()
}

func displayBanner() {
	banner := `
-- Ncore CLI tool for search and download (Go Version)
-- Author: Balazs Nadasdi <yitsushi@gmail.com>
-- Licence: Do what you want, but do not publish directly
`
	fmt.Println(banner)
}

func categoryHelp() string {
	text := `Categories; Coma separated list
           xvid_hun => Film SD/HU            xvid => Film SD/EN
           dvd_hun => Film DVDR/HU           dvd => Film DVDR/EN
           dvd9_hun => Film DVD9/HU          dvd9 => Film DVD9/EN
           hd_hun => Film HD/HU              hd => Film HD/EN

           xvidser_hun => Sorozat SD/HU      xvidser => Sorozat SD/EN
           dvdser_hun => Sorozat DVDR/HU     dvdser => Sorozat DVDR/EN
           hdser_hun => Sorozat HD/HU        hdser => Sorozat HD/EN

           mp3_hun => MP3/HU                 mp3 => MP3/EN
           lossless_hun => Lossless/HU       lossless => Lossless/EN
           clip => Klip

           game_iso => Jatek PC/ISO          game_rip => Jatek PC/RIP
           console => Konzol

           ebook_hun => eBook/HU             ebook => eBook/EN

           iso => APP/ISO                    misc => APP/RIP
           mobil => APP/Mobil

           xxx_xvid => XXX SD                xxx_dvd => XXX DVDR/DVD9
           xxx_imageset => XXX Imageset      xxx_hd => XXX HD`

	return text
}
