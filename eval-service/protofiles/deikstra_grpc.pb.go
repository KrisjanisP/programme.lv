// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: deikstra.proto

package protofiles

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SchedulerClient is the client API for Scheduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchedulerClient interface {
	GetJobs(ctx context.Context, in *RegisterWorker, opts ...grpc.CallOption) (Scheduler_GetJobsClient, error)
	// rpc ReportWorkerStatus(stream ReportWorkerStatusRequest) returns (google.protobuf.Empty) {}
	ReportJobStatus(ctx context.Context, opts ...grpc.CallOption) (Scheduler_ReportJobStatusClient, error)
}

type schedulerClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerClient(cc grpc.ClientConnInterface) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) GetJobs(ctx context.Context, in *RegisterWorker, opts ...grpc.CallOption) (Scheduler_GetJobsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Scheduler_ServiceDesc.Streams[0], "/protofiles.Scheduler/GetJobs", opts...)
	if err != nil {
		return nil, err
	}
	x := &schedulerGetJobsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Scheduler_GetJobsClient interface {
	Recv() (*Job, error)
	grpc.ClientStream
}

type schedulerGetJobsClient struct {
	grpc.ClientStream
}

func (x *schedulerGetJobsClient) Recv() (*Job, error) {
	m := new(Job)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *schedulerClient) ReportJobStatus(ctx context.Context, opts ...grpc.CallOption) (Scheduler_ReportJobStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &Scheduler_ServiceDesc.Streams[1], "/protofiles.Scheduler/ReportJobStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &schedulerReportJobStatusClient{stream}
	return x, nil
}

type Scheduler_ReportJobStatusClient interface {
	Send(*JobStatusUpdate) error
	CloseAndRecv() (*JobAction, error)
	grpc.ClientStream
}

type schedulerReportJobStatusClient struct {
	grpc.ClientStream
}

func (x *schedulerReportJobStatusClient) Send(m *JobStatusUpdate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *schedulerReportJobStatusClient) CloseAndRecv() (*JobAction, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(JobAction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SchedulerServer is the server API for Scheduler service.
// All implementations must embed UnimplementedSchedulerServer
// for forward compatibility
type SchedulerServer interface {
	GetJobs(*RegisterWorker, Scheduler_GetJobsServer) error
	// rpc ReportWorkerStatus(stream ReportWorkerStatusRequest) returns (google.protobuf.Empty) {}
	ReportJobStatus(Scheduler_ReportJobStatusServer) error
	mustEmbedUnimplementedSchedulerServer()
}

// UnimplementedSchedulerServer must be embedded to have forward compatible implementations.
type UnimplementedSchedulerServer struct {
}

func (UnimplementedSchedulerServer) GetJobs(*RegisterWorker, Scheduler_GetJobsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetJobs not implemented")
}
func (UnimplementedSchedulerServer) ReportJobStatus(Scheduler_ReportJobStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method ReportJobStatus not implemented")
}
func (UnimplementedSchedulerServer) mustEmbedUnimplementedSchedulerServer() {}

// UnsafeSchedulerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchedulerServer will
// result in compilation errors.
type UnsafeSchedulerServer interface {
	mustEmbedUnimplementedSchedulerServer()
}

func RegisterSchedulerServer(s grpc.ServiceRegistrar, srv SchedulerServer) {
	s.RegisterService(&Scheduler_ServiceDesc, srv)
}

func _Scheduler_GetJobs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RegisterWorker)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SchedulerServer).GetJobs(m, &schedulerGetJobsServer{stream})
}

type Scheduler_GetJobsServer interface {
	Send(*Job) error
	grpc.ServerStream
}

type schedulerGetJobsServer struct {
	grpc.ServerStream
}

func (x *schedulerGetJobsServer) Send(m *Job) error {
	return x.ServerStream.SendMsg(m)
}

func _Scheduler_ReportJobStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SchedulerServer).ReportJobStatus(&schedulerReportJobStatusServer{stream})
}

type Scheduler_ReportJobStatusServer interface {
	SendAndClose(*JobAction) error
	Recv() (*JobStatusUpdate, error)
	grpc.ServerStream
}

type schedulerReportJobStatusServer struct {
	grpc.ServerStream
}

func (x *schedulerReportJobStatusServer) SendAndClose(m *JobAction) error {
	return x.ServerStream.SendMsg(m)
}

func (x *schedulerReportJobStatusServer) Recv() (*JobStatusUpdate, error) {
	m := new(JobStatusUpdate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Scheduler_ServiceDesc is the grpc.ServiceDesc for Scheduler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Scheduler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetJobs",
			Handler:       _Scheduler_GetJobs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReportJobStatus",
			Handler:       _Scheduler_ReportJobStatus_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "deikstra.proto",
}
