package main

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

func main() {
	fmt.Println(addPng("test"))
}

func addPng(name string) string {
	return name + ".png"
}

// this program takes an Invoice and writes it into a text file.

type Invoice struct {
	Id         int
	CustomerId int
	Raised     time.Time
	Due        time.Time
	Paid       bool
	Note       string
	Items      []*Item
}

type Item struct {
	Id       string
	Price    float64
	Quantity int
	Note     string
}

//make it easy to use a reader or writer for a specifc format in a generic way
type InvoiceMarshaler interface {
	MarshalInvoices(writer io.Writer, invoices []*Invoice) error
}

type TxtMarshaler struct{}

func (TxtMarshaler) MarshalInvoices(writer io.Writer, invoices []*Invoice) error {
	bufferedWriter := bufio.NewWriter(writer)
	defer bufferedWriter.Flush()
	// closure: defined as an anonymous function. Has access to bufferedWriter
	var write writerFunc = func(format string, args ...interface{}) error {
		_, err := fmt.Fprintf(bufferedWriter, format, args...)
		return err
	}
	if err := write("%s %d\n", fileType, fileVersion); err != nil {
		return err
	}

	//iterate over every invoice and call writeInvoice method on each
	for _, invoice := range invoices {
		if err := write.writeInvoice(invoice); err != nil {
			return err
		}
	}
	return nil
}

type writerFunc func(string, ...interface{}) error

// writeInvoice is a method of the write() function b/c we gave the write() function the type writerFunc
func (write writerFunc) writeInvoice(invoice *Invoice) error {
	note := ""
	if invoice.Note != "" {
		note = ": " + invoice.Note
	}
	// %t used for Boolean
	if err := write("INVOICE ID=%d CUSTOMER=%d RAISED=%s DUE=%s PAID=%t%s\n", invoice.Id, invoice.CustoemrId, invoice.Raised.Format(dateFormat),
		invoice.Due.Format(dateFormat), invoice.Paid, note); error != nil {
		return err
	}

	// call writeItems on the items in the invoice
	if err := write.writeItems(invoice.Items); err != nil {
		return err
	}
	return write("\f\n")
}

func (write writerFunc) writeItems(items []*Item) error {
	for _, item := range items {
		note := ""
		if item.Note != "" {
			note = ": " + item.Note
		}
		if err := write("ITEM ID=%s PRICE=%.2f QUANTITY=%d%s\n", item.Id, item.Price, item.Quantity, note); err != nil {
			return err
		}
	}
	return nil
}
