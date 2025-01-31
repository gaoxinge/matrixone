// Copyright 2021 - 2022 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package debug

import (
	"context"
	"strings"

	pb "github.com/matrixorigin/matrixone/pkg/pb/debug"
	"github.com/matrixorigin/matrixone/pkg/pb/txn"
	"github.com/matrixorigin/matrixone/pkg/vm/engine"
)

type serviceType string

var (
	dn serviceType = "DN"

	supportedServiceTypes = map[serviceType]struct{}{
		dn: {},
	}
)

var (
	// register all supported debug command here
	supportedCmds = map[string]handleFunc{
		strings.ToUpper(pb.CmdMethod_Ping.String()): handlePing(),
	}
)

type requestSender = func(context.Context, []txn.CNOpRequest) ([]txn.CNOpResponse, error)

type handleFunc func(ctx context.Context,
	service serviceType,
	parameter string,
	sender requestSender,
	clusterDetailsGetter engine.GetClusterDetailsFunc) (pb.DebugResult, error)
