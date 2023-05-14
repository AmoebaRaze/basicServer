package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Run with
//		go run .
// Send request with:
//

// написать хендлеры для обработки файла.
func handleEcho(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string
	for _, row := range records {
		response = fmt.Sprintf("\n%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Println(response)
	fmt.Fprint(w, response)
}

func transpose(matrix [][]string) [][]string {
	rows := len(matrix)
	cols := len(matrix[0])
	transposed := make([][]string, cols)

	for i := range transposed {
		transposed[i] = make([]string, rows)
		for j := range transposed[i] {
			transposed[i][j] = matrix[j][i]
		}
	}
	return transposed
}

func handleInvert(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	//var response string

	matrix := make([][]string, len(records))
	for i, record := range records {
		matrix[i] = record
	}

	// транспонируем матрицу
	transposed := transpose(matrix)

	for _, row := range transposed {
		fmt.Printf("\n%s,%s,%s", row[0], row[1], row[2])
		fmt.Fprintf(w, "\n%s,%s,%s", row[0], row[1], row[2])
	}
}

func handleFlatten(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	var response string
	for _, row := range records {
		response = fmt.Sprintf("\n%s\n%s\n", response, strings.Join(row, ","))
	}
	flatString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(response)), ","), "]")

	fmt.Println("\n" + flatString)
	fmt.Fprint(w, "\n"+flatString)
}
func handleSum(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	var sum int
	for _, row := range records {
		for _, val := range row {
			count, _ := strconv.Atoi(val)
			sum = sum + count
		}
	}
	fmt.Fprintf(w, "\n%d", sum)
	fmt.Println(sum)

}
func handleMultiply(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	res := 1
	for _, row := range records {
		for _, val := range row {
			count, _ := strconv.Atoi(val)
			res = res * count
		}
	}
	fmt.Fprintf(w, "\n%d", res)
	fmt.Println(res)
}

// curl -F 'file=@matrix.csv' "localhost:8080/echo"
// curl -F 'file=@matrix.csv' "localhost:8080/flatten"
// curl -F 'file=@matrix.csv' "localhost:8080/invert"
// curl -F 'file=@matrix.csv' "localhost:8080/sum"
// curl -F 'file=@matrix.csv' "localhost:8080/multiply"

func main() {
	http.HandleFunc("/echo", handleEcho) // роутер эхо
	http.HandleFunc("/invert", handleInvert)
	http.HandleFunc("/flatten", handleFlatten)
	http.HandleFunc("/sum", handleSum)
	http.HandleFunc("/multiply", handleMultiply)
	http.ListenAndServe(":8080", nil)
}
