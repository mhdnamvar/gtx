package cmd

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type IsoField struct {
	Id     string `xml:"id,attr"`
	Length string `xml:"length,attr"`
	Name   string `xml:"name,attr"`
	Pad    string `xml:"pad,attr"`
	Class  string `xml:"class,attr"`
}

type IsoPackager struct {
	Name      string
	IsoFields []IsoField `xml:"isofield"`
}

func (isoField *IsoField) String() string {
	l, err := strconv.Atoi(isoField.Length)
	if err != nil {
		log.Fatal(err)
	}
	// &IsoCodec{"DE000", "MTI - MESSAGE TYPE INDICATOR", IsoFixed, IsoAscii, 4, IsoNumeric, IsoLeftPad},
	return fmt.Sprintf("&IsoCodec{\"DE%03s\", \"%s\", IsoFixed, IsoAscii, %d, IsoNumeric, IsoLeftPad}", isoField.Id, isoField.Name, l)
}

func (isoPackager *IsoPackager) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("package iso8583\n\n")
	buffer.WriteString(fmt.Sprintf("var %s = IsoSpec{\n", isoPackager.Name))
	for _, f := range isoPackager.IsoFields {
		buffer.WriteString(fmt.Sprintf("\t%v,\n", &f))
	}
	buffer.WriteString("}")
	return buffer.String()
}

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import jPos ISO8583 protocols",
	Long:  `Use import command to create IsoSpec for GTX using jPos ISO8583 XML packager files.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Fatal("Usage: gtx import <jpos protocol file> <gtx spec file>")
		}
		fmt.Println("Importing ", args[0])
		xmlFile, err := os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}
		defer xmlFile.Close()

		b, _ := ioutil.ReadAll(xmlFile)
		var isoPackager IsoPackager
		isoPackager.Name = args[1]
		xml.Unmarshal(b, &isoPackager)

		fileName := fmt.Sprintf("iso8583/%s.go", strings.ToLower(isoPackager.Name))
		if _, err := os.Stat(fileName); err == nil {
			log.Fatalf("Iso spec %s already exists!", fileName)
		} else if os.IsNotExist(err) {
			f, err := os.Create(fileName)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			_, err = f.WriteString(isoPackager.String())
		} else {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}
