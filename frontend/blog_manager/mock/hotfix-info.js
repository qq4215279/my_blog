import Mock from 'mockjs'

const List = []
const count = 100

for (let i = 0; i < count; i++) {
  List.push(Mock.mock({
    id: '@increment',
    gameId: 'zjzr',
    updateGroup: 'android',
    oldVersion: '1.0.0.0',
    newVersion: '1.1.1.1',
    packType: '@integer(1, 2)',
    updateFile: '1.0.0.0/1.1.1.1',
    fileSize: '123456',
    md5: 'sdfw2x1200sdx'
  }))
}

const KeyList = []
for (let i = 0; i < count; i++) {
  KeyList.push(Mock.mock({
    id: '@increment',
    gameId: 'zjzr',
    updateGroup: 'android',
    version: '1.0.0.' + '@integer(1, 100)',
    packType: '@integer(1, 2)',
    execFile: '',
    isKeyVersion: '@integer(0, 1)',
    isBaseVersion: '@integer(0, 1)',
    updateTips: '更新提示',
    createTime: new Date()
  }))
}

const CDNList = []
for (let i = 0; i < count; i++) {
  CDNList.push(Mock.mock({
    id: '@increment',
    gameId: 'zjzr',
    updateGroup: 'android',
    url: 'http://baidu.com',
    priority: '@integer(1, 2)'
  }))
}

export default [
  {
    url: '/hotfixInfo.action',
    type: 'get',
    response: config => {
      const { importance, type, title, page = 1, limit = 20, sort } = config.query

      let mockList = List.filter(item => {
        return true
      })

      if (sort === '-id') {
        mockList = mockList.reverse()
      }

      const pageList = mockList.filter((item, index) => index < limit * page && index >= limit * (page - 1))

      return {
        code: 20000,
        data: {
          total: mockList.length,
          items: pageList
        }
      }
    }
  },
  {
    url: '/hotfixKeyInfo.action',
    type: 'get',
    response: config => {
      const { importance, type, title, page = 1, limit = 20, sort } = config.query

      let mockList = KeyList.filter(item => {
        return true
      })

      if (sort === '-id') {
        mockList = mockList.reverse()
      }

      const pageList = mockList.filter((item, index) => index < limit * page && index >= limit * (page - 1))

      return {
        code: 20000,
        data: {
          total: mockList.length,
          items: pageList
        }
      }
    }
  },
  {
    url: '/hotfixCdnList.action',
    type: 'get',
    response: config => {
      const { importance, type, title, page = 1, limit = 20, sort } = config.query

      let mockList = CDNList.filter(item => {
        return true
      })

      if (sort === '-id') {
        mockList = mockList.reverse()
      }

      const pageList = mockList.filter((item, index) => index < limit * page && index >= limit * (page - 1))

      return {
        code: 20000,
        data: {
          total: mockList.length,
          items: pageList
        }
      }
    }
  },
  {
    url: '/hotfixUpdateGroupList.action',
    type: 'get',
    response: config => {
      const { importance, type, title, page = 1, limit = 20, sort } = config.query

      let mockList = [
        { id: 1, updateGroup: 'android' },
        { id: 2, updateGroup: 'appstore' },
        { id: 3, updateGroup: 'jailbreak' }
      ]

      if (sort === '-id') {
        mockList = mockList.reverse()
      }

      const pageList = mockList.filter((item, index) => index < limit * page && index >= limit * (page - 1))

      return {
        code: 20000,
        data: {
          total: mockList.length,
          items: pageList
        }
      }
    }
  },
  {
    url: '/hotfixLatestVersionList.action',
    type: 'get',
    response: config => {
      const { importance, type, title, page = 1, limit = 20, sort } = config.query

      let mockList = [
        { gameId: 'zjzr', updateGroup: 1, packType: 1, latestVersion: '1.0.0.0' },
        { gameId: 'zjzr', updateGroup: 1, packType: 1, latestVersion: '1.0.0.0' },
        { gameId: 'zjzr', updateGroup: 1, packType: 1, latestVersion: '1.0.0.0' }
      ]

      if (sort === '-id') {
        mockList = mockList.reverse()
      }

      const pageList = mockList.filter((item, index) => index < limit * page && index >= limit * (page - 1))

      return {
        code: 20000,
        data: {
          total: mockList.length,
          items: pageList
        }
      }
    }
  },
  {
    url: '/hotfixForceUpdateList.action',
    type: 'get',
    response: config => {
      const { importance, type, title, page = 1, limit = 20, sort } = config.query

      let mockList = [
        { gameId: 'zjzr', yx: 'hario', yxSource: 'hario', url: 'https://www.baidu.com', forceTips: '必须强更了啊', forceUpdateVersion: '1.0.0.0', forceSuggestVersion: '1.0.0.0' },
        { gameId: 'zjzr', yx: 'hario', yxSource: 'hario', url: 'https://www.baidu.com', forceTips: '必须强更了啊', forceUpdateVersion: '1.0.0.0', forceSuggestVersion: '1.0.0.0' },
        { gameId: 'zjzr', yx: 'hario', yxSource: 'hario', url: 'https://www.baidu.com', forceTips: '必须强更了啊', forceUpdateVersion: '1.0.0.0', forceSuggestVersion: '1.0.0.0' }
      ]

      if (sort === '-id') {
        mockList = mockList.reverse()
      }

      const pageList = mockList.filter((item, index) => index < limit * page && index >= limit * (page - 1))

      return {
        code: 20000,
        data: {
          total: mockList.length,
          items: pageList
        }
      }
    }
  }
]
