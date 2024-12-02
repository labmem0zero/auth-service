package services

type Service interface {
	Start(string)
	Stop(string)
}

type Services struct {
	services []Service
}

func StartServices(reqID string, services ...Service) Services {
	srv := Services{}
	for _, s := range services {
		s.Start(reqID)
		srv.services = append(srv.services, s)
	}
	return srv
}

func (srv Services) StopServices(reqID string) {
	for _, s := range srv.services {
		s.Stop(reqID)
	}
}
