package lexer

import (
	"strconv"
)


type Lexer struct {
	exp       string
	expLength int
	number    int
	index     int
}

func ConstructLexer(exp string) Lexer {
	return Lexer{exp: exp, expLength: len(exp), index: 0}
}

func (lexer *Lexer) GetToken() Token {
	token := ILLEGAL_TOKEN

	if lexer.index == lexer.expLength {
		return TOK_EOF
	}

	for lexer.getcurrentByte() == ' ' || lexer.getcurrentByte() == '\t' {
		lexer.index = lexer.index + 1
	}

	if lexer.index == lexer.expLength {
		return TOK_EOF
	}

	currentByte := lexer.getcurrentByte()
	if currentByte == '+' {
		lexer.index++
		token = TOK_PLUS
	} else if currentByte == '-' {
		lexer.index++
		token = TOK_SUB
	} else if currentByte == '/' {
		lexer.index++
		token = TOK_DIV
	} else if currentByte == '*' {
		lexer.index++
		token = TOK_MUL
	} else if currentByte == '(' {
		lexer.index++
		token = TOK_OPAREN
	} else if currentByte == ')' {
		lexer.index++
		token = TOK_CPAREN
	} else if currentByte == '.' {
		lexer.index++
		token = TOK_DOT
	} else if checkIfANumber(currentByte) {
		token = TOK_NUM
		var numberByteList []byte

		numberByteList = append(numberByteList, currentByte)
		lexer.index++
		if lexer.index < lexer.expLength {
			currentByte = lexer.getcurrentByte()
			for checkIfANumber(currentByte) {
				numberByteList = append(numberByteList, currentByte)
				lexer.index++
				if lexer.index == lexer.expLength {
					break
				}
				currentByte = lexer.getcurrentByte()
			}
		}
		lexer.number, _ = strconv.Atoi(string(numberByteList))
	}
	return token
}

func (lexer *Lexer) GetNumber() int {
	return lexer.number
}

func (lexer *Lexer) getcurrentByte() byte {
	return lexer.exp[lexer.index]
}

func checkIfANumber(currentByte byte) bool {
	return currentByte == '0' || currentByte == '1' || currentByte == '2' || currentByte == '3' || currentByte == '4' || currentByte == '5' || currentByte == '6' || currentByte == '7' || currentByte == '8' || currentByte == '9'
}
