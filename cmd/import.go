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
	protocol          string
	force             bool
	scanTypes         bool
	output            bool
	scanTypesMissing  bool
	isoMap            map[string]string
	supportedIsoTypes map[string]bool
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
	} else if isoField.Class == "org.jpos.iso.IFA_LLLNUM" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLLA", "IsoAscii", length, "IsoNumeric", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFA_LLCHAR" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLA", "IsoAscii", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFA_BINARY" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoAscii", length*2, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFA_NUMERIC" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoAscii", length, "IsoNumeric", "IsoLeftPad")
	} else if isoField.Class == "org.jpos.iso.IFA_AMOUNT" {
		// TODO: Amount field should be defined
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoAscii", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFA_LLLCHAR" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLLA", "IsoAscii", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IF_CHAR" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoAscii", length, "IsoText", "IsoRightPad")
	} else if isoField.Class == "org.jpos.iso.IFA_LLBINARY" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLA", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFA_LLLBINARY" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLLA", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFA_LLLLBINARY" {
		//TODO: Should be revised, not sure about LLLL encoding!
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLLA", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFE_BITMAP" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoEbcdic", length*2, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFE_LLNUM" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLE", "IsoEbcdic", length, "IsoNumeric", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFE_LLCHAR" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLE", "IsoEbcdic", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFE_BINARY" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoEbcdic", length*2, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFE_NUMERIC" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoEbcdic", length, "IsoNumeric", "IsoLeftPad")
	} else if isoField.Class == "org.jpos.iso.IFE_AMOUNT" {
		// TODO: Amount field should be defined
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoEbcdic", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFE_CHAR" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoEbcdic", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IF_ECHAR" { //deprecated type use IFE_CHAR instead
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoEbcdic", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFE_LLLCHAR" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLLE", "IsoEbcdic", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFE_LLBINARY" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLE", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFE_LLLBINARY" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLLE", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFE_LLLLBINARY" {
		//TODO: Should be revised, not sure about LLLL encoding!
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLLE", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_BITMAP" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_LLNUM" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLB", "IsoBinary", length, "IsoNumeric", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_LLLNUM" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLLB", "IsoBinary", length, "IsoNumeric", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_LLCHAR" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLB", "IsoAscii", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_BINARY" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_NUMERIC" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoBinary", length/2, "IsoNumeric", "IsoLeftPad")
	} else if isoField.Class == "org.jpos.iso.IFB_AMOUNT" {
		// TODO: Amount field should be defined
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoBinary", length, "IsoText", "IsoLeftPad")
	} else if isoField.Class == "org.jpos.iso.IFB_CHAR" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoFixed", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_LLHCHAR" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLB", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_LLLCHAR" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLLB", "IsoAscii", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_LLBINARY" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLB", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_LLLBINARY" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLLB", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_LLLLBINARY" {
		//TODO: Should be revised, not sure about LLLL encoding!
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLLB", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_LLHECHAR" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLB", "IsoEbcdic", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_LLHBINARY" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLB", "IsoBinary", length, "IsoText", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_LLHNUM" {
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLB", "IsoBinary", length, "IsoNumeric", "IsoNoPad")
	} else if isoField.Class == "org.jpos.iso.IFB_FLLNUM" {
		//TODO: Right pad with F, padding should be defined
		l := length
		if length%2 != 0 {
			l = l/2 + 1
		}
		return fmt.Sprintf("&IsoCodec{\"%s\", \"%s\", %s, %s, %d, %s, %s}", id, isoField.Name, "IsoLLB", "IsoBinary", l, "IsoNumeric", "IsoRightPad")
	} else {
		log.Fatalf("IsoField type %s not supported!", isoField.Class)
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
	importCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing protocol if exists")
	importCmd.Flags().BoolVarP(&scanTypes, "scan-types", "s", false, "Find iso types from jpos xml protocol files")
	importCmd.Flags().BoolVarP(&scanTypesMissing, "scan-types-missing", "m", false, "Find only missing iso types")
	importCmd.Flags().BoolVarP(&output, "scan-types-output", "o", false, "Write iso types to specific file")

	isoMap = make(map[string]string)
	supportedIsoTypes = make(map[string]bool)
	supportedIsoTypes["org.jpos.iso.IFA_BITMAP"] = true
	supportedIsoTypes["org.jpos.iso.IFA_LLNUM"] = true
	supportedIsoTypes["org.jpos.iso.IFA_LLLNUM"] = true
	supportedIsoTypes["org.jpos.iso.IFA_LLCHAR"] = true
	supportedIsoTypes["org.jpos.iso.IFA_BINARY"] = true
	supportedIsoTypes["org.jpos.iso.IFA_NUMERIC"] = true
	supportedIsoTypes["org.jpos.iso.IFA_AMOUNT"] = true
	supportedIsoTypes["org.jpos.iso.IFA_LLLCHAR"] = true
	supportedIsoTypes["org.jpos.iso.IF_CHAR"] = true
	supportedIsoTypes["org.jpos.iso.IFA_LLBINARY"] = true
	supportedIsoTypes["org.jpos.iso.IFA_LLLBINARY"] = true
	supportedIsoTypes["org.jpos.iso.IFA_LLLLBINARY"] = true
	supportedIsoTypes["org.jpos.iso.IFE_BITMAP"] = true
	supportedIsoTypes["org.jpos.iso.IFE_LLNUM"] = true
	supportedIsoTypes["org.jpos.iso.IFE_LLCHAR"] = true
	supportedIsoTypes["org.jpos.iso.IFE_BINARY"] = true
	supportedIsoTypes["org.jpos.iso.IFE_NUMERIC"] = true
	supportedIsoTypes["org.jpos.iso.IFE_AMOUNT"] = true
	supportedIsoTypes["org.jpos.iso.IFE_CHAR"] = true
	supportedIsoTypes["org.jpos.iso.IF_ECHAR"] = true
	supportedIsoTypes["org.jpos.iso.IFE_LLLCHAR"] = true
	supportedIsoTypes["org.jpos.iso.IFE_LLBINARY"] = true
	supportedIsoTypes["org.jpos.iso.IFE_LLLBINARY"] = true
	supportedIsoTypes["org.jpos.iso.IFE_LLLLBINARY"] = true
	supportedIsoTypes["org.jpos.iso.IFB_BITMAP"] = true
	supportedIsoTypes["org.jpos.iso.IFB_LLNUM"] = true
	supportedIsoTypes["org.jpos.iso.IFB_LLLNUM"] = true
	supportedIsoTypes["org.jpos.iso.IFB_LLCHAR"] = true
	supportedIsoTypes["org.jpos.iso.IFB_BINARY"] = true
	supportedIsoTypes["org.jpos.iso.IFB_NUMERIC"] = true
	supportedIsoTypes["org.jpos.iso.IFB_AMOUNT"] = true
	supportedIsoTypes["org.jpos.iso.IFB_CHAR"] = true
	supportedIsoTypes["org.jpos.iso.IFB_LLLCHAR"] = true
	supportedIsoTypes["org.jpos.iso.IFB_LLBINARY"] = true
	supportedIsoTypes["org.jpos.iso.IFB_LLLBINARY"] = true
	supportedIsoTypes["org.jpos.iso.IFB_LLLLBINARY"] = true
	supportedIsoTypes["org.jpos.iso.IFB_LLHECHAR"] = true
	supportedIsoTypes["org.jpos.iso.IFB_LLHBINARY"] = true
	supportedIsoTypes["org.jpos.iso.IFB_LLHNUM"] = true
	supportedIsoTypes["org.jpos.iso.IFB_FLLNUM"] = true
	supportedIsoTypes["org.jpos.iso.IFB_LLHCHAR"] = true

}

func checkFlags(cmd *cobra.Command, args []string) {
	if protocol != "" {
		importProtocol(cmd, args)
	} else if scanTypes {
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
	if scanTypes && output {
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
		noMissing := 0
		for _, s := range isoMap {
			s = strings.TrimRight(s, " ")
			if scanTypesMissing {
				if !supportedIsoTypes[s] {
					noMissing++
					_, err = f.WriteString(s)
					_, err = f.WriteString("\r\n")
				}
			} else {
				_, err = f.WriteString(s)
				_, err = f.WriteString("\r\n")
			}
		}
		log.Printf("Result: %d jpos type(s) found, %d type(s) missing, for more details check %s", len(isoMap), noMissing, args[1])
	} else {
		for _, s := range isoMap {
			s = strings.TrimRight(s, " ")
			if scanTypesMissing {
				if !supportedIsoTypes[s] {
					log.Println("missing", s)
				}
			} else {
				log.Println(s)
			}
		}
	}
}
