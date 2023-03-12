import axios from 'axios'

class Axios {

  _axios = null;

  constructor(addr) {
    const axiosConfig = {
      baseURL: addr || '/',
      timeout: 5 * 1000,
      crossDomain: true,
      validateStatus(status) {
        return status >= 200 && status < 510
      }
    }

    const _axios = axios.create(axiosConfig)
    _axios.interceptors.response.use(function (res) {
      if (res.status.toString().charAt(0) === '2') {
        return res.data
      }
      return Promise.reject()
    }, function (error) {
      if(!error.response) {
        console.log('error', error)
      }
      return Promise.reject(error)
    })
    this._axios = _axios
  }

  get(url, params = {}) {
    return this._axios({
      method: 'get',
      url,
      params,
    })
  }

  post(url, params = {}) {
    return this._axios({
      method: 'post',
      url,
      params,
    })
  }
}

export default Axios