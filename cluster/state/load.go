// Copyright © 2022-2023 Obol Labs Inc. Licensed under the terms of a Business Source License 1.1

package state

import (
	"encoding/json"
	"os"

	"github.com/obolnetwork/charon/app/errors"
	"github.com/obolnetwork/charon/cluster"
)

// Load loads a cluster state from disk. It supports both legacy lock files and raw DAG files.
func Load(file string, lockCallback func(cluster.Lock) error) (Cluster, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return Cluster{}, errors.Wrap(err, "read file")
	}

	var rawDAG RawDAG
	if err := json.Unmarshal(b, &rawDAG); err != nil {
		return loadLegacyLock(b, lockCallback)
	}

	return Materialise(rawDAG)
}

func loadLegacyLock(input []byte, lockCallback func(cluster.Lock) error) (Cluster, error) {
	var lock cluster.Lock
	if err := json.Unmarshal(input, &lock); err != nil {
		return Cluster{}, errors.Wrap(err, "unmarshal legacy lock")
	}

	if lockCallback != nil {
		if err := lockCallback(lock); err != nil {
			return Cluster{}, err
		}
	}

	legacy, err := NewLegacyLock(lock)
	if err != nil {
		return Cluster{}, errors.Wrap(err, "create legacy lock")
	}

	return Materialise([]SignedMutation{legacy})
}