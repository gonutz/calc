package calc

import (
	"fmt"
	"strconv"
	"strings"
)

// Calculator accepts user input and holds the state of the calculation.
// Use Input to update the equation, input "=" to have the result calculated.
// This will solve the whole equation, multiplication and division before
// addition and subtraction.
type Calculator struct {
	short, long string
	lastWasOp   bool
	nums, ops   []string
}

// NewCalculator returns a fresh Calculator. Its inital state is a "0" as the
// short output, and "" as the long output.
func NewCalculator() *Calculator {
	var c Calculator
	c.reset()
	return &c
}

func (c *Calculator) reset() {
	c.short = "0"
	c.long = ""
	c.lastWasOp = false
	c.nums = nil
	c.ops = nil
}

// LongOutput returns the equation entered so far. It does not contain the
// number/operator that is currently being edited, that is contained in the
// ShortOutput().
func (c *Calculator) LongOutput() string {
	return c.long
}

// ShortOutput returns the number or operator that is currently being edited.
// The previous equation inputs can be retrieved with LongOutput().
func (c *Calculator) ShortOutput() string {
	return c.short
}

const divByZeroErr = "Error: div by 0"

// Input accepts the following characters:
//
// "0123456789." to edit the current number
// "+-*/"        for basic math operations
// "="           to display the result of the current equation
// "C"           to clear the current calculation
// "N"           to negate the current number
//
// After each input the short and long outputs are updated.
func (c *Calculator) Input(r rune) {
	const validInputs = "0123456789.+-*/=CN"
	if !strings.ContainsRune(validInputs, r) {
		return
	}

	if c.short == divByZeroErr {
		c.reset()
	}

	if r == 'C' {
		c.reset()
		return
	}

	if (r == '+' || r == '*' || r == '/') && c.short == "0" && c.long == "" {
		return
	}

	if r == 'N' {
		if strings.HasSuffix(c.long, "=") {
			// at the end of a calculation, reset the previous equation
			c.long = ""
			c.lastWasOp = false
		}
		if !(c.short == "0" || c.lastWasOp) {
			c.short = negate(c.short)
		}
		return
	}

	if c.short == "0" && r == '-' {
		c.short = "-"
		return
	}

	if strings.HasSuffix(c.long, "=") {
		// if after a calculation we continue
		// - with an operator, we use the last calculation result as the first
		//   number
		// - with a number, we forget the calculation result and start a new
		//   calculation
		if isOp(r) {
			c.long = c.short
			c.nums = []string{c.short}
		} else {
			c.long = ""
		}
		c.short = string(r)
	} else if isOp(r) && c.lastWasOp {
		// replace the operator with the new one
		c.short = string(r)
	} else if isOp(r) && !c.lastWasOp {
		// operator after number, append the number to long string and set the
		// short string to the operator
		c.short = strings.TrimSuffix(c.short, ".")
		c.nums = append(c.nums, c.short)
		c.long += c.short
		c.short = string(r)
	} else if !isOp(r) && c.lastWasOp {
		// number after operator, append operator to long and set short to new
		// number
		c.ops = append(c.ops, c.short)
		c.long += c.short
		c.short = string(r)
	} else if !isOp(r) && !c.lastWasOp {
		// continue writing the number
		if c.short == "0" {
			c.short = ""
		}
		if !(r == '.' && strings.Contains(c.short, ".")) {
			c.short += string(r)
		}
	}

	if c.short == "." {
		c.short = "0."
	}

	if r == '=' {
		c.solve()
	}

	c.lastWasOp = isOp(r)
}

func isOp(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/' || r == '='
}

func negate(s string) string {
	if strings.HasPrefix(s, "-") {
		return s[1:]
	} else {
		return "-" + s
	}
}

func (c *Calculator) solve() {
	c.long += "="
	// solveFor returns false on error.
	solveFor := func(op1, op2 string) bool {
		for i := 0; i < len(c.ops); i++ {
			op := c.ops[i]
			if op == op1 || op == op2 {
				a, b := toNum(c.nums[i]), toNum(c.nums[i+1])
				// remove this op from c.ops
				copy(c.ops[i:], c.ops[i+1:])
				c.ops = c.ops[:len(c.ops)-1]
				// remove the second number from c.nums, the first will be
				// overwritten with the result
				copy(c.nums[i+1:], c.nums[i+2:])
				c.nums = c.nums[:len(c.nums)-1]
				// solve this op
				var result float64
				if op == "*" {
					result = a * b
				} else if op == "/" {
					if b == 0 {
						c.short = divByZeroErr
						return false
					}
					result = a / b
				} else if op == "+" {
					result = a + b
				} else if op == "-" {
					result = a - b
				}
				c.nums[i] = toString(result)
				i-- // we remove one op and num, next is the same index again
			}
		}
		return true
	}
	if !solveFor("*", "/") {
		return
	}
	solveFor("+", "-")
	c.short = c.nums[0]
	c.nums = nil
	c.ops = nil
}

func toNum(s string) float64 {
	n, _ := strconv.ParseFloat(s, 64)
	return n
}

func toString(n float64) string {
	s := fmt.Sprintf("%f", n)
	// trim trailing 0s and, if a single decimal point remains, trim that
	for len(s) > 0 && strings.HasSuffix(s, "0") {
		s = s[:len(s)-1]
	}
	s = strings.TrimSuffix(s, ".")
	return s
}
