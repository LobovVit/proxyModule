package model

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type DocumentCard struct {
	Source string

	Guid uuid.UUID
	Date time.Time
	Num  string
	Type string

	BaseGuid uuid.UUID
	BaseDate time.Time
	BaseNum  string
	BaseType string

	CreationDate time.Time
	CreatedBy    string
	LastUpdDate  time.Time
	LastUpdBy    string

	OrgFrom    string
	OrgTo      string
	PaFrom     string
	PaTo       string
	BudgetFrom string
	BudgetTo   string

	CurrencyCode string
}

type DocumentHeader struct {
	HAttribute1  string
	HAttribute2  string
	HAttribute3  string
	HAttribute4  string
	HAttribute5  string
	HAttribute6  string
	HAttribute7  string
	HAttribute8  string
	HAttribute9  string
	HAttribute10 string
}

type DocumentLine struct {
	LEnteredAmount   decimal.Decimal
	lAccountedAmount decimal.Decimal

	LAttribute1  string
	LAttribute2  string
	LAttribute3  string
	LAttribute4  string
	LAttribute5  string
	LAttribute6  string
	LAttribute7  string
	LAttribute8  string
	LAttribute9  string
	LAttribute10 string
}

type Document struct {
	Card   DocumentCard
	Header DocumentHeader
	Lines  []DocumentLine
}

func NewUuid() {
	var x DocumentCard
	x.Guid = uuid.New()
	fmt.Print(x.Guid)
	var y = DocumentLine{decimal.NewFromInt(10), decimal.NewFromInt(7), "", "", "", "", "", "", "", "", "", ""}
	z := y.LEnteredAmount.Div(y.lAccountedAmount)
	fmt.Print(z)
}
