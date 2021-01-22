package higor

import (
	"math"
	"reflect"
	"testing"
)

func TestPrintHelloHigor(t *testing.T) {
	value := PrintHelloHigor()
	if value != "Hello from higor" {
		t.Errorf("HellowHigor failed")
	}
}

func TestHead(t *testing.T) {
	valuesExpected := Book{}
	valuesExpected["id"] = Page{1, 2, 3, 4, 5}
	valuesExpected["name"] = Page{"Hamish", "Anson", "Willie", "Eimile", "Rawley"}
	valuesExpected["work_remotely"] = Page{false, math.NaN(), true, true, true}
	valuesExpected["salary"] = Page{4528.90, 1418.86, 1311.34, 3895.20, 2350.92}
	valuesExpected["age"] = Page{96, math.NaN(), math.NaN(), 80, math.NaN()}
	valuesExpected["country_code"] = Page{"PE", math.NaN(), "PH", "ID", "ZA"}

	dfHigor := NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()
	dfHigorHead := dfHigor.Head()
	valuesResult := dfHigorHead.Values

	for k, v := range valuesResult {
		for i, element := range v {
			switch element.(type) {
			case float64:

				if math.IsNaN(element.(float64)) {
					if !math.IsNaN(valuesExpected[k][i].(float64)) {
						t.Errorf("Column: \"%s\", expected: %v recived: %v", k, valuesExpected[k], v)
					}
				} else if element != valuesExpected[k][i] {

					t.Errorf("Column: \"%s\", expected: %v recived: %v", k, valuesExpected[k], v)
				}
			default:

				if element != valuesExpected[k][i] {

					t.Errorf("Column: \"%s\", expected: %v recived: %v", k, valuesExpected[k], v)
				}
			}

		}
	}
}

func TestTail(t *testing.T) {

	// Values expected
	valuesExpected := Book{}
	valuesExpected["id"] = Page{96, 97, 98, 99, 100}
	valuesExpected["name"] = Page{math.NaN(), "Novelia", "Maegan", "Andreana", "Freeman"}
	valuesExpected["work_remotely"] = Page{false, true, false, true, false}
	valuesExpected["salary"] = Page{math.NaN(), 3948.23, 2905.48, 3732.29, 2850.99}
	valuesExpected["age"] = Page{54, math.NaN(), 48, 73, 39}
	valuesExpected["country_code"] = Page{"GF", "JP", "UA", "CN", "TH"}

	// Index expected
	indexExpected := []int{95, 96, 97, 98, 99}

	// Get Result
	dfHigor := NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()
	dfHigorTail := dfHigor.Tail()
	valuesResult := dfHigorTail.Values
	indexResult := dfHigorTail.Index

	// Values test
	for k, v := range valuesResult {
		for i, element := range v {
			switch element.(type) {
			case float64:

				if math.IsNaN(element.(float64)) {
					if !math.IsNaN(valuesExpected[k][i].(float64)) {
						t.Errorf("Column: \"%s\", expected: %v recived: %v", k, valuesExpected[k], v)
					}
				} else if element != valuesExpected[k][i] {

					t.Errorf("Column: \"%s\", expected: %v recived: %v", k, valuesExpected[k], v)
				}
			default:

				if element != valuesExpected[k][i] {

					t.Errorf("Column: \"%s\", expected: %v recived: %v", k, valuesExpected[k], v)
				}
			}

		}
	}

	// Index test
	if !reflect.DeepEqual(indexExpected, indexResult) {
		t.Errorf("Index error, expected: %v recived: %v", indexExpected, indexResult)
	}

}

func TestDrop(t *testing.T) {

	// Result expected
	expectedColumns := []string{"id", "work_remotely", "salary", "country_code"}

	// Get result
	dfHigor := NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()
	dfHigor.Drop("name", "age")

	if !reflect.DeepEqual(expectedColumns, dfHigor.Columns) {
		t.Errorf("Columns error, expected: %v recived: %v", expectedColumns, dfHigor.Columns)
	}

}

func TestMean(t *testing.T) {
	dfHigor := NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()

	// Test float64 number
	valueResult := dfHigor.Values["salary"].Mean()
	valueExpected := 2963.707882352941

	if valueResult != valueExpected {
		t.Errorf("Median error, expected: %v, result: %v", valueExpected, valueResult)
	}

	// Test int
	valueIntResult := math.Round(dfHigor.Values["id"].Mean()*100) / 100
	valueIntExpected := 49.54

	if valueIntResult != valueIntExpected {
		t.Errorf("Mean error, expected: %v, result: %v", valueIntExpected, valueIntResult)
	}

	// Test string
	valueStringResult := dfHigor.Values["name"].Mean()
	if !math.IsNaN(valueStringResult) {
		t.Errorf("Mean error, String not calculate corretly")
	}

}

func TestMax(t *testing.T) {

	// Load data
	dfHigor := NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()

	// Int test
	valueIntExpected := 100.0
	valueIntResult := dfHigor.Values["id"].Max()

	if valueIntExpected != valueIntResult {
		t.Errorf("Max value error, expected: %v, result: %v", valueIntExpected, valueIntResult)
	}

	// Float test
	valueFloatExpected := 4971.74
	valueFloatResult := dfHigor.Values["salary"].Max()

	if valueFloatExpected != valueFloatResult {
		t.Errorf("Max value error, expected: %v, result: %v", valueFloatExpected, valueFloatResult)
	}

	// Test string
	valueStringResult := dfHigor.Values["name"].Max()
	if !math.IsNaN(valueStringResult) {
		t.Errorf("Mean error, String not calculate corretly")
	}
}

func TestMin(t *testing.T) {
	// Load data
	dfHigor := NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()

	// Float test
	valueFloatExpected := 217.69
	valueFloatResult := dfHigor.Values["salary"].Min()
	if valueFloatExpected != valueFloatResult {
		t.Errorf("Min value error, expected: %v, result: %v", valueFloatExpected, valueFloatResult)
	}

	// Int test
	valueIntExpected := 1.0
	valueIntResult := dfHigor.Values["id"].Min()

	if valueIntExpected != valueIntResult {
		t.Errorf("Min value error, expected: %v, result: %v", valueIntExpected, valueIntResult)
	}

	// Test string
	valueStringResult := dfHigor.Values["name"].Min()
	if !math.IsNaN(valueStringResult) {
		t.Errorf("Min error, String not calculate corretly")
	}

}

func TestDescribe(t *testing.T) {
	// Get data
	dfHigor := NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()
	dfHigorDescribe := dfHigor.Describe()

	// Order
	// [Mean, Max, Min]
	columns := []string{"id", "name", "work_remotely", "salary", "age", "country_code"}
	book := Book{}
	book["id"] = Page{49.53932584269663, 100, 1}
	book["name"] = Page{math.NaN(), math.NaN(), math.NaN()}
	book["work_remotely"] = Page{math.NaN(), math.NaN(), math.NaN()}
	book["salary"] = Page{2963.707882352941, 4971.74, 217.69}
	book["age"] = Page{48.916666666666664, 100, 2}
	book["country_code"] = Page{math.NaN(), math.NaN(), math.NaN()}

	df := DataFrame{
		Columns: columns,
		Values:  book,
		Index:   []int{0, 1, 2},
	}

	// Comparative Index
	if !reflect.DeepEqual(dfHigorDescribe.Index, df.Index) {
		t.Errorf("Index error=> Expected: %v, Result: %v", df.Index, dfHigorDescribe.Index)
	}

	// Comparative Values
	for columnName, columnValues := range df.Values {
		for index, element := range columnValues {
			switch element.(type) {
			case float64:
				if math.IsNaN(element.(float64)) {
					if !math.IsNaN(dfHigorDescribe.Values[columnName][index].(float64)) {
						t.Errorf("Column: \"%s\", expected: %v recived: %v", columnName, element, dfHigorDescribe.Values[columnName][index])
					}
				} else if element != dfHigorDescribe.Values[columnName][index].(float64) {

					t.Errorf("Column: \"%s\", expected: %v recived: %v", columnName, element, dfHigorDescribe.Values[columnName][index])
				}
			case int:
				if float64(element.(int)) != dfHigorDescribe.Values[columnName][index] {

					t.Errorf("Column: \"%s\", expected: %v (%T) recived: %v (%T)", columnName, element, element, dfHigorDescribe.Values[columnName][index], dfHigorDescribe.Values[columnName][index])
				}
			case string:

				if element != dfHigorDescribe.Values[columnName][index].(string) {

					t.Errorf("Column: \"%s\", expected: %v recived: %v", columnName, element, dfHigorDescribe.Values[columnName][index])
				}
			case bool:
				if element != dfHigorDescribe.Values[columnName][index].(bool) {

					t.Errorf("Column: \"%s\", expected: %v recived: %v", columnName, element, dfHigorDescribe.Values[columnName][index])
				}
			}
		}
	}

}
