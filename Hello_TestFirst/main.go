package main

func main() {

}

const engPref = "hello"

type LangType struct {
	spanish string
	french  string
}

func Languages() LangType {
	const langSpanish = "Spanish"
	const langFrench = "French"

	var languages = LangType{spanish: langSpanish, french: langFrench}

	return languages
}

func Hello(name string, language ...string) string {

	pref := GetLangPref(language...)

	if name == "" {
		name = "world"
	}

	return pref + ", " + name
}

func GetLangPref(language ...string) string {

	pref := engPref
	languages := Languages()

	if len(language) > 0 && language[0] == languages.spanish {
		pref = pref + " in Spanish"
	}
	if len(language) > 0 && language[0] == languages.french {
		pref = pref + " in French"
	}
	return pref
}
