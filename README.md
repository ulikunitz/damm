# Package damm

The Go language package damm computes a check digit for a string of
decimal digits. It uses an algorithm proposed by H. Michael Damm in
2004.

## Examples

The function CheckDigit computes and returns it as string.

    c, err := CheckDigit("12345678901")

The argument must contain only decimal digits. Otherwise an error is
reported.

The function Validate verifies digits with the check digit appended.

    ok := Validate("123456789018")

The function returns true if the argument contains only decimal digits and
the appended check digit is correct.

## Algorithm

More information about the algorithm can be found in Wikipedia.

http://en.wikipedia.org/wiki/Damm_algorithm
