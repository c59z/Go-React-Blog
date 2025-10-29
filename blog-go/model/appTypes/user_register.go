package appTypes

import "encoding/json"

type Register int

const (
	Email  Register = iota // Register via email
	Github                 // Register via GitHub
)

// MarshalJSON implements json.Marshaler
func (r Register) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// UnmarshalJSON implements json.Unmarshaler
func (r *Register) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*r = ToRegister(str)
	return nil
}

// String returns string representation of Register
func (r Register) String() string {
	var str string
	switch r {
	case Email:
		str = "Email"
	case Github:
		str = "Github"
	default:
		str = "未知"
	}
	return str
}

// ToRegister converts string to Register
func ToRegister(str string) Register {
	switch str {
	case "Email":
		return Email
	case "Github":
		return Github
	default:
		return -1
	}
}
