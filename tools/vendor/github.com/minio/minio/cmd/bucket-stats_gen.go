package cmd

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BucketReplicationStats) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "PendingSize":
			z.PendingSize, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "PendingSize")
				return
			}
		case "ReplicatedSize":
			z.ReplicatedSize, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "ReplicatedSize")
				return
			}
		case "ReplicaSize":
			z.ReplicaSize, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "ReplicaSize")
				return
			}
		case "FailedSize":
			z.FailedSize, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "FailedSize")
				return
			}
		case "PendingCount":
			z.PendingCount, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "PendingCount")
				return
			}
		case "FailedCount":
			z.FailedCount, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "FailedCount")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BucketReplicationStats) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "PendingSize"
	err = en.Append(0x86, 0xab, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.PendingSize)
	if err != nil {
		err = msgp.WrapError(err, "PendingSize")
		return
	}
	// write "ReplicatedSize"
	err = en.Append(0xae, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.ReplicatedSize)
	if err != nil {
		err = msgp.WrapError(err, "ReplicatedSize")
		return
	}
	// write "ReplicaSize"
	err = en.Append(0xab, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.ReplicaSize)
	if err != nil {
		err = msgp.WrapError(err, "ReplicaSize")
		return
	}
	// write "FailedSize"
	err = en.Append(0xaa, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.FailedSize)
	if err != nil {
		err = msgp.WrapError(err, "FailedSize")
		return
	}
	// write "PendingCount"
	err = en.Append(0xac, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.PendingCount)
	if err != nil {
		err = msgp.WrapError(err, "PendingCount")
		return
	}
	// write "FailedCount"
	err = en.Append(0xab, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.FailedCount)
	if err != nil {
		err = msgp.WrapError(err, "FailedCount")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BucketReplicationStats) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "PendingSize"
	o = append(o, 0x86, 0xab, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendUint64(o, z.PendingSize)
	// string "ReplicatedSize"
	o = append(o, 0xae, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendUint64(o, z.ReplicatedSize)
	// string "ReplicaSize"
	o = append(o, 0xab, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendUint64(o, z.ReplicaSize)
	// string "FailedSize"
	o = append(o, 0xaa, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendUint64(o, z.FailedSize)
	// string "PendingCount"
	o = append(o, 0xac, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendUint64(o, z.PendingCount)
	// string "FailedCount"
	o = append(o, 0xab, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendUint64(o, z.FailedCount)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BucketReplicationStats) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "PendingSize":
			z.PendingSize, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PendingSize")
				return
			}
		case "ReplicatedSize":
			z.ReplicatedSize, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ReplicatedSize")
				return
			}
		case "ReplicaSize":
			z.ReplicaSize, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ReplicaSize")
				return
			}
		case "FailedSize":
			z.FailedSize, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "FailedSize")
				return
			}
		case "PendingCount":
			z.PendingCount, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PendingCount")
				return
			}
		case "FailedCount":
			z.FailedCount, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "FailedCount")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BucketReplicationStats) Msgsize() (s int) {
	s = 1 + 12 + msgp.Uint64Size + 15 + msgp.Uint64Size + 12 + msgp.Uint64Size + 11 + msgp.Uint64Size + 13 + msgp.Uint64Size + 12 + msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BucketStats) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ReplicationStats":
			err = z.ReplicationStats.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "ReplicationStats")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BucketStats) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "ReplicationStats"
	err = en.Append(0x81, 0xb0, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73)
	if err != nil {
		return
	}
	err = z.ReplicationStats.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "ReplicationStats")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BucketStats) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "ReplicationStats"
	o = append(o, 0x81, 0xb0, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73)
	o, err = z.ReplicationStats.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "ReplicationStats")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BucketStats) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ReplicationStats":
			bts, err = z.ReplicationStats.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "ReplicationStats")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BucketStats) Msgsize() (s int) {
	s = 1 + 17 + z.ReplicationStats.Msgsize()
	return
}