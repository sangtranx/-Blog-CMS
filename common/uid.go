package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"strconv"
	"strings"
)

// UID is method to generate an virtual unique identifer for whose system
// its structure contains 62 bits : LocalID - ObjectType - ShardId
// 32 bits for LocalID, max (2^32)
// 10 bits for ObjectType
// 18 bits for ShardID

type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(localID uint32, objectType int, shardID uint32) UID {
	return UID{
		localID:    localID,
		objectType: objectType,
		shardID:    shardID,
	}
}

// shardID : 1, Object: 1. ID : 1 => 0001 0001 0001
// 1 << 8 = 0001 0000 0000
// 1 << 4 =         1 0000
// 1 << 0 =              1

func (uid UID) String() string {
	val := uint64(uid.localID)<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardID)<<0
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func (uid UID) GetLocalID() uint32 {
	return uid.localID
}

func (uid UID) GetObjectType() int {
	return uid.objectType
}

func (uid UID) GetShardID() uint32 {
	return uid.shardID
}

func DecomposeUID(s string) (UID, error) {

	uid, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return UID{}, err
	}

	if (1 << 18) > uid {
		return UID{}, errors.New("Wrong uid")
	}

	// x = 1110 1110 0101 => x >> 4 = 1110 1110 & 0000 1111 = 1110
	u := UID{
		localID:    uint32(uid >> 28),
		objectType: int(uid >> 18 & 0x3FF),
		shardID:    uint32(uid >> 0 & 0x3FFFF),
	}

	return u, nil
}

func FromBase58(s string) (UID, error) {
	return DecomposeUID(string(base58.Decode(s)))
}

func (uid *UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uid.String())), nil
}

func (uid *UID) UnmarshalJSON(data []byte) error {

	decodeUID, err := FromBase58(strings.Replace(string(data), "\"", "\"", -1))

	if err != nil {
		return err
	}

	uid.localID = decodeUID.GetLocalID()
	uid.objectType = int(decodeUID.GetObjectType())
	uid.shardID = decodeUID.GetShardID()

	return nil
}

func (uid *UID) Value() (driver.Value, error) {

	if uid == nil {
		return nil, nil
	}

	return int64(uid.localID), nil
}

func (uid *UID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var i uint32

	switch t := value.(type) {
	case int:
		i = uint32(t)
	case int8:
		i = uint32(t)
	case int16:
		i = uint32(t)
	case int32:
		i = uint32(t)
	case int64:
		if t < 0 || t > int64(^uint32(0)) {
			return fmt.Errorf("int64 value out of range for uint32: %d", t)
		}
		i = uint32(t)
	case uint8:
		i = uint32(t)
	case uint16:
		i = uint32(t)
	case uint32:
		i = t
	case uint64:
		if t > uint64(^uint32(0)) {
			return fmt.Errorf("uint64 value out of range for uint32: %d", t)
		}
		i = uint32(t)
	case []byte:
		// Assuming the byte slice represents a string-encoded integer
		val, err := strconv.ParseUint(string(t), 10, 32)
		if err != nil {
			return fmt.Errorf("failed to parse []byte as uint32: %w", err)
		}
		i = uint32(val)
	case string:
		// Assuming the string represents an integer
		val, err := strconv.ParseUint(t, 10, 32)
		if err != nil {
			return fmt.Errorf("failed to parse string as uint32: %w", err)
		}
		i = uint32(val)
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}

	uid.localID = i

	return nil
}
