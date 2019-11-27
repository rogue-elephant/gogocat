# GoGoCat
[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/gogocat)

[![gogocat](https://snapcraft.io//gogocat/badge.svg)](https://snapcraft.io/gogocat)

```$sudo snap install gogocat```

Then

```$curl https://raw.githubusercontent.com/rogue-elephant/gogocat/master/readme.md | gogocat```

![](https://img.shields.io/github/languages/code-size/rogue-elephant/gogocat)
![](https://img.shields.io/github/release-date/rogue-elephant/gogocat)
![](https://img.shields.io/github/last-commit/rogue-elephant/gogocat)
![](https://img.shields.io/github/issues-raw/rogue-elephant/gogocat)
![](https://img.shields.io/github/issues-closed-raw/rogue-elephant/gogocat)
![](https://img.shields.io/badge/using-golang-008866?style=flat&logo=go)

pipe CLI utility for prettifying output.
Based on the excellent tutorial here https://flaviocopes.com/go-tutorial-lolcat/

## Use
```cat main.go | docker run -i rogueelephant/gogocat```


## Example
_(Assuming you have cloned the directory and are in the top folder context)_
```$cat main.go | go run main.go```

```$cat main.go | docker run -i gogocat```

![](readme.gif)