package FileUtil

import (
	"strings"
)

type FileType string

const (
	JPG  FileType = ".jpg"
	PNG  FileType = ".png"
	GIF  FileType = ".gif"
	JPEG FileType = ".jpeg"
	BMP  FileType = ".bmp"
	TIFF FileType = ".tiff"
	PDF  FileType = ".pdf"
	DOC  FileType = ".doc"
	DOCX FileType = ".docx"
	XLS  FileType = ".xls"
	XLSX FileType = ".xlsx"
	PPT  FileType = ".ppt"
	PPTX FileType = ".pptx"
	MP4  FileType = ".mp4"
	MP3  FileType = ".mp3"
	WAV  FileType = ".wav"
	WMA  FileType = ".wma"
	AVI  FileType = ".avi"
	RMVB FileType = ".rmvb"
	FLV  FileType = ".flv"
	MKV  FileType = ".mkv"
	MOV  FileType = ".mov"
	MPG  FileType = ".mpg"
	MPEG FileType = ".mpeg"
	VOB  FileType = ".vob"
	SWF  FileType = ".swf"
	RAR  FileType = ".rar"
	ZIP  FileType = ".zip"
	TAR  FileType = ".tar"
	GZ   FileType = ".gz"
	_7Z  FileType = ".7z"
	BZ2  FileType = ".bz2"
	ISO  FileType = ".iso"
	CSS  FileType = ".css"
	HTML FileType = ".html"
	JS   FileType = ".js"
	PHP  FileType = ".php"
	JSP  FileType = ".jsp"
	C    FileType = ".c"
	CPP  FileType = ".cpp"
	CXX  FileType = ".cxx"
	H    FileType = ".h"
	HPP  FileType = ".hpp"
	JAVA FileType = ".java"
	PY   FileType = ".py"
	PYC  FileType = ".pyc"
	PYO  FileType = ".pyo"
	PYW  FileType = ".pyw"
	PYX  FileType = ".pyx"
	RB   FileType = ".rb"
	RBW  FileType = ".rbw"
	RBO  FileType = ".rbo"
	RBX  FileType = ".rbx"
	SH   FileType = ".sh"
	TXT  FileType = ".txt"
	XML  FileType = ".xml"
	YAML FileType = ".yaml"
	YML  FileType = ".yml"
	CSV  FileType = ".csv"
)

var TypeNames = map[string]FileType{
	".zip":  ZIP,
	".tar":  TAR,
	".gz":   GZ,
	".7z":   _7Z,
	".bz2":  BZ2,
	".iso":  ISO,
	".css":  CSS,
	".html": HTML,
	".js":   JS,
	".php":  PHP,
	".jsp":  JSP,
	".c":    C,
	".cpp":  CPP,
	".cxx":  CXX,
	".h":    H,
	".hpp":  HPP,
	".java": JAVA,
	".py":   PY,
	".pyc":  PYC,
	".pyo":  PYO,
	".pyw":  PYW,
	".pyx":  PYX,
	".rb":   RB,
	".rbw":  RBW,
	".rbo":  RBO,
	".rbx":  RBX,
	".sh":   SH,
	".txt":  TXT,
	".xml":  XML,
	".yaml": YAML,
	".yml":  YML,
	".csv":  CSV,
	".jpg":  JPG,
	".png":  PNG,
	".gif":  GIF,
	".jpeg": JPEG,
	".bmp":  BMP,
	".tiff": TIFF,
	".pdf":  PDF,
	".doc":  DOC,
	".docx": DOCX,
	".xls":  XLS,
	".xlsx": XLSX,
	".ppt":  PPT,
	".pptx": PPTX,
	".mp4":  MP4,
	".mp3":  MP3,
	".wav":  WAV,
	".wma":  WMA,
	".avi":  AVI,
	".rmvb": RMVB,
	".flv":  FLV,
	".mkv":  MKV,
	".mov":  MOV,
	".mpg":  MPG,
	".mpeg": MPEG,
	".vob":  VOB,
	".swf":  SWF,
	".rar":  RAR,
}

// FileTypeFromExt returns the FileType from the given file extension.
func fileTypeFromExt(ext string) FileType {
	return TypeNames[strings.ToLower(ext)]
}
