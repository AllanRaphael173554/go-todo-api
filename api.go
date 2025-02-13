package main

import "net/http"

type APIserver struct {
	listenAddress string
}

func NewApiServer(listenAdd string) *APIserver {
	return &APIserver{
		listenAddress: listenAdd,
	}
}

func (s *APIserver) Run() {

}

func (s *APIserver) handleToDoList(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIserver) handleGetToDoList(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIserver) handleCreateToDoList(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIserver) handleDeleteToDoList(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// redo in future
func (s *APIserver) handleTransferToDoList(w http.ResponseWriter, r *http.Request) error {
	return nil
}
