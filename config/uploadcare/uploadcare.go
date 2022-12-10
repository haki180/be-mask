package uploadcare

import (
	"github.com/uploadcare/uploadcare-go/ucare"
)

func Initialize() (ucare.Client, error) {
	creds := ucare.APICreds{
		SecretKey: "969f3769f7a68443119c",
		PublicKey: "80ede2a9d4ed893fd565",
	}

	conf := &ucare.Config{
		SignBasedAuthentication: true,
		APIVersion:              ucare.APIv06,
	}

	return ucare.NewClient(creds, conf)
}
