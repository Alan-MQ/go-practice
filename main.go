package main

type RPCdata struct {
	Name string        // name of the function
	Args []interface{} // request's or response's body expect error.
	Err  string        // Error any executing remote server
}

func Encode(data RPCdata) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {
	print()
}
