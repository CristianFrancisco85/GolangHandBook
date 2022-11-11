package models

type WebSocketMessage struct {
	Type    string      `json:"type"`    // Metodo HTTP de la peticion
	Payload interface{} `json:"payload"` // Payload de la peticion este lleva el contenido del mensaje
}
