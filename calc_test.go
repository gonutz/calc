package calc_test

import (
	"testing"

	"github.com/gonutz/calc"
)

func makeCalcTest(t *testing.T) func(input, longOutput, shortOutput string) {
	c := calc.NewCalculator()
	return func(input, longOutput, shortOutput string) {
		t.Helper()
		for _, r := range input {
			c.Input(r)
		}
		if c.LongOutput() != longOutput || c.ShortOutput() != shortOutput {
			t.Errorf(
				"want %q %q but have %q %q",
				longOutput, shortOutput,
				c.LongOutput(), c.ShortOutput(),
			)
		}
	}
}

func TestTwoNumbersCanBeAdded(t *testing.T) {
	inout := makeCalcTest(t)
	inout("", "", "0")
	inout("1", "", "1")
	inout("+", "1", "+")
	inout("2", "1+", "2")
	inout("=", "1+2=", "3")
}

func TestMulAndDivComeBeforePlusAndMinus(t *testing.T) {
	inout := makeCalcTest(t)
	inout("1+2*3=", "1+2*3=", "7")
}

func TestEqualsAfterOnlyANumberResultsInThatNumber(t *testing.T) {
	makeCalcTest(t)("=", "0=", "0")
	makeCalcTest(t)("1=", "1=", "1")
}

func TestLastEnteredOperatorIsUsed(t *testing.T) {
	makeCalcTest(t)("5+*-7=", "5-7=", "-2")
}

func TestFirstNumberCanBeNegative(t *testing.T) {
	makeCalcTest(t)("-22=", "-22=", "-22")
}

func TestFirstNumberCanBePositive(t *testing.T) {
	inout := makeCalcTest(t)
	inout("+", "", "0")
	inout("3=", "3=", "3")
}

func TestDivByZeroIsError(t *testing.T) {
	makeCalcTest(t)("1/0=", "1/0=", "Error: div by 0")
}

func TestDecimalsCanBeUsed(t *testing.T) {
	inout := makeCalcTest(t)
	inout("1.", "", "1.")
	inout("5", "", "1.5")
	inout(".6", "", "1.56") // second decimal point is ignored
	inout("*2.0=", "1.56*2.0=", "3.12")
}

func TestLastOperatorIsIgnoredIfNotFollowedByNumber(t *testing.T) {
	makeCalcTest(t)("1+2+=", "1+2=", "3")
}

func TestNumberAfterEqualsResetsCalculation(t *testing.T) {
	inout := makeCalcTest(t)
	inout("1+2=", "1+2=", "3")
	inout("50", "", "50")
}

func TestOperatorAfterEqualsContinuesWithLastResult(t *testing.T) {
	inout := makeCalcTest(t)
	inout("1+2=", "1+2=", "3")
	inout("*", "3", "*")
}

func TestLargeNumbersDoNotWorkDueToFloatingPointPrecision(t *testing.T) {
	// This test fails because 64 floats are too small to accurately represent
	// the result

	//makeCalcTest(t)("1111111111111111111111111111111111111111111+"+
	//	"2222222222222222222222222222222222222222222=",
	//	"1111111111111111111111111111111111111111111+2222222222222222222222222222222222222222222=",
	//	"3333333333333333333333333333333333333333333")
}

func TestPlusMinusMultiplyDivideAllWork(t *testing.T) {
	inout := makeCalcTest(t)
	inout("3+4=", "3+4=", "7")
	inout("3-4=", "3-4=", "-1")
	inout("3*4=", "3*4=", "12")
	inout("3/4=", "3/4=", "0.75")
}

func TestAfterDivErrorNewCalculationStarts(t *testing.T) {
	makeCalcTest(t)("1/0=5+6=", "5+6=", "11")
	makeCalcTest(t)("1/0=-5=", "-5=", "-5")
}

func TestOperatorsAtStartAreIgnoredExceptForMinus(t *testing.T) {
	inout := makeCalcTest(t)
	inout("*", "", "0")
	inout("/", "", "0")
	inout("+", "", "0")
}

func TestClearingWillResetCalculator(t *testing.T) {
	makeCalcTest(t)("1+C", "", "0")
}

func TestLastResultIsUsedForNextCalculation(t *testing.T) {
	inout := makeCalcTest(t)
	inout("2*3=", "2*3=", "6")
	inout("+3=", "6+3=", "9")
}

func TestSingleDecimalPointAfterNumberIsRemoved(t *testing.T) {
	makeCalcTest(t)("1.+2=", "1+2=", "3")
}

func TestSubtractionWorksLeftToRight(t *testing.T) {
	makeCalcTest(t)("1-1-1=", "1-1-1=", "-1")
}

func TestNumbersStartingWithDecimalArePrependedWithZero(t *testing.T) {
	inout := makeCalcTest(t)
	inout(".", "", "0.")
	inout("5", "", "0.5")
	inout("=", "0.5=", "0.5")
	inout("+.1-.2=", "0.5+0.1-0.2=", "0.4")
}

func TestInvalidInputIsIgnored(t *testing.T) {
	makeCalcTest(t)("abc", "", "0")
}

func TestNumbersCanBeNegated(t *testing.T) {
	inout := makeCalcTest(t)
	inout("1", "", "1")
	inout("N", "", "-1")
	inout("N", "", "1")
	inout("N", "", "-1")
}

func TestNegatingZeroDoesNothing(t *testing.T) {
	makeCalcTest(t)("N", "", "0")
	makeCalcTest(t)("1-1=N", "", "0")
}

func TestOperatorsCannotBeBegated(t *testing.T) {
	makeCalcTest(t)("1+N", "1", "+")
}

func TestLastResultCanBeNegated(t *testing.T) {
	inout := makeCalcTest(t)
	inout("1+2=", "1+2=", "3")
	inout("N", "", "-3")
	inout("+4=", "-3+4=", "1")

	inout("C", "", "0")
	inout("1-2=", "1-2=", "-1")
	inout("N", "", "1")
	inout("+4=", "1+4=", "5")
}
