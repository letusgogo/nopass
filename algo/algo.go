package algo

type Algorithm interface {
	Generate(password string, salt string) string
}

type SHA256Algo struct {
}

func (s *SHA256Algo) Generate(password string, salt string) string {
	// ...
	return ""
}

type MD5Algo struct {
}

func (m *MD5Algo) Generate(password string, salt string) string {
	return ""
}
