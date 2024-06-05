package utils

import (
	"archive/tar"
	"compress/bzip2"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// Check whether file or directory exists or not
// Args:
//
//	file_path: path of file or directory
//
// Returns:
//
//	true | false, true stands for exists, false stands for not exists
func FileExists(file_path string) bool {
	_, err := os.Stat(file_path)
	return err == nil || os.IsExist(err)
}

// Get file size
// Args:
//
//	file_path: path of file
//
// Returns:
//
//	size: file size in Bytes
//	err: nil stands for success, other stands for not exists
func GetFileSize(file_path string) (int64, error) {
	if !FileExists(file_path) {
		return 0, errors.New(fmt.Sprintf("File(%s) not exists.", file_path))
	}
	file_info, err := os.Stat(file_path)
	if err != nil {
		return 0, err
	}
	return file_info.Size(), nil
}

// check whether file is symbol link
// Args:
//
//	filepath: file path
//
// Returns:
//
//	is_symbol_link: true or false
//	err           : nil stands for success, otherwise stands for fail
func IsSymbolLink(file_path string) (is_symbol_link bool, err error) {
	file_info, err := os.Lstat(file_path)
	if err != nil {
		err_msg := fmt.Sprintf("Fail to get file info of %s: %s", file_path, err.Error())
		return false, errors.New(err_msg)
	}
	return file_info.Mode()&os.ModeSymlink != 0, nil
}

// 判断一个本地路径是否是目录
// Args:
//
//	path: absolute path
//
// Return:
//
//	true : path is a directory
//	false: path isn't exist or path isn't a directory
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// traverse and create pakcaged file according to tar_reader
// Args:
//
//	tar_reader: pointer of tar.Reader
//	dst_dir   : destination directory
//
// Returns:
//
//	nil stands for success, otherwise stands for fail
func traversePackagedFile(tar_reader *tar.Reader, dst_dir string) error {
	for header, err := tar_reader.Next(); err != io.EOF; header, err = tar_reader.Next() {
		if err != nil {
			return err
		}
		// 读取文件信息
		file_path := filepath.Join(dst_dir, header.Name)
		file_info := header.FileInfo()
		file_mode := file_info.Mode()

		if file_mode.IsDir() {
			if !FileExists(file_path) {
				err := os.MkdirAll(file_path, file_mode)
				if err != nil {
					return err
				}
			}
			continue
		} else if file_mode.IsRegular() {
			parent_dir := filepath.Dir(file_path)
			if !FileExists(parent_dir) {
				err := os.MkdirAll(parent_dir, 0775)
				if err != nil {
					return err
				}
			}
			file_writer, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file_mode)
			if err != nil {
				return err
			}
			_, err = io.Copy(file_writer, tar_reader)
			if err != nil {
				file_writer.Close()
				return err
			}
			file_writer.Close()
			continue
		} else if file_mode&os.ModeSymlink != 0 {
			os.Symlink(header.Linkname, file_path)
			os.Chmod(file_path, file_mode)
			continue
		}
	}
	return nil
}

// 把指定文件列表打包成tar包，文件列表可以包含文件或文件夹
// Args:
//
//	current_dir  : 进程的当前工作目录
//	work_dir     : 生成tar文件的工作目录，inputs 文件列表都是相对于work_dir
//	inputs       : 相对于work_dir 的文件列表，支持普通文件和目录
//	tar_file_path: 要生成的tar 文件路径
//
// Return:
//
//	err: nil stands for success, otherwise stands for fail
func Tar(current_dir string, work_dir string, inputs []string, tar_file_path string) error {
	cur_stat, err := os.Stat(current_dir)
	if err != nil {
		return err
	}
	if !cur_stat.IsDir() {
		err_msg := fmt.Sprintf("%s is not a directory", current_dir)
		return errors.New(err_msg)
	}

	work_stat, err := os.Stat(work_dir)
	if err != nil {
		return err
	}
	if !work_stat.IsDir() {
		err_msg := fmt.Sprintf("%s is not a directory", work_dir)
		return errors.New(err_msg)
	}

	// 首先跳转工作目录，打包完成后再切换到当前目录
	os.Chdir(work_dir)
	defer os.Chdir(current_dir)

	file_writer, err := os.Create(tar_file_path)
	if err != nil {
		return err
	}
	defer file_writer.Close()
	tar_writer := tar.NewWriter(file_writer)
	defer tar_writer.Close()
	walk := func(path string, info os.FileInfo, err error) error {
		file_info, err := os.Lstat(path)
		if err != nil {
			return err
		}
		// file 如果是软链，获取软链的引用文件信息
		link_to := ""
		if file_info.Mode()&os.ModeSymlink != 0 {
			link_to, _ = os.Readlink(path)
		}
		// 把file 信息记录到tar文件头中
		file_info_header, err := tar.FileInfoHeader(file_info, link_to)
		// 为了解压后文件保持打包前的相对目录和位置，需要把单独的文件名修改为更具体的文件路径。比如  test.cpp -> src/auth
		file_info_header.Name = path
		if tar_writer.WriteHeader(file_info_header); err != nil {
			return err
		}
		// 文件信息已写到tar文件头中，如果是空目录或软链，直接返回
		if file_info.IsDir() || file_info.Mode()&os.ModeSymlink != 0 {
			return nil
		}
		// 如果是普通文件，拷贝文件内容到tar 文件中
		file_reader, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file_reader.Close()
		if _, err = io.Copy(tar_writer, file_reader); err != nil {
			return err
		}
		return err
	}
	for _, file_or_dir := range inputs {
		filepath.Walk(file_or_dir, walk)
	}
	return nil
}

// untar xxx.tar
// Args:
//
//	tarfile: tar file path
//	dst_dir: destination directory
//
// Returns:
//
//	nil stands for success, otherwise stands for fail
func UnTar(tar_file_path string, dst_dir string) (err error) {
	if !FileExists(tar_file_path) {
		err_msg := fmt.Sprintf("%s does not exist", tar_file_path)
		return errors.New(err_msg)
	}

	file_info, err := os.Stat(dst_dir)
	if err == nil || os.IsExist(err) {
		// dst_dir存在，但是个文件
		if !file_info.IsDir() {
			err_msg := fmt.Sprintf("%s has been exist and is not a directory", dst_dir)
			return errors.New(err_msg)
		}
	} else if os.IsNotExist(err) {
		// dst_dir不存在，创建该目录
		err = os.MkdirAll(dst_dir, 0755)
		if err != nil {
			err_msg := fmt.Sprintf("Fail to create directory %s: %s", dst_dir, err.Error())
			return errors.New(err_msg)
		}
	} else {
		// 其他类型错误
		return err
	}

	file_reader, err := os.Open(tar_file_path)
	if err != nil {
		return err
	}
	defer file_reader.Close()

	tar_reader := tar.NewReader(file_reader)
	return traversePackagedFile(tar_reader, dst_dir)
}

// decompress xxx.tar.gz
// Args:
//
//	tarfile: tar file path
//	dst_dir: destination directory
//
// Returns:
//
//	nil stands for success, otherwise stands for fail
func DeCompressGzip(tar_file_path string, dst_dir string) (err error) {
	if !FileExists(tar_file_path) {
		err_msg := fmt.Sprintf("%s does not exist", tar_file_path)
		return errors.New(err_msg)
	}

	file_info, err := os.Stat(dst_dir)
	if err == nil || os.IsExist(err) {
		// dst_dir存在，但是个文件
		if !file_info.IsDir() {
			err_msg := fmt.Sprintf("%s has been exist and is not a directory", dst_dir)
			return errors.New(err_msg)
		}
	} else if os.IsNotExist(err) {
		// dst_dir不存在，创建该目录
		err = os.MkdirAll(dst_dir, 0755)
		if err != nil {
			err_msg := fmt.Sprintf("Fail to create directory %s: %s", dst_dir, err.Error())
			return errors.New(err_msg)
		}
	} else {
		// 其他类型错误
		return err
	}

	file_reader, err := os.Open(tar_file_path)
	if err != nil {
		return err
	}
	defer file_reader.Close()

	gzip_reader, err := gzip.NewReader(file_reader)
	if err != nil {
		return err
	}

	tar_reader := tar.NewReader(gzip_reader)
	return traversePackagedFile(tar_reader, dst_dir)
}

// decompress xxx.tar.bz2
// Args:
//
//	tarfile: tar file path
//	dst_dir: destination directory
//
// Returns:
//
//	nil stands for success, otherwise stands for fail
func DeCompressBzip2(tar_file_path string, dst_dir string) (err error) {
	if !FileExists(tar_file_path) {
		err_msg := fmt.Sprintf("%s does not exist", tar_file_path)
		return errors.New(err_msg)
	}

	file_info, err := os.Stat(dst_dir)
	if err == nil || os.IsExist(err) {
		// dst_dir存在，但是个文件
		if !file_info.IsDir() {
			err_msg := fmt.Sprintf("%s has been exist and is not a directory", dst_dir)
			return errors.New(err_msg)
		}
	} else if os.IsNotExist(err) {
		// dst_dir不存在，创建该目录
		err = os.MkdirAll(dst_dir, 0755)
		if err != nil {
			err_msg := fmt.Sprintf("Fail to create directory %s: %s", dst_dir, err.Error())
			return errors.New(err_msg)
		}
	} else {
		// 其他类型错误
		return err
	}

	file_reader, err := os.Open(tar_file_path)
	if err != nil {
		return err
	}
	defer file_reader.Close()

	bzip2_reader := bzip2.NewReader(file_reader)
	tar_reader := tar.NewReader(bzip2_reader)
	return traversePackagedFile(tar_reader, dst_dir)
}

// Get file's modify time
// Args:
//
//	file_path: file path
//
// Returns:
//
//	modify_time: file's modify time
//	err        : nil stands for success, otherwise stands for fail
func GetFileModifyTime(file_path string) (modify_time time.Time, err error) {
	file_info, err := os.Stat(file_path)
	if err != nil {
		return modify_time, err
	}
	modify_time = file_info.ModTime()
	return modify_time, nil
}

// Get all files in specified directory (contains files in sub directory)
// Args:
//
//	dir_path: path of directory
//
// Returns:
//
//	file_path_list: list of file path
//	err           : nil stands for success, otherwise stands for fail
func GetFilesInDir(dir_path string) (file_path_list []string, err error) {
	if !FileExists(dir_path) {
		err_msg := fmt.Sprintf("%s doesn't exists", dir_path)
		return file_path_list, errors.New(err_msg)
	}

	err = filepath.Walk(dir_path, func(file_path string, file_info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if file_info.IsDir() {
			return nil
		}

		abs_file_path, _ := filepath.Abs(file_path)
		file_path_list = append(file_path_list, abs_file_path)
		return nil
	})
	return file_path_list, err
}

// get Info Map
func GetInfoMap() map[string]any {
	var infoMap map[string]any

	infoMap["name"] = "Gin Framework"
	infoMap["version"] = "1.0.0"
	infoMap["author"] = "zhangsan"

	return infoMap
}
