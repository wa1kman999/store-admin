import service from '@/utils/request'


// 获取硬盘信息
export const getDiskList = (data) => {
  return service({
      url: "/disk/getDiskList",
      method: 'post',
      data: data
  })
}

//新增硬盘
export const createDisk = (data) => {
  return service({
      url: "/disk/createDisk",
      method: 'post',
      data: data
  })
}

//修改硬盘
export const updateDisk = (data) => {
  return service({
      url: "/disk/updateDisk",
      method: 'post',
      data: data
  })
}

//  获取盘通过id
export const getDiskById = (data) => {
  return service({
      url: "/disk/getDiskById",
      method: 'post',
      data: data
  })
}

//删除硬盘
export const deleteDisk = (data) => {
  return service({
      url: "/disk/deleteDisk",
      method: 'post',
      data: data
  })
}
