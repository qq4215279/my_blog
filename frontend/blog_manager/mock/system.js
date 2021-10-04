// import Mock from 'mockjs'

export default [
  {
    url: '/dbVersion.action',
    type: 'get',
    response: config => {
      const mockList = [{ dbVersion: '28' }]

      return {
        code: 20000,
        data: {
          total: mockList.length,
          items: mockList
        }
      }
    }
  },
  {
    url: '/getSystemInfo.action',
    type: 'get',
    response: config => {
      const mockList = { 
        sysVersion: '1.1.1',
        dbVersion: '1.1.1',
        updateInfo: '1.1.1\n-------\n1. 版本新增\n2. 版本新增' 
      }

      return {
        code: 20000,
        data: {
          items: mockList
        }
      }
    }
  }
]
