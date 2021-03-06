package pp

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/l3lackShark/gosumemory/memory"
	"github.com/tcolgate/mp3"
)

func hashFileMD5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil

}
func calculateMP3Time() (int32, error) {
	t := 0.0
	r, err := os.Open(memory.MenuData.Bm.Path.FullMP3Path)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	d := mp3.NewDecoder(r)
	var f mp3.Frame
	skipped := 0

	for {

		if err := d.Decode(&f, &skipped); err != nil {
			if err == io.EOF {
				break
			}
		}

		t = t + f.Duration().Seconds()
	}

	return int32(t * 1000), nil
}
