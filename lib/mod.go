package lib

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/modder"
)

func getModder(lang string) (modder.Modder, error) {
	// TODO try to detect language by looking for
	// a [lang].mod file
	mod, ok := modder.ModderMap[lang]
	if !ok {
		return nil, fmt.Errorf("Unknown language %q. Add configuration at https://github.com/hofstadter-io/mvs/blob/master/lib/modder/langs.go", lang)
	}

	return mod, nil
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
