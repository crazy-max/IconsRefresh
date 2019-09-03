// +build ignore

package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

var versioninfoTpl = template.Must(template.New("").Parse(`{
	"FixedFileInfo":
	{
		"FileFlagsMask": "3f",
		"FileFlags ": "00",
		"FileOS": "040004",
		"FileType": "01",
		"FileSubType": "00"
	},
	"StringFileInfo":
	{
		"Comments": "",
		"CompanyName": "",
		"FileDescription": "Refresh icons on Desktop, Start Menu and Taskbar",
		"FileVersion": "{{ .Version }}.0",
		"InternalName": "",
		"LegalCopyright": "https://github.com/{{ .Repository }}",
		"LegalTrademarks": "",
		"OriginalFilename": "IconsRefresh.exe",
		"PrivateBuild": "",
		"ProductName": "IconsRefresh",
		"ProductVersion": "{{ .Version }}.0",
		"SpecialBuild": ""
	},
	"VarFileInfo":
	{
		"Translation": {
			"LangID": "0409",
			"CharsetID": "04B0"
		}
	}
}`))

func main() {
	var version, repository string
	var ok bool

	if version, ok = os.LookupEnv("VERSION"); !ok {
		version = "0.0.0.0"
	}
	fmt.Println("gen.version:", version)
	if repository, ok = os.LookupEnv("GITHUB_REPOSITORY"); !ok {
		repository = "crazy-max/IconsRefresh"
	}
	fmt.Println("gen.repository:", repository)

	f, err := os.Create("versioninfo.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = versioninfoTpl.Execute(f, struct {
		Repository string
		Version    string
	}{
		Repository: repository,
		Version:    version,
	})
	if err != nil {
		log.Fatal(err)
	}
}
