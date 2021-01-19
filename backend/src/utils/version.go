/**
 * @Author xieed
 * @Description 版本号工具类
 * @Date 2020/9/24 19:33
 **/

package utils

import "strings"

type Version struct {
	versions []int
}

var defaultVersionSeparator = "."

// 版本号自增
func IncrementVersion(versionStr string) (res string) {
	versionArray := ToIntArrayBySeparator(versionStr, defaultVersionSeparator)
	lastIndex := len(versionArray) - 1
	versionArray[lastIndex] = versionArray[lastIndex] + 1
	return IntArray2String(versionArray, defaultVersionSeparator)
}

// 分支名称格式化
func BranchNameFormat(branchNameFmt string, repoName string, lang string, branchVersion string) (res string) {
	res = branchNameFmt
	res = strings.ReplaceAll(res, "${repo_name}", repoName)
	res = strings.ReplaceAll(res, "${lang}", lang)
	res = strings.ReplaceAll(res, "${branch_version}", branchVersion)
	return res
}
