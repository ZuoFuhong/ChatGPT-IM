import { invoke } from "@tauri-apps/api/tauri";

// 默认配置
const defaultSettings = {
    http_url: 'http://127.0.0.1:8080',
    ws_url: 'ws:/127.0.0.1:8080/ws',
    user_id: '1629770111088857088'
}

// 查询用户设置
export async function getUserSettings() {
    try {
        const data = await invoke('get_config_content');
        return JSON.parse(data)
    } catch(e) {
        console.log(e)
    }
    return defaultSettings
}

// 保存用户设置
export async function saveUserSettings(data) {
    try {
        const res = await invoke('set_config_content', data)
        console.log(res)
    } catch(e) {
        console.log(e)
    }
}