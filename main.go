package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var result map[string]interface{}

func main() {
	result = map[string]interface{}{}

	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(stdin), &data); err != nil {
		log.Fatal(err)
	}

	err = flatten(data, "")
	if err != nil {
		log.Fatal(err)
	}
	resultByte, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(resultByte))
}

func flatten(data map[string]interface{}, key string) error {
	for k, v := range data {
		if key != "" {
			k = fmt.Sprintf("%s_%s", key, k)
		}
		switch vv := v.(type) {
		case []interface{}:
			result[k] = join(vv, ",")
		case map[string]interface{}:
			err := flatten(vv, k)
			if err != nil {
				return err
			}
		default:
			result[k] = vv
		}
	}
	return nil
}

func join(a []interface{}, sep string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(a); i++ {
		buffer.WriteString(fmt.Sprintf("%s", a[i]))
		if i != len(a)-1 {
			buffer.WriteString(sep)
		}
	}

	return buffer.String()
}
