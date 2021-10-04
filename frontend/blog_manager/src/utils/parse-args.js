/*
 * Copyright 2018-2020,上海哈里奥网络科技有限公司
 * All Right Reserved.
 */
import { isEmpty } from '@/utils/validate'

const PARAM = '_param_'

// 4.
export function getIndex(argsName) { // 0_param_1
  if (isEmpty(argsName)) {
    return []
  }
  return [Number(argsName.substring(0, argsName.indexOf('_'))), Number(argsName.substring(argsName.lastIndexOf('_') + 1, argsName.length))]
}

// 3.生成唯一name
export function generateArgsNames(finalArgsArray) {
  const argsNames = []
  finalArgsArray.forEach((arr, i) => {
    const objArr = {}
    arr.forEach((str, j) => {
      const type = itemType(str)
      if (type === 'select' || type === 'input') {
        const key = i + PARAM + j
        objArr[key] = ''
      }
    })
    argsNames.push(objArr)
  })
  return argsNames
}

// item 类型
export function itemType(item) {
  if (isEmpty(item)) {
    return ''
  }

  const leftIndex = item.indexOf('${')
  const rightIndex = item.lastIndexOf('}')

  if (leftIndex !== -1 && rightIndex !== -1 && leftIndex < rightIndex) {
    return 'select'
  } else if (leftIndex === -1 && rightIndex === -1 && item !== ',' && item !== '.') {
    return 'input'
  } else {
    return 'string'
  }
}

// 2.将数组中的每个数组内容分别拼接为字符串
export function handleArr2Str(finalArgsArray) {
  // array = [['${', '${repo_select}', '.', '${env_select}', '}', '/lib', '${', '${repo_select}', '}', '/lib', ',', 'target_class_name', '${xxx}'], ['${repo_select}', ',', 'target_class_name']]

  const resultArray = []
  finalArgsArray.forEach(split$Arr => {
    let str = ''
    split$Arr.forEach(item => {
      str = str + item
    })
    resultArray.push(str)
  })

  // console.log('finalResult===>', finalResult)
  return resultArray
}

// 1.将 ['','',...]  =>  [['','',''...],['','',''...],...]
export function parseArgsArray(argsArray) {
  // argsArray = ['${${repo_select}.${env_select}}/lib${${repo_select}}/lib,target_class_name${xxx}', '${repo_select},target_class_name']

  const finalArgsArray = []
  argsArray.forEach((argsStr) => {
    finalArgsArray.push(parseArgsStr(argsStr))
  })

  return finalArgsArray
}

export function parseArgsStr(argsStr) {
  // str = '${repo_select},target_class_name'

  let resultArr = []
  const strArr = argsStr.split(',')
  strArr.forEach((str, index) => {
    const afterParseArr = parseArgsComma(str)
    if (index !== strArr.length - 1) {
      afterParseArr.push(',')
    }

    resultArr = resultArr.concat(afterParseArr)
  })

  return resultArr
}

export function parseArgsComma(commaStr) {
  // commaStr = '${${repo_select}${env_select}}/lib${${repo_select}}/lib'
  if (isEmpty(commaStr)) {
    return []
  }

  // 1.给$Arr 赋值
  const $Arr = []
  let i = 0
  while (i !== -1) {
    let index = -1
    if (i === 0) {
      index = commaStr.indexOf('${')
    } else {
      index = commaStr.indexOf('${', i)
    }

    if (index !== -1) {
      $Arr.push(index)
    }

    if (index === -1) {
      i = index
    } else {
      i = index + 1
    }
  }

  // 2.遍历$Arr，切割${}，生成数组
  const split$Arr = []
  let lastRIndex = -1
  for (let j = 0; j < $Arr.length; j++) {
    const index = $Arr[j] // ${ 在commaStr中的索引
    let nextIndex = -1
    if (j < $Arr.length - 1) {
      nextIndex = $Arr[j + 1]
    } else {
      nextIndex = commaStr.length
    }

    let RIndex = -1 // } 在commaStr中的索引
    RIndex = commaStr.indexOf('}', index) // [15,28,29,47]  commaStr = '${${repo_select}${env_select}}/lib${${repo_select}}/lib'

    if (nextIndex < RIndex) { // 情况： ${  ${repo_select}
      continue
    } else { // 情况：  ${repo_select} ${
      if (lastRIndex < index - 1) {
        split$Arr.push(commaStr.substring(lastRIndex + 1, index)) // input
      }
      split$Arr.push(commaStr.substring(index, RIndex + 1)) // select
      lastRIndex = RIndex
    }
  }

  // 末尾字符： }/lib
  if (lastRIndex < commaStr.length - 1) {
    split$Arr.push(commaStr.substring(lastRIndex + 1, commaStr.length))
  }

  return handleSplit$Arr(split$Arr)
}

// 处理有 ${  } 部分的字符串
export function handleSplit$Arr(split$Arr) {
  // split$Arr = ['${', '${repo_select}', '.', '${env_select}', '}/lib${', '${repo_select}', '}/lib']

  for (let i = 0; i < split$Arr.length; i++) {
    const item = split$Arr[i]
    if (isEmpty(item)) {
      continue
    }
    const leftIndex = item.indexOf('${')
    const rightIndex = item.lastIndexOf('}')
    if (leftIndex !== -1 && rightIndex !== -1 && leftIndex < rightIndex) {
      continue
    } else {
      if (leftIndex !== -1 && rightIndex === -1 && item.length > 2) { // 有'${' 无'}'
        split$Arr.splice(i, 1, item.substring(0, leftIndex), '${')
      } else if (rightIndex !== -1 && leftIndex === -1 && item.length > 1) { // 有'}' 无'${'
        split$Arr.splice(i, 1, '}', item.substring(rightIndex + 1, item.length))
      } else if (leftIndex !== -1 && rightIndex !== -1 && item.length > 3) { // 有'${' 有'}'
        split$Arr.splice(i, 1, '}', item.substring(rightIndex + 1, leftIndex), '${')
      }
    }
  }

  return split$Arr
}
