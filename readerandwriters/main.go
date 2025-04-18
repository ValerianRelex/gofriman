package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func processData(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			writer.Write(b[0:count])
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func processDataCopy(reader io.Reader, writer io.Writer) {
	written, err := io.Copy(writer, reader)
	if err != nil {
		Printfln("Ошибка во время копирования данных, %v", err)
	} else {
		Printfln("Записано %v байт данных", written)
	}
}

func scanSingle(reader io.Reader, val interface{}) (int, error) {
	return fmt.Fscan(reader, val)
}

func main() {
	r := strings.NewReader("Kayak")
	var builder strings.Builder
	processData(r, &builder)
	Printfln("String builder contents: %s", builder.String())

	//
	fmt.Printf("\n")
	r2 := strings.NewReader("Карл у клары украл кораллы")
	builder.Reset()
	processDataCopy(r2, &builder)
	Printfln("String builder contents: %s", builder.String())

	//
	fmt.Printf("\n")
	pipeReader, pipeWriter := io.Pipe()
	go func() {
		GenerateData(pipeWriter)
		pipeWriter.Close()
	}()
	ConsumeData(pipeReader)

	//
	fmt.Printf("\n")
	reader3 := strings.NewReader("Kayak Watersports $279.00")

	for {
		var str string
		_, err := scanSingle(reader3, &str)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		Printfln("Value: %v", str)
	}

	// JSON
	fmt.Printf("\n")

	var b bool = true
	var str string = "Hello"
	var fval float64 = 99.99
	var ival int = 200
	var pointer *int = &ival

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	for _, val := range []interface{}{b, str, fval, ival, pointer} {
		encoder.Encode(val)
	}

	fmt.Print(writer.String()) // энкодер сам добавляет символ перевода на новую строку

	//
	fmt.Printf("\n")
	writer.Reset() // сброс, для нового примера

	names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	numbers := [3]int{10, 20, 30}
	var byteArray [5]byte

	copy(byteArray[0:], []byte(names[0]))
	byteSlice := []byte(names[0])

	encoder.Encode(names)
	encoder.Encode(numbers)
	encoder.Encode(byteArray)
	encoder.Encode(byteSlice)

	fmt.Print(writer.String())

	//
	fmt.Printf("\n")
	writer.Reset() // сброс, для нового примера

	m := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 49.95,
	}

	encoder.Encode(m)
	fmt.Print(writer.String())

	//  Кодирование структуры
	fmt.Printf("\n")
	writer.Reset() // сброс, для нового примера

	encoder.Encode(Kayak)
	fmt.Print(writer.String())

	// Кодирование структуры со встроенным полем
	fmt.Printf("\n")
	writer.Reset() // сброс, для нового примера

	dp := DiscountedProduct{
		Product:  &Kayak,
		Discount: 10.50,
	}
	encoder.Encode(&dp)
	fmt.Print(writer.String())

	// Кодирование среза интерфейса
	fmt.Printf("\n Кодирование среза интерфейса\n")
	writer.Reset() // сброс, для нового примера

	namedItems := []Named{&dp, &Person{PersonName: "Alice"}}
	encoder.Encode(namedItems)

	fmt.Print(writer.String())

	// Декодирование основных типов данных
	fmt.Printf("\n Декодирование основных типов данных\n")
	writer.Reset() // сброс, для нового примера

	reader := strings.NewReader(`true "Hello" 99.99 200`) // int и float64 будут декодированы в float64
    vals := []interface{} { }
    decoder := json.NewDecoder(reader)
    for {
        var decodedVal interface{}
        err := decoder.Decode(&decodedVal)
        if (err != nil) {
            if (err != io.EOF) {
                Printfln("Error: %v", err.Error())
            }
            break
        }
        vals = append(vals, decodedVal)
    }
    for _, val := range vals {
        Printfln("Decoded (%T): %v", val, val)
    }

	// Указание типов для декодирования
	fmt.Printf("\n Указание типов для декодирования\n")
	writer.Reset() // сброс, для нового примера

	reader22 := strings.NewReader(`true "Hello" 99.99 200`)
    var bval bool
    var sval string
    var fpval float64
    var ival22 int
    vals22 := []interface{} { &bval, &sval, &fpval, &ival } // порядок имеет значение!!!
    decoder22 := json.NewDecoder(reader22)
    for i := 0; i < len(vals22); i++ {
        err := decoder22.Decode(vals22[i])
        if err != nil {
            Printfln("Error: %v", err.Error())
            break
        }
    }
    Printfln("Decoded (%T): %v", bval, bval)
    Printfln("Decoded (%T): %v", sval, sval)
    Printfln("Decoded (%T): %v", fpval, fpval)
    Printfln("Decoded (%T): %v", ival22, ival22)

	//
	fmt.Printf("\n Декодирование массива\n")
	writer.Reset() // сброс, для нового примера

	readerArray := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)
    valsA := []interface{} { }
    decoderA := json.NewDecoder(readerArray)
    for {
        var decodedVal interface{}
        err := decoderA.Decode(&decodedVal)
        if (err != nil) {
            if (err != io.EOF) {
                Printfln("Error: %v", err.Error())
            }
            break
        }
        valsA = append(valsA, decodedVal)
    }
    for _, val := range valsA {
        Printfln("Decoded (%T): %v", val, val)
    }

	//
	fmt.Printf("\n Декодирование карт\n")
	writer.Reset() // сброс, для нового примера

	readerM := strings.NewReader(`{"Kayak" : 279, "Lifejacket" : 49.95}`)
    mSI := map[string]interface{} {}
    decoderM := json.NewDecoder(readerM)
    err := decoderM.Decode(&mSI)

    if err != nil {
        Printfln("Error: %v", err.Error())
    } else {
        Printfln("Map: %T, %v", m, m)
        for k, v := range mSI {
            Printfln("Key: %v, Value: %v", k, v)
        }
    }
}
