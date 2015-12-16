package server

func (s *Server) Ping() (string, error) {
	return "Pong!", nil
}
