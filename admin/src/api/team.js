import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/team/list',
    method: 'post',
    data: query
  })
}

export function submit(data) {
  return request({
    url: '/team/submit',
    method: 'post',
    data
  })
}
