package appTypes

import "encoding/json"

type Storage int

const (
	Local Storage = iota
	SMMS
)

func (s Storage) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *Storage) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*s = ToStorage(str)
	return nil
}

func (s Storage) String() string {
	var str string
	switch s {
	case Local:
		str = "local"
	case SMMS:
		str = "SMMS"
	default:
		str = "local"
	}
	return str
}

func ToStorage(str string) Storage {
	switch str {
	case "local":
		return Local
	case "SMMS":
		return SMMS
	default:
		return -1
	}
}
