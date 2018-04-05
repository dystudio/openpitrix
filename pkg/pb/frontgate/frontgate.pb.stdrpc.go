// Code generated by protoc-gen-stdrpc. DO NOT EDIT.
//
// plugin: https://github.com/chai2010/protorpc/tree/master/protoc-gen-plugin
// plugin: https://github.com/chai2010/protorpc/tree/master/protoc-gen-stdrpc
//
// source: frontgate.proto

package openpitrix_frontgate

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/golang/protobuf/proto"
)

var (
	_ = fmt.Sprint
	_ = io.Reader(nil)
	_ = log.Print
	_ = net.Addr(nil)
	_ = rpc.Call{}
	_ = time.Second

	_ = proto.String
)

type FrontgateService interface {
	CloseChannel(in *Empty, out *Empty) error
	GetInfo(in *Empty, out *Info) error
	GetConfdInfo(in *GetConfdInfoRequest, out *ConfdInfo) error
	StartConfd(in *StartConfdRequest, out *Empty) error
	StopConfd(in *StopConfdRequest, out *Empty) error
	RegisterMetadata(in *RegisterMetadataRequest, out *Empty) error
	DeregisterMetadata(in *DeregisterMetadataRequest, out *Empty) error
	RegisterCmd(in *RegisterCmdRequest, out *Empty) error
	DeregisterCmd(in *DeregisterCmdRequest, out *Empty) error
	GetSubTaskResult(in *GetSubTaskResultRequest, out *SubTaskResult) error
	ReportSubTaskResult(in *SubTaskResult, out *Empty) error
}

// AcceptFrontgateServiceClient accepts connections on the listener and serves requests
// for each incoming connection.  Accept blocks; the caller typically
// invokes it in a go statement.
func AcceptFrontgateServiceClient(lis net.Listener, x FrontgateService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("FrontgateService", x); err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeConn(conn)
	}
}

// RegisterFrontgateService publish the given FrontgateService implementation on the server.
func RegisterFrontgateService(srv *rpc.Server, x FrontgateService) error {
	if err := srv.RegisterName("FrontgateService", x); err != nil {
		return err
	}
	return nil
}

// NewFrontgateServiceServer returns a new FrontgateService Server.
func NewFrontgateServiceServer(x FrontgateService) *rpc.Server {
	srv := rpc.NewServer()
	if err := srv.RegisterName("FrontgateService", x); err != nil {
		log.Fatal(err)
	}
	return srv
}

// ListenAndServeFrontgateService listen announces on the local network address laddr
// and serves the given FrontgateService implementation.
func ListenAndServeFrontgateService(network, addr string, x FrontgateService) error {
	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	srv := rpc.NewServer()
	if err := srv.RegisterName("FrontgateService", x); err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeConn(conn)
	}
}

// ServeFrontgateService serves the given FrontgateService implementation.
func ServeFrontgateService(conn io.ReadWriteCloser, x FrontgateService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("FrontgateService", x); err != nil {
		log.Fatal(err)
	}
	srv.ServeConn(conn)
}

type FrontgateServiceClient struct {
	*rpc.Client
}

// NewFrontgateServiceClient returns a FrontgateService stub to handle
// requests to the set of FrontgateService at the other end of the connection.
func NewFrontgateServiceClient(conn io.ReadWriteCloser) *FrontgateServiceClient {
	c := rpc.NewClient(conn)
	return &FrontgateServiceClient{c}
}

func (c *FrontgateServiceClient) CloseChannel(in *Empty) (out *Empty, err error) {
	if in == nil {
		in = new(Empty)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Empty)
	if err = c.Call("FrontgateService.CloseChannel", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncCloseChannel(in *Empty, out *Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Empty)
	}
	return c.Go(
		"FrontgateService.CloseChannel",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) GetInfo(in *Empty) (out *Info, err error) {
	if in == nil {
		in = new(Empty)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Info)
	if err = c.Call("FrontgateService.GetInfo", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncGetInfo(in *Empty, out *Info, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Empty)
	}
	return c.Go(
		"FrontgateService.GetInfo",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) GetConfdInfo(in *GetConfdInfoRequest) (out *ConfdInfo, err error) {
	if in == nil {
		in = new(GetConfdInfoRequest)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ConfdInfo)
	if err = c.Call("FrontgateService.GetConfdInfo", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncGetConfdInfo(in *GetConfdInfoRequest, out *ConfdInfo, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(GetConfdInfoRequest)
	}
	return c.Go(
		"FrontgateService.GetConfdInfo",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) StartConfd(in *StartConfdRequest) (out *Empty, err error) {
	if in == nil {
		in = new(StartConfdRequest)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Empty)
	if err = c.Call("FrontgateService.StartConfd", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncStartConfd(in *StartConfdRequest, out *Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(StartConfdRequest)
	}
	return c.Go(
		"FrontgateService.StartConfd",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) StopConfd(in *StopConfdRequest) (out *Empty, err error) {
	if in == nil {
		in = new(StopConfdRequest)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Empty)
	if err = c.Call("FrontgateService.StopConfd", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncStopConfd(in *StopConfdRequest, out *Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(StopConfdRequest)
	}
	return c.Go(
		"FrontgateService.StopConfd",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) RegisterMetadata(in *RegisterMetadataRequest) (out *Empty, err error) {
	if in == nil {
		in = new(RegisterMetadataRequest)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Empty)
	if err = c.Call("FrontgateService.RegisterMetadata", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncRegisterMetadata(in *RegisterMetadataRequest, out *Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(RegisterMetadataRequest)
	}
	return c.Go(
		"FrontgateService.RegisterMetadata",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) DeregisterMetadata(in *DeregisterMetadataRequest) (out *Empty, err error) {
	if in == nil {
		in = new(DeregisterMetadataRequest)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Empty)
	if err = c.Call("FrontgateService.DeregisterMetadata", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncDeregisterMetadata(in *DeregisterMetadataRequest, out *Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(DeregisterMetadataRequest)
	}
	return c.Go(
		"FrontgateService.DeregisterMetadata",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) RegisterCmd(in *RegisterCmdRequest) (out *Empty, err error) {
	if in == nil {
		in = new(RegisterCmdRequest)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Empty)
	if err = c.Call("FrontgateService.RegisterCmd", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncRegisterCmd(in *RegisterCmdRequest, out *Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(RegisterCmdRequest)
	}
	return c.Go(
		"FrontgateService.RegisterCmd",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) DeregisterCmd(in *DeregisterCmdRequest) (out *Empty, err error) {
	if in == nil {
		in = new(DeregisterCmdRequest)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Empty)
	if err = c.Call("FrontgateService.DeregisterCmd", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncDeregisterCmd(in *DeregisterCmdRequest, out *Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(DeregisterCmdRequest)
	}
	return c.Go(
		"FrontgateService.DeregisterCmd",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) GetSubTaskResult(in *GetSubTaskResultRequest) (out *SubTaskResult, err error) {
	if in == nil {
		in = new(GetSubTaskResultRequest)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(SubTaskResult)
	if err = c.Call("FrontgateService.GetSubTaskResult", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncGetSubTaskResult(in *GetSubTaskResultRequest, out *SubTaskResult, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(GetSubTaskResultRequest)
	}
	return c.Go(
		"FrontgateService.GetSubTaskResult",
		in, out,
		done,
	)
}

func (c *FrontgateServiceClient) ReportSubTaskResult(in *SubTaskResult) (out *Empty, err error) {
	if in == nil {
		in = new(SubTaskResult)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Empty)
	if err = c.Call("FrontgateService.ReportSubTaskResult", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *FrontgateServiceClient) AsyncReportSubTaskResult(in *SubTaskResult, out *Empty, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(SubTaskResult)
	}
	return c.Go(
		"FrontgateService.ReportSubTaskResult",
		in, out,
		done,
	)
}

// DialFrontgateService connects to an FrontgateService at the specified network address.
func DialFrontgateService(network, addr string) (*FrontgateServiceClient, error) {
	c, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &FrontgateServiceClient{c}, nil
}

// DialFrontgateServiceTimeout connects to an FrontgateService at the specified network address.
func DialFrontgateServiceTimeout(network, addr string, timeout time.Duration) (*FrontgateServiceClient, error) {
	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, err
	}
	return &FrontgateServiceClient{rpc.NewClient(conn)}, nil
}