import request from '@/utils/request'

export function list(query) {
  return request({
    url: '/player/list',
    method: 'post',
    data: query
  })
}

export function submit(data) {
  return request({
    url: '/player/submit',
    method: 'post',
    data
  })
}
