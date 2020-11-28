# NilMusic 服务端分离与功能扩充技术方案

## 一、需求

​	原有应用所有数据存储在本地relamDB中，且只根据json配置文件提供简单的专辑与播放功能，无法满足肥豪应付老师的需求，故对原有应用进行改造，将服务端与客户端分离，并对现有专辑功能进行扩充，视情况使用kotlin重构客户端。

## 二、功能与流程

- 提供默认专辑，对所有用户可见
- 用户可创建个人专辑，提供歌名、url、歌手、封面，保存并实现白嫖
- 歌曲播放功能迁移到service
- 实现歌曲播放页面进度条功能
- 视情况添加混淆文件

## 三、框架选型与存储设计

- http服务端使用gin
- 垃圾项目，不需要缓存，直接mysql，视情况考虑用redis和mysql做读写分离
- 客户端使用relam做本地持久化

### 数据字典

#### user

| 字段     | 数据类型      | 是否为空 | 默认值         | 描述                |
| -------- | ------------- | -------- | -------------- | ------------------- |
| id       | int           | not null | auto increment | 用户ID， 主键，自增 |
| pnum     | nvarchar(20)  | not null | 11111111111    | 手机号              |
| password | nvarchar(200) | not null | poiuytrewq     | 密码                |

#### album

| 字段    | 数据类型      | 是否为空 | 默认值                                    | 描述                |
| ------- | ------------- | -------- | ----------------------------------------- | ------------------- |
| id      | int           | not null | auto increment                            | 专辑ID， 主键，自增 |
| name    | nvarchar(200) | not null | new album                                 | 专辑名称            |
| poster  | nvarchar(500) | not null | https://www.suvvm.work/images/ortrait.jpg | 专辑封面            |
| playnum | nvarchar(20)  | not null | 0万                                       | 播放量              |

#### music

| 字段   | 数据类型      | 是否为空 | 默认值                                                       | 描述               |
| ------ | ------------- | -------- | ------------------------------------------------------------ | ------------------ |
| id     | int           | not null | auto increment                                               | 歌曲ID，主键，自增 |
| name   | nvarchar(200) | not null | new music                                                    | 歌曲名称           |
| poster | nvarchar(500) | not null | https://www.suvvm.work/images/ortrait.jpg                    | 歌曲封面           |
| path   | nvarchar(500) | not null | http://m8.music.126.net/20201119220648/17233129086daaf596237f43b218beb5/ymusic/1a32/22d0/301e/3964f63dc993257f280cb214cefc403a.mp3 | 歌曲外链           |
| author | nvarchar(200) | not null | suvvm                                               | 歌手               |

#### user_album

| 字段 | 数据类型 | 是否为空 | 默认值         | 描述                       |
| ---- | -------- | -------- | -------------- | -------------------------- |
| id   | int      | not null | auto increment | 用户专辑关系ID，主键，自增 |
| uid  | int      | not null |                | 用户id                     |
| aid  | int      | not null |                | 专辑id                     |

#### album_music

| 字段 | 数据类型 | 是否为空 | 默认值         | 描述                       |
| ---- | -------- | -------- | -------------- | -------------------------- |
| id   | int      | not null | auto increment | 专辑音乐关系ID，主键，自增 |
| aid  | int      | not null |                | 专辑id                     |
| mid  | int      | not null |                | 音乐id                     |



## 四、核心逻辑

### 系统架构图

![image](https://raw.githubusercontent.com/suvvm/nilmusic_service/master/resources/nilMusic%E7%B3%BB%E7%BB%9F%E6%9E%B6%E6%9E%84.png)

### 1、服务端

- 用户信息存储迁移至服务端
- 服务端api层提供与客户端交互接口

#### 接口idl

互联网垃圾，懒得做token也懒得做验证，凑活着用吧。心情好的话以后会接jwt

1. 注册接口

```
Path:/nilmusic/user/register
Method:POST
Headers:
	content-type:application/json
UrlParam:
BodyParam:
	pnum(string)* 
	password(string)*
response:
	code(int)
	msg(string)
```

2. 登陆接口

```
Path:/nilmusic/user/login
Method:POST
Headers:
	content-type:application/json
UrlParam:
BodyParam:
	pnum(string)*
	password(string)*
response:
	code(int)
	msg(string)
	uid(int)
```

3. 获取当前用户全部专辑

```
Path:/nilmusic/list/all
Method:GET
Headers:
	content-type:application/json
UrlParam:
	uid(int)*
BodyParam:
response:
	code(int)
	msg(string)
	album_list(array)
		(object)
			id(int)
			name(string)
			poster(string)
			play_num(string)
```

4. 创建专辑

```
Path:/nilmusic/album/create
Method:POST
Headers:
	content-type:application/json
UrlParam:
BodyParam:
	uid(int)*
	name(string)*
	poster(string)*
	play_num(string)*
response:
	code(int)
	msg(string)
    aid(int)
```

5. 删除专辑

```
Path:/nilmusic/album/delete
Method:DELETE
Headers:
	content-type:application/json
UrlParam:
BodyParam:
	uid(int)*
	aid(int)*
response:
	code(int)
	msg(string)
```

6. 获取指定专辑中的音乐

```
Path:/nilmusic/album/music
Method:GET
Headers:
	content-type:application/json
UrlParam:
	aid(int)*
BodyParam:
response:
	code(int)
	msg(string)
	music_list(array)
		(object)
			id(int)
			name(string)
			poster(string)
			path(string)
			author(string)
```

7. 向指定专辑中添加音乐

```
Path:/nilmusic/album/music/add
Method:POST
Headers:
	content-type:application/json
UrlParam:
BodyParam:
    aid(int)*
	name(string)*
    poster(string)*
    path(string)*
    author(string)*
response:
	code(int)
	msg(string)
    mid(int)
```

8. 修改指定音乐

```
Path:/nilmusic/album/music/mdf
Method:PUT
Headers:
	content-type:application/json
UrlParam:
BodyParam:
	mid(int)*
	name(string)
    poster(string)
    path(string)
    author(string)
response:
	code(int)
	msg(string)
```

9. 删除专辑中的音乐

```
Path:/nilmusic/album/music/delete
Method:DELETE
Headers:
	content-type:application/json
UrlParam:
BodyParam:
	aid(int)*
	mid(int)*
response:
	code(int)
	msg(string)
```



### 2、客户端

- 提供http工具与服务端交互
- 提供自定专辑相关UI
- 歌曲播放功能迁移到service

