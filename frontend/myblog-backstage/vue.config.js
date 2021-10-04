// const path = require('path');

module.exports = {
  css: {
    loaderOptions: {
      sass: {
        additionalData: `@import "./src/style/theme.scss";`,
        // sass 版本 9 中使用 additionalData 版本 8 中使用 prependData
      },
    },
  },
  devServer: {
      port:8083,
      proxy: {
        "/api": {
          //如果要代理 websockets，配置这个参数
          ws: true, 
          // logLevel:"debug",
          target:"http://localhost:8001",
          // target:"http://baidu.com",
          // target: "http://mock",
          changeOrigin: true,
          pathRewrite: {
            "^/api": ""
          }
        }
      }
  }
}