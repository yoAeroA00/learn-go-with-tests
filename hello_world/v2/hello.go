package hello

import "fmt"

var englishHelloPrefix = "Hello, "

const (
	hiLang           = "Hindi"
	englishEndSuffix = "!"
)

func Hello(name, language string) string {
	if name == "" {
		englishHelloPrefix, name = chooseLanguage(language)
	}
	return englishHelloPrefix + name + englishEndSuffix
}

func chooseLanguage(language string) (ehp string, n string) {
	switch language {
	case hiLang:
		ehp = "Namaste, "
		n = "Duniya"
	default:
		ehp = englishHelloPrefix
		n = "World"
	}
	return
}

func main() {
	fmt.Println(Hello("Adi", ""))
}
