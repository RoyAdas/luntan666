import axios from 'axios'
const API = 'http://localhost:8080'

export function login(username, password) {
  return axios.post(API + '/login', { username, password })
}

export function getUser() {
  const token = localStorage.getItem('token') || ''
  return axios.get(API + '/user', { headers: { Authorization: 'Bearer ' + token } })
}

export function updateUser(data) {
  const token = localStorage.getItem('token') || ''
  return axios.put(API + '/user', data, { headers: { Authorization: 'Bearer ' + token } })
}

export function fetchPosts(search='') {
  const q = search ? '?search=' + encodeURIComponent(search) : ''
  return axios.get(API + '/posts' + q)
}

export function fetchPost(id) {
  return axios.get(API + '/posts/' + id)
}

export function createPost(post) {
  const token = localStorage.getItem('token') || ''
  return axios.post(API + '/posts', post, { headers: { Authorization: 'Bearer ' + token } })
}

export function likePost(id) {
  const token = localStorage.getItem('token') || ''
  return axios.post(API + '/posts/' + id + '/like', {}, { headers: { Authorization: 'Bearer ' + token } })
}

export function deletePost(id) {
  const token = localStorage.getItem('token') || ''
  return axios.delete(API + '/posts/' + id, { headers: { Authorization: 'Bearer ' + token } })
}
