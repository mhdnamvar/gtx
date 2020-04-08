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

	. "../iso8583/isocodec"
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
	force             bool
)

func (isoField *IsoField) String() string {
	isoType, err := isoField.GetIsoType()
	if err != nil {
		log.Fatalf("[%s] %v", isoField.Class, err)
		return ""
	}
	return fmt.Sprintf("\n\t/* DE%03s */\n\t %v", isoField.Id, isoType)

}

func (isoPackager *IsoPackager) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("package isocodec\n\n")
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
	Short: "Import protocol",
	Long: `Use import command to import pos ISO8583 protocol e.g. gtx import protocol-iso87binary.xml Binary87`,
	Run: checkFlags,
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing protocol if exists")
}

func checkFlags(cmd *cobra.Command, args []string) {
	importProtocol(cmd, args)
}

func importProtocol(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		_ = cmd.Usage()
		os.Exit(1)
	}
	log.Println("Importing jpos protocol", args[0])
	xmlFile, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)
	var isoPackager IsoPackager
	isoPackager.Name = args[1]
	_ = xml.Unmarshal(b, &isoPackager)

	fileName := fmt.Sprintf("iso8583/isocodec/%s.go", strings.ToLower(isoPackager.Name))
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

func (isoField *IsoField) GetIsoType() (*IsoType, error){
	var isoType *IsoType
	length, err := strconv.Atoi(isoField.Length)
	if err != nil {
		log.Fatal(err)
		return isoType, InvalidData
	}

	if strings.EqualFold(isoField.Class, "org.jpos.iso.IFA_BITMAP") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         length,
				Max:         length,
				ContentType: IsoBitmap,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFA_BINARY") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         length,
				Max:         length,
				ContentType: IsoHexString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IF_CHAR") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         length,
				Max:         length,
				ContentType: IsoString,
				Padding:     IsoRightPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFA_LLCHAR") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         0,
				Max:         2,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         0,
				Max:         length,
				ContentType: IsoString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFA_LLBINARY") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         1,
				Max:         1,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         0,
				Max:         length,
				ContentType: IsoHexString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFA_LLLCHAR") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         2,
				Max:         2,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         0,
				Max:         length,
				ContentType: IsoString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFA_LLLBINARY") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         2,
				Max:         2,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         0,
				Max:         length,
				ContentType: IsoHexString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFA_NUMERIC") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         length,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFA_LLNUM") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         1,
				Max:         1,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         0,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFA_LLLNUM") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         2,
				Max:         2,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         0,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFA_AMOUNT") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoAscii,
				Min:         1,
				Max:         1,
				ContentType: IsoString,
				Padding:     IsoNoPad,
			},
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         length,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFA_AMOUNT2003") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         0,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFA_LLLLBINARY") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         2,
				Max:         2,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         length,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFE_BITMAP") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         length,
				Max:         length,
				ContentType: IsoBitmap,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFE_LLNUM") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         1,
				Max:         1,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
			Value: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         0,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFE_LLCHAR") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         0,
				Max:         2,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         0,
				Max:         length,
				ContentType: IsoString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFE_BINARY") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         length,
				Max:         length,
				ContentType: IsoHexString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFE_NUMERIC") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         length,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFE_CHAR") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         length,
				Max:         length,
				ContentType: IsoString,
				Padding:     IsoRightPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IF_ECHAR") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         length,
				Max:         length,
				ContentType: IsoString,
				Padding:     IsoRightPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFE_LLLCHAR") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         2,
				Max:         2,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         0,
				Max:         length,
				ContentType: IsoString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFE_LLBINARY") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         1,
				Max:         1,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
			Value: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         0,
				Max:         length,
				ContentType: IsoHexString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFE_LLLBINARY") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         2,
				Max:         2,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         0,
				Max:         length,
				ContentType: IsoHexString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFE_LLLLBINARY") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         2,
				Max:         2,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         0,
				Max:         length,
				ContentType: IsoHexString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFE_AMOUNT") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         1,
				Max:         1,
				ContentType: IsoString,
				Padding:     IsoNoPad,
			},
			Value: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         length,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_BITMAP") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         length,
				Max:         length,
				ContentType: IsoBitmap,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_LLNUM") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         1,
				Max:         1,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         0,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_LLLNUM") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         2,
				Max:         2,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         0,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_LLCHAR") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         0,
				Max:         2,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         0,
				Max:         length,
				ContentType: IsoString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_BINARY") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         length,
				Max:         length,
				ContentType: IsoHexString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_NUMERIC") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         length,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_CHAR") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         length,
				Max:         length,
				ContentType: IsoString,
				Padding:     IsoRightPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_LLLCHAR") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         2,
				Max:         2,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         0,
				Max:         length,
				ContentType: IsoString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_LLBINARY") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         1,
				Max:         1,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         0,
				Max:         length,
				ContentType: IsoHexString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_LLLBINARY") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         2,
				Max:         2,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         0,
				Max:         length,
				ContentType: IsoHexString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_LLLLBINARY") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         2,
				Max:         2,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         0,
				Max:         length,
				ContentType: IsoHexString,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_AMOUNT") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         1,
				Max:         1,
				ContentType: IsoString,
				Padding:     IsoNoPad,
			},
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         length,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_LLHECHAR") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         2,
				Max:         2,
				ContentType: IsoHexString,
				Padding:     IsoNoPad,
			},
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         length,
				Max:         length,
				ContentType: IsoString,
				Padding:     IsoLeftPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_LLHCHAR") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoAscii,
				Min:         2,
				Max:         2,
				ContentType: IsoHexString,
				Padding:     IsoNoPad,
			},
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         length,
				Max:         length,
				ContentType: IsoString,
				Padding:     IsoLeftPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_LLHBINARY") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         2,
				Max:         2,
				ContentType: IsoHexString,
				Padding:     IsoNoPad,
			},
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         length,
				Max:         length,
				ContentType: IsoHexString,
				Padding:     IsoLeftPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_LLHNUM") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         2,
				Max:         2,
				ContentType: IsoHexString,
				Padding:     IsoLeftPad,
			},
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         0,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_FLLNUM") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoBinary,
				Min:         1,
				Max:         1,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         0,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoRightPadF,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFA_FLLNUM") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoAscii,
				Min:         1,
				Max:         1,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         0,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoRightPadF,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFA_FLLCHAR") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoAscii,
				Min:         1,
				Max:         1,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
			Value: &IsoData{
				Encoding:    IsoAscii,
				Min:         0,
				Max:         length,
				ContentType: IsoString,
				Padding:     IsoRightPadF,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFB_AMOUNT2003") {
		isoType = &IsoType{
			Len: nil,
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         0,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoLeftPad,
			},
		}
	} else if strings.EqualFold(isoField.Class, "org.jpos.iso.IFEB_LLNUM") {
		isoType = &IsoType{
			Len: &IsoData{
				Encoding:    IsoEbcdic,
				Min:         1,
				Max:         1,
				ContentType: IsoNumeric,
				Padding:     IsoNoPad,
			},
			Value: &IsoData{
				Encoding:    IsoBinary,
				Min:         0,
				Max:         length,
				ContentType: IsoNumeric,
				Padding:     IsoRightPadF,
			},
		}
	}

	if isoType != nil {
		return isoType, nil
	}
	return isoType, NotSupported
}