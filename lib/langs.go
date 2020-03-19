package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/hofstadter-io/mvs/lib/modder"
	"github.com/hofstadter-io/mvs/lib/modder/custom"
	"github.com/hofstadter-io/mvs/lib/modder/exec"
	"github.com/hofstadter-io/mvs/lib/util"
)

// FROM the USER's HOME dir
const GLOBAL_MVS_CONFIG = ".mvs/config.yaml"
const LOCAL_MVS_CONFIG = ".mvsconfig.yaml"

var (
	// Default known modderr
	LangModderMap = map[string]modder.Modder{
		"go":  GolangModder,
		"cue": CuelangModder,
		"hof": HoflangModder,
	}
	// TODO, add custom Modders here (for simple) read from a ./.mvsconfig file

	// Common files to copy from modules, also includes the .md version of the filename
	CommonCopies = []string{
		"README",
		"README.md",
		"LICENSE",
		"LICENSE.md",
		"PATENTS",
		"PATENTS.md",
		"CONTRIBUTORS",
		"CONTRIBUTORS.md",
		"SECURITY",
		"SECURITY.md",
	}

	GolangModder = &exec.Modder{
		Name:     "go",
		Version:  "1.14",
		ModFile:  "go.mod",
		SumFile:  "go.sum",
		ModsDir:  "vendor",
		Checksum: "vendor/modules.txt",
		Commands: map[string][]string{
			"init":   []string{"go", "mod", "init"},
			"graph":  []string{"go", "mod", "graph"},
			"tidy":   []string{"go", "mod", "tidy"},
			"vendor": []string{"go", "mod", "vendor"},
			"verify": []string{"go", "mod", "verify"},
		},
	}

	CuelangModder = &custom.Modder{
		Name:     "cue",
		Version:  "0.0.15",
		ModFile:  "cue.mods",
		SumFile:  "cue.sums",
		ModsDir:  "cue.mod/pkg",
		Checksum: "cue.mod/modules.txt",
		InitTemplates: map[string]string{
			"cue.mod/module.cue": `module: "{{ .Module }}"
`,
		},
		VendorIncludeGlobs: []string{
			"cue.mods",
			"cue.sums",
			"cue.mod/module.cue",
			"cue.mod/modules.txt",
			"**/*.cue",
		},
		VendorExcludeGlobs: []string{
			"cue.mod/pkg",
		},
	}

	HoflangModder = &custom.Modder{
		Name:     "hof",
		Version:  "0.0.0",
		ModFile:  "hof.mods",
		SumFile:  "hof.sums",
		ModsDir:  "hof.mod/pkg",
		Checksum: "hof.mod/modules.txt",
		InitTemplates: map[string]string{
			"hof.mod/module.cue": `module: "{{ .Module }}"
`,
		},
	}
)

const knownLangMessage = `
Known Languages:

  %s

For more info on a language:

  mvs info <lang>
`

func DiscoverLangs() (langs []string) {

	for lang, mdrI := range LangModderMap {
		mdr, ok := mdrI.(*custom.Modder)
		if ok {
			// Let's check for a custom
			_, err := os.Lstat(mdr.ModFile)
			if err != nil {
				if _, ok := err.(*os.PathError); !ok {
					fmt.Println(err)
					// return err
				}
				// file not found, but error
				continue
			}
			// we found a mod file
			langs = append(langs, lang)
		} else {

			// Let's try an exev modder
			emdr, ok := mdrI.(*exec.Modder)
			if ok {
				_, err := os.Lstat(emdr.ModFile)
				if err != nil {
					if _, ok := err.(*os.PathError); !ok {
						fmt.Println(err)
						// return err
					}
					// file not found, but error
					continue
				}
				// we found a mod file
				langs = append(langs, lang)
			}
			continue
		}

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
		if _, ok := err.(*os.PathError); !ok {
			return err
		}
		// The user has not setup a global $HOME/.mvs/mvsconfig file
		return nil
	}

	var mdrMap map[string]custom.Modder
	err = yaml.Unmarshal(bytes, &mdrMap)
	if err != nil {
		return err
	}

	for lang, _ := range mdrMap {
		mdr := mdrMap[lang]
		LangModderMap[lang] = &mdr
	}

	return nil
}
