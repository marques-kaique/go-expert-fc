package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{} //or any
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error //registra um handler para o evento
	Dispatch(event EventInterface) error                            //faz com que o evento seja tratado pelos handlers registrados
	Remove(eventName string, handler EventHandlerInterface) error   //remove do dispatcher o handler registrado para o evento
	Has(eventName string, handler EventHandlerInterface) bool       //verifica se o handler está registrado para o evento
	Clear() error                                                   //remove todos os handlers registrados, limpando o dispatcher
}
