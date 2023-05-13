package data

import "fmt"

// Мой сервис, который работает только с json

type JsonDocument struct {
}

func (doc JsonDocument) ConvertToXml() string {
	return "<xml></xml>"
}

// структура адаптера

type JsonDocumentAdapter struct {

	// инжектим структуру JsonDocument
	jsonDocument *JsonDocument
}

/*
Уже не реализуем SendXmlData интерфейса AnalyticalDataService, а реализуем SendXmlData
у адаптера
*/

func (adapter JsonDocumentAdapter) SendXmlData() {
	adapter.jsonDocument.ConvertToXml()
	fmt.Println("Отправка xml данных!")

}
