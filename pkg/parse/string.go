// Code generated by "stringer -type=RedirMode,PrimaryType,DQSegmentType -output=string.go"; DO NOT EDIT.

package parse

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RedirInvalid-0]
	_ = x[RedirInput-1]
	_ = x[RedirOutput-2]
	_ = x[RedirInputOutput-3]
	_ = x[RedirAppend-4]
	_ = x[RedirHeredoc-5]
}

const _RedirMode_name = "RedirInvalidRedirInputRedirOutputRedirInputOutputRedirAppendRedirHeredoc"

var _RedirMode_index = [...]uint8{0, 12, 22, 33, 49, 60, 72}

func (i RedirMode) String() string {
	if i < 0 || i >= RedirMode(len(_RedirMode_index)-1) {
		return "RedirMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RedirMode_name[_RedirMode_index[i]:_RedirMode_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InvalidPrimary-0]
	_ = x[BarewordPrimary-1]
	_ = x[SingleQuotedPrimary-2]
	_ = x[DoubleQuotedPrimary-3]
	_ = x[ArithmeticPrimary-4]
	_ = x[WildcardCharPrimary-5]
	_ = x[OutputCapturePrimary-6]
	_ = x[VariablePrimary-7]
}

const _PrimaryType_name = "InvalidPrimaryBarewordPrimarySingleQuotedPrimaryDoubleQuotedPrimaryArithmeticPrimaryWildcardCharPrimaryOutputCapturePrimaryVariablePrimary"

var _PrimaryType_index = [...]uint8{0, 14, 29, 48, 67, 84, 103, 123, 138}

func (i PrimaryType) String() string {
	if i < 0 || i >= PrimaryType(len(_PrimaryType_index)-1) {
		return "PrimaryType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PrimaryType_name[_PrimaryType_index[i]:_PrimaryType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DQInvalidSegment-0]
	_ = x[DQStringSegment-1]
	_ = x[DQExpansionSegment-2]
}

const _DQSegmentType_name = "DQInvalidSegmentDQStringSegmentDQExpansionSegment"

var _DQSegmentType_index = [...]uint8{0, 16, 31, 49}

func (i DQSegmentType) String() string {
	if i < 0 || i >= DQSegmentType(len(_DQSegmentType_index)-1) {
		return "DQSegmentType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DQSegmentType_name[_DQSegmentType_index[i]:_DQSegmentType_index[i+1]]
}
