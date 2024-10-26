package generator

type Config struct {
	Length     int    `yaml:"length" default:"5"`
	Characters string `yaml:"characters" default:"abcdefghijklmnopqrstuvwxyz"`
	ButchSize  int    `yaml:"butch_size" default:"100"`
}

func generatePasswords(characters []rune, length int, currentPassword []rune, ch chan<- string) {
	if length == 0 {
		ch <- string(currentPassword)
		return
	}

	for _, char := range characters {
		generatePasswords(characters, length-1, append(currentPassword, char), ch)
	}
}

func GeneratePasswords(cfg *Config) <-chan string {
	ch := make(chan string, cfg.ButchSize)
	go func() {
		defer close(ch)
		generatePasswords([]rune(cfg.Characters), cfg.Length, []rune{}, ch)
	}()
	return ch
}
