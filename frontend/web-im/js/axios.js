const axiosConfig = {
    baseURL: config.BASE_URL,
    timeout: 5 * 1000,
    crossDomain: true,
    validateStatus(status) {
        return status >= 200 && status < 510
    }
}

const _axios = axios.create(axiosConfig)

_axios.interceptors.request.use(function(originConfig) {
    return originConfig;
}, function(error) {
    return Promise.reject(error);
})

_axios.interceptors.response.use(function (res) {
    if (res.status.toString().charAt(0) === '2') {
        return res.data
    }
    return Promise.reject()
}, function(error) {
    if(!error.response) {
        console.log('error', error)
    }
    return Promise.reject(error)
})

function AsyncGet(url, params = {}) {
    return _axios({
        method: 'get',
        url,
        params
    })
}

function AsyncPost(url, data = {}, params = {}) {
    return _axios({
        method: 'post',
        url,
        data,
        params,
    })
}

function AsyncPut(url, data = {}, params = {}) {
    return _axios({
        method: 'put',
        url,
        params,
        data,
    })
}

function AsyncDelete(url, params = {}) {
    return _axios({
        method: 'delete',
        url,
        params,
    })
}

