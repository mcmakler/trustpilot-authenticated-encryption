package example

import (
	"fmt"
	"github/mcmakler/trustpilot-authenticated-encryption/trustpilotLinkGen"
)

var (
	//Insert your keys and domain
	authenticationKey = "AAECAwQFBgcICQABAgMEBQYHCAkAAQIDBAUGBwgJAAE="
	encriptionKey     = "AAECAwQFBgcICQABAgMEBQYHCAkAAQIDBAUGBwgJAAE="
	domain            = "domain.com"
)

func ExampleLinkGeneration() {
	//Creating of link generator
	tLG, err := trustpilotLinkGen.NewTrustpilotLinkGenerator(
		encriptionKey,
		authenticationKey,
		domain)
	if err != nil {
		panic("TrustpilotLinkGenerator creating error: " + err.Error())
	}

	//Creating of data with skus and tags
	trustpilotData := &trustpilotLinkGen.TrustpilotUserData{
		Email: "test@email.example",
		Name:  "Name Lastname",
		Ref:   "Reference Number",
		Skus:  []string{"Sku1", "Sku2"}, //Can be empty
		Tags:  []string{"Tag1", "Tag2"}, //Can be empty
	}

	//Generating of a link
	link, err := tLG.GenerateBusinessLink(trustpilotData)
	if err != nil {
		panic("TrustpilotLinkGeneration error: " + err.Error())
	}
	fmt.Println(link)

	//Creating of data without only required fields
	trustpilotShortData := &trustpilotLinkGen.TrustpilotUserData{
		Email: "test@email.example",
		Name:  "Name Lastname",
		Ref:   "Reference Number",
	}

	//Generating of a link
	link, err = tLG.GenerateBusinessLink(trustpilotShortData)
	if err != nil {
		panic("TrustpilotLinkGeneration error: " + err.Error())
	}
	fmt.Println(link)
}
