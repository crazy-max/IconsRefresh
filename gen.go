// +build ignore

package main

import (
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
		"FileVersion": "{{ .Version }}",
		"InternalName": "",
		"LegalCopyright": "https://github.com/{{ .Repository }}",
		"LegalTrademarks": "",
		"OriginalFilename": "IconsRefresh.exe",
		"PrivateBuild": "",
		"ProductName": "IconsRefresh",
		"ProductVersion": "{{ .Version }}",
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
	if repository, ok = os.LookupEnv("GITHUB_REPOSITORY"); !ok {
		repository = "crazy-max/IconsRefresh"
	}

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
