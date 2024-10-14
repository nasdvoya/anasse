package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Partner struct {
	Nome          string `json:"nome"`
	Morada        string `json:"morada"`
	Freguesia     string `json:"freguesia"`
	Conselho      string `json:"conselho"`
	LocalCobranca string `json:"local_cobranca"`
	Doador        string `json:"doador"`
	Observacoes   string `json:"observacoes"`
	CP            string `json:"cp"`
	TelResid      string `json:"tel_resid"`
	TelTrab       string `json:"tel_trab"`
	Telemovel     string `json:"telemovel"`
	Email         string `json:"email"`
}

var DB *sql.DB

func IntializeDatabase() {
	var err error

	DB, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS partners (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT,
		morada TEXT,
		freguesia TEXT,
		conselho TEXT,
		local_cobranca TEXT,
		doador TEXT,
		observacoes TEXT,
		cp TEXT,
		tel_resid TEXT,
		tel_trab TEXT,
		telemovel TEXT,
		email TEXT
	)`

	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func CreatePartner(partner Partner) (int64, error) {
	query := `
	INSERT INTO partners (
		nome, morada, freguesia, conselho, local_cobranca, doador,
		observacoes, cp, tel_resid, tel_trab, telemovel, email
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := DB.Exec(
		query,
		partner.Nome, partner.Morada, partner.Freguesia, partner.Conselho,
		partner.LocalCobranca, partner.Doador, partner.Observacoes, partner.CP,
		partner.TelResid, partner.TelTrab, partner.Telemovel, partner.Email,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeletePartner(id int64) error {
	query := `
	DELETE FROM partners
	WHERE id = ?
	`

	_, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func ReadPartnerList() []Partner {
	rows, _ := DB.Query("SELECT Nome, Morada FROM partners")
	defer rows.Close()

	partners := make([]Partner, 0)
	for rows.Next() {
		var partner Partner
		rows.Scan(&partner.Nome, &partner.Morada)
		partners = append(partners, partner)
	}
	return partners
}
