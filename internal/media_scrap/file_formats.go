package mediascrap

import "strings"

type FileFormats struct {
	List   []string
	Lookup map[string]bool
}

func (s *FileFormats) String() string {
	return strings.Join(s.List, ",")
}

func (s *FileFormats) Set(value string) error {
	if s.Lookup == nil {
		s.Lookup = make(map[string]bool)
	}
	parts := strings.Split(value, ",")
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			s.List = append(s.List, trimmed)
			s.Lookup[trimmed] = true
		}
	}
	return nil
}
