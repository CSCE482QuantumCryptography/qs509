package qs509

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

func Benchmark(startTime, endTime time.Time) {

	executionTime := endTime.Sub(startTime)

	// BENCHMARK

	saveFolderPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
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

func BenchmarkMap(sa string, ka string, timeMap map[string][]time.Time, outFile string) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	index, err := f.NewSheet("Sheet 1")

	if err != nil {
		fmt.Println(err)
		return
	}

	f.SetActiveSheet(index)

	f.SetCellValue("Sheet1", "A1", "Measurement")
	f.SetCellValue("Sheet1", "B1", "Time (s)")

	spot := 2
	for key, value := range timeMap {
		executionTime := value[1].Sub(value[0])

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(spot), key)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(spot), executionTime)

		spot++
	}

	if err := f.SaveAs(outFile); err != nil {
		fmt.Println(err)
	}

}
