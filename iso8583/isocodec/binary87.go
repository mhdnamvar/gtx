package isocodec

var Binary87 = IsoSpec{
	
	/* DE000 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE001 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 16,
			Max: 16,
			ContentType: IsoBitmap,
			Padding: IsoNoPad,
		},
	},
	
	/* DE002 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 0,
			Max: 19,
			ContentType: IsoNumeric,
			Padding: IsoRightPadF,
		},
	},
	
	/* DE003 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 3,
			Max: 3,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE004 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 6,
			Max: 6,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE005 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 6,
			Max: 6,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE006 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 6,
			Max: 6,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE007 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 5,
			Max: 5,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE008 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 4,
			Max: 4,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE009 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 4,
			Max: 4,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE010 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 4,
			Max: 4,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE011 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 3,
			Max: 3,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE012 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 3,
			Max: 3,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE013 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE014 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE015 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE016 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE017 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE018 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE019 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE020 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE021 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE022 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE023 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE024 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE025 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE026 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE027 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE028 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoString,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 0,
			Max: 9,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE029 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoString,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 0,
			Max: 9,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE030 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoString,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 0,
			Max: 9,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE031 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoString,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 0,
			Max: 9,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE032 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 0,
			Max: 11,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE033 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 0,
			Max: 11,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE034 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 28,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE035 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 0,
			Max: 37,
			ContentType: IsoNumeric,
			Padding: IsoRightPadF,
		},
	},
	
	/* DE036 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 104,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE037 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 12,
			Max: 12,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE038 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 6,
			Max: 6,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE039 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 2,
			Max: 2,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE040 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 3,
			Max: 3,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE041 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 8,
			Max: 8,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE042 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 15,
			Max: 15,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE043 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 40,
			Max: 40,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE044 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 25,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE045 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 76,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE046 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE047 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE048 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE049 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE050 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 3,
			Max: 3,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE051 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 3,
			Max: 3,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE052 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE053 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE054 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 120,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE055 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE056 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE057 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE058 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE059 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE060 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE061 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE062 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE063 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE064 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE065 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE066 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE067 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE068 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE069 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE070 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE071 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE072 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE073 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 3,
			Max: 3,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE074 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 5,
			Max: 5,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE075 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 5,
			Max: 5,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE076 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 5,
			Max: 5,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE077 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 5,
			Max: 5,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE078 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 5,
			Max: 5,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE079 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 5,
			Max: 5,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE080 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 5,
			Max: 5,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE081 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 5,
			Max: 5,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE082 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 6,
			Max: 6,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE083 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 6,
			Max: 6,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE084 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 6,
			Max: 6,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE085 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 6,
			Max: 6,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE086 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE087 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE088 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE089 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE090 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 21,
			Max: 21,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE091 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 1,
			Max: 1,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE092 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 2,
			Max: 2,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE093 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 6,
			Max: 6,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE094 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 7,
			Max: 7,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE095 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 42,
			Max: 42,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE096 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE097 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoString,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 0,
			Max: 17,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE098 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 25,
			Max: 25,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE099 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 0,
			Max: 11,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE100 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 0,
			Max: 11,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE101 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 17,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE102 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 28,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE103 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 28,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE104 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 100,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE105 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE106 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE107 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE108 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE109 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE110 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE111 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE112 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE113 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE114 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE115 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE116 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE117 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE118 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE119 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE120 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE121 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE122 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE123 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE124 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE125 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE126 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE127 */
	 &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding: IsoAscii, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE128 */
	 &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	},
}