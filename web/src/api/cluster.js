import service from '@/utils/request'

export const getClusterList = (data) => {
    return service({
        url: '/cluster/getClusterList',
        method: 'post',
        data
    })
}

export const createCluster = (data) => {
    return service({
        url: '/cluster/createCluster',
        method: 'post',
        data
    })
}

export const deleteCluster = (data) => {
    return service({
        url: '/cluster/deleteCluster',
        method: 'post',
        data
    })
}

export const getAllClusters = (data) => {
    return service({
        url: '/cluster/getAllClusters',
        method: 'post',
        data
    })
}