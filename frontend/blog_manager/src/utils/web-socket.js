/*
 * Copyright 2018-2020,上海哈里奥网络科技有限公司
 * All Right Reserved.
 */

let webSocket = null

// 日志相关
// 初始化
export function initWebSocket() {
  const wsUri = config.wsUri
  webSocket = new WebSocket(wsUri)
  webSocket.onmessage = webSocketOnMessage
  webSocket.onopen = webSocketOnOpen
  webSocket.onerror = webSocketOnError
  webSocket.onclose = webSocketClose

  console.log('初始化webSocket=>', webSocket)
  return webSocket
}

// 数据发送
export function webSocketSend(data) {
  // console.log('webSocket=>', webSocket)
  webSocket.send(JSON.stringify(data))
}

// 关闭websocket
export function close() {
  webSocket.close()
  webSocket.onclose = webSocketClose
}

// 连接建立之后执行send方法发送数据
function webSocketOnOpen(e) {
  console.log('webSocketOnOpen', e)
}

// 连接建立失败重连
function webSocketOnError(e) {
  console.log('webSocketOnError', e)
  initWebSocket()
}

// 关闭
function webSocketClose(e) {
  console.log('断开连接', e)
}

// 数据接收
function webSocketOnMessage(e) {
  const data = JSON.parse(e.data)
  console.log('接受的数据', data)
}
