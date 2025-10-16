// Package router fournit un simple routeur de messages.
package router

import "fmt"

// Message représente un message simple avec un contenu.
type Message struct {
	Contenu string
}

// Routeur représente un routeur de messages qui peut enregistrer des gestionnaires pour des types de messages spécifiques.
type Routeur struct {
	handlers map[string]func(Message)
}

// NouveauRouteur crée une nouvelle instance de Routeur.
func NouveauRouteur() *Routeur {
	return &Routeur{
		handlers: make(map[string]func(Message)),
	}
}

// EnregistrerGestionnaire enregistre une fonction de gestionnaire pour un type de message spécifique.
func (r *Routeur) EnregistrerGestionnaire(messageType string, handler func(Message)) {
	r.handlers[messageType] = handler
}

// RouterMessage achemine un message vers le gestionnaire approprié en fonction de son type.
func (r *Routeur) RouterMessage(messageType string, message Message) {
	if handler, ok := r.handlers[messageType]; ok {
		handler(message)
	} else {
		fmt.Printf("Aucun gestionnaire trouvé pour le type de message : %s\n", messageType)
	}
}