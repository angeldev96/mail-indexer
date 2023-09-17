package models

type Email struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func FindEmailsByTerm(term string) []Email {
	// Aquí deberías conectar con tu base de datos o lógica real para buscar correos electrónicos.
	// Por ahora, devolvemos un resultado ficticio.
	return []Email{
		{
			ID:      "1",
			Subject: "Test Email",
			Body:    "This is a test email.",
		},
	}
}
