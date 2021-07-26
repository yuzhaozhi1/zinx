package znet

//Server 接口实现，定义一个Server服务类
type Server struct {
	Name      string // 服务器的名称
	IpVersion string // tcp4 or other
	Host      string // host
	Port      int    // port

}

//============== 实现 ziface.IServer 里的全部接口方法 ========

// Start 开启服务器的方法
func (s *Server) Start() {

}

// Stop 停止服务器的方法
func (s *Server) Stop() {

}

// Serve 运行服务
func (s *Server) Serve() {

}

// NewServer 实例化一个server
func NewServer(name string) *Server {
	return &Server{
		Name:      name,
		IpVersion: "ipv4",
		Port:      8080,
		Host:      "0.0.0.0",
	}
}
