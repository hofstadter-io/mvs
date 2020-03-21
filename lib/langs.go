package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"

	"cuelang.org/go/cue"
	"github.com/hofstadter-io/mvs/lib/langs"
	"github.com/hofstadter-io/mvs/lib/modder"
	"github.com/hofstadter-io/mvs/lib/util"
)

// FROM the USER's HOME dir
const GLOBAL_MVS_CONFIG = ".mvs/config.cue"
const LOCAL_MVS_CONFIG = ".mvsconfig.cue"

var (
	// Default known modderr
	LangModderMap = langs.DefaultModders
)

const knownLangMessage = `
Known Languages:

  %s

For more info on a language:

  mvs info <lang>
`

func DiscoverLangs() (langs []string) {

	for lang, mdr := range LangModderMap {
		// Let's check for a custom
		_, err := os.Lstat(mdr.ModFile)
		if err != nil {
			if _, ok := err.(*os.PathError); !ok && err.Error() != "file does not exist" {
				fmt.Println(err)
				// return err // more of a warning right now
			}
			// file not found
			continue
		}
		// we found a mod file
		langs = append(langs, lang)
	}

	return langs
}

func KnownLangs() string {
	langs := []string{}

	for lang, _ := range LangModderMap {
		langs = append(langs, lang)
	}

	sort.Strings(langs)
	langStr := strings.Join(langs, "\n  ")

	msg := fmt.Sprintf(knownLangMessage, langStr)

	return msg
}

const unknownLangMessage = `
Unknown language %q.

Please check the following files for definitions
  %s  (in the current directory)
	$HOME/%s

To see a list of known languages from the current directory:

  mvs info
`

func LangInfo(lang string) (string, error) {

	if lang == "" {
		return KnownLangs(), nil
	}

	modder, ok := LangModderMap[lang]
	if !ok {
		return "", fmt.Errorf(unknownLangMessage, lang, LOCAL_MVS_CONFIG, GLOBAL_MVS_CONFIG)
	}

	// TODO output as cue
	bytes, err := yaml.Marshal(modder)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func InitLangs() {
	var err error
	homedir := util.UserHomeDir()

	// Global Language Modder Config
	err = initFromFile(path.Join(homedir, GLOBAL_MVS_CONFIG))
	if err != nil {
		fmt.Println(err)
	}

	// Local Language Modder Config
	err = initFromFile(LOCAL_MVS_CONFIG)
	if err != nil {
		fmt.Println(err)
	}

}

func initFromFile(filepath string) error {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		if _, ok := err.(*os.PathError); !ok && err.Error() != "file does not exist" {
			return err
		}
		// The user has not setup a global $HOME/.mvs/mvsconfig file
		return nil
	}

	var mdrMap map[string]*modder.Modder

	var r cue.Runtime
	i, err := r.Compile(filepath, string(bytes))
	if err != nil {
		return err
	}
	err = i.Value().Decode(&mdrMap)
	if err != nil {
		return err
	}

	for lang, _ := range mdrMap {
		mdr := mdrMap[lang]
		LangModderMap[lang] = mdr
	}

	return nil
}
