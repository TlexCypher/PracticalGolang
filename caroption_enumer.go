// Code generated by "enumer -type=CarOption"; DO NOT EDIT.

package main

import (
	"fmt"
	"strings"
)

const (
	_CarOptionName_0      = "GPSAWD"
	_CarOptionLowerName_0 = "gpsawd"
	_CarOptionName_1      = "SunRoof"
	_CarOptionLowerName_1 = "sunroof"
	_CarOptionName_2      = "HeatedSeat"
	_CarOptionLowerName_2 = "heatedseat"
	_CarOptionName_3      = "DriverAssist"
	_CarOptionLowerName_3 = "driverassist"
)

var (
	_CarOptionIndex_0 = [...]uint8{0, 3, 6}
	_CarOptionIndex_1 = [...]uint8{0, 7}
	_CarOptionIndex_2 = [...]uint8{0, 10}
	_CarOptionIndex_3 = [...]uint8{0, 12}
)

func (i CarOption) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _CarOptionName_0[_CarOptionIndex_0[i]:_CarOptionIndex_0[i+1]]
	case i == 4:
		return _CarOptionName_1
	case i == 8:
		return _CarOptionName_2
	case i == 16:
		return _CarOptionName_3
	default:
		return fmt.Sprintf("CarOption(%d)", i)
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _CarOptionNoOp() {
	var x [1]struct{}
	_ = x[GPS-(1)]
	_ = x[AWD-(2)]
	_ = x[SunRoof-(4)]
	_ = x[HeatedSeat-(8)]
	_ = x[DriverAssist-(16)]
}

var _CarOptionValues = []CarOption{GPS, AWD, SunRoof, HeatedSeat, DriverAssist}

var _CarOptionNameToValueMap = map[string]CarOption{
	_CarOptionName_0[0:3]:       GPS,
	_CarOptionLowerName_0[0:3]:  GPS,
	_CarOptionName_0[3:6]:       AWD,
	_CarOptionLowerName_0[3:6]:  AWD,
	_CarOptionName_1[0:7]:       SunRoof,
	_CarOptionLowerName_1[0:7]:  SunRoof,
	_CarOptionName_2[0:10]:      HeatedSeat,
	_CarOptionLowerName_2[0:10]: HeatedSeat,
	_CarOptionName_3[0:12]:      DriverAssist,
	_CarOptionLowerName_3[0:12]: DriverAssist,
}

var _CarOptionNames = []string{
	_CarOptionName_0[0:3],
	_CarOptionName_0[3:6],
	_CarOptionName_1[0:7],
	_CarOptionName_2[0:10],
	_CarOptionName_3[0:12],
}

// CarOptionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CarOptionString(s string) (CarOption, error) {
	if val, ok := _CarOptionNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _CarOptionNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CarOption values", s)
}

// CarOptionValues returns all values of the enum
func CarOptionValues() []CarOption {
	return _CarOptionValues
}

// CarOptionStrings returns a slice of all String values of the enum
func CarOptionStrings() []string {
	strs := make([]string, len(_CarOptionNames))
	copy(strs, _CarOptionNames)
	return strs
}

// IsACarOption returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CarOption) IsACarOption() bool {
	for _, v := range _CarOptionValues {
		if i == v {
			return true
		}
	}
	return false
}