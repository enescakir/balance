package main

func (s *Server) routes() {

	s.router.HandleFunc("/", s.handleDashboard())
	s.router.HandleFunc("/isbalanced", s.handleCheck())
	s.router.HandleFunc("/logs", s.handleAllLogs())
	//s.router.HandleFunc("/admin", s.adminOnly(s.handleAdminIndex()))

}