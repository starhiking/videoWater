package deal

import (
	"fmt"
	"myTool/ffmpeg"
	"myTool/file"
	"myTool/sys"
	"os"
	"strconv"
	"strings"
)

func GetFCmd(system int) string {

	if system == 0 {
		info := sys.GetSysInfo()
		system = info.PlatForm
	}
	if system == sys.MacOS {
		return "./source/mac/tool"
	} else if system == sys.Win64 {
		return "./source/win/64/tool.exe"
	} else if system == sys.Win32 {
		return "./source/win/32/tool.exe"
	}

	fmt.Println("系统类型无法识别，请在配置中指定：1:mac 3: win32  4:win64")
	return ""

}

func StringToValue(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		return v
	}

	if strings.HasPrefix(str, "+") {
		str = strings.TrimPrefix(str, "+")
		v, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return v
	}

	if strings.HasPrefix(str, "-") {
		str = strings.TrimPrefix(str, "-")
		v, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return v
	}
	return 0
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func createResultDir(videoDir string) string {
	result := videoDir + "/result"
	_ = os.RemoveAll(result)
	_ = os.MkdirAll(result, os.ModePerm)

	return result
}

func GetAllBgm(dir string) []string  {
	files , err := file.GetAllFiles(dir)
	if err != nil {
		return nil
	}
	var res []string

	for _, f := range  files {
		v := ffmpeg.IsMusic(f)
		if v {
			res = append(res, f)
		}

	}
	return res
}