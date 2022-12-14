// Code generated by array/numeric.gen.go.tmpl. DO NOT EDIT.

// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package array

import (
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v11/arrow"
)

// A type which represents an immutable sequence of int64 values.
type Int64 struct {
	array
	values []int64
}

// NewInt64Data creates a new Int64.
func NewInt64Data(data *Data) *Int64 {
	a := &Int64{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Int64) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Int64) Value(i int) int64 { return a.values[i] }

// Values returns the values.
func (a *Int64) Int64Values() []int64 { return a.values }

// String returns a string representation of the array.
func (a *Int64) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Int64) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Int64Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualInt64(left, right *Int64) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of uint64 values.
type Uint64 struct {
	array
	values []uint64
}

// NewUint64Data creates a new Uint64.
func NewUint64Data(data *Data) *Uint64 {
	a := &Uint64{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Uint64) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Uint64) Value(i int) uint64 { return a.values[i] }

// Values returns the values.
func (a *Uint64) Uint64Values() []uint64 { return a.values }

// String returns a string representation of the array.
func (a *Uint64) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Uint64) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Uint64Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualUint64(left, right *Uint64) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of float64 values.
type Float64 struct {
	array
	values []float64
}

// NewFloat64Data creates a new Float64.
func NewFloat64Data(data *Data) *Float64 {
	a := &Float64{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Float64) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Float64) Value(i int) float64 { return a.values[i] }

// Values returns the values.
func (a *Float64) Float64Values() []float64 { return a.values }

// String returns a string representation of the array.
func (a *Float64) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Float64) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Float64Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualFloat64(left, right *Float64) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of int32 values.
type Int32 struct {
	array
	values []int32
}

// NewInt32Data creates a new Int32.
func NewInt32Data(data *Data) *Int32 {
	a := &Int32{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Int32) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Int32) Value(i int) int32 { return a.values[i] }

// Values returns the values.
func (a *Int32) Int32Values() []int32 { return a.values }

// String returns a string representation of the array.
func (a *Int32) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Int32) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Int32Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualInt32(left, right *Int32) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of uint32 values.
type Uint32 struct {
	array
	values []uint32
}

// NewUint32Data creates a new Uint32.
func NewUint32Data(data *Data) *Uint32 {
	a := &Uint32{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Uint32) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Uint32) Value(i int) uint32 { return a.values[i] }

// Values returns the values.
func (a *Uint32) Uint32Values() []uint32 { return a.values }

// String returns a string representation of the array.
func (a *Uint32) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Uint32) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Uint32Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualUint32(left, right *Uint32) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of float32 values.
type Float32 struct {
	array
	values []float32
}

// NewFloat32Data creates a new Float32.
func NewFloat32Data(data *Data) *Float32 {
	a := &Float32{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Float32) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Float32) Value(i int) float32 { return a.values[i] }

// Values returns the values.
func (a *Float32) Float32Values() []float32 { return a.values }

// String returns a string representation of the array.
func (a *Float32) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Float32) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Float32Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualFloat32(left, right *Float32) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of int16 values.
type Int16 struct {
	array
	values []int16
}

// NewInt16Data creates a new Int16.
func NewInt16Data(data *Data) *Int16 {
	a := &Int16{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Int16) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Int16) Value(i int) int16 { return a.values[i] }

// Values returns the values.
func (a *Int16) Int16Values() []int16 { return a.values }

// String returns a string representation of the array.
func (a *Int16) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Int16) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Int16Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualInt16(left, right *Int16) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of uint16 values.
type Uint16 struct {
	array
	values []uint16
}

// NewUint16Data creates a new Uint16.
func NewUint16Data(data *Data) *Uint16 {
	a := &Uint16{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Uint16) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Uint16) Value(i int) uint16 { return a.values[i] }

// Values returns the values.
func (a *Uint16) Uint16Values() []uint16 { return a.values }

// String returns a string representation of the array.
func (a *Uint16) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Uint16) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Uint16Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualUint16(left, right *Uint16) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of int8 values.
type Int8 struct {
	array
	values []int8
}

// NewInt8Data creates a new Int8.
func NewInt8Data(data *Data) *Int8 {
	a := &Int8{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Int8) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Int8) Value(i int) int8 { return a.values[i] }

// Values returns the values.
func (a *Int8) Int8Values() []int8 { return a.values }

// String returns a string representation of the array.
func (a *Int8) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Int8) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Int8Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualInt8(left, right *Int8) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of uint8 values.
type Uint8 struct {
	array
	values []uint8
}

// NewUint8Data creates a new Uint8.
func NewUint8Data(data *Data) *Uint8 {
	a := &Uint8{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Uint8) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Uint8) Value(i int) uint8 { return a.values[i] }

// Values returns the values.
func (a *Uint8) Uint8Values() []uint8 { return a.values }

// String returns a string representation of the array.
func (a *Uint8) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Uint8) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Uint8Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualUint8(left, right *Uint8) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of arrow.Timestamp values.
type Timestamp struct {
	array
	values []arrow.Timestamp
}

// NewTimestampData creates a new Timestamp.
func NewTimestampData(data *Data) *Timestamp {
	a := &Timestamp{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Timestamp) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Timestamp) Value(i int) arrow.Timestamp { return a.values[i] }

// Values returns the values.
func (a *Timestamp) TimestampValues() []arrow.Timestamp { return a.values }

// String returns a string representation of the array.
func (a *Timestamp) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Timestamp) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.TimestampTraits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualTimestamp(left, right *Timestamp) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of arrow.Time32 values.
type Time32 struct {
	array
	values []arrow.Time32
}

// NewTime32Data creates a new Time32.
func NewTime32Data(data *Data) *Time32 {
	a := &Time32{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Time32) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Time32) Value(i int) arrow.Time32 { return a.values[i] }

// Values returns the values.
func (a *Time32) Time32Values() []arrow.Time32 { return a.values }

// String returns a string representation of the array.
func (a *Time32) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Time32) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Time32Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualTime32(left, right *Time32) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of arrow.Time64 values.
type Time64 struct {
	array
	values []arrow.Time64
}

// NewTime64Data creates a new Time64.
func NewTime64Data(data *Data) *Time64 {
	a := &Time64{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Time64) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Time64) Value(i int) arrow.Time64 { return a.values[i] }

// Values returns the values.
func (a *Time64) Time64Values() []arrow.Time64 { return a.values }

// String returns a string representation of the array.
func (a *Time64) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Time64) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Time64Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualTime64(left, right *Time64) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of arrow.Date32 values.
type Date32 struct {
	array
	values []arrow.Date32
}

// NewDate32Data creates a new Date32.
func NewDate32Data(data *Data) *Date32 {
	a := &Date32{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Date32) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Date32) Value(i int) arrow.Date32 { return a.values[i] }

// Values returns the values.
func (a *Date32) Date32Values() []arrow.Date32 { return a.values }

// String returns a string representation of the array.
func (a *Date32) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Date32) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Date32Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualDate32(left, right *Date32) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of arrow.Date64 values.
type Date64 struct {
	array
	values []arrow.Date64
}

// NewDate64Data creates a new Date64.
func NewDate64Data(data *Data) *Date64 {
	a := &Date64{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Date64) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Date64) Value(i int) arrow.Date64 { return a.values[i] }

// Values returns the values.
func (a *Date64) Date64Values() []arrow.Date64 { return a.values }

// String returns a string representation of the array.
func (a *Date64) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Date64) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.Date64Traits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualDate64(left, right *Date64) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}

// A type which represents an immutable sequence of arrow.Duration values.
type Duration struct {
	array
	values []arrow.Duration
}

// NewDurationData creates a new Duration.
func NewDurationData(data *Data) *Duration {
	a := &Duration{}
	a.refCount = 1
	a.setData(data)
	return a
}

// Reset resets the array for re-use.
func (a *Duration) Reset(data *Data) {
	a.setData(data)
}

// Value returns the value at the specified index.
func (a *Duration) Value(i int) arrow.Duration { return a.values[i] }

// Values returns the values.
func (a *Duration) DurationValues() []arrow.Duration { return a.values }

// String returns a string representation of the array.
func (a *Duration) String() string {
	o := new(strings.Builder)
	o.WriteString("[")
	for i, v := range a.values {
		if i > 0 {
			fmt.Fprintf(o, " ")
		}
		switch {
		case a.IsNull(i):
			o.WriteString("(null)")
		default:
			fmt.Fprintf(o, "%v", v)
		}
	}
	o.WriteString("]")
	return o.String()
}

func (a *Duration) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.DurationTraits.CastFromBytes(vals.Bytes())
		beg := a.array.data.offset
		end := beg + a.array.data.length
		a.values = a.values[beg:end]
	}
}

func arrayEqualDuration(left, right *Duration) bool {
	for i := 0; i < left.Len(); i++ {
		if left.IsNull(i) {
			continue
		}
		if left.Value(i) != right.Value(i) {
			return false
		}
	}
	return true
}
