package functions

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
)

func GetResponse(responseCode string, responseMessage string) interface{} {
	type Response struct {
		ResponseCode    string `json:"response_code"`
		ResponseMessage string `json:"response_message"`
	}

	result := Response{
		responseCode,
		responseMessage,
	}
	AddLogResponse(result)
	return result
}

func MysqlConnect(host, port, uname, pass, dbname string) error {
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		uname,
		pass,
		host,
		port,
		dbname,
	)

	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

//noinspection GoPrintFunctions
func AddLogResponse(result interface{}) {
	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("-------------------------RESPONSE---------------------")
	fmt.Println(string(b))
	fmt.Println("----------------------------END------------------------\n")
}

func PrettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

/*
func GetResponseWithData(responseCode string, responseMessage string, responseData interface{}) interface{} {
	type Response struct {
		ResponseCode    string      `json:"response_code"`
		ResponseMessage string      `json:"response_message"`
		ResponseData    interface{} `json:"response_data"`
	}

	result := Response{
		responseCode,
		responseMessage,
		responseData,
	}
	AddLogResponse(result)
	return result
}
*/
