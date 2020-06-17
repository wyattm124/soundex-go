package main

import "fmt"
import "os"

/*
 * This function will drop any occurance of the character c from the
 * input byteslice and put the result in the out byte slice.
 * The first character is always preserved
 */

func loseCharExceptFirst (in *[]byte, out *[]byte, c byte) {
	var o []byte = *out
	var i []byte = *in
	k := 1
	o[0] = i[0]
	for _, e := range i[1:] {
		if e != c {
			o[k] = e
			k++
		}
	}
}

/* THis function will remove consecutive occurances of a character in the
 * in byte slice and replace the consecutive string of characters with one
 * occurance of the character. for example if in = ['a', 'a', 'b', 'b', 'c']
 * then the output will be out = ['a', 'b', 'c'].
 */

func loseRepeated (in *[]byte, out *[]byte) {
	var o []byte = *out
	var i []byte = *in
	curr := i[0]
	next := i[0]
	c_ind := 0
	n_ind := 0
	for n_ind < len(i) && c_ind < len(o) {
		o[c_ind] = curr
		c_ind++
		for next == curr && n_ind < len(i) - 1 {
			n_ind++
			next = i[n_ind]
		}
		curr = next
	}
}

/* This function will take in the character c and return v if it is a value
 * the number c gets in Soundex, or 'c' if the character that is a consonant
 * that does not have a value in Soundex
 */
func parse (c byte) byte {
	switch c {
	case 'a','A','e','E','i','I','o','O','u','U','y','Y':
		return 'v'
	case 'h','H','w','W':
		return 'c'
	case 'b','B','f','F','p','P','v','V':
		return '1'
	case 'c','C','g','G','j','J','k','K','q','Q','s','S','x','X','z','Z':
		return '2'
	case 'd', 'D', 't', 'T':
		return '3'
	case 'l', 'L':
		return '4'
	case 'm', 'M', 'n', 'N':
		return '5'
	case 'r', 'R':
		return '6'
	default:
		return 'X'
	}
}

/* Sndx computes the Sounex code of in and puts the result in out.
 * I did not follow the algorithim exactly as specified on wiki
 * instead I did what seemed to be more succinct but equivalent :
 *    1. drop the consonants that do not have a Soundex value (except the first character)
 *    2. replace consecutive occurances of a character with one occurance
 *    3. drop all the vowels, except for the first character
 *    4. replace the first character with its real character value instead of its soundex value
 */

func Sndx (in *[]byte, out *[]byte) {
	var i []byte = *in
	var o []byte = *out
	o[0] = '0'
	o[1] = '0'
	o[2] = '0'
	o[3] = '0'
	if len(i) == 0 {
		return
	}
	var first byte = i[0];
	parseOut := make([]byte, len(i))
	for i, e := range i {
		parseOut[i] = parse(e)
	}
	consOut := make([]byte, len(parseOut))
	loseCharExceptFirst(&parseOut, &consOut, 'c')
	repeatOut := make([]byte, len(consOut))
	loseRepeated(&consOut, &repeatOut)
	loseCharExceptFirst(&repeatOut, out, 'v')
	o[0] = first
}

func main () {
	inputs := os.Args[1:]
	if len(inputs) < 1 {
		fmt.Println("Must put in a word to convert to Soundex");
	}
	out := make([]byte, 256)
	for _, e := range inputs {
		in := []byte(e)
		o := out[:len(in)]
		Sndx(&in, &o)
		fmt.Println(string(out[:4]))
	}
}

