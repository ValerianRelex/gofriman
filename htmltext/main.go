package main

import (
	"html/template"
	"os"
)

func main() {
	t1, err1 := template.ParseFiles("templates/template.html")
	t2, err2 := template.ParseFiles("templates/extras.html")

	if err1 == nil && err2 == nil {
		t1.Execute(os.Stdout, &Kayak)
		os.Stdout.WriteString("\n") // важный момент, переход на новую строку. ведь в шаблонах нет этого символа!
		t2.Execute(os.Stdout, &Kayak)
	} else {
		Printfln("Error: %v %v", err1.Error(), err2.Error())
	}

	//
	Printfln("\n\n использование комбинированного шаблона^")

	allTemplates, err1 := template.ParseFiles("templates/template.html", "templates/extras.html")

	if err1 == nil {
		allTemplates.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
		os.Stdout.WriteString("\n")
		allTemplates.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
		os.Stdout.WriteString("\n")
		allTemplates.Execute(os.Stdout, &Kayak) // выведет первый попавшийся шаблон
	} else {
		Printfln("Error: %v %v", err1.Error())
	}

	//
	Printfln("\n\n Перечисление загруженных шаблонов!")

	allTemplates1, err := template.ParseGlob("templates/*.html") // сразу все загрузили
	if err == nil {
		for _, t := range allTemplates1.Templates() {
			Printfln("Template name: %v", t.Name())
		}
	} else {
		Printfln("Error: %v %v", err.Error())
	}

	//
	Printfln("\n\n Поиск шаблона и работа со срезом в шаблоне")

	allTemplates2, err2 := template.ParseGlob("templates/*.html")

	if err2 == nil {
		selectedTemplated := allTemplates2.Lookup("templateS.html")
		err2 = Exec(selectedTemplated)
	}

	if err2 != nil {
		Printfln("Error: %v ", err2.Error())
	}


	//
	Printfln("\n\n Опыт с мапой")
	catMap := map[string]string {}

	for _, p := range Products {
		if catMap[p.Category]=="" {
			catMap[p.Category] = p.Category
			Printfln(catMap[p.Category])
		}
	}
}

func Exec(selectedTemplated *template.Template) error {
	return selectedTemplated.Execute(os.Stdout, Products)
}
