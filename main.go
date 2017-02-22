package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	RPC_PORT           = 9332
	RPC_USERNAME       = "user"
	RPC_PASSWORD       = "pass"
	RPC_HOST           = "127.0.0.1"
	ACTIVATION_PERIOD  = 750
	ENFORCEMENT_PERIOD = 950
	TARGET_WINDOW      = 1000
)

var (
	BlockIndex map[int]int
)

func BuildURL() string {
	return fmt.Sprintf("http://%s:%s@%s:%d", RPC_USERNAME, RPC_PASSWORD, RPC_HOST, RPC_PORT)
}

func SendHTTPGetRequest(url string, jsonDecode bool, result interface{}) (err error) {
	res, err := http.Get(url)

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		log.Printf("HTTP status code: %d\n", res.StatusCode)
		return errors.New("Status code was not 200.")
	}

	contents, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if jsonDecode {
		err := JSONDecode(contents, &result)

		if err != nil {
			return err
		}
	} else {
		result = &contents
	}

	return nil
}

func JSONDecode(data []byte, to interface{}) error {
	err := json.Unmarshal(data, &to)

	if err != nil {
		return err
	}

	return nil
}

func SendRPCRequest(method, req interface{}) (map[string]interface{}, error) {
	var params []interface{}
	if req != nil {
		params = append(params, req)
	} else {
		params = nil
	}

	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     1,
		"params": params,
	})

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(BuildURL(), "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result["error"] != nil {
		errorMsg := result["error"].(map[string]interface{})
		return nil, fmt.Errorf("Error code: %v, message: %v\n", errorMsg["code"], errorMsg["message"])
	}
	return result, nil
}

func GetBlockVersion(block int) (int, error) {
	result, err := SendRPCRequest("getblockhash", block)
	if err != nil {
		return 0, err
	}

	blockHash := result["result"].(string)

	result, err = SendRPCRequest("getblock", blockHash)
	if err != nil {
		return 0, err
	}
	result = result["result"].(map[string]interface{})
	version := result["version"].(float64)
	return int(version), nil
}

func CheckBlocks(version, height, threshold int) (bool, int) {
	nFound := 0
	blockVer := 0
	var ok bool
	var err error

	for i := height - TARGET_WINDOW; i < height; i++ {
		blockVer, ok = BlockIndex[i]
		if !ok {
			blockVer, err = GetBlockVersion(i)
			if err != nil {
				log.Fatal("Failed to obtain block version.")
			}
			BlockIndex[i] = blockVer
		}
		if version == blockVer {
			nFound++
		}
	}
	if nFound >= threshold {
		return true, nFound
	}
	return false, nFound
}

func main() {
	BlockIndex = make(map[int]int)
	startHeight := 800000
	targetVer := 3
	bActivated := false
	height := startHeight

	for {
		success, found := CheckBlocks(targetVer, height, ACTIVATION_PERIOD)
		percentage := float64(found) / TARGET_WINDOW * 100 / 1
		log.Printf("Height: %d Percentage: %f%%\n", height, percentage)

		if !bActivated && success {
			log.Printf("Block %d reached v%d activation.\n", height, targetVer)
			bActivated = true
		}

		success, found = CheckBlocks(targetVer, height, ENFORCEMENT_PERIOD)
		if success {
			log.Printf("Block %d reached v%d enforcement.\n", height, targetVer)
			break
		}
		height++
	}
}

// 811879 v3
// 918684 v4
