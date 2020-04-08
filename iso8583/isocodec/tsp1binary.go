package isocodec

var Tsp1Binary = IsoSpec{
	
	/* DE000 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 4,
			Max: 4,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE001 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 16,
			Max: 16,
			ContentType: IsoBitmap,
			Padding: IsoNoPad,
		},
	},
	
	/* DE002 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 19,
			ContentType: IsoNumeric,
			Padding: IsoRightPadF,
		},
	},
	
	/* DE003 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 6,
			Max: 6,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE004 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 12,
			Max: 12,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE005 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 12,
			Max: 12,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE006 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 12,
			Max: 12,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE007 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 10,
			Max: 10,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE008 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE009 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE010 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE011 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 6,
			Max: 6,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE012 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 6,
			Max: 6,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE013 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 4,
			Max: 4,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE014 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 4,
			Max: 4,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE015 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 4,
			Max: 4,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE016 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 4,
			Max: 4,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE017 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 4,
			Max: 4,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE018 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 4,
			Max: 4,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE019 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 3,
			Max: 3,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE020 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 3,
			Max: 3,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE021 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 3,
			Max: 3,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE022 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 3,
			Max: 3,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE023 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 3,
			Max: 3,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE024 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 3,
			Max: 3,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE025 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE026 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE027 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE028 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoString,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 9,
			Max: 9,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE029 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoString,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 9,
			Max: 9,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE030 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoString,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 9,
			Max: 9,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE031 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoString,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 9,
			Max: 9,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE032 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 11,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE033 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 11,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE034 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 28,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE035 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 37,
			ContentType: IsoNumeric,
			Padding: IsoRightPadF,
		},
	},
	
	/* DE036 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 104,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE037 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 12,
			Max: 12,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE038 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 6,
			Max: 6,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE039 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 2,
			Max: 2,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE040 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 3,
			Max: 3,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE041 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 8,
			Max: 8,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE042 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 15,
			Max: 15,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE043 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 40,
			Max: 40,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE044 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 25,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE045 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 76,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE046 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE047 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE048 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE049 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 3,
			Max: 3,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE050 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 3,
			Max: 3,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE051 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 3,
			Max: 3,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE052 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE053 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 16,
			Max: 16,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE054 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 120,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE055 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE056 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE057 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE058 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE059 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE060 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE061 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE062 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE063 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE064 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE065 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE066 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE067 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE068 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 3,
			Max: 3,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE069 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 3,
			Max: 3,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE070 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 3,
			Max: 3,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE071 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 4,
			Max: 4,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE072 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 4,
			Max: 4,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE073 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 6,
			Max: 6,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE074 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 10,
			Max: 10,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE075 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 10,
			Max: 10,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE076 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 10,
			Max: 10,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE077 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 10,
			Max: 10,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE078 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 10,
			Max: 10,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE079 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 10,
			Max: 10,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE080 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 10,
			Max: 10,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE081 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 10,
			Max: 10,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE082 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 12,
			Max: 12,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE083 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 12,
			Max: 12,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE084 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 12,
			Max: 12,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE085 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 12,
			Max: 12,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE086 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 16,
			Max: 16,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE087 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 16,
			Max: 16,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE088 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 16,
			Max: 16,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE089 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 16,
			Max: 16,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE090 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 42,
			Max: 42,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE091 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 1,
			Max: 1,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE092 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 2,
			Max: 2,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE093 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 6,
			Max: 6,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE094 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 7,
			Max: 7,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE095 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 42,
			Max: 42,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE096 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE097 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoString,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 17,
			Max: 17,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
	},
	
	/* DE098 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoAscii, 
			Min: 25,
			Max: 25,
			ContentType: IsoString,
			Padding: IsoRightPad,
		},
	},
	
	/* DE099 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 11,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE100 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 11,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	},
	
	/* DE101 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 17,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE102 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 28,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE103 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 28,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE104 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 100,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE105 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE106 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE107 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE108 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE109 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE110 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE111 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE112 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE113 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE114 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE115 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE116 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE117 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE118 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE119 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE120 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE121 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE122 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE123 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE124 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE125 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE126 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE127 */
	 &IsoType{
		Len: &IsoData{
			Encoding:IsoBinary, 
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		}, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 0,
			Max: 999,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	},
	
	/* DE128 */
	 &IsoType{
		Len: nil, 
		Value: &IsoData{
			Encoding:IsoBinary, 
			Min: 8,
			Max: 8,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	},
}