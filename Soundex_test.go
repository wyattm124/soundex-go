package main

import "testing"

/* This is a unit test to make sure that the implementation of Soundex
 * works on the test cases given in the American Soundex wiki page
 */
func TestSndx(t *testing.T) {
	out := make([]byte, 16)
	case1 := []byte{'R', 'o', 'b', 'e', 'r', 't'}
	case2 := []byte{'R', 'u', 'p', 'e', 'r', 't'}
	case3 := []byte{'R', 'u', 'b', 'i', 'n'}
	case4 := []byte{'A', 's', 'h', 'c', 'r', 'a', 'f', 't'}
	case5 := []byte{'A', 's', 'h', 'c', 'r', 'o', 'f', 't'}
	case6 := []byte{'t', 'y', 'm', 'c', 'z', 'a', 'k'}
	case7 := []byte{'p', 'f', 'i', 's', 't', 'e', 'r'}
	case8 := []byte{'H', 'o', 'n', 'e', 'y', 'm', 'a', 'n'}
	ans1 := "R163"
	ans2 := "R150"
	ans3 := "A261"
	ans4 := "t522"
	ans5 := "p236"
	ans6 := "H555"

        Sndx(&case1, &out)
	if string(out[:4]) != ans1 {
		t.Errorf("For input %s output %s != %s", string(case1), string(out[:4]), ans1)
	}

	Sndx(&case2, &out)
	if string(out[:4]) != ans1 {
		t.Errorf("For input %s output %s != %s", string(case2), string(out[:4]), ans1)
	}

	Sndx(&case3, &out)
	if string(out[:4]) != ans2 {
		t.Errorf("For input %s output %s != %s", string(case3), string(out[:4]), ans2)
	}

	Sndx(&case4, &out)
	if string(out[:4]) != ans3 {
		t.Errorf("For input %s output %s != %s", string(case4), string(out[:4]), ans3)
	}

	Sndx(&case5, &out)
	if string(out[:4]) != ans3 {
		t.Errorf("For input %s output %s != %s", string(case5), string(out[:4]), ans3)
	}

	Sndx(&case6, &out)
	if string(out[:4]) != ans4 {
		t.Errorf("For input %s output %s != %s", string(case6), string(out[:4]), ans4)
	}

	Sndx(&case7, &out)
	if string(out[:4]) != ans5 {
		t.Errorf("For input %s output %s != %s", string(case7), string(out[:4]), ans5)
	}

	Sndx(&case8, &out)
	if string(out[:4]) != ans6 {
		t.Errorf("For input %s output %s != %s", string(case8), string(out[:4]), ans6)
	}
}
