package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ctrsploit/sploit-spec/pkg/env/vt"
)

func AppendDev( info *vt.Basic, Type vt.DeviceType, name string ) {
	for _, it := range info.DevList {
		if it.Type == Type && it.Name == name {
			return
		}
	}
    //fmt.Printf("add device Type : %d, Name : %s \n", Type, name)
	info.DevList = append(info.DevList, vt.DeviceInfo { Type: Type, Name: name} )
}

func RemoveDev( info *vt.Basic, Type vt.DeviceType, name string ) {
    
	for i, it := range info.DevList {
		if it.Type == Type && it.Name == name {
			info.DevList = append(info.DevList[:i], info.DevList[i+1:]... )
            return
		}
	}
}

func ContainsDev( info vt.Basic , name string ) ( found bool ) {
    found = true
	for _, it := range info.DevList {
		if it.Name == name {
            return
		}
	}
    return false
}

func Trim( str string ) string {
    return strings.Trim(str, " \r\n\t")
}

func RunCmd(cmdName string, args[] string) (string, error) {
    //fmt.Printf("launch %s %s\n",cmdName, strings.Join(args, " "))
    cmd := exec.Command(cmdName , args...)
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("RunCmd error : %s\n",err);
        return "",err
    } 
    return Trim(string(out)),nil
}


func FindFile(root, targetFile string) (bool) {
	var foundFilePath string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == targetFile {
			foundFilePath = path
			return filepath.SkipDir // Skip further traversal since we found the file
		}
		return nil
	})

    //fmt.Printf("%s\n",foundFilePath);

	if err != nil || foundFilePath == "" {
		return false
	}

	return true
}

func EnumFile( src  string  ) ( [] string ) {
    matches, err := filepath.Glob(src)
    if err != nil {
        fmt.Println("Error:", err)
        return []string{}
    }
    return matches
} 