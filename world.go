package main

import (
	"github.com/lonnng/nano/component"
	"github.com/lonnng/nano"
	"github.com/google/uuid"
	"github.com/lonnng/nano/session"
	"./protocol"
)

type World struct {
	component.Base
	*nano.Group
}

func NewWorld() *World{
	return &World{
		Group:nano.NewGroup(uuid.New().String()),
	}
}

func (w *World) Init(){
	nano.OnSessionClosed(func(session *session.Session) {
		w.Leave(session)
		w.Broadcast("leave",&protocol.LeaveWorldResponse{ID:session.ID()})
	})
}

func (w *World) Enter(s *session.Session,msg []byte) error{
	w.Group.Add(s)
	return s.Response(&protocol.EnterWorldResponse{ID:s.ID()})
}

func(w *World) Update(s *session.Session,msg []byte) error{
	return w.Broadcast("update",msg)
}