/// 查询用户设置
#[tauri::command]
pub fn get_config_content() -> String {
    let data = r#"
        {
            "user_id": "1629770111088857088",
            "http_url": "http://127.0.0.1:8080",
            "ws_url": "ws:/127.0.0.1:8080/ws"
        }
    "#;
    return data.to_string();
}

/// 保存用户设置
#[tauri::command]
pub fn set_config_content() -> String {
    return "ok".to_string()
}