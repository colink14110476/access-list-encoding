// Copyright © 2020 Attestant Limited.
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

package types

// Need to `go install github.com/ferranbt/fastssz/sszgen@v0.1.3` for this to work.
// Assumes 	github.com/attestantio/go-builder-client v0.6.1	github.com/attestantio/go-eth2-client v0.25.0 github.com/holiman/uint256 v1.3.2

//go:generate rm -f ./accesslist_submit_block_request_encoding.go ./accesslist_submit_block_request_v4_encoding.go
//nolint:revive

// go:generate sszgen -include ../../go-eth2-client/spec/deneb,../../go-eth2-client/spec/electra,../../go-eth2-client/spec/phase0,../../go-builder-client/api/deneb,../../go-builder-client/api/v1,../../go-eth2-client/spec/bellatrix,../../go-eth2-client/spec/capella,../../uint256 --path . --objs AccesslistSubmitBlockRequest
// go:generate goimports -w accesslist_submit_block_request_encoding.go accesslist_submit_block_request_v4_encoding.go
