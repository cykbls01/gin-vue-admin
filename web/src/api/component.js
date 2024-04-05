import service from '@/utils/request'

export const getComponentList = (data) => {
  return service({
    url: '/component/getComponentList',
    method: 'post',
    data
  })
}

export const createComponent = (data) => {
  return service({
    url: '/component/createComponent',
    method: 'post',
    data
  })
}

export const deleteComponent = (data) => {
  return service({
    url: '/component/deleteComponent',
    method: 'post',
    data
  })
}

export const getAllComponents = (data) => {
  return service({
    url: '/component/getAllComponents',
    method: 'post',
    data
  })
}

export const getGrafanaLink = (data) => {
  return service({
    url: '/component/getGrafanaLink',
    method: 'post',
    data
  })
}

export const getHelmConfig = (data) => {
  return service({
    url: '/component/getHelmConfig',
    method: 'post',
    data
  })
}
