package main

type SingleFieldStringStruct struct {
	field string
}

//go:noinline
func (s SingleFieldStringStruct) CopyFrom(other SingleFieldStringStruct) {}

//go:noinline
func (s *SingleFieldStringStruct) ReferenceFrom(
	other *SingleFieldStringStruct,
) {
}

// -----------------------------------------------------------------------------

type TwoFieldStringStruct struct {
	field1 string
	field2 string
}

//go:noinline
func (s TwoFieldStringStruct) CopyFrom(other TwoFieldStringStruct) {}

//go:noinline
func (s *TwoFieldStringStruct) ReferenceFrom(other *TwoFieldStringStruct) {}

// -----------------------------------------------------------------------------

type FourFieldStringStruct struct {
	field1 string
	field2 string
	field3 string
	field4 string
}

//go:noinline
func (s FourFieldStringStruct) CopyFrom(other FourFieldStringStruct) {}

//go:noinline
func (s *FourFieldStringStruct) ReferenceFrom(other *FourFieldStringStruct) {}

// -----------------------------------------------------------------------------

type EightFieldStringStruct struct {
	field1 string
	field2 string
	field3 string
	field4 string
	field5 string
	field6 string
	field7 string
	field8 string
}

//go:noinline
func (s EightFieldStringStruct) CopyFrom(other EightFieldStringStruct) {}

//go:noinline
func (s *EightFieldStringStruct) ReferenceFrom(other *EightFieldStringStruct) {}

// -----------------------------------------------------------------------------

type SixteenFieldStringStruct struct {
	field1  string
	field2  string
	field3  string
	field4  string
	field5  string
	field6  string
	field7  string
	field8  string
	field9  string
	field10 string
	field11 string
	field12 string
	field13 string
	field14 string
	field15 string
	field16 string
}

//go:noinline
func (s SixteenFieldStringStruct) CopyFrom(other SixteenFieldStringStruct) {}

//go:noinline
func (s *SixteenFieldStringStruct) ReferenceFrom(
	other *SixteenFieldStringStruct,
) {
}
