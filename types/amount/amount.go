// This file contains code modified from the btcd project,
// which is licensed under the ISC License.
//
// Original license: https://github.com/btcsuite/btcd/blob/master/LICENSE
//

package amount

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/aerium-network/aerium/crypto"
)

const (
	// NanoAUMPerAUM is the number of NanoAUM in one AUM.
	NanoAUMPerAUM = 1e9

	// DefaultMaxNanoAUM is the default maximum transaction amount allowed in NanoAUM (42 million for mainnet).
	DefaultMaxNanoAUM = 42e6 * NanoAUMPerAUM

	// TestnetMaxNanoAUM is the maximum amount for testnet (9 billion).
	TestnetMaxNanoAUM = 9e9 * NanoAUMPerAUM
)

// Unit describes a method of converting an Amount to something
// other than the base unit of a Aerium.  The value of the Unit
// is the exponent component of the decadic multiple to convert from
// an amount in Aerium to an amount counted in units.
type Unit int

// These constants define various units used when describing a Aerium
// monetary amount.
const (
	UnitMegaAUM  Unit = 6
	UnitKiloAUM  Unit = 3
	UnitAUM      Unit = 0
	UnitMilliAUM Unit = -3
	UnitMicroAUM Unit = -6
	UnitNanoAUM  Unit = -9
)

// String returns the unit as a string.  For recognized units, the SI
// prefix is used, or "NanoAUM" for the base unit.  For all unrecognized
// units, "1eN AUM" is returned, where N is the AmountUnit.
func (u Unit) String() string {
	switch u {
	case UnitMegaAUM:
		return "MAUM"
	case UnitKiloAUM:
		return "kAUM"
	case UnitAUM:
		return "AUM"
	case UnitMilliAUM:
		return "mAUM"
	case UnitMicroAUM:
		return "μAUM"
	case UnitNanoAUM:
		return "NanoAUM"
	default:
		return "1e" + strconv.FormatInt(int64(u), 10) + " AUM"
	}
}

// Amount represents the atomic unit in Aerium blockchain.
// Each unit equals to 1e-9 of a AUM.
type Amount int64

// round converts a floating point number, which may or may not be representable
// as an integer, to the Amount integer type by rounding to the nearest integer.
// This is performed by adding or subtracting 0.5 depending on the sign, and
// relying on integer truncation to round the value to the nearest Amount.
func round(f float64) Amount {
	if f < 0 {
		return Amount(f - 0.5)
	}

	return Amount(f + 0.5)
}

// NewAmount creates an Amount from a floating-point value representing
// an amount in AUM.  NewAmount returns an error if f is NaN or +-Infinity,
// but it does not check whether the amount is within the total amount of AUM
// producible, as it may not refer to an amount at a single moment in time.
//
// NewAmount is specifically for converting AUM to NanoAUM.
// For creating a new Amount with an int64 value which denotes a quantity of NanoAUM,
// do a simple type conversion from type int64 to Amount.
func NewAmount(aum float64) (Amount, error) {
	// The amount is only considered invalid if it cannot be represented
	// as an integer type. This may happen if f is NaN or +-Infinity.
	switch {
	case math.IsNaN(aum),
		math.IsInf(aum, 1),
		math.IsInf(aum, -1):
		return 0, errors.New("invalid aum amount")
	}

	return round(aum * float64(NanoAUMPerAUM)), nil
}

// FromString parses a string representing a value in AUM.
// It then uses NewAmount to create an Amount based on the parsed
// floating-point value.
// If the parsing of the string fails, it returns an error.
func FromString(str string) (Amount, error) {
	str = strings.Replace(str, "AUM", "", 1)
	str = strings.TrimSpace(str)
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}

	return NewAmount(f)
}

// MaxNanoAUM returns the maximum amount based on chain type
func MaxNanoAUM() Amount {
	if crypto.AddressHRP == "tae" { // Testnet
		return TestnetMaxNanoAUM
	}
	return DefaultMaxNanoAUM // Mainnet
}

// ToUnit converts a monetary amount counted in Aerium base units to a
// floating-point value representing an amount of Aerium (AUM).
func (a Amount) ToUnit(u Unit) float64 {
	return float64(a) / math.Pow10(int(u+9))
}

// ToAUM is equivalent to calling ToUnit with AmountAUM.
func (a Amount) ToAUM() float64 {
	return a.ToUnit(UnitAUM)
}

// ToNanoAUM is equivalent to calling ToUnit with AmountNanoAUM.
// It returns the amount of NanoAUM or atomic unit as a 64-bit integer.
func (a Amount) ToNanoAUM() int64 {
	return int64(a)
}

// Format formats a monetary amount counted in Aerium base units as a
// string for a given unit.  The conversion will succeed for any unit,
// however, known units will be formatted with an appended label describing
// the units with SI notation, and "NanoAUM" for the base unit.
func (a Amount) Format(u Unit) string {
	units := " " + u.String()
	formatted := strconv.FormatFloat(a.ToUnit(u), 'f', -int(u+9), 64)

	return formatted + units
}

// String is the equivalent of calling Format with AmountAUM.
func (a Amount) String() string {
	return a.Format(UnitAUM)
}

// MulF64 multiplies an Amount by a floating point value.
func (a Amount) MulF64(f float64) Amount {
	return round(float64(a) * f)
}
