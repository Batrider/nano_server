package main

import (
	"github.com/lonnng/nano/component"
	"github.com/lonnng/nano/session"
	"github.com/lonnng/nano/examples/demo/tadpole/logic/protocol"
)

type Manager struct {
	component.Base
}

func NewManager() *Manager{
	return &Manager{}
}

func (m *Manager) Login(s *session.Session,msg *protocol.JoyLoginRequest) error {

	id := s.ID()
	s.Bind(id)

	return s.Response(protocol.LoginResponse{
		Status: protocol.LoginStatusSucc,
		ID:     id,
	})
}