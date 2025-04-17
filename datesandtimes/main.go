package main

import (
	"fmt"
	"time"
)

func PrintTime(label string, t *time.Time) {
	// Printfln("%s: Day: %v: Month: %v Year: %v", label, t.Day(), t.Month(), t.Year())
	fmt.Println(label, t.Format(time.RFC822Z))
}
func main() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)
	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("UNIX", &unix)

	for i := 1; i < 10; i++ {
		Printfln("Время в наносекундах %v", time.Nanosecond)
	}

	// анализ строки даты в файле
	layout := "2006-Jan-02"
	dates := []string{
		"1995-Jun-09",
		"2015-Jun-02",
	}
	for _, d := range dates {
		time, err := time.Parse(layout, d)
		if err == nil {
			PrintTime("Parsed", &time)
		} else {
			Printfln("Error: %s", err.Error())
		}
	}

	Printfln(time.Now().Local().GoString())

	//  Создание и проверка продолжительности
	Printfln("\n Создание и проверка продолжительности")
	var d time.Duration = time.Hour + (30 * time.Minute)
	Printfln("Hours: %v", d.Hours())
	Printfln("Mins: %v", d.Minutes())
	Printfln("Seconds: %v", d.Seconds())
	Printfln("Millseconds: %v", d.Milliseconds())
	rounded := d.Round(time.Hour)
	Printfln("Rounded Hours: %v", rounded.Hours())
	Printfln("Rounded Mins: %v", rounded.Minutes())
	trunc := d.Truncate(time.Hour)
	Printfln("Truncated  Hours: %v", trunc.Hours())
	Printfln("Rounded Mins: %v", trunc.Minutes())


	// Создание продолжительности относительно времени
	toYears := func(d time.Duration) int {
		return int(d.Hours() / (24 * 365))
	}
	future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
	past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)
	Printfln("Future: %v", toYears(time.Until(future)))
	Printfln("Past: %v", toYears(time.Since(past)))


	// функции времени для работы с каналами и горутинами
	fmt.Println("")
	ch := make( chan string)

	// time.AfterFunc(time.Second * 1, func () {
    //     writeToChannel(ch)
    // })


	go writeToChannel(ch)

	channelOpen := true
    for channelOpen {
        Printfln("Starting channel read")
        select {
            case name, ok := <- ch:
                if (!ok) {
                    channelOpen = false
                    break
                } else {
                    Printfln("Read name: %v", name)
                }
            case <- time.After(time.Second * 2):
                Printfln("Timeout")
        }
    }

	// здесь бесконечный вывод каждую секунду.
	nameChannel := make (chan string)
	go writeToChannel2(nameChannel)
    for name := range nameChannel {
        Printfln("Read name: %v", name)
    }

	// for data := range ch {
	// 	Printfln("Read -> %v, from channel", data)
	// }

	// time.Sleep(time.Second * 3)
}

func writeToChannel(ch chan <- string) {
	Printfln("Waiting for initial duration...")
    _ = <- time.After(time.Second * 2)
    Printfln("Initial duration elapsed.")

	names := []string {"Федор", "Никита", "Клеменс"}
	for _, n := range names {
		ch <- n
		time.Sleep(time.Second * 3)
	}
	close(ch)
}

func writeToChannel2(nameChannel chan <- string) {
    names := []string { "Alice", "Bob", "Charlie", "Dora" }
    tickChannel := time.Tick(time.Second)
    index := 0
    for {
        <- tickChannel // неважно что передает канало, главное его периодичность, в данном случае - паузка в одну секунду
        nameChannel <- names[index]
        index++
        if (index == len(names)) {
            index = 0
        }
    }
 }
