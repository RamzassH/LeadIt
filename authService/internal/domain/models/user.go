package models

import "time"

type User struct {
	ID         int64
	Name       string
	Surname    string
	MiddleName string
	AboutMe    string
	Messengers map[string]string
	Email      string
	BirthDate  time.Time
	PassHash   []byte
}

type UpdateUserPayload struct {
	UserId     int64             `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name       string            `json:"name,omitempty" validate:"omitempty,min=1,max=100"`
	Surname    string            `json:"surname,omitempty" validate:"omitempty,min=1,max=100"`
	MiddleName string            `json:"middle_name,omitempty" validate:"omitempty,min=1,max=100"`
	AboutMe    string            `json:"about_me,omitempty" validate:"omitempty,min=1,max=100"`
	Messengers map[string]string `json:"messengers,omitempty" validate:"omitempty"`
	Email      string            `json:"email,omitempty" validate:"omitempty,email"`
	Password   string            `json:"password,omitempty" validate:"omitempty"`
	BirthDate  time.Time         `json:"birth,omitempty" validate:"omitempty"`
}

type RegisterUserPayload struct {
	Name     string `validate:"required,min=1,max=100"`
	Surname  string `validate:"required,min=1,max=100"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

type LoginUserPayload struct {
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

type VerifyUserPayload struct {
	Code string `json:"code,omitempty" validate:"required"`
}

type IsAdminPayload struct {
	UserId int64 `validate:"required,min=1,max=100"`
}
