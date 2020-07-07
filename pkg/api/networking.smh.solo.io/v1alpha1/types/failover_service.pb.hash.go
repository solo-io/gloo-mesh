// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha1/failover_service.proto

package types

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/mitchellh/hashstructure"
	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
func (m *FailoverServiceSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.smh.solo.io.github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/types.FailoverServiceSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetHostname())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetNamespace())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetPort()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetPort(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetCluster())); err != nil {
		return 0, err
	}

	for _, v := range m.GetFailoverServices() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *FailoverServiceStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.smh.solo.io.github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/types.FailoverServiceStatus")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetObservedGeneration())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTranslationStatus()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetTranslationStatus(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetTranslatorErrors() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	if h, ok := interface{}(m.GetValidationStatus()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetValidationStatus(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *FailoverServiceSpec_Port) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.smh.solo.io.github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/types.FailoverServiceSpec_Port")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetPort())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetProtocol())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *FailoverServiceStatus_TranslatorError) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("networking.smh.solo.io.github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/types.FailoverServiceStatus_TranslatorError")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetTranslatorId())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetErrorMessage())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
