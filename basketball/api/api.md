# 接口文档
## 基本协议
### 请求域名 
   - 测试环境：`http://localhost:8081`
   - 生产环境：`https://api.example.com`
### 请求协议
   - 所有请求均采用 HTTP 协议
   - 请求方法：`POST`
   - 请求头：
      - `Content-Type: application/json`
      - `Authorization: Bearer <token>`
      > 注意：在每个需要认证的接口中，都需要在请求头中包含 `Authorization` 字段，值为 `Bearer <token>`，其中 `<token>` 为用户登录后获取的认证令牌。

### 响应协议
   - 所有响应均采用 JSON 格式
   - 响应参数：
      | 参数名 | 类型 | 描述 |
      | ------ | ---- | ---- |
      | `code` | int | 状态码, 200为成功 |
      | `msg` | string | 状态描述 |
      | `data` | object | 响应数据，jsonObject |
   - 响应示例：
      ```json
      {
         "code": 200,
         "msg": "success",
         "data": {
            "token": "example_token",
            "expire": 1694560000
         }
      }
      ```

### 公共参数
   | 参数名 | 类型 | 是否必填 | 描述 |
   | ------ | ---- | -------- | ---- |
   | `timestamp` | int64 | 是 | 时间戳（单位：秒） |
   | `sign` | string | 是 | 签名 |

### 签名算法
   - 签名算法：`MD5(timestamp + app_secret)`
   - 签名示例：`MD5(1694560000 + example_app_secret) = 1234567890abcdef1234567890abcdef`

## 用户模块
### 登陆接口
   - 接口路径：`/user/login`
   - 请求方法：`POST`
   - 私有参数：
      | 参数名 | 类型 | 是否必填 | 描述 |
      | ------ | ---- | -------- | ---- |
      | `username` | string | 是 | 用户名 |
      | `password` | string | 是 | 密码 |
   - 响应参数：
      | 参数名 | 类型 | 描述 |
      | ------ | ---- | ---- |
      | `token` | string | 认证令牌 |
      | `expire` | int64 | 过期时间（单位：秒） |
      | `refresh_token` | string | 刷新令牌 |
      | `refresh_expire` | int64 | 刷新令牌过期时间（单位：秒） |

### 注册接口
   - 接口路径：`/user/register`
   - 请求方法：`POST`
   - 私有参数：
      | 参数名 | 类型 | 是否必填 | 描述 |
      | ------ | ---- | -------- | ---- |
      | `username` | string | 是 | 用户名 |
      | `password` | string | 是 | 密码 |
      | `confirm_password` | string | 是 | 确认密码 |
   - 响应参数：
      | 参数名 | 类型 | 描述 |
      | ------ | ---- | ---- |
      | `id` | string | 用户ID |

## 联赛模块
### 创建联赛接口
   - 接口路径：`/league/submit`
   - 请求方法：`POST`
   - 私有参数：
      | 参数名 | 类型 | 是否必填 | 描述 |
      | ------ | ---- | -------- | ---- |
      | `name` | string | 是 | 联赛名称 |
      | `description` | string | 是 | 联赛描述 |

## 球队模块
### 创建球队接口
   - 接口路径：`/team/submit`
   - 请求方法：`POST`
   - 私有参数：
      | 参数名 | 类型 | 是否必填 | 描述 |
      | ------ | ---- | -------- | ---- |
      | `name` | string | 是 | 球队名称 |
      | `description` | string | 是 | 球队描述 |

## 球员模块
### 创建球员接口
   - 接口路径：`/player/submit`
   - 请求方法：`POST`
   - 私有参数：
      | 参数名 | 类型 | 是否必填 | 描述 |
      | ------ | ---- | -------- | ---- |
      | `name` | string | 是 | 球员名称 |
      | `description` | string | 否 | 球员描述 |
      | `team_id` | int | 否 | 球队ID |
      | `position` | string | 否 | 球员位置 |
      | `number` | int | 否 | 球员号码 |
      | `height` | float | 否 | 球员身高 |
      | `weight` | float | 否 | 球员体重 |
