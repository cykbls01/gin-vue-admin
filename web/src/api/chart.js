import service from '@/utils/request'

export const getChartList = (data) => {
  return service({
    url: '/chart/getChartList',
    method: 'post',
    data
  })
}

export const createChart = (data) => {
  return service({
    url: '/chart/createChart',
    method: 'post',
    data
  })
}

export const deleteChart = (data) => {
  return service({
    url: '/chart/deleteChart',
    method: 'post',
    data
  })
}

export const getAllCharts = (data) => {
  return service({
    url: '/chart/getAllCharts',
    method: 'post',
    data
  })
}
