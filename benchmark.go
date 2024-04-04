package qs509

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"strings"

	"github.com/xuri/excelize/v2"
)

func Benchmark(startTime, endTime time.Time) {
	if startTime.After(endTime){
		return
	}


	executionTime := endTime.Sub(startTime)

	// BENCHMARK

	saveFolderPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else if strings.Contains(saveFolderPath, "/testing") {
		saveFolderPath = strings.Replace(saveFolderPath, "/testing", "", -1)
	}

	// parentDir := filepath.Dir(saveFolderPath)
	// saveParentFolderPath := parentDir

	file, err := excelize.OpenFile(saveFolderPath + "/benchmarkLog/benchmarkTime.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	// new data to add
	dataExcel := [][]interface{}{
		{startTime, "test1", executionTime},
		{startTime, "test2", executionTime},
	}
	for i, row := range rows {
		dataRow := i + 1
		for j, col := range row {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
		}
	}

	for i, row := range dataExcel {
		dataRow := i + len(rows) + 1
		for j, col := range row {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
		}
	}

	if err := file.Save(); err != nil {
		log.Fatal(err)
	}

	// NEW FILE CREATION

	// benchmark logs one instance
	fileNew := excelize.NewFile()

	headers := []string{"Access Time", "Algorithm", "Runtime"}
	for i, header := range headers {
		fileNew.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}

	for i, row := range dataExcel {
		dataRow := i + 2
		for j, col := range row {
			fileNew.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
		}
	}

	if err := fileNew.SaveAs(saveFolderPath + "/benchmarkLog/benchmarkInstance.xlsx"); err != nil {
		log.Fatal(err)
	}
}

func CreateFile(fileName string) {
	f := excelize.NewFile()
	f.NewSheet("client")
	f.NewSheet("server")
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}

}

func BenchmarkMap(timeMap map[string][]time.Time, sa string, ka string, outFile string, sheet string) {
	f, err := excelize.OpenFile(outFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		fmt.Println(err)
		return
	}

	f.SetCellValue(sheet, "A1", "SA: "+sa)
	f.SetCellValue(sheet, "B1", "KA: "+ka)

	f.SetCellValue(sheet, "A2", "Measurement")
	f.SetCellValue(sheet, "B2", "Time (us)")

	spot := 3
	for key, value := range timeMap {
		executionTime := value[1].Sub(value[0])

		f.SetCellValue(sheet, "A"+strconv.Itoa(spot), key)
		f.SetCellValue(sheet, "B"+strconv.Itoa(spot), executionTime.Microseconds())

		spot++
	}

	if err := f.SaveAs(outFile); err != nil {
		fmt.Println(err)
	}

}
