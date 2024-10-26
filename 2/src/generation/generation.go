package generation

type Generator struct {
	length int
	chars  []rune
}

func (g *Generator) Generate(position int) string {
	length := g.length
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = g.chars[0]
	}
	i := 0
	base := len(g.chars)
	for position > 0 {
		mod := position % base
		position = position / base
		runes[length-i-1] = g.chars[mod]
		i++
	}
	return string(runes)
}

func New(chars string, length int) *Generator {
	return &Generator{
		chars:  []rune(chars),
		length: length,
	}
}
