package localization

import (
	"log"

	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"

	validator "gopkg.in/go-playground/validator.v9"
)

func init() {

	pt := pt.New()
	uni = ut.New(pt, pt)

	err := uni.Import(ut.FormatJSON, "translations")
	if err != nil {
		log.Fatal(err)
	}

	err = uni.VerifyTranslations()
	if err != nil {
		log.Fatal(err)
	}

}

var uni *ut.UniversalTranslator
var validate *validator.Validate

func T(language string, key interface{}, params ...string) string {

	trans, _ := uni.GetTranslator(language)

	translation, _ := trans.T(key, params...)

	return translation

}
