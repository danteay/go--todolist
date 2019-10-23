package serialization

type JSON map[string]interface{}

func (j JSON) Algo() string {
	return "algo"
}
