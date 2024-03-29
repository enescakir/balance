package internal

// routes adds handlers for endpoints
func (s *Server) routes() {
	s.router.HandleFunc("/", s.handleDashboard())
	s.router.HandleFunc("/isbalanced", s.log(s.handleCheck()))
	s.router.HandleFunc("/logs", s.handleLogIndex())
	s.router.HandleFunc("/logs/status", s.handleLogStatusCounts())
	s.router.HandleFunc("/logs/histogram", s.handleLogResponseHistogram())
}
