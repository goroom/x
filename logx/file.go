package logx

import (
	"fmt"
	"os"
)

type File struct {
	File *os.File
	Size Unit
}

func newFile(dir, name string, symlinkName string) (*File, error) {
	f, err := openFile(dir, name)
	if err != nil {
		return nil, err
	}

	if symlinkName != "" {
		_ = os.Remove(dir + "/" + symlinkName)
		err = os.Symlink(name, dir+"/"+symlinkName)
		if err != nil {
			fmt.Println("error,", err)
		}
	}

	return &File{
		File: f,
	}, nil
}

func (f *File) write(data []byte) error {
	_, err := f.File.Write(data)
	if err != nil {
		return err
	}
	f.Size += Unit(len(data))

	return nil
}

func (f *File) Close() {
	if f == nil {
		return
	}

	err := f.File.Close()
	if err != nil {
		fmt.Println("close log file error", err)
	}
}

func DefaultFileSplitBySize(lg *Logger, f *File, format *Entry) (*File, error) {
	if f != nil && f.Size < lg.opt.FileSizeMax {
		return f, nil
	}

	if f == nil {
		f, err := newFile(lg.opt.FileDir, lg.opt.FileName+".log", "")
		if err != nil {
			return nil, err
		}

		return f, nil
	}
	f.Close()

	// delete last file
	filePath := lg.opt.GetFilePath(int(lg.opt.FileCountMax - 1))
	if isFileExit(filePath) {
		err := os.Remove(filePath)
		if err != nil {
			return nil, err
		}
	}

	for i := lg.opt.FileCountMax - 2; i >= 0; i-- {
		filePath = lg.opt.GetFilePath(int(i))
		if isFileExit(filePath) {
			err := os.Rename(filePath, lg.opt.GetFilePath(int(i+1)))
			if err != nil {
				return nil, err
			}
		}
	}
	return newFile(lg.opt.FileDir, lg.opt.FileName+".log", "")
}

func DefaultFileSplitByMinute(lg *Logger, f *File, format *Entry) (*File, error) {
	newFileName := lg.opt.FileName + "_" + format.Time.Format("20060102_1504") + ".log"
	if f != nil && f.File.Name() == lg.opt.FileDir+"/"+newFileName {
		return f, nil
	}

	if f != nil {
		f.Close()
	}

	return newFile(lg.opt.FileDir, newFileName, lg.opt.FileName+".log")
}

func DefaultFileSplitByHour(lg *Logger, f *File, format *Entry) (*File, error) {
	newFileName := lg.opt.FileName + "_" + format.Time.Format("20060102_15") + ".log"
	if f != nil && f.File.Name() == lg.opt.FileDir+"/"+newFileName {
		return f, nil
	}

	if f != nil {
		f.Close()
	}

	return newFile(lg.opt.FileDir, newFileName, lg.opt.FileName+".log")
}

func DefaultFileNoSplit(lg *Logger, f *File, format *Entry) (*File, error) {
	if f == nil {
		f, err := newFile(lg.opt.FileDir, lg.opt.FileName+".log", "")
		if err != nil {
			return nil, err
		}

		return f, nil
	}
	return f, nil
}

func openFile(dir, name string) (*os.File, error) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return os.OpenFile(dir+"/"+name, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
}
