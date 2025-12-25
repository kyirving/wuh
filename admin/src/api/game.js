import request from '@/utils/request'

export function list(query) {
  return request({
    url: '/game/list',
    method: 'post',
    data: query
  })
}

export function submit(data) {
  return request({
    url: '/game/submit',
    method: 'post',
    data
  })
}
