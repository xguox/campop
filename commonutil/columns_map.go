package commonutil

import "sort"

// CameraColumnsMapConst 相机字段中英映射

var CameraColumnsMapConstEnFirst = map[string]string{
	"Name":                   "名称",
	"ProductType":            "产品类型",
	"SensorType":             "影像传感器",
	"ReleaseDate":            "发布日期",
	"Discontinued":           "停产",
	"ImageSensorSize":        "影像传感器尺寸",
	"TotalPixels":            "总像素",
	"EffectivePixels":        "有效像素",
	"ImageProcessor":         "影像处理器",
	"StaticImageFormat":      "静态图像格式",
	"StaticImageSize":        "静态图像尺寸",
	"FrameRatio":             "画幅比",
	"MovieFormat":            "短片格式",
	"MovieSize":              "短片尺寸",
	"AudioFormat":            "音频格式",
	"DustRemovalSystem":      "除尘系统",
	"LensStructure":          "镜头结构",
	"AsphericalLenses":       "非球面镜片",
	"FocalLength":            "焦距",
	"EqFullFrameFocalLength": "等效35mm画幅焦距",
	"MaximumAperture":        "最大光圈",
	"OpticalZoom":            "光学变焦",
	"RecentFocusingDistance": "最近对焦距离",
	"FocusMode":              "对焦模式",
	"LensMount":              "镜头卡口",
	"CompatibleLens":         "兼容镜头",
	"FocusSystemType":        "对焦系统类型",
	"Focus":                  "对焦点",
	"MeteringSystem":         "测光系统",
	"MeteringMode":           "测光模式",
	"AfAssistLight":          "AF辅助灯/光",
	"ExposureMode":           "曝光模式",
	"ShootingMode":           "拍摄模式",
	"FilmSimulationMode":     "胶片模拟模式",
	"AdvancedFilter":         "创意滤镜",
	"ShutterType":            "快门类型",
	"ShutterSpeed":           "快门速度",
	"FlashSyncSpeed":         "闪光同步速度",
	"ExposureCompensation":   "曝光补偿",
	"IsoSensitivity":         "ISO感光度",
	"WhiteBalance":           "白平衡",
	"SelfTimer":              "自拍",
	"ContinuousShot":         "连拍",
	"Viewfinder":             "取景器",
	"ViewingRange":           "取景范围",
	"ViewingMagnification":   "取景放大倍率",
	"DiopterAdjustmentRange": "屈光度调节范围",
	"FocusScreen":            "对焦屏",
	"Reflector":              "反光镜",
	"DeepDepthPreview":       "景深预视",
	"LcdType":                "LCD类型",
	"LcdSize":                "LCD尺寸",
	"LcdPixels":              "LCD像素",
	"FieldOfView":            "视野率",
	"LcdCharacteristics":     "LCD特性",
	"StorageCard":            "存储卡",
	"DataInterface":          "数据接口",
	"VideoAudioInterface":    "视频/音频接口",
	"WirelessFunction":       "无线功能",
	"Battery":                "电池",
	"BatteryLife":            "续航能力",
	"Colour":                 "颜色",
	"Size":                   "尺寸",
	"Weight":                 "重量",
	"OperatingTemperature":   "工作温度",
	"WorkingHumidity":        "工作湿度",
	"Annex":                  "附件",
}

// LensColumnsMapConst 镜头字段中英映射
var LensColumnsMapConstEnFirst = map[string]string{
	"Name":                   "名称",
	"LensType":               "镜头类型",
	"ReleaseDate":            "发布日期",
	"FocalLength":            "焦距",
	"EqFullFrameFocalLength": "等效35mm画幅焦距",
	"MaximumAperture":        "最大光圈",
	"MinimumAperture":        "最小光圈",
	"ViewingAngle":           "视角（对角）",
	"LensStructure":          "镜头结构",
	"AsphericalLens":         "非球面镜片",
	"LowDispersionLens":      "低/超低色散镜片",
	"ApertureBlade":          "光圈叶片",
	"RecentFocusDistance":    "最近对焦距离",
	"MaximumMagnification":   "最大放大倍率",
	"FilterAperture":         "滤镜口径",
	"LensMount":              "镜头卡口",
	"ImageStabilizer":        "防抖",
	"Colour":                 "颜色",
	"Size":                   "尺寸",
	"Weight":                 "重量",
	"Annex":                  "附件",
	"Remarks":                "备注",
}

func CameraColumnsMapKeys() []string {
	return []string{"id",
		"名称",
		"产品类型",
		"影像传感器",
		"发布日期",
		"停产",
		"影像传感器尺寸",
		"总像素",
		"有效像素",
		"影像处理器",
		"静态图像格式",
		"静态图像尺寸",
		"画幅比",
		"短片格式",
		"短片尺寸",
		"音频格式",
		"除尘系统",
		"镜头结构",
		"非球面镜片",
		"焦距",
		"等效35mm画幅焦距",
		"最大光圈",
		"光学变焦",
		"最近对焦距离",
		"对焦模式",
		"镜头卡口",
		"兼容镜头",
		"对焦系统类型",
		"对焦点",
		"测光系统",
		"测光模式",
		"AF辅助灯",
		"曝光模式",
		"拍摄模式",
		"胶片模拟模式",
		"创意滤镜",
		"快门类型",
		"快门速度",
		"闪光同步速度",
		"曝光补偿",
		"ISO感光度",
		"白平衡",
		"自拍",
		"连拍",
		"取景器",
		"取景范围",
		"取景放大倍率",
		"屈光度调节范围",
		"对焦屏",
		"反光镜",
		"景深预视",
		"LCD类型",
		"LCD尺寸",
		"LCD像素",
		"视野率",
		"LCD特性",
		"存储卡",
		"数据接口",
		"视频",
		"无线功能",
		"电池",
		"续航能力",
		"颜色",
		"尺寸",
		"重量",
		"工作温度",
		"工作湿度",
		"附件"}
}

func LensColumnsMapKeys() []string {
	return []string{"id",
		"名称",
		"镜头类型",
		"发布日期",
		"焦距",
		"等效35mm画幅焦距",
		"最大光圈",
		"最小光圈",
		"视角（对角）",
		"镜头结构",
		"非球面镜片",
		"低",
		"光圈叶片",
		"最近对焦距离",
		"最大放大倍率",
		"滤镜口径",
		"镜头卡口",
		"防抖",
		"颜色",
		"尺寸",
		"重量",
		"附件",
		"备注"}
}

func getKeysFromStringsMap(anyMap map[string]string) (attrs []string) {
	for k := range anyMap {
		attrs = append(attrs, k)
	}
	sort.Strings(attrs)
	attrs = append([]string{"名称"}, attrs...)
	return
}

func getValuesFromStringsMap(anyMap map[string]string) (attrs []string) {
	for _, v := range anyMap {
		attrs = append(attrs, v)
	}
	sort.Strings(attrs)
	attrs = append([]string{"名称"}, attrs...)
	return
}