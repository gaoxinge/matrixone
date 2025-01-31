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
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/sql/plan"
)

type Number interface {
	int32 | int64 | types.Datetime
}

type Param struct {
	Attrs []string
	Cols  []*plan.ColDef
	start string
	end   string
	step  string
}

type Argument struct {
	Es *Param
}

var (

//	timeStampCol = &plan.ColDef{
//		Name: "generate_series",
//		Typ: &plan.Type{
//			Id:       int32(types.T_timestamp),
//			Nullable: false,
//		},
//	}
//
//	intCol = &plan.ColDef{
//		Name: "generate_series",
//		Typ: &plan.Type{
//			Id:       int32(types.T_int64),
//			Nullable: false,
//		},
//	}
)
