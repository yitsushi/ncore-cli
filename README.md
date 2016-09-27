# Download latest version

Jump to [Tags](http://gitlab.cheppers.com/yitsushi/ncore-cli/tags/) and
find your binary depends on your operation system / processor architecture.

# How to use

```
  -1  Oneline display mode
  -b  Black & White; aka no color
  -c string
      Categories; Coma separated list
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
           xxx_imageset => XXX Imageset      xxx_hd => XXX HD
  -d int
      Download torrent by ID
  -l int
      Limit results; Max 25 (default 25)
  -r  HitAndRun list
  -s string
      Search keyword
```

## Examples

### Search

```
❯ ncore-cli -s "Doctor Who S08" -1

-- Ncore CLI tool for search and download (Go Version)
-- Author: Balazs Nadasdi <yitsushi@gmail.com>
-- Licence: Do what you want, but do not publish directly

[ 1717908] [         SD/EN] (  14↑     0↓) Doctor.Who.S08.EXTRAS.BDRIP.x264-Krissz
[ 1717907] [         SD/EN] (  86↑    14↓) Doctor.Who.S08.BDRIP.x264-Krissz
[ 1684087] [         HD/EN] (  10↑     3↓) Doctor.Who.2005.S08.1080p.WEB-DL.DD5.1.H.264-ECI
[ 1683915] [         HD/EN] (  15↑     8↓) Doctor.Who.2005.S08.720p.HDTV.x264-MIXGROUP
[ 1683873] [         SD/EN] (   4↑     3↓) Doctor.Who.2005.S08.HDTV.XviD-MIXDROUP
[ 1683778] [         SD/EN] (  25↑     4↓) Doctor.Who.S08.WEB-DL.x264-Krissz
[ 1683777] [         SD/EN] (   4↑     0↓) Doctor.Who.2005.S08.HDTV.x264-MIXGROUP
[ 1633817] [         SD/EN] (   2↑     0↓) Doctor Who S08 - Live Preshow
[ 1633815] [         SD/EN] (   2↑     0↓) Doctor Who S08 - After Who Live
```

### Download

```
❯ ncore-cli -d 1684087

-- Ncore CLI tool for search and download (Go Version)
-- Author: Balazs Nadasdi <yitsushi@gmail.com>
-- Licence: Do what you want, but do not publish directly

Torrent file save as '[nCore][hdser]Doctor.Who.2005.S08.1080p.WEB-DL.DD5.1.H.264-ECI.torrent'.
```

# Build

```
make build
```

Done. :D

# Collaborate

Just send a PR and we can talk about it. Or just simply create new issues about
your wishes :)
