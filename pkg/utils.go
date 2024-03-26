
package pkg

import (
    "os"
    "os/exec"
	"path/filepath"
    "fmt"
	"strings"
)


func RunCmd(cmdName string, args[] string) (string, error) {
    cmd := exec.Command(cmdName , args...)
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("RunCmd error : %s\n",err);
        return "",err
    } 
    return strings.Trim(string(out)," \r\n\t"),nil
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

    fmt.Printf("%s\n",foundFilePath);

	if err != nil || foundFilePath == "" {
		return false
	}

	return true
}