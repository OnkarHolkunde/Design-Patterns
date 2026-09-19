type Server struct {
	Host    string
	Port    int
	Timeout int
}

type Option func(*Server)

func WithTimeout(timeout int) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func NewServer(host string, port int, opts ...Option) *Server {
	s := &Server{Host: host, Port: port, Timeout: 10} // Defaults
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Usage:
// s := NewServer("localhost", 8080, WithTimeout(30))