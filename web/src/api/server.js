import service from '@/utils/request'


// 获取服务器信息
export const getServerList = (data) => {
  return service({
      url: "/server/getServerList",
      method: 'post',
      data: data
  })
}

//新增服务器
export const createServer = (data) => {
  return service({
      url: "/server/createServer",
      method: 'post',
      data: data
  })
}

//修改服务器
export const updateServer = (data) => {
  return service({
      url: "/server/updateServer",
      method: 'post',
      data: data
  })
}

//  获取服务器通过id
export const getServerById = (data) => {
  return service({
      url: "/server/getServerById",
      method: 'post',
      data: data
  })
}

//删除服务器
export const deleteServer = (data) => {
  return service({
      url: "/server/deleteServer",
      method: 'post',
      data: data
  })
}
