import request from '@/utils/request'

// 查询公告列表
export function listNotice(query) {
  return request({
    url: '/api/v1/notice/list',
    method: 'get',
    params: query
  })
}

// 查询公告详细
export function getNotice(noticeId) {
  return request({
    url: '/api/v1/notice/' + noticeId,
    method: 'get'
  })
}

// 新增公告
export function addNotice(data) {
  return request({
    url: '/api/v1/notice/add',
    method: 'post',
    data: data
  })
}

// 修改公告
export function updateNotice(data) {
  return request({
    url: '/api/v1/notice/edit',
    method: 'put',
    data: data
  })
}

// 删除公告
export function delNotice(noticeId) {
  return request({
    url: '/api/v1/notice/' + noticeId,
    method: 'delete'
  })
}
