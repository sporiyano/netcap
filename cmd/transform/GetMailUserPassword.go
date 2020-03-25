package main

import (
	maltego "github.com/dreadl0ck/netcap/maltego"
	"github.com/dreadl0ck/netcap/types"
)

func GetMailUserPassword() {
	maltego.POP3Transform(
		nil,
		func(lt maltego.LocalTransform, trx *maltego.MaltegoTransform, pop3  *types.POP3, min, max uint64, profilesFile string, ipaddr string) {
			if pop3.Client == ipaddr {
				user := lt.Value
				if pop3.User == user && pop3.Pass != "" {
					escapedName := maltego.EscapeText(pop3.Pass)
					ent := trx.AddEntity("maltego.Password", escapedName)
					ent.SetType("maltego.Password")
					ent.SetValue(escapedName)

					// di := "<h3>Mail User</h3><p>Timestamp First: " + pop3.Timestamp + "</p>"
					// ent.AddDisplayInformation(di, "Netcap Info")
					ent.SetLinkColor("#000000")
					//ent.SetLinkThickness(maltego.GetThickness(uint64(count), min, max))

					ent.AddProperty("ipaddr", "IPAddress", "strict", ipaddr)
					ent.AddProperty("path", "Path", "strict", profilesFile)
				}
			}
		},
	)
}