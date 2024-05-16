// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package domain

import (
	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Reward         protoreflect.MessageDescriptor
	fd_Reward_domain  protoreflect.FieldDescriptor
	fd_Reward_isClaim protoreflect.FieldDescriptor
	fd_Reward_amount  protoreflect.FieldDescriptor
)

func init() {
	file_opkit_domain_reward_proto_init()
	md_Reward = File_opkit_domain_reward_proto.Messages().ByName("Reward")
	fd_Reward_domain = md_Reward.Fields().ByName("domain")
	fd_Reward_isClaim = md_Reward.Fields().ByName("isClaim")
	fd_Reward_amount = md_Reward.Fields().ByName("amount")
}

var _ protoreflect.Message = (*fastReflection_Reward)(nil)

type fastReflection_Reward Reward

func (x *Reward) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Reward)(x)
}

func (x *Reward) slowProtoReflect() protoreflect.Message {
	mi := &file_opkit_domain_reward_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Reward_messageType fastReflection_Reward_messageType
var _ protoreflect.MessageType = fastReflection_Reward_messageType{}

type fastReflection_Reward_messageType struct{}

func (x fastReflection_Reward_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Reward)(nil)
}
func (x fastReflection_Reward_messageType) New() protoreflect.Message {
	return new(fastReflection_Reward)
}
func (x fastReflection_Reward_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Reward
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Reward) Descriptor() protoreflect.MessageDescriptor {
	return md_Reward
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Reward) Type() protoreflect.MessageType {
	return _fastReflection_Reward_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Reward) New() protoreflect.Message {
	return new(fastReflection_Reward)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Reward) Interface() protoreflect.ProtoMessage {
	return (*Reward)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Reward) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Domain != "" {
		value := protoreflect.ValueOfString(x.Domain)
		if !f(fd_Reward_domain, value) {
			return
		}
	}
	if x.IsClaim != false {
		value := protoreflect.ValueOfBool(x.IsClaim)
		if !f(fd_Reward_isClaim, value) {
			return
		}
	}
	if x.Amount != nil {
		value := protoreflect.ValueOfMessage(x.Amount.ProtoReflect())
		if !f(fd_Reward_amount, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Reward) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "opkit.domain.Reward.domain":
		return x.Domain != ""
	case "opkit.domain.Reward.isClaim":
		return x.IsClaim != false
	case "opkit.domain.Reward.amount":
		return x.Amount != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: opkit.domain.Reward"))
		}
		panic(fmt.Errorf("message opkit.domain.Reward does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Reward) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "opkit.domain.Reward.domain":
		x.Domain = ""
	case "opkit.domain.Reward.isClaim":
		x.IsClaim = false
	case "opkit.domain.Reward.amount":
		x.Amount = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: opkit.domain.Reward"))
		}
		panic(fmt.Errorf("message opkit.domain.Reward does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Reward) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "opkit.domain.Reward.domain":
		value := x.Domain
		return protoreflect.ValueOfString(value)
	case "opkit.domain.Reward.isClaim":
		value := x.IsClaim
		return protoreflect.ValueOfBool(value)
	case "opkit.domain.Reward.amount":
		value := x.Amount
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: opkit.domain.Reward"))
		}
		panic(fmt.Errorf("message opkit.domain.Reward does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Reward) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "opkit.domain.Reward.domain":
		x.Domain = value.Interface().(string)
	case "opkit.domain.Reward.isClaim":
		x.IsClaim = value.Bool()
	case "opkit.domain.Reward.amount":
		x.Amount = value.Message().Interface().(*v1beta1.Coin)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: opkit.domain.Reward"))
		}
		panic(fmt.Errorf("message opkit.domain.Reward does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Reward) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "opkit.domain.Reward.amount":
		if x.Amount == nil {
			x.Amount = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.Amount.ProtoReflect())
	case "opkit.domain.Reward.domain":
		panic(fmt.Errorf("field domain of message opkit.domain.Reward is not mutable"))
	case "opkit.domain.Reward.isClaim":
		panic(fmt.Errorf("field isClaim of message opkit.domain.Reward is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: opkit.domain.Reward"))
		}
		panic(fmt.Errorf("message opkit.domain.Reward does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Reward) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "opkit.domain.Reward.domain":
		return protoreflect.ValueOfString("")
	case "opkit.domain.Reward.isClaim":
		return protoreflect.ValueOfBool(false)
	case "opkit.domain.Reward.amount":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: opkit.domain.Reward"))
		}
		panic(fmt.Errorf("message opkit.domain.Reward does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Reward) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in opkit.domain.Reward", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Reward) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Reward) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Reward) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Reward) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Reward)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Domain)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.IsClaim {
			n += 2
		}
		if x.Amount != nil {
			l = options.Size(x.Amount)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Reward)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Amount != nil {
			encoded, err := options.Marshal(x.Amount)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if x.IsClaim {
			i--
			if x.IsClaim {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
		}
		if len(x.Domain) > 0 {
			i -= len(x.Domain)
			copy(dAtA[i:], x.Domain)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Domain)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Reward)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Reward: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Reward: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Domain = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field IsClaim", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.IsClaim = bool(v != 0)
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Amount == nil {
					x.Amount = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Amount); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: opkit/domain/reward.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Reward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain  string        `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	IsClaim bool          `protobuf:"varint,2,opt,name=isClaim,proto3" json:"isClaim,omitempty"`
	Amount  *v1beta1.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Reward) Reset() {
	*x = Reward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opkit_domain_reward_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reward) ProtoMessage() {}

// Deprecated: Use Reward.ProtoReflect.Descriptor instead.
func (*Reward) Descriptor() ([]byte, []int) {
	return file_opkit_domain_reward_proto_rawDescGZIP(), []int{0}
}

func (x *Reward) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Reward) GetIsClaim() bool {
	if x != nil {
		return x.IsClaim
	}
	return false
}

func (x *Reward) GetAmount() *v1beta1.Coin {
	if x != nil {
		return x.Amount
	}
	return nil
}

var File_opkit_domain_reward_proto protoreflect.FileDescriptor

var file_opkit_domain_reward_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x70, 0x6b, 0x69, 0x74, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x70, 0x6b,
	0x69, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x73, 0x0a, 0x06, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x37, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x8f, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x70, 0x6b,
	0x69, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x0b, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x6b, 0x69, 0x74,
	0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0xa2, 0x02, 0x03, 0x4f, 0x44, 0x58, 0xaa, 0x02, 0x0c,
	0x4f, 0x70, 0x6b, 0x69, 0x74, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0xca, 0x02, 0x0c, 0x4f,
	0x70, 0x6b, 0x69, 0x74, 0x5c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0xe2, 0x02, 0x18, 0x4f, 0x70,
	0x6b, 0x69, 0x74, 0x5c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x4f, 0x70, 0x6b, 0x69, 0x74, 0x3a, 0x3a,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_opkit_domain_reward_proto_rawDescOnce sync.Once
	file_opkit_domain_reward_proto_rawDescData = file_opkit_domain_reward_proto_rawDesc
)

func file_opkit_domain_reward_proto_rawDescGZIP() []byte {
	file_opkit_domain_reward_proto_rawDescOnce.Do(func() {
		file_opkit_domain_reward_proto_rawDescData = protoimpl.X.CompressGZIP(file_opkit_domain_reward_proto_rawDescData)
	})
	return file_opkit_domain_reward_proto_rawDescData
}

var file_opkit_domain_reward_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_opkit_domain_reward_proto_goTypes = []interface{}{
	(*Reward)(nil),       // 0: opkit.domain.Reward
	(*v1beta1.Coin)(nil), // 1: cosmos.base.v1beta1.Coin
}
var file_opkit_domain_reward_proto_depIdxs = []int32{
	1, // 0: opkit.domain.Reward.amount:type_name -> cosmos.base.v1beta1.Coin
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_opkit_domain_reward_proto_init() }
func file_opkit_domain_reward_proto_init() {
	if File_opkit_domain_reward_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opkit_domain_reward_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reward); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_opkit_domain_reward_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_opkit_domain_reward_proto_goTypes,
		DependencyIndexes: file_opkit_domain_reward_proto_depIdxs,
		MessageInfos:      file_opkit_domain_reward_proto_msgTypes,
	}.Build()
	File_opkit_domain_reward_proto = out.File
	file_opkit_domain_reward_proto_rawDesc = nil
	file_opkit_domain_reward_proto_goTypes = nil
	file_opkit_domain_reward_proto_depIdxs = nil
}
