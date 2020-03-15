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
	folder   bool
	output   bool
	isoMap   map[string]string
)

func (isoField *IsoField) String() string {
	length, err := strconv.Atoi(isoField.Length)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	id := fmt.Sprintf("DE%03s", isoField.Id)
	if isoField.Class == "org.jpos.iso.IFA_BITMAP" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoAscii", length*2, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFA_LLNUM" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLA", "IsoAscii", length, "IsoNumeric", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFA_LLCHAR" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLA", "IsoAscii", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFA_BINARY" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoAscii", length*2, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFA_NUMERIC" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoAscii", length, "IsoNumeric", "IsoLeftPad")
	} else if isoField.Class == "org.jpos.iso.IFA_AMOUNT" {
		// Not correct - Amout field should be defined
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoAscii", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFA_LLLCHAR" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLLA", "IsoAscii", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IF_CHAR" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoAscii", length, "IsoText", "IsoRightPad")
	} else if isoField.Class == "org.jpos.iso.IFA_LLLLBINARY" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLLA", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else {
		log.Fatalf("IsoField class type %s not defined", isoField.Class)
		return ""
	}
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
	importCmd.Flags().StringVarP(&protocol, "protocol", "p", "", "Import jpos XML protocol")
	importCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing protocol otherwise create it")
	importCmd.Flags().BoolVarP(&folder, "find-types", "t", false, "Find iso types from jpos xml protocol files. You need to specify the protocols directory")
	importCmd.Flags().BoolVarP(&output, "find-types-output", "o", false, "Extracted iso types file name")

	isoMap = make(map[string]string)
}

func checkFlags(cmd *cobra.Command, args []string) {
	if protocol != "" {
		importProtocol(cmd, args)
	} else if folder {
		findIsoTypes(cmd, args)
	} else {
		_ = cmd.Usage()
	}
}

func importProtocol(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Println(args)
		_ = cmd.Usage()
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
	_ = xml.Unmarshal(b, &isoPackager)

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

func findIsoTypes(cmd *cobra.Command, args []string) {
	noOfArgs := 1
	if folder && output {
		noOfArgs = 2
	}

	if len(args) < noOfArgs {
		log.Println(args)
		_ = cmd.Usage()
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fileName := fmt.Sprintf("%s/%s", args[0], f.Name())
		log.Printf("check %s...", fileName)
		xmlFile, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}

		b, _ := ioutil.ReadAll(xmlFile)
		var isoPackager IsoPackager
		isoPackager.Name = f.Name()
		_ = xml.Unmarshal(b, &isoPackager)

		for _, f := range isoPackager.IsoFields {
			isoMap[f.Class] = f.Class
		}
		xmlFile.Close()
	}

	if output {
		f, err := os.Create(args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		for _, s := range isoMap {
			_, err = f.WriteString(s)
			_, err = f.WriteString("\r\n")
		}
		log.Printf("Result: %d jpos type(s) found, for more details check %s", len(isoMap), args[1])
	} else {
		for _, s := range isoMap {
			fmt.Println(s)
		}
	}
}
