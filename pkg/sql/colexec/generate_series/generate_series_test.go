// Copyright 2022 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generate_series

import (
	"bytes"
	"encoding/json"
	"github.com/matrixorigin/matrixone/pkg/container/batch"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/logutil"
	"github.com/matrixorigin/matrixone/pkg/sql/plan"
	"github.com/matrixorigin/matrixone/pkg/testutil"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
	"github.com/stretchr/testify/require"
	"math"
	"strings"
	"testing"
)

type Kase[T Number] struct {
	start T
	end   T
	step  T
	res   []T
	err   bool
}

func TestDoGenerateInt32(t *testing.T) {
	kases := []Kase[int32]{
		{
			start: 1,
			end:   10,
			step:  1,
			res:   []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			start: 1,
			end:   10,
			step:  2,
			res:   []int32{1, 3, 5, 7, 9},
		},
		{
			start: 1,
			end:   10,
			step:  -1,
			res:   []int32{},
		},
		{
			start: 10,
			end:   1,
			step:  -1,
			res:   []int32{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			start: 10,
			end:   1,
			step:  -2,
			res:   []int32{10, 8, 6, 4, 2},
		},
		{
			start: 10,
			end:   1,
			step:  1,
			res:   []int32{},
		},
		{
			start: 1,
			end:   10,
			step:  0,
			err:   true,
		},
		{
			start: 1,
			end:   10,
			step:  -1,
			res:   []int32{},
		},
		{
			start: 1,
			end:   1,
			step:  0,
			err:   true,
		},
		{
			start: 1,
			end:   1,
			step:  1,
			res:   []int32{1},
		},
		{
			start: math.MaxInt32 - 1,
			end:   math.MaxInt32,
			step:  1,
			res:   []int32{math.MaxInt32 - 1, math.MaxInt32},
		},
		{
			start: math.MaxInt32,
			end:   math.MaxInt32 - 1,
			step:  -1,
			res:   []int32{math.MaxInt32, math.MaxInt32 - 1},
		},
		{
			start: math.MinInt32,
			end:   math.MinInt32 + 100,
			step:  19,
			res:   []int32{math.MinInt32, math.MinInt32 + 19, math.MinInt32 + 38, math.MinInt32 + 57, math.MinInt32 + 76, math.MinInt32 + 95},
		},
		{
			start: math.MinInt32 + 100,
			end:   math.MinInt32,
			step:  -19,
			res:   []int32{math.MinInt32 + 100, math.MinInt32 + 81, math.MinInt32 + 62, math.MinInt32 + 43, math.MinInt32 + 24, math.MinInt32 + 5},
		},
	}
	for _, kase := range kases {
		res, err := generateInt32(kase.start, kase.end, kase.step)
		if kase.err {
			require.NotNil(t, err)
			continue
		}
		require.Nil(t, err)
		require.Equal(t, kase.res, res)
	}
}

func TestDoGenerateInt64(t *testing.T) {
	kases := []Kase[int64]{
		{
			start: 1,
			end:   10,
			step:  1,
			res:   []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			start: 1,
			end:   10,
			step:  2,
			res:   []int64{1, 3, 5, 7, 9},
		},
		{
			start: 1,
			end:   10,
			step:  -1,
			res:   []int64{},
		},
		{
			start: 10,
			end:   1,
			step:  -1,
			res:   []int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			start: 10,
			end:   1,
			step:  -2,
			res:   []int64{10, 8, 6, 4, 2},
		},
		{
			start: 10,
			end:   1,
			step:  1,
			res:   []int64{},
		},
		{
			start: 1,
			end:   10,
			step:  0,
			err:   true,
		},
		{
			start: 1,
			end:   10,
			step:  -1,
			res:   []int64{},
		},
		{
			start: 1,
			end:   1,
			step:  0,
			err:   true,
		},
		{
			start: 1,
			end:   1,
			step:  1,
			res:   []int64{1},
		},
		{
			start: math.MaxInt32 - 1,
			end:   math.MaxInt32,
			step:  1,
			res:   []int64{math.MaxInt32 - 1, math.MaxInt32},
		},
		{
			start: math.MaxInt32,
			end:   math.MaxInt32 - 1,
			step:  -1,
			res:   []int64{math.MaxInt32, math.MaxInt32 - 1},
		},
		{
			start: math.MinInt32,
			end:   math.MinInt32 + 100,
			step:  19,
			res:   []int64{math.MinInt32, math.MinInt32 + 19, math.MinInt32 + 38, math.MinInt32 + 57, math.MinInt32 + 76, math.MinInt32 + 95},
		},
		{
			start: math.MinInt32 + 100,
			end:   math.MinInt32,
			step:  -19,
			res:   []int64{math.MinInt32 + 100, math.MinInt32 + 81, math.MinInt32 + 62, math.MinInt32 + 43, math.MinInt32 + 24, math.MinInt32 + 5},
		},
		// int64
		{
			start: math.MaxInt32,
			end:   math.MaxInt32 + 100,
			step:  19,
			res:   []int64{math.MaxInt32, math.MaxInt32 + 19, math.MaxInt32 + 38, math.MaxInt32 + 57, math.MaxInt32 + 76, math.MaxInt32 + 95},
		},
		{
			start: math.MaxInt32 + 100,
			end:   math.MaxInt32,
			step:  -19,
			res:   []int64{math.MaxInt32 + 100, math.MaxInt32 + 81, math.MaxInt32 + 62, math.MaxInt32 + 43, math.MaxInt32 + 24, math.MaxInt32 + 5},
		},
		{
			start: math.MinInt32,
			end:   math.MinInt32 - 100,
			step:  -19,
			res:   []int64{math.MinInt32, math.MinInt32 - 19, math.MinInt32 - 38, math.MinInt32 - 57, math.MinInt32 - 76, math.MinInt32 - 95},
		},
		{
			start: math.MinInt32 - 100,
			end:   math.MinInt32,
			step:  19,
			res:   []int64{math.MinInt32 - 100, math.MinInt32 - 81, math.MinInt32 - 62, math.MinInt32 - 43, math.MinInt32 - 24, math.MinInt32 - 5},
		},
		{
			start: math.MaxInt64 - 1,
			end:   math.MaxInt64,
			step:  1,
			res:   []int64{math.MaxInt64 - 1, math.MaxInt64},
		},
		{
			start: math.MaxInt64,
			end:   math.MaxInt64 - 1,
			step:  -1,
			res:   []int64{math.MaxInt64, math.MaxInt64 - 1},
		},
		{
			start: math.MaxInt64 - 100,
			end:   math.MaxInt64,
			step:  19,
			res:   []int64{math.MaxInt64 - 100, math.MaxInt64 - 81, math.MaxInt64 - 62, math.MaxInt64 - 43, math.MaxInt64 - 24, math.MaxInt64 - 5},
		},
		{
			start: math.MaxInt64,
			end:   math.MaxInt64 - 100,
			step:  -19,
			res:   []int64{math.MaxInt64, math.MaxInt64 - 19, math.MaxInt64 - 38, math.MaxInt64 - 57, math.MaxInt64 - 76, math.MaxInt64 - 95},
		},
		{
			start: math.MinInt64,
			end:   math.MinInt64 + 100,
			step:  19,
			res:   []int64{math.MinInt64, math.MinInt64 + 19, math.MinInt64 + 38, math.MinInt64 + 57, math.MinInt64 + 76, math.MinInt64 + 95},
		},
		{
			start: math.MinInt64 + 100,
			end:   math.MinInt64,
			step:  -19,
			res:   []int64{math.MinInt64 + 100, math.MinInt64 + 81, math.MinInt64 + 62, math.MinInt64 + 43, math.MinInt64 + 24, math.MinInt64 + 5},
		},
	}
	for _, kase := range kases {
		res, err := generateInt64(kase.start, kase.end, kase.step)
		if kase.err {
			require.NotNil(t, err)
			continue
		}
		require.Nil(t, err)
		require.Equal(t, kase.res, res)
	}
}

func TestGenerateTimestamp(t *testing.T) {
	kases := []struct {
		start string
		end   string
		step  string
		res   []types.Datetime
		err   bool
	}{
		{
			start: "2019-01-01 00:00:00",
			end:   "2019-01-01 00:01:00",
			step:  "30 second",
			res: []types.Datetime{
				transStr2Datetime("2019-01-01 00:00:00"),
				transStr2Datetime("2019-01-01 00:00:30"),
				transStr2Datetime("2019-01-01 00:01:00"),
			},
		},
		{
			start: "2019-01-01 00:01:00",
			end:   "2019-01-01 00:00:00",
			step:  "-30 second",
			res: []types.Datetime{
				transStr2Datetime("2019-01-01 00:01:00"),
				transStr2Datetime("2019-01-01 00:00:30"),
				transStr2Datetime("2019-01-01 00:00:00"),
			},
		},
		{
			start: "2019-01-01 00:00:00",
			end:   "2019-01-01 00:01:00",
			step:  "30 minute",
			res: []types.Datetime{
				transStr2Datetime("2019-01-01 00:00:00"),
			},
		},
		{
			start: "2020-02-29 00:01:00",
			end:   "2021-03-01 00:00:00",
			step:  "1 year",
			res: []types.Datetime{
				transStr2Datetime("2020-02-29 00:01:00"),
				transStr2Datetime("2021-02-28 00:01:00"),
			},
		},
		{
			start: "2020-02-29 00:01:00",
			end:   "2021-03-01 00:00:00",
			step:  "1 month",
			res: []types.Datetime{
				transStr2Datetime("2020-02-29 00:01:00"),
				transStr2Datetime("2020-03-29 00:01:00"),
				transStr2Datetime("2020-04-29 00:01:00"),
				transStr2Datetime("2020-05-29 00:01:00"),
				transStr2Datetime("2020-06-29 00:01:00"),
				transStr2Datetime("2020-07-29 00:01:00"),
				transStr2Datetime("2020-08-29 00:01:00"),
				transStr2Datetime("2020-09-29 00:01:00"),
				transStr2Datetime("2020-10-29 00:01:00"),
				transStr2Datetime("2020-11-29 00:01:00"),
				transStr2Datetime("2020-12-29 00:01:00"),
				transStr2Datetime("2021-01-29 00:01:00"),
				transStr2Datetime("2021-02-28 00:01:00"),
			},
		},
		{
			start: "2020-02-29 00:01:00",
			end:   "2021-03-01 00:00:00",
			step:  "1 year",
			res: []types.Datetime{
				transStr2Datetime("2020-02-29 00:01:00"),
				transStr2Datetime("2021-02-28 00:01:00"),
			},
		},
		{
			start: "2020-02-28 00:01:00",
			end:   "2021-03-01 00:00:00",
			step:  "1 year",
			res: []types.Datetime{
				transStr2Datetime("2020-02-28 00:01:00"),
				transStr2Datetime("2021-02-28 00:01:00"),
			},
		},
	}
	for _, kase := range kases {
		var precision int32
		p1, p2 := getPrecision(kase.start), getPrecision(kase.end)
		if p1 > p2 {
			precision = p1
		} else {
			precision = p2
		}
		res, err := generateDatetime(kase.start, kase.end, kase.step, precision)
		if kase.err {
			require.NotNil(t, err)
			continue
		}
		require.Nil(t, err)
		require.Equal(t, kase.res, res)
	}
}

func transStr2Datetime(s string) types.Datetime {
	precision := getPrecision(s)
	t, err := types.ParseDatetime(s, precision)
	if err != nil {
		logutil.Errorf("parse timestamp '%s' failed", s)
	}
	return t
}

func getPrecision(s string) int32 {
	var precision int32
	ss := strings.Split(s, ".")
	if len(ss) > 1 {
		precision = int32(len(ss[1]))
	}
	return precision
}

func TestString(t *testing.T) {
	String(nil, new(bytes.Buffer))
}

func TestPrepare(t *testing.T) {
	err := Prepare(nil, new(bytes.Buffer))
	require.Nil(t, err)
}
func TestGenStep(t *testing.T) {
	kase := "10 hour"
	num, tp, err := genStep(kase)
	require.Nil(t, err)
	require.Equal(t, int64(10), num)
	require.Equal(t, types.Hour, tp)
	kase = "10 houx"
	_, _, err = genStep(kase)
	require.NotNil(t, err)
	kase = "hour"
	_, _, err = genStep(kase)
	require.NotNil(t, err)
	kase = "989829829129131939147193 hour"
	_, _, err = genStep(kase)
	require.NotNil(t, err)
}

func TestCall(t *testing.T) {
	proc := testutil.NewProc()
	beforeCall := proc.Mp().CurrNB()
	arg := &Argument{
		Es: &Param{
			Attrs: []string{"result"},
			Cols: []*plan.ColDef{
				{
					Name: "result",
					Typ:  &plan.Type{},
				},
			},
		},
	}
	param := arg.Es
	proc.SetInputBatch(nil)
	end, err := Call(0, proc, arg)
	require.Nil(t, err)
	require.Equal(t, true, end)

	param.Cols[0].Typ.Id = int32(types.T_int32)
	bat := makeBatch([]string{"1", "10"}, proc)
	proc.SetInputBatch(bat)
	end, err = Call(0, proc, arg)
	require.Nil(t, err)
	require.Equal(t, false, end)
	require.Equal(t, 10, proc.InputBatch().GetVector(0).Length())
	proc.InputBatch().Clean(proc.Mp())
	bat.Clean(proc.Mp())

	param.Cols[0].Typ.Id = int32(types.T_int64)
	bat = makeBatch([]string{"654345676543", "654345676549"}, proc)
	proc.SetInputBatch(bat)
	end, err = Call(0, proc, arg)
	require.Nil(t, err)
	require.Equal(t, false, end)
	require.Equal(t, 7, proc.InputBatch().GetVector(0).Length())
	proc.InputBatch().Clean(proc.Mp())
	bat.Clean(proc.Mp())

	param.Cols[0].Typ.Id = int32(types.T_datetime)
	param.Cols[0].Typ.Precision = 0
	bat = makeBatch([]string{"2020-01-01 00:00:00", "2020-01-01 00:00:59", "1 second"}, proc)
	proc.SetInputBatch(bat)
	end, err = Call(0, proc, arg)
	require.Nil(t, err)
	require.Equal(t, false, end)
	require.Equal(t, 60, proc.InputBatch().GetVector(0).Length())
	proc.InputBatch().Clean(proc.Mp())
	bat.Clean(proc.Mp())

	bat = makeBatch([]string{"2020-01-01 00:00:-1", "2020-01-01 00:00:59", "1 second"}, proc)
	proc.SetInputBatch(bat)
	end, err = Call(0, proc, arg)
	require.NotNil(t, err)
	require.Equal(t, false, end)
	proc.InputBatch().Clean(proc.Mp())

	require.Equal(t, beforeCall, proc.Mp().CurrNB())

}

func makeBatch(arg []string, proc *process.Process) *batch.Batch {
	dt, _ := json.Marshal(arg)
	b := batch.New(true, []string{"result"})
	b.Cnt = 1
	b.InitZsOne(1)
	b.Vecs[0] = vector.New(types.Type{Oid: types.T_varchar})
	err := b.Vecs[0].Append(dt, false, proc.Mp())
	if err != nil {
		b.Clean(proc.Mp())
		logutil.Errorf("set bytes at failed, err: %v", err)
	}
	return b
}
