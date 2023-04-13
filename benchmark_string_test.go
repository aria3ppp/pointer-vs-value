package main

import "testing"

func BenchmarkSingleFieldStringStruct(b *testing.B) {
	b.Run("value", func(b *testing.B) {
		s := SingleFieldStringStruct{
			field: "field_string_value",
		}
		other := SingleFieldStringStruct{
			field: "other_field_string_value",
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.CopyFrom(other)
		}
	})

	b.Run("pointer", func(b *testing.B) {
		s := &SingleFieldStringStruct{
			field: "field_string_value",
		}
		other := &SingleFieldStringStruct{
			field: "other_field_string_value",
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.ReferenceFrom(other)
		}
	})
}

func BenchmarkTwoFieldStringStruct(b *testing.B) {
	b.Run("value", func(b *testing.B) {
		s := TwoFieldStringStruct{
			field1: "field1_string_value",
			field2: "field2_string_value",
		}
		other := TwoFieldStringStruct{
			field1: "other_field1_string_value",
			field2: "other_field2_string_value",
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.CopyFrom(other)
		}
	})

	b.Run("pointer", func(b *testing.B) {
		s := &TwoFieldStringStruct{
			field1: "field1_string_value",
			field2: "field2_string_value",
		}
		other := &TwoFieldStringStruct{
			field1: "other_field1_string_value",
			field2: "other_field2_string_value",
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.ReferenceFrom(other)
		}
	})
}

func BenchmarkFourFieldStringStruct(b *testing.B) {
	b.Run("value", func(b *testing.B) {
		s := FourFieldStringStruct{
			field1: "field1_string_value",
			field2: "field2_string_value",
			field3: "field3_string_value",
			field4: "field4_string_value",
		}
		other := FourFieldStringStruct{
			field1: "other_field1_string_value",
			field2: "other_field2_string_value",
			field3: "other_field3_string_value",
			field4: "other_field4_string_value",
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.CopyFrom(other)
		}
	})

	b.Run("pointer", func(b *testing.B) {
		s := &FourFieldStringStruct{
			field1: "field1_string_value",
			field2: "field2_string_value",
			field3: "field3_string_value",
			field4: "field4_string_value",
		}
		other := &FourFieldStringStruct{
			field1: "other_field1_string_value",
			field2: "other_field2_string_value",
			field3: "other_field3_string_value",
			field4: "other_field4_string_value",
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.ReferenceFrom(other)
		}
	})
}

func BenchmarkEightFieldStringStruct(b *testing.B) {
	b.Run("value", func(b *testing.B) {
		s := EightFieldStringStruct{
			field1: "field1_string_value",
			field2: "field2_string_value",
			field3: "field3_string_value",
			field4: "field4_string_value",
			field5: "field5_string_value",
			field6: "field6_string_value",
			field7: "field7_string_value",
			field8: "field8_string_value",
		}
		other := EightFieldStringStruct{
			field1: "other_field1_string_value",
			field2: "other_field2_string_value",
			field3: "other_field3_string_value",
			field4: "other_field4_string_value",
			field5: "other_field5_string_value",
			field6: "other_field6_string_value",
			field7: "other_field7_string_value",
			field8: "other_field8_string_value",
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.CopyFrom(other)
		}
	})

	b.Run("pointer", func(b *testing.B) {
		s := &EightFieldStringStruct{
			field1: "field1_string_value",
			field2: "field2_string_value",
			field3: "field3_string_value",
			field4: "field4_string_value",
			field5: "field5_string_value",
			field6: "field6_string_value",
			field7: "field7_string_value",
			field8: "field8_string_value",
		}
		other := &EightFieldStringStruct{
			field1: "other_field1_string_value",
			field2: "other_field2_string_value",
			field3: "other_field3_string_value",
			field4: "other_field4_string_value",
			field5: "other_field5_string_value",
			field6: "other_field6_string_value",
			field7: "other_field7_string_value",
			field8: "other_field8_string_value",
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.ReferenceFrom(other)
		}
	})
}

func BenchmarkSixteenFieldStringStruct(b *testing.B) {
	b.Run("value", func(b *testing.B) {
		s := SixteenFieldStringStruct{
			field1:  "field1_string_value",
			field2:  "field2_string_value",
			field3:  "field3_string_value",
			field4:  "field4_string_value",
			field5:  "field5_string_value",
			field6:  "field6_string_value",
			field7:  "field7_string_value",
			field8:  "field8_string_value",
			field9:  "field9_string_value",
			field10: "field10_string_value",
			field11: "field11_string_value",
			field12: "field12_string_value",
			field13: "field13_string_value",
			field14: "field14_string_value",
			field15: "field15_string_value",
			field16: "field16_string_value",
		}
		other := SixteenFieldStringStruct{
			field1:  "other_field1_string_value",
			field2:  "other_field2_string_value",
			field3:  "other_field3_string_value",
			field4:  "other_field4_string_value",
			field5:  "other_field5_string_value",
			field6:  "other_field6_string_value",
			field7:  "other_field7_string_value",
			field8:  "other_field8_string_value",
			field9:  "other_field9_string_value",
			field10: "other_field10_string_value",
			field11: "other_field11_string_value",
			field12: "other_field12_string_value",
			field13: "other_field13_string_value",
			field14: "other_field14_string_value",
			field15: "other_field15_string_value",
			field16: "other_field16_string_value",
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.CopyFrom(other)
		}
	})

	b.Run("pointer", func(b *testing.B) {
		s := &SixteenFieldStringStruct{
			field1:  "field1_string_value",
			field2:  "field2_string_value",
			field3:  "field3_string_value",
			field4:  "field4_string_value",
			field5:  "field5_string_value",
			field6:  "field6_string_value",
			field7:  "field7_string_value",
			field8:  "field8_string_value",
			field9:  "field9_string_value",
			field10: "field10_string_value",
			field11: "field11_string_value",
			field12: "field12_string_value",
			field13: "field13_string_value",
			field14: "field14_string_value",
			field15: "field15_string_value",
			field16: "field16_string_value",
		}
		other := &SixteenFieldStringStruct{
			field1:  "other_field1_string_value",
			field2:  "other_field2_string_value",
			field3:  "other_field3_string_value",
			field4:  "other_field4_string_value",
			field5:  "other_field5_string_value",
			field6:  "other_field6_string_value",
			field7:  "other_field7_string_value",
			field8:  "other_field8_string_value",
			field9:  "other_field9_string_value",
			field10: "other_field10_string_value",
			field11: "other_field11_string_value",
			field12: "other_field12_string_value",
			field13: "other_field13_string_value",
			field14: "other_field14_string_value",
			field15: "other_field15_string_value",
			field16: "other_field16_string_value",
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.ReferenceFrom(other)
		}
	})
}
