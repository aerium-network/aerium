// This file contains code modified from the btcd project,
// which is licensed under the ISC License.
//
// Original license: https://github.com/btcsuite/btcd/blob/master/LICENSE
//

package amount_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/aerium-network/aerium/types/amount"
	"github.com/stretchr/testify/assert"
)

func TestAmountCreation(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		valid    bool
		expected amount.Amount
	}{
		// Positive tests.
		{
			name:     "zero",
			amount:   0,
			valid:    true,
			expected: 0,
		},
		{
			name:     "max producible",
			amount:   42e6,
			valid:    true,
			expected: amount.DefaultMaxNanoAUM,
		},
		{
			name:     "min producible",
			amount:   -42e6,
			valid:    true,
			expected: -amount.DefaultMaxNanoAUM,
		},
		{
			name:     "exceeds max producible",
			amount:   42e6 + 8e-9,
			valid:    true,
			expected: amount.DefaultMaxNanoAUM + 8,
		},
		{
			name:     "exceeds min producible",
			amount:   -42e6 - 8e-9,
			valid:    true,
			expected: -amount.DefaultMaxNanoAUM - 8,
		},
		{
			name:     "one hundred",
			amount:   100,
			valid:    true,
			expected: 100 * amount.NanoAUMPerAUM,
		},
		{
			name:     "fraction",
			amount:   0.012345678,
			valid:    true,
			expected: 12345678,
		},
		{
			name:     "rounding up",
			amount:   54.999999999999943157,
			valid:    true,
			expected: 55 * amount.NanoAUMPerAUM,
		},
		{
			name:     "rounding down",
			amount:   55.000000000000056843,
			valid:    true,
			expected: 55 * amount.NanoAUMPerAUM,
		},

		// Negative tests.
		{
			name:   "not-a-number",
			amount: math.NaN(),
			valid:  false,
		},
		{
			name:   "-infinity",
			amount: math.Inf(-1),
			valid:  false,
		},
		{
			name:   "+infinity",
			amount: math.Inf(1),
			valid:  false,
		},
	}

	for _, tt := range tests {
		amt, err := amount.NewAmount(tt.amount)
		if tt.valid {
			assert.NoErrorf(t, err,
				"%v: Positive test Amount creation failed with: %v", tt.name, err)
		} else {
			assert.Errorf(t, err,
				"%v: Negative test Amount creation succeeded (value %v) when should fail", tt.name, amt)
		}

		assert.Equal(t, tt.expected, amt,
			"%v: Created amount %v does not match expected %v", tt.name, amt, tt.expected)
	}
}

func TestAmountUnitConversions(t *testing.T) {
	tests := []struct {
		name      string
		amount    amount.Amount
		unit      amount.Unit
		converted float64
		str       string
	}{
		{
			name:      "MAUM",
			amount:    amount.DefaultMaxNanoAUM,
			unit:      amount.UnitMegaAUM,
			converted: 42,
			str:       "42 MAUM",
		},
		{
			name:      "kAUM",
			amount:    444_333_222_111_000,
			unit:      amount.UnitKiloAUM,
			converted: 444.333_222_111_000,
			str:       "444.333222111 kAUM",
		},
		{
			name:      "AUM",
			amount:    444_333_222_111_000,
			unit:      amount.UnitAUM,
			converted: 444_333.222_111,
			str:       "444333.222111 AUM",
		},
		{
			name:      "a thousand NanoAUM as AUM",
			amount:    1_000,
			unit:      amount.UnitAUM,
			converted: 0.000_001,
			str:       "0.000001 AUM",
		},
		{
			name:      "a single NanoAUM as AUM",
			amount:    1,
			unit:      amount.UnitAUM,
			converted: 0.000_000_001,
			str:       "0.000000001 AUM",
		},
		{
			name:      "amount with trailing zero but no decimals",
			amount:    10_000_000_000,
			unit:      amount.UnitAUM,
			converted: 10,
			str:       "10 AUM",
		},
		{
			name:      "mAUM",
			amount:    444_333_222_111_000,
			unit:      amount.UnitMilliAUM,
			converted: 444_333_222.111_000,
			str:       "444333222.111 mAUM",
		},
		{
			name:      "μAUM",
			amount:    444_333_222_111_000,
			unit:      amount.UnitMicroAUM,
			converted: 444_333_222_111.000,
			str:       "444333222111 μAUM",
		},
		{
			name:      "NanoAUM",
			amount:    444_333_222_111_000,
			unit:      amount.UnitNanoAUM,
			converted: 444_333_222_111_000,
			str:       "444333222111000 NanoAUM",
		},
		{
			name:      "non-standard unit",
			amount:    444_333_222_111_000,
			unit:      amount.Unit(-1),
			converted: 4_443_332.221_110_00,
			str:       "4443332.22111 1e-1 AUM",
		},
	}

	for _, tt := range tests {
		f := tt.amount.ToUnit(tt.unit)
		assert.Equal(t, tt.converted, f,
			"%v: converted value %v does not match expected %v", tt.name, f, tt.converted)

		str := tt.amount.Format(tt.unit)
		assert.Equal(t, tt.str, str,
			"%v: format '%v' does not match expected '%v'", tt.name, str, tt.str)

		// Verify that Amount.ToAUM works as advertised.
		f1 := tt.amount.ToUnit(amount.UnitAUM)
		f2 := tt.amount.ToAUM()
		assert.Equal(t, f1, f2,
			"%v: ToAUM does not match ToUnit(AmountAUM): %v != %v", tt.name, f1, f2)

		// Verify that Amount.String works as advertised.
		s1 := tt.amount.Format(amount.UnitAUM)
		s2 := tt.amount.String()
		assert.Equal(t, s1, s2,
			"%v: String does not match Format(AmountAUM): %v != %v", tt.name, s1, s2)
	}
}

func TestAmountMulF64(t *testing.T) {
	tests := []struct {
		name string
		amt  amount.Amount
		mul  float64
		res  amount.Amount
	}{
		{
			name: "Multiply 0.1 AUM by 2",
			amt:  100e6, // 0.1 AUM
			mul:  2,
			res:  200e6, // 0.2 AUM
		},
		{
			name: "Multiply 0.2 AUM by 0.02",
			amt:  200e6, // 0.2 AUM
			mul:  1.02,
			res:  204e6, // 0.204 AUM
		},
		{
			name: "Multiply 0.1 AUM by -2",
			amt:  100e6, // 0.1 AUM
			mul:  -2,
			res:  -200e6, // -0.2 AUM
		},
		{
			name: "Multiply 0.2 AUM by -0.02",
			amt:  200e6, // 0.2 AUM
			mul:  -1.02,
			res:  -204e6, // -0.204 AUM
		},
		{
			name: "Multiply -0.1 AUM by 2",
			amt:  -100e6, // -0.1 AUM
			mul:  2,
			res:  -200e6, // -0.2 AUM
		},
		{
			name: "Multiply -0.2 AUM by 0.02",
			amt:  -200e6, // -0.2 AUM
			mul:  1.02,
			res:  -204e6, // -0.204 AUM
		},
		{
			name: "Multiply -0.1 AUM by -2",
			amt:  -100e6, // -0.1 AUM
			mul:  -2,
			res:  200e6, // 0.2 AUM
		},
		{
			name: "Multiply -0.2 AUM by -0.02",
			amt:  -200e6, // -0.2 AUM
			mul:  -1.02,
			res:  204e6, // 0.204 AUM
		},
		{
			name: "Round down",
			amt:  49, // 49 NanoAUMs
			mul:  0.01,
			res:  0,
		},
		{
			name: "Round up",
			amt:  50, // 50 NanoAUMs
			mul:  0.01,
			res:  1, // 1 NanoAUM
		},
		{
			name: "Multiply by 0",
			amt:  1e9, // 1 AUM
			mul:  0,
			res:  0, // 0 AUM
		},
		{
			name: "Multiply 1 by 0.5",
			amt:  1, // 1 NanoAUM
			mul:  0.5,
			res:  1, // 1 NanoAUM
		},
		{
			name: "Multiply 100 by 66%",
			amt:  100, // 100 NanoAUMs
			mul:  0.66,
			res:  66, // 66 NanoAUMs
		},
		{
			name: "Multiply 100 by 66.6%",
			amt:  100, // 100 NanoAUMs
			mul:  0.666,
			res:  67, // 67 NanoAUMs
		},
		{
			name: "Multiply 100 by 2/3",
			amt:  100, // 100 NanoAUMs
			mul:  2.0 / 3,
			res:  67, // 67 NanoAUMs
		},
	}

	for _, tt := range tests {
		a := tt.amt.MulF64(tt.mul)
		if a != tt.res {
			t.Errorf("%v: expected %v got %v", tt.name, tt.res, a)
		}
	}
}

func TestCoinToChangeConversion(t *testing.T) {
	tests := []struct {
		amount  string
		AUM     float64
		NanoAUM int64
		str     string
		parsErr error
	}{
		{"0", 0, 0, "0 AUM", nil},
		{"1", 1, 1000000000, "1 AUM", nil},
		{"1 AUM", 1, 1000000000, "1 AUM", nil},
		{"123.123", 123.123, 123123000000, "123.123 AUM", nil},
		{"123.0123", 123.0123, 123012300000, "123.0123 AUM", nil},
		{"123.01230", 123.0123, 123012300000, "123.0123 AUM", nil},
		{"123.000123", 123.000123, 123000123000, "123.000123 AUM", nil},
		{"123.000000123", 123.000000123, 123000000123, "123.000000123 AUM", nil},
		{"-123.000000123", -123.000000123, -123000000123, "-123.000000123 AUM", nil},
		{"0123.000000123", 123.000000123, 123000000123, "123.000000123 AUM", nil},
		{"+123.000000123", 123.000000123, 123000000123, "123.000000123 AUM", nil},
		{"123.0000001234", 123.000000123, 123000000123, "123.000000123 AUM", nil},
		{"1coin", 0, 0, "0", strconv.ErrSyntax},
	}
	for _, tt := range tests {
		amt, err := amount.FromString(tt.amount)
		if tt.parsErr == nil {
			assert.NoError(t, err)
			assert.Equal(t, tt.NanoAUM, amt.ToNanoAUM())
			assert.Equal(t, tt.AUM, amt.ToAUM())
			assert.Equal(t, tt.str, amt.String())
		} else {
			assert.ErrorIs(t, err, tt.parsErr)
		}
	}
}
