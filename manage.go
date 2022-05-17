package nsink

type Client struct {
	Total uint
	Trees map[string]uint
}

func NewClient() Client {
	return Client{
		Total: 0,
		Trees: make(map[string]uint),
	}
}

func (s Server) SetClient(ip string, ids []string) {
	client := NewClient()
	s.Clients[ip] = client
	for _, id := range ids {
		client.Trees[id] = 0
	}
}

func (s Server) GetClientStats(ip string) *Client {
	if client, ok := s.Clients[ip]; ok {
		return &client
	}
	return nil
}

func (s Server) DeleteClient(ip string) {
	delete(s.Clients, ip)
}

func (s Server) SetTreeAddress(id string, name string, ip string) {
	tree, ok := s.Trees[id]
	if !ok {
		tree = NullTree()
		s.Trees[id] = tree
	}
	tree.Insert(name, ip)
}

func (s Server) SetTreeBlock(id string, name string) {
	s.SetTreeAddress(id, name, "0.0.0.0")
}

func (s Server) DeleteTree(id string) {
	delete(s.Trees, id)
}

func (s Server) DeleteTreeAddress(id string, name string) {
	if tree, ok := s.Trees[id]; ok {
		tree.Delete(name)
	}
}
