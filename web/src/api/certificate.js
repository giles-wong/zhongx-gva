import service from '@/utils/request'

export const getReadingList = (data) => {
    return service({
        url: '/cert/getReadingList',
        method: 'post',
        data: data
    })
}

export const createReading = (data) => {
    return service({
        url: '/cert/addReading',
        method: 'post',
        data: data
    })
}

export const editReading = (data) => {
    return service({
        url: '/cert/editReading',
        method: 'post',
        data: data
    })
}

export const deleteReading = (data) => {
    return service({
        url: '/cert/deleteReading',
        method: 'delete',
        data: data
    })
}