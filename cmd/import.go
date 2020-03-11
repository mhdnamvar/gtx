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

var (
	protocol string
	force    bool
)

func (isoField *IsoField) String() string {
	length, err := strconv.Atoi(isoField.Length)
	if err != nil {
		log.Fatal(err)
	}
	contentType := "IsoText"
	lenType := "IsoFixed"
	encodingType := "IsoAscii"
	paddingType := "IsoNoPad"

	// Encoding
	if strings.Contains(isoField.Class, "IFB") {
		encodingType = "IsoBinary"
		// Length
		if strings.Contains(isoField.Class, "LLL") {
			lenType = "IsoLLLB"
		} else if strings.Contains(isoField.Class, "LL") {
			lenType = "IsoLLB"
		}
	} else if strings.Contains(isoField.Class, "IFE") {
		encodingType = "IsoEbcdic"
		// Length
		if strings.Contains(isoField.Class, "LLL") {
			lenType = "IsoLLLE"
		} else if strings.Contains(isoField.Class, "LL") {
			lenType = "IsoLLE"
		}
	}

	// ContentType
	if strings.Contains(isoField.Class, "NUMERIC") || strings.Contains(isoField.Class, "NUM") {
		contentType = "IsoNumeric"
	}

	// Padding
	if strings.ToLower(isoField.Pad) == "true" {
		paddingType = "IsoLeftPad"
	} else if strings.ToLower(isoField.Pad) == "false" {
		paddingType = "IsoNoPad"
	} else {
		if lenType == "IsoFixed" && encodingType != "IsoBinary" {
			if contentType == "IsoNumeric" {
				paddingType = "IsoLeftPad"
			} else {
				paddingType = "IsoRightPad"
			}
		}
	}

	return fmt.Sprintf("&IsoCodec{\"DE%03s\", \"%s\", %s, %s, %d, %s, %s}",
		isoField.Id, isoField.Name, lenType, encodingType, length, contentType, paddingType)
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
	Short: "Import configuration",
	Long: `Use import command to import external configuration like jpos ISO8583 protocols.
e.g. gtx import --protocol protocol-iso87binary.xml Binary87`,
	Run: checkFlags,
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().StringVarP(&protocol, "protocol", "p", "", "Import jpos protocol.")
	importCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite protocol if exists.")
}

func checkFlags(cmd *cobra.Command, args []string) {
	if protocol != "" {
		importProtocol(cmd, args)
	} else {
		cmd.Usage()
	}
}

func importProtocol(cmd *cobra.Command, args []string) {

	if len(args) < 1 {
		log.Println(args)
		cmd.Usage()
		os.Exit(1)
	}
	log.Println("Importing jpos protocol", protocol)
	xmlFile, err := os.Open(protocol)
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)
	var isoPackager IsoPackager
	isoPackager.Name = args[0]
	xml.Unmarshal(b, &isoPackager)

	fileName := fmt.Sprintf("iso8583/%s.go", strings.ToLower(isoPackager.Name))
	if _, err := os.Stat(fileName); err == nil {
		if force {
			log.Println(fmt.Sprintf("Iso spec %s already exists, it will be overwritten!", fileName))
			create(isoPackager, fileName)
		} else {
			log.Fatalf("Import failed. Iso spec \"%s\" already exists!", fileName)
		}
	} else if os.IsNotExist(err) {
		create(isoPackager, fileName)
	} else {
		log.Fatal(err)
	}
}

func create(isoPackager IsoPackager, fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(isoPackager.String())
	log.Println("The protocol imported successfully.")
}
