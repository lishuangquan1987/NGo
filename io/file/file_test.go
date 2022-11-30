package file

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

// ok
func TestAppendAllText(t *testing.T) {
	dir, _ := os.Getwd()
	path := "./test.txt"
	path = filepath.Join(dir, path)
	t.Log(path)
	err := AppendAllText(path, "test")
	if err != nil {
		t.Error(err)
	}
	t.Log("test finish")
}
func TestAppendAllLines(t *testing.T) {
	dir, _ := os.Getwd()
	path := "./test.txt"
	path = filepath.Join(dir, path)
	t.Log(path)
	contents := []string{"1", "2", "3", "4", "5"}
	err := AppendAllLines(path, contents)
	if err != nil {
		t.Error(err)
	}
	t.Log("test finish")
}

// ok
func TestWriteAllText(t *testing.T) {
	dir, _ := os.Getwd()
	path := "./test.txt"
	path = filepath.Join(dir, path)
	t.Log(path)
	err := WriteAllText(path, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		t.Error(err)
	}
	t.Log("test finish")
}
func TestWriteAllLines(t *testing.T) {
	dir, _ := os.Getwd()
	path := "./test.txt"
	path = filepath.Join(dir, path)
	t.Log(path)
	contents := []string{"hello", "world", "I", "am", "go"}
	err := WriteAllLines(path, contents)
	if err != nil {
		t.Error(err)
	}
	t.Log("test finish")
}

func TestCopy(t *testing.T) {
	dir, _ := os.Getwd()
	srcPath := "./test.txt"
	srcPath = filepath.Join(dir, srcPath)
	dstPath := "./test1.txt"
	dstPath = filepath.Join(dir, dstPath)
	err := Copy(srcPath, dstPath)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Log("test finish")
}
func TestCopyError(t *testing.T) {
	dir, _ := os.Getwd()
	srcPath := "./test111.txt" //not exist file
	srcPath = filepath.Join(dir, srcPath)
	dstPath := "./test1.txt"
	dstPath = filepath.Join(dir, dstPath)
	err := Copy(srcPath, dstPath)
	if err == nil {
		t.Error("copy a file that not exist success")
	} else {
		t.Logf("test ok with error:%v", err)
	}

}

func TestCopyWithOverwrite(t *testing.T) {
	dir, _ := os.Getwd()
	srcPath := "./test.txt"
	srcPath = filepath.Join(dir, srcPath)
	dstPath := "./test1.txt"
	dstPath = filepath.Join(dir, dstPath)
	err := CopyWithOverwrite(srcPath, dstPath)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Log("test finish")
}
func TestMove(t *testing.T) {
	dir, _ := os.Getwd()
	srcPath := "./test.txt"
	srcPath = filepath.Join(dir, srcPath)
	dstPath := "./test_move.txt"
	dstPath = filepath.Join(dir, dstPath)
	err := Move(srcPath, dstPath)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Log("test finish")
}
func TestMoveError(t *testing.T) {
	dir, _ := os.Getwd()
	srcPath := "./test1.txt"
	srcPath = filepath.Join(dir, srcPath)
	dstPath := "./test_move.txt"
	dstPath = filepath.Join(dir, dstPath)
	err := Move(srcPath, dstPath)
	if err == nil {
		t.Errorf("%v", err)
		return
	}
	t.Log("test finish")
}
func TestMoveWithOverwrite(t *testing.T) {
	dir, _ := os.Getwd()
	srcPath := "./test.txt"
	srcPath = filepath.Join(dir, srcPath)
	dstPath := "./test_move.txt"
	dstPath = filepath.Join(dir, dstPath)
	err := Move(srcPath, dstPath)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Log("test finish")
}
