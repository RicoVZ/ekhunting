// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file.proto

package onemon

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FileAccess int32

const (
	FileAccess_Unknown       FileAccess = 0
	FileAccess_OpenRead      FileAccess = 1
	FileAccess_OpenModify    FileAccess = 2
	FileAccess_CreateRead    FileAccess = 3
	FileAccess_CreateModify  FileAccess = 4
	FileAccess_Delete        FileAccess = 5
	FileAccess_DeleteDelayed FileAccess = 6
	FileAccess_Rename        FileAccess = 7
	FileAccess_Truncate      FileAccess = 8
)

var FileAccess_name = map[int32]string{
	0: "Unknown",
	1: "OpenRead",
	2: "OpenModify",
	3: "CreateRead",
	4: "CreateModify",
	5: "Delete",
	6: "DeleteDelayed",
	7: "Rename",
	8: "Truncate",
}
var FileAccess_value = map[string]int32{
	"Unknown":       0,
	"OpenRead":      1,
	"OpenModify":    2,
	"CreateRead":    3,
	"CreateModify":  4,
	"Delete":        5,
	"DeleteDelayed": 6,
	"Rename":        7,
	"Truncate":      8,
}

func (x FileAccess) String() string {
	return proto.EnumName(FileAccess_name, int32(x))
}
func (FileAccess) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// ID: 8
type File struct {
	Ts      uint32     `protobuf:"varint,7,opt,name=ts" json:"ts,omitempty"`
	Pid     uint64     `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
	Kind    FileAccess `protobuf:"varint,2,opt,name=kind,enum=onemon.FileAccess" json:"kind,omitempty"`
	Status  uint32     `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	Id      uint64     `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	Srcpath string     `protobuf:"bytes,5,opt,name=srcpath" json:"srcpath,omitempty"`
	Dstpath string     `protobuf:"bytes,6,opt,name=dstpath" json:"dstpath,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *File) GetTs() uint32 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *File) GetPid() uint64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *File) GetKind() FileAccess {
	if m != nil {
		return m.Kind
	}
	return FileAccess_Unknown
}

func (m *File) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *File) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *File) GetSrcpath() string {
	if m != nil {
		return m.Srcpath
	}
	return ""
}

func (m *File) GetDstpath() string {
	if m != nil {
		return m.Dstpath
	}
	return ""
}

func init() {
	proto.RegisterType((*File)(nil), "onemon.File")
	proto.RegisterEnum("onemon.FileAccess", FileAccess_name, FileAccess_value)
}

func init() { proto.RegisterFile("file.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
}