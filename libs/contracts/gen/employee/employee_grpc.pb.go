// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: employee/employee.proto

package employeev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Employee_AddEmployee_FullMethodName    = "/employee.Employee/addEmployee"
	Employee_GetEmployee_FullMethodName    = "/employee.Employee/getEmployee"
	Employee_GetEmployees_FullMethodName   = "/employee.Employee/getEmployees"
	Employee_UpdateEmployee_FullMethodName = "/employee.Employee/updateEmployee"
	Employee_DeleteEmployee_FullMethodName = "/employee.Employee/deleteEmployee"
	Employee_InviteEmployee_FullMethodName = "/employee.Employee/inviteEmployee"
)

// EmployeeClient is the client API for Employee service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeClient interface {
	AddEmployee(ctx context.Context, in *AddEmployeeRequest, opts ...grpc.CallOption) (*AddEmployeeResponse, error)
	GetEmployee(ctx context.Context, in *GetEmployeeRequest, opts ...grpc.CallOption) (*GetEmployeeResponse, error)
	GetEmployees(ctx context.Context, in *GetEmployeesRequest, opts ...grpc.CallOption) (*GetEmployeesResponse, error)
	UpdateEmployee(ctx context.Context, in *UpdateEmployeeRequest, opts ...grpc.CallOption) (*UpdateEmployeeResponse, error)
	DeleteEmployee(ctx context.Context, in *DeleteEmployeeRequest, opts ...grpc.CallOption) (*DeleteEmployeeResponse, error)
	InviteEmployee(ctx context.Context, in *InviteEmployeeRequest, opts ...grpc.CallOption) (*InviteEmployeeRequest, error)
}

type employeeClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeClient(cc grpc.ClientConnInterface) EmployeeClient {
	return &employeeClient{cc}
}

func (c *employeeClient) AddEmployee(ctx context.Context, in *AddEmployeeRequest, opts ...grpc.CallOption) (*AddEmployeeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddEmployeeResponse)
	err := c.cc.Invoke(ctx, Employee_AddEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeClient) GetEmployee(ctx context.Context, in *GetEmployeeRequest, opts ...grpc.CallOption) (*GetEmployeeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEmployeeResponse)
	err := c.cc.Invoke(ctx, Employee_GetEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeClient) GetEmployees(ctx context.Context, in *GetEmployeesRequest, opts ...grpc.CallOption) (*GetEmployeesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEmployeesResponse)
	err := c.cc.Invoke(ctx, Employee_GetEmployees_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeClient) UpdateEmployee(ctx context.Context, in *UpdateEmployeeRequest, opts ...grpc.CallOption) (*UpdateEmployeeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateEmployeeResponse)
	err := c.cc.Invoke(ctx, Employee_UpdateEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeClient) DeleteEmployee(ctx context.Context, in *DeleteEmployeeRequest, opts ...grpc.CallOption) (*DeleteEmployeeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteEmployeeResponse)
	err := c.cc.Invoke(ctx, Employee_DeleteEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeClient) InviteEmployee(ctx context.Context, in *InviteEmployeeRequest, opts ...grpc.CallOption) (*InviteEmployeeRequest, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InviteEmployeeRequest)
	err := c.cc.Invoke(ctx, Employee_InviteEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServer is the server API for Employee service.
// All implementations must embed UnimplementedEmployeeServer
// for forward compatibility.
type EmployeeServer interface {
	AddEmployee(context.Context, *AddEmployeeRequest) (*AddEmployeeResponse, error)
	GetEmployee(context.Context, *GetEmployeeRequest) (*GetEmployeeResponse, error)
	GetEmployees(context.Context, *GetEmployeesRequest) (*GetEmployeesResponse, error)
	UpdateEmployee(context.Context, *UpdateEmployeeRequest) (*UpdateEmployeeResponse, error)
	DeleteEmployee(context.Context, *DeleteEmployeeRequest) (*DeleteEmployeeResponse, error)
	InviteEmployee(context.Context, *InviteEmployeeRequest) (*InviteEmployeeRequest, error)
	mustEmbedUnimplementedEmployeeServer()
}

// UnimplementedEmployeeServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEmployeeServer struct{}

func (UnimplementedEmployeeServer) AddEmployee(context.Context, *AddEmployeeRequest) (*AddEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEmployee not implemented")
}
func (UnimplementedEmployeeServer) GetEmployee(context.Context, *GetEmployeeRequest) (*GetEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployee not implemented")
}
func (UnimplementedEmployeeServer) GetEmployees(context.Context, *GetEmployeesRequest) (*GetEmployeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployees not implemented")
}
func (UnimplementedEmployeeServer) UpdateEmployee(context.Context, *UpdateEmployeeRequest) (*UpdateEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployee not implemented")
}
func (UnimplementedEmployeeServer) DeleteEmployee(context.Context, *DeleteEmployeeRequest) (*DeleteEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployee not implemented")
}
func (UnimplementedEmployeeServer) InviteEmployee(context.Context, *InviteEmployeeRequest) (*InviteEmployeeRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteEmployee not implemented")
}
func (UnimplementedEmployeeServer) mustEmbedUnimplementedEmployeeServer() {}
func (UnimplementedEmployeeServer) testEmbeddedByValue()                  {}

// UnsafeEmployeeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServer will
// result in compilation errors.
type UnsafeEmployeeServer interface {
	mustEmbedUnimplementedEmployeeServer()
}

func RegisterEmployeeServer(s grpc.ServiceRegistrar, srv EmployeeServer) {
	// If the following call pancis, it indicates UnimplementedEmployeeServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Employee_ServiceDesc, srv)
}

func _Employee_AddEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServer).AddEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employee_AddEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServer).AddEmployee(ctx, req.(*AddEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employee_GetEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServer).GetEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employee_GetEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServer).GetEmployee(ctx, req.(*GetEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employee_GetEmployees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmployeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServer).GetEmployees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employee_GetEmployees_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServer).GetEmployees(ctx, req.(*GetEmployeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employee_UpdateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServer).UpdateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employee_UpdateEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServer).UpdateEmployee(ctx, req.(*UpdateEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employee_DeleteEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServer).DeleteEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employee_DeleteEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServer).DeleteEmployee(ctx, req.(*DeleteEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employee_InviteEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServer).InviteEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employee_InviteEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServer).InviteEmployee(ctx, req.(*InviteEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Employee_ServiceDesc is the grpc.ServiceDesc for Employee service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Employee_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "employee.Employee",
	HandlerType: (*EmployeeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addEmployee",
			Handler:    _Employee_AddEmployee_Handler,
		},
		{
			MethodName: "getEmployee",
			Handler:    _Employee_GetEmployee_Handler,
		},
		{
			MethodName: "getEmployees",
			Handler:    _Employee_GetEmployees_Handler,
		},
		{
			MethodName: "updateEmployee",
			Handler:    _Employee_UpdateEmployee_Handler,
		},
		{
			MethodName: "deleteEmployee",
			Handler:    _Employee_DeleteEmployee_Handler,
		},
		{
			MethodName: "inviteEmployee",
			Handler:    _Employee_InviteEmployee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "employee/employee.proto",
}
