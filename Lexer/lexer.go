package lexer


type lexer struct {
	input string
	// current position in input (points to current char)
	pos int
	// current reading position in input (after current char)
	readPosition int
	// current char under examination
	currentchar byte
}

func (l lexer) andvanced(){
	l.pos++
	if l.pos < len(l.input){
		l.currentchar = l.input[pos]
	}
}