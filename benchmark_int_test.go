package main

import "testing"

func BenchmarkSingleFieldIntStruct(b *testing.B) {
	b.Run("value", func(b *testing.B) {
		s := SingleFieldIntStruct{
			field: 1,
		}
		other := SingleFieldIntStruct{
			field: 11,
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.CopyFrom(other)
		}
	})

	b.Run("pointer", func(b *testing.B) {
		s := &SingleFieldIntStruct{
			field: 1,
		}
		other := &SingleFieldIntStruct{
			field: 11,
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.ReferenceFrom(other)
		}
	})
}

func BenchmarkTwoFieldIntStruct(b *testing.B) {
	b.Run("value", func(b *testing.B) {
		s := TwoFieldIntStruct{
			field1: 1,
			field2: 2,
		}
		other := TwoFieldIntStruct{
			field1: 11,
			field2: 22,
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.CopyFrom(other)
		}
	})

	b.Run("pointer", func(b *testing.B) {
		s := &TwoFieldIntStruct{
			field1: 1,
			field2: 2,
		}
		other := &TwoFieldIntStruct{
			field1: 11,
			field2: 22,
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.ReferenceFrom(other)
		}
	})
}

func BenchmarkFourFieldIntStruct(b *testing.B) {
	b.Run("value", func(b *testing.B) {
		s := FourFieldIntStruct{
			field1: 1,
			field2: 2,
			field3: 3,
			field4: 4,
		}
		other := FourFieldIntStruct{
			field1: 11,
			field2: 22,
			field3: 33,
			field4: 44,
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.CopyFrom(other)
		}
	})

	b.Run("pointer", func(b *testing.B) {
		s := &FourFieldIntStruct{
			field1: 1,
			field2: 2,
			field3: 3,
			field4: 4,
		}
		other := &FourFieldIntStruct{
			field1: 11,
			field2: 22,
			field3: 33,
			field4: 44,
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.ReferenceFrom(other)
		}
	})
}

func BenchmarkEightFieldIntStruct(b *testing.B) {
	b.Run("value", func(b *testing.B) {
		s := EightFieldIntStruct{
			field1: 1,
			field2: 2,
			field3: 3,
			field4: 4,
			field5: 5,
			field6: 6,
			field7: 7,
			field8: 8,
		}
		other := EightFieldIntStruct{
			field1: 11,
			field2: 22,
			field3: 33,
			field4: 44,
			field5: 55,
			field6: 66,
			field7: 77,
			field8: 88,
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.CopyFrom(other)
		}
	})

	b.Run("pointer", func(b *testing.B) {
		s := &EightFieldIntStruct{
			field1: 1,
			field2: 2,
			field3: 3,
			field4: 4,
			field5: 5,
			field6: 6,
			field7: 7,
			field8: 8,
		}
		other := &EightFieldIntStruct{
			field1: 11,
			field2: 22,
			field3: 33,
			field4: 44,
			field5: 55,
			field6: 66,
			field7: 77,
			field8: 88,
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.ReferenceFrom(other)
		}
	})
}

func BenchmarkSixteenFieldIntStruct(b *testing.B) {
	b.Run("value", func(b *testing.B) {
		s := SixteenFieldIntStruct{
			field1:  1,
			field2:  2,
			field3:  3,
			field4:  4,
			field5:  5,
			field6:  6,
			field7:  7,
			field8:  8,
			field9:  9,
			field10: 10,
			field11: 11,
			field12: 12,
			field13: 13,
			field14: 14,
			field15: 15,
			field16: 16,
		}
		other := SixteenFieldIntStruct{
			field1:  11,
			field2:  22,
			field3:  33,
			field4:  44,
			field5:  55,
			field6:  66,
			field7:  77,
			field8:  88,
			field9:  99,
			field10: 1010,
			field11: 1111,
			field12: 1212,
			field13: 1313,
			field14: 1414,
			field15: 1515,
			field16: 1616,
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.CopyFrom(other)
		}
	})

	b.Run("pointer", func(b *testing.B) {
		s := &SixteenFieldIntStruct{
			field1:  1,
			field2:  2,
			field3:  3,
			field4:  4,
			field5:  5,
			field6:  6,
			field7:  7,
			field8:  8,
			field9:  9,
			field10: 10,
			field11: 11,
			field12: 12,
			field13: 13,
			field14: 14,
			field15: 15,
			field16: 16,
		}
		other := &SixteenFieldIntStruct{
			field1:  11,
			field2:  22,
			field3:  33,
			field4:  44,
			field5:  55,
			field6:  66,
			field7:  77,
			field8:  88,
			field9:  99,
			field10: 1010,
			field11: 1111,
			field12: 1212,
			field13: 1313,
			field14: 1414,
			field15: 1515,
			field16: 1616,
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.ReferenceFrom(other)
		}
	})
}
