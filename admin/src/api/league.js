import request from '@/utils/request'

export function getList() {
  return request({
    url: '/league/list',
    method: 'post'
  })
}

export function submit(data) {
  return request({
    url: '/league/submit',
    method: 'post',
    data
  })
}
