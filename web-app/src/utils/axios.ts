import axios from 'axios'

const instance = axios.create({
  baseURL: 'http://124.220.48.103:3000/api'
})

export default instance