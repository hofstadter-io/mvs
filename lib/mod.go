package lib

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/modder"
)

func getModder(lang string) (*modder.Modder, error) {
	// TODO try to detect language by looking for
	// a [lang].mod file
	mod, ok := LangModderMap[lang]
	if !ok {
		return nil, fmt.Errorf("Unknown language %q. Add configuration at https://github.com/hofstadter-io/mvs/blob/master/lib/modder/langs.go", lang)
	}

	return mod, nil
}

// This is a convienence function for calling the other mod functions with a list of languages
func ProcessLangs(method string, langs []string) error {

	// discover and update slice
	if len(langs) == 0 {
		langs = DiscoverLangs()
	}

	var err error

	for _, lang := range langs {
		switch method {
		case "graph":
			err = Graph(lang)
		case "tidy":
			err = Tidy(lang)
		case "vendor":
			err = Vendor(lang)
		case "verify":
			err = Verify(lang)
		default:
			panic("unimplemented language in ProcessLangs for " + lang)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func Init(lang, module string) error {
	mod, err := getModder(lang)
	if err != nil {
		return err
	}
	return mod.Init(module)
}

func Graph(lang string) error {
	mod, err := getModder(lang)
	if err != nil {
		return err
	}
	return mod.Graph()
}

func Tidy(lang string) error {
	mod, err := getModder(lang)
	if err != nil {
		return err
	}
	return mod.Tidy()
}

func Vendor(lang string) error {
	// TODO, if lang == "" { look for all and process }
	mod, err := getModder(lang)
	if err != nil {
		return err
	}
	return mod.Vendor()
}

func Verify(lang string) error {
	mod, err := getModder(lang)
	if err != nil {
		return err
	}
	return mod.Verify()
}
