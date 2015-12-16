package server

func (s *Server) Ping() ([]string, error) {
	return []string{"Pong!"}, nil
}
