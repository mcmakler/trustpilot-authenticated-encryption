# trustpilot-authenticated-encryption

Go implementation for Trustpilot Business Generated Links:

https://support.trustpilot.com/hc/en-us/articles/115004145087--Business-Generated-Links-for-developers-

## Usage

Creating of TrustpilotLinkGenerator interface exemplar:

```go
tLG, err := trustpilotLinkGen.NewTrustpilotLinkGenerator(
	encriptionKey,      //your base64 encoded encryption key 
	authenticationKey,  //your base64 encoded description key
	domain)             //your domain
```

Creating of payload:

```go
trustpilotData := &trustpilotLinkGen.TrustpilotUserData{
		Email: "test@email.example",
		Name:  "Name Lastname",
		Ref:   "Reference Number",
		Skus:  []string{"Sku1", "Sku2"}, //Can be omitted
		Tags:  []string{"Tag1", "Tag2"}, //Can be omitted
	}
```

Generating of a link:

```go
link, err := tLG.GenerateBusinessLink(trustpilotData)
```