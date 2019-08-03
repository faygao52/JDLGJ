package models

import (
	"jdlgj/models/base"

	uuid "github.com/satori/go.uuid"
)

type (
	//Message describe a message model
	Message struct {
		base.Base
		Name        string `json:"name" binding:"required"`
		Contact     string `json:"contact" binding:"required"`
		Description string `json:"description"`
		Answered    bool   `json:"visible"`
	}

	//MessageResource represents message for external usage
	MessageResource struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name" binding:"required"`
		Contact     string    `json:"contact" binding:"required"`
		Description string    `json:"description"`
		Answered    bool      `json:"visible"`
	}
)

// ToResource Convert Message to external representation
func (m Message) ToResource() interface{} {
	return MessageResource{
		ID:          m.ID,
		Name:        m.Name,
		Contact:     m.Contact,
		Description: m.Description,
		Answered:    m.Answered,
	}
}
