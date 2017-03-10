// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/server/status/status.proto
// DO NOT EDIT!

/*
	Package status is a generated protocol buffer package.

	It is generated from these files:
		cockroach/pkg/server/status/status.proto

	It has these top-level messages:
		StoreStatus
		NodeStatus
*/
package status

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_build "github.com/cockroachdb/cockroach/pkg/build"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// StoreStatus records the most recent values of metrics for a store.
type StoreStatus struct {
	Desc    cockroach_roachpb.StoreDescriptor `protobuf:"bytes,1,opt,name=desc" json:"desc"`
	Metrics map[string]float64                `protobuf:"bytes,2,rep,name=metrics" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
}

func (m *StoreStatus) Reset()                    { *m = StoreStatus{} }
func (m *StoreStatus) String() string            { return proto.CompactTextString(m) }
func (*StoreStatus) ProtoMessage()               {}
func (*StoreStatus) Descriptor() ([]byte, []int) { return fileDescriptorStatus, []int{0} }

// NodeStatus records the most recent values of metrics for a node.
type NodeStatus struct {
	Desc          cockroach_roachpb.NodeDescriptor `protobuf:"bytes,1,opt,name=desc" json:"desc"`
	BuildInfo     cockroach_build.Info             `protobuf:"bytes,2,opt,name=build_info,json=buildInfo" json:"build_info"`
	StartedAt     int64                            `protobuf:"varint,3,opt,name=started_at,json=startedAt" json:"started_at"`
	UpdatedAt     int64                            `protobuf:"varint,4,opt,name=updated_at,json=updatedAt" json:"updated_at"`
	Metrics       map[string]float64               `protobuf:"bytes,5,rep,name=metrics" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	StoreStatuses []StoreStatus                    `protobuf:"bytes,6,rep,name=store_statuses,json=storeStatuses" json:"store_statuses"`
	Args          []string                         `protobuf:"bytes,7,rep,name=args" json:"args,omitempty"`
	Env           []string                         `protobuf:"bytes,8,rep,name=env" json:"env,omitempty"`
}

func (m *NodeStatus) Reset()                    { *m = NodeStatus{} }
func (m *NodeStatus) String() string            { return proto.CompactTextString(m) }
func (*NodeStatus) ProtoMessage()               {}
func (*NodeStatus) Descriptor() ([]byte, []int) { return fileDescriptorStatus, []int{1} }

func init() {
	proto.RegisterType((*StoreStatus)(nil), "cockroach.server.status.StoreStatus")
	proto.RegisterType((*NodeStatus)(nil), "cockroach.server.status.NodeStatus")
}
func (m *StoreStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintStatus(dAtA, i, uint64(m.Desc.Size()))
	n1, err := m.Desc.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Metrics) > 0 {
		keysForMetrics := make([]string, 0, len(m.Metrics))
		for k := range m.Metrics {
			keysForMetrics = append(keysForMetrics, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForMetrics)
		for _, k := range keysForMetrics {
			dAtA[i] = 0x12
			i++
			v := m.Metrics[string(k)]
			mapSize := 1 + len(k) + sovStatus(uint64(len(k))) + 1 + 8
			i = encodeVarintStatus(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintStatus(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x11
			i++
			i = encodeFixed64Status(dAtA, i, uint64(math.Float64bits(float64(v))))
		}
	}
	return i, nil
}

func (m *NodeStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintStatus(dAtA, i, uint64(m.Desc.Size()))
	n2, err := m.Desc.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x12
	i++
	i = encodeVarintStatus(dAtA, i, uint64(m.BuildInfo.Size()))
	n3, err := m.BuildInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x18
	i++
	i = encodeVarintStatus(dAtA, i, uint64(m.StartedAt))
	dAtA[i] = 0x20
	i++
	i = encodeVarintStatus(dAtA, i, uint64(m.UpdatedAt))
	if len(m.Metrics) > 0 {
		keysForMetrics := make([]string, 0, len(m.Metrics))
		for k := range m.Metrics {
			keysForMetrics = append(keysForMetrics, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForMetrics)
		for _, k := range keysForMetrics {
			dAtA[i] = 0x2a
			i++
			v := m.Metrics[string(k)]
			mapSize := 1 + len(k) + sovStatus(uint64(len(k))) + 1 + 8
			i = encodeVarintStatus(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintStatus(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x11
			i++
			i = encodeFixed64Status(dAtA, i, uint64(math.Float64bits(float64(v))))
		}
	}
	if len(m.StoreStatuses) > 0 {
		for _, msg := range m.StoreStatuses {
			dAtA[i] = 0x32
			i++
			i = encodeVarintStatus(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			dAtA[i] = 0x3a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Env) > 0 {
		for _, s := range m.Env {
			dAtA[i] = 0x42
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeFixed64Status(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Status(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintStatus(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StoreStatus) Size() (n int) {
	var l int
	_ = l
	l = m.Desc.Size()
	n += 1 + l + sovStatus(uint64(l))
	if len(m.Metrics) > 0 {
		for k, v := range m.Metrics {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovStatus(uint64(len(k))) + 1 + 8
			n += mapEntrySize + 1 + sovStatus(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *NodeStatus) Size() (n int) {
	var l int
	_ = l
	l = m.Desc.Size()
	n += 1 + l + sovStatus(uint64(l))
	l = m.BuildInfo.Size()
	n += 1 + l + sovStatus(uint64(l))
	n += 1 + sovStatus(uint64(m.StartedAt))
	n += 1 + sovStatus(uint64(m.UpdatedAt))
	if len(m.Metrics) > 0 {
		for k, v := range m.Metrics {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovStatus(uint64(len(k))) + 1 + 8
			n += mapEntrySize + 1 + sovStatus(uint64(mapEntrySize))
		}
	}
	if len(m.StoreStatuses) > 0 {
		for _, e := range m.StoreStatuses {
			l = e.Size()
			n += 1 + l + sovStatus(uint64(l))
		}
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			l = len(s)
			n += 1 + l + sovStatus(uint64(l))
		}
	}
	if len(m.Env) > 0 {
		for _, s := range m.Env {
			l = len(s)
			n += 1 + l + sovStatus(uint64(l))
		}
	}
	return n
}

func sovStatus(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStatus(x uint64) (n int) {
	return sovStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoreStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatus
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoreStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Desc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthStatus
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Metrics == nil {
				m.Metrics = make(map[string]float64)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStatus
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapvaluetemp uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				iNdEx += 8
				mapvaluetemp = uint64(dAtA[iNdEx-8])
				mapvaluetemp |= uint64(dAtA[iNdEx-7]) << 8
				mapvaluetemp |= uint64(dAtA[iNdEx-6]) << 16
				mapvaluetemp |= uint64(dAtA[iNdEx-5]) << 24
				mapvaluetemp |= uint64(dAtA[iNdEx-4]) << 32
				mapvaluetemp |= uint64(dAtA[iNdEx-3]) << 40
				mapvaluetemp |= uint64(dAtA[iNdEx-2]) << 48
				mapvaluetemp |= uint64(dAtA[iNdEx-1]) << 56
				mapvalue := math.Float64frombits(mapvaluetemp)
				m.Metrics[mapkey] = mapvalue
			} else {
				var mapvalue float64
				m.Metrics[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatus
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodeStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatus
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Desc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuildInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BuildInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartedAt", wireType)
			}
			m.StartedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartedAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthStatus
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Metrics == nil {
				m.Metrics = make(map[string]float64)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStatus
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapvaluetemp uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				iNdEx += 8
				mapvaluetemp = uint64(dAtA[iNdEx-8])
				mapvaluetemp |= uint64(dAtA[iNdEx-7]) << 8
				mapvaluetemp |= uint64(dAtA[iNdEx-6]) << 16
				mapvaluetemp |= uint64(dAtA[iNdEx-5]) << 24
				mapvaluetemp |= uint64(dAtA[iNdEx-4]) << 32
				mapvaluetemp |= uint64(dAtA[iNdEx-3]) << 40
				mapvaluetemp |= uint64(dAtA[iNdEx-2]) << 48
				mapvaluetemp |= uint64(dAtA[iNdEx-1]) << 56
				mapvalue := math.Float64frombits(mapvaluetemp)
				m.Metrics[mapkey] = mapvalue
			} else {
				var mapvalue float64
				m.Metrics[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreStatuses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoreStatuses = append(m.StoreStatuses, StoreStatus{})
			if err := m.StoreStatuses[len(m.StoreStatuses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Env", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Env = append(m.Env, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatus
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStatus
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStatus
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStatus
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStatus(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStatus = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStatus   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/server/status/status.proto", fileDescriptorStatus) }

var fileDescriptorStatus = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xcd, 0xc6, 0x69, 0x4b, 0x26, 0x80, 0xd0, 0xaa, 0xc0, 0x2a, 0x07, 0x63, 0x42, 0x0f, 0x3e,
	0xad, 0x21, 0x27, 0x14, 0xb8, 0xb4, 0x82, 0x03, 0x20, 0x90, 0x48, 0x6f, 0x5c, 0xa2, 0xad, 0xbd,
	0x0d, 0x56, 0x5a, 0xaf, 0xb5, 0xbb, 0x8e, 0xd4, 0x7f, 0xc1, 0xcf, 0xca, 0x91, 0x03, 0x07, 0x0e,
	0x08, 0x81, 0xf9, 0x23, 0x68, 0x3f, 0x8c, 0x1d, 0xa9, 0x91, 0x90, 0x7a, 0xf2, 0xf8, 0xcd, 0x7b,
	0x33, 0xf3, 0xde, 0x42, 0x9c, 0x8a, 0x74, 0x25, 0x05, 0x4b, 0x3f, 0x27, 0xe5, 0x6a, 0x99, 0x28,
	0x2e, 0xd7, 0x5c, 0x26, 0x4a, 0x33, 0x5d, 0x29, 0xff, 0xa1, 0xa5, 0x14, 0x5a, 0xe0, 0x87, 0xff,
	0x98, 0xd4, 0xb1, 0xa8, 0x6b, 0x8f, 0x8f, 0xb6, 0x47, 0xd8, 0xaa, 0x3c, 0x4b, 0x2e, 0xb9, 0x66,
	0x19, 0xd3, 0xcc, 0xc9, 0xc7, 0xe1, 0x36, 0xeb, 0xac, 0xca, 0x2f, 0xb2, 0x24, 0x2f, 0xce, 0x85,
	0xef, 0x1f, 0x2e, 0xc5, 0x52, 0xd8, 0x32, 0x31, 0x95, 0x43, 0x27, 0xdf, 0x10, 0x8c, 0x4e, 0xb5,
	0x90, 0xfc, 0xd4, 0xee, 0xc2, 0x2f, 0x61, 0x90, 0x71, 0x95, 0x12, 0x14, 0xa1, 0x78, 0x34, 0x9d,
	0xd0, 0xf6, 0x26, 0xbf, 0x96, 0x5a, 0xf6, 0x2b, 0xae, 0x52, 0x99, 0x97, 0x5a, 0xc8, 0x93, 0xc1,
	0xe6, 0xe7, 0xa3, 0xde, 0xdc, 0xaa, 0xf0, 0x3b, 0x38, 0xb8, 0xe4, 0x5a, 0xe6, 0xa9, 0x22, 0xfd,
	0x28, 0x88, 0x47, 0xd3, 0x67, 0x74, 0x87, 0x29, 0xda, 0x59, 0x4a, 0xdf, 0x3b, 0xcd, 0xeb, 0x42,
	0xcb, 0xab, 0x79, 0x33, 0x61, 0x3c, 0x83, 0xdb, 0xdd, 0x06, 0xbe, 0x07, 0xc1, 0x8a, 0x5f, 0xd9,
	0xcb, 0x86, 0x73, 0x53, 0xe2, 0x43, 0xd8, 0x5b, 0xb3, 0x8b, 0x8a, 0x93, 0x7e, 0x84, 0x62, 0x34,
	0x77, 0x3f, 0xb3, 0xfe, 0x73, 0x34, 0xf9, 0x11, 0x00, 0x7c, 0x10, 0x59, 0xe3, 0xea, 0xc5, 0x96,
	0xab, 0xc7, 0xd7, 0xb8, 0x32, 0xe4, 0x1d, 0xa6, 0x66, 0x00, 0x36, 0xcc, 0x85, 0x09, 0xd3, 0xae,
	0x1a, 0x4d, 0xef, 0x77, 0x46, 0xd8, 0x26, 0x7d, 0x53, 0x9c, 0x0b, 0x2f, 0x1b, 0x5a, 0xc4, 0x00,
	0xf8, 0x09, 0x80, 0xd2, 0x4c, 0x6a, 0x9e, 0x2d, 0x98, 0x26, 0x41, 0x84, 0xe2, 0xa0, 0x21, 0x79,
	0xfc, 0x58, 0x1b, 0x52, 0x55, 0x66, 0xcc, 0x93, 0x06, 0x5d, 0x92, 0xc7, 0x8f, 0x35, 0x7e, 0xdb,
	0x46, 0xbb, 0x67, 0xa3, 0x7d, 0xba, 0x33, 0xda, 0xd6, 0xf8, 0xf5, 0xc9, 0xe2, 0x8f, 0x70, 0x57,
	0x99, 0xf8, 0x17, 0x4e, 0xc0, 0x15, 0xd9, 0xb7, 0x23, 0x8f, 0xfe, 0xe7, 0xb5, 0xfc, 0x69, 0x77,
	0x54, 0x0b, 0x71, 0x85, 0x09, 0x0c, 0x98, 0x5c, 0x2a, 0x72, 0x10, 0x05, 0xf1, 0xb0, 0x89, 0xcf,
	0x20, 0xf8, 0x01, 0x04, 0xbc, 0x58, 0x93, 0x5b, 0x9d, 0x86, 0x01, 0x6e, 0xf2, 0xbc, 0x27, 0xd1,
	0xe6, 0x77, 0xd8, 0xdb, 0xd4, 0x21, 0xfa, 0x5a, 0x87, 0xe8, 0x7b, 0x1d, 0xa2, 0x5f, 0x75, 0x88,
	0xbe, 0xfc, 0x09, 0x7b, 0x9f, 0xf6, 0xdd, 0xc1, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x11, 0xe5,
	0xa8, 0x24, 0x77, 0x03, 0x00, 0x00,
}
