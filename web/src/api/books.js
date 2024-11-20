import service from '@/utils/request'

export const getBooksList = (data) => {
    return service({
        url: '/books/getBookList',
        method: 'post',
        data: data
    })
}

export const createBook = (data) => {
    return service({
        url: '/books/addBook',
        method: 'post',
        data: data
    })
}

export const editBook = (data) => {
    return service({
        url: '/books/editBook',
        method: 'post',
        data: data
    })
}

export const deleteBook = (data) => {
    return service({
        url: '/books/deleteBook',
        method: 'delete',
        data: data
    })
}