package apiserver

type ApiServer struct {
}

func New() *ApiServer {
	return &ApiServer{}
}
func (s *ApiServer) Start() error {
	return nil
}
