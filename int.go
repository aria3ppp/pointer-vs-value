package main

type SingleFieldIntStruct struct {
	field int
}

//go:noinline
func (s SingleFieldIntStruct) CopyFrom(other SingleFieldIntStruct) {}

//go:noinline
func (s *SingleFieldIntStruct) ReferenceFrom(
	other *SingleFieldIntStruct,
) {
}

// -----------------------------------------------------------------------------

type TwoFieldIntStruct struct {
	field1 int
	field2 int
}

//go:noinline
func (s TwoFieldIntStruct) CopyFrom(other TwoFieldIntStruct) {}

//go:noinline
func (s *TwoFieldIntStruct) ReferenceFrom(other *TwoFieldIntStruct) {}

// -----------------------------------------------------------------------------

type FourFieldIntStruct struct {
	field1 int
	field2 int
	field3 int
	field4 int
}

//go:noinline
func (s FourFieldIntStruct) CopyFrom(other FourFieldIntStruct) {}

//go:noinline
func (s *FourFieldIntStruct) ReferenceFrom(other *FourFieldIntStruct) {}

// -----------------------------------------------------------------------------

type EightFieldIntStruct struct {
	field1 int
	field2 int
	field3 int
	field4 int
	field5 int
	field6 int
	field7 int
	field8 int
}

//go:noinline
func (s EightFieldIntStruct) CopyFrom(other EightFieldIntStruct) {}

//go:noinline
func (s *EightFieldIntStruct) ReferenceFrom(other *EightFieldIntStruct) {}

// -----------------------------------------------------------------------------

type SixteenFieldIntStruct struct {
	field1  int
	field2  int
	field3  int
	field4  int
	field5  int
	field6  int
	field7  int
	field8  int
	field9  int
	field10 int
	field11 int
	field12 int
	field13 int
	field14 int
	field15 int
	field16 int
}

//go:noinline
func (s SixteenFieldIntStruct) CopyFrom(other SixteenFieldIntStruct) {}

//go:noinline
func (s *SixteenFieldIntStruct) ReferenceFrom(
	other *SixteenFieldIntStruct,
) {
}
