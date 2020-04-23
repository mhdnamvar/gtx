package newtlv

import (
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestTLV_Encode(t *testing.T) {
//	type fields struct {
//		Tag    []byte
//		Length []byte
//		Value  []byte
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		want    []byte
//		wantErr bool
//	}{
//		{"Test1",
//			fields{[]byte("9F10"),
//				nil,
//				[]byte("1F"),
//			},
//			[]byte("9F10021F"),
//			false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tlv := &TLV{
//				Tag:    tt.fields.Tag,
//				Length: tt.fields.Length,
//				Value:  tt.fields.Value,
//			}
//			got, err := tlv.Encode()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Encode() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestTLV_Encode(t *testing.T) {
	tlv := &TLV{
		Tag:        "90",
		Value:      "B790C6A1786C13F10D836E1951B57FD3FF42B913A3BB8A5BAA4BA47B2CF2B642AB683840D8BF571D93907" +
			"E109B7FF066BA09D9C75CE43F1C5F0F798273A864181F22491FF085DCA3EEDA1089BC7C22CE23CAD473A86A9100EC" +
			"1715262CC1F1C256C41C3B93A92D962FF94221C13C96213821C5854A7C77DDFBB49BC4351153FAF08736464837C104" +
			"A37A42037BF22DBE5DE1FC78733ABA4A22D1559A0859051625FEB13378184B19D17B48BF7AC5CB65",
	}
	actual, err := tlv.Encode()
	assert.Equal(t, nil, err)
	//assert.Equal(t, []byte(""), actual)
	//color.Green("%x", actual)
	color.Green("%X", actual)
}

func TestTLV_Parse(t *testing.T) {
	data := "B790C6A1786C13F10D836E1951B57FD3FF42B913A3BB8A5BAA4BA47B2CF2B642AB683840D8BF571D93907E109B7FF066BA09D9C75CE43F1C5F0F798273A864181F22491FF085DCA3EEDA1089BC7C22CE23CAD473A86A9100EC1715262CC1F1C256C41C3B93A92D962FF94221C13C96213821C5854A7C77DDFBB49BC4351153FAF08736464837C104A37A42037BF22DBE5DE1FC78733ABA4A22D1559A0859051625FEB13378184B19D17B48BF7AC5CB65"
	var tlv TLV
	err := tlv.Parse(data)
	assert.Equal(t, nil, err)
	//assert.Equal(t, []byte(""), actual)
	//color.Green("%x", actual)
	color.Green("Len(%X)", tlv.Length)
}
