package db

import (
	"fmt"
	"os"
	"path/filepath"
	"suvvm.work/nilmusic_service/config"
	"suvvm.work/nilmusic_service/model"
	"testing"
)

var (
	dbConfig = "../../conf/db_config.yaml"
)

func init() {
	str, err := os.Getwd() // 获取相对路径
	if err != nil {
		panic(fmt.Sprintf("filepath failed, err=%v", err))
	}
	filename, err := filepath.Abs(filepath.Join(str, dbConfig)) // 获取db配置文件路径
	if err != nil {
		panic(fmt.Sprintf("filepath failed, err=%v", err))
	}
	conf := config.Init(filename)                    // 读取db配置文件
	if err = InitDB(&conf.DBConfig); err != nil { // 初始化db链接
		panic(fmt.Sprintf("init db conn err=%v", err))
	}
}

// TestAddUser 插入用户测试
func TestAddUser(t *testing.T) {
	user := &model.User{
		Pnum: "15098997526",
		Password: "Poiuytrewq1",
	}
	user, err := AddUser(user)
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("resp user:%v", user)
}

// TestGetUser 查询用户测试
func TestGetUser(t *testing.T) {
	user := &model.User{
		Pnum: "15098997526",
	}
	user, err := GetUser(user)
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("resp user:%v", user)
}

// TestMdfUser 修改用户测试
func TestMdfUser(t *testing.T) {
	user, err := GetUser(&model.User{Pnum: "15098997526"})
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("resp user:%v", user)
	tusr := &model.User{
		ID: user.ID,
		Password: "!!!!!!!!!!!!",
	}
	err = MdfUser(tusr)
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("mdf user success")
}

// TestDelUser 删除用户测试
func TestDelUser(t *testing.T) {
	user, err := GetUser(&model.User{Pnum: "15098997526"})
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("resp user:%v", user)
	tusr := &model.User{
		Pnum: user.Pnum,
	}
	err = DelUser(tusr)
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("del user success")
}

// TestAddAlbum 插入专辑测试
func TestAddAlbum(t *testing.T) {
	album := &model.Album{
		Name: "武器A",
		Poster: "https://www.suvvm.work/images/ortrait.jpg",
		Playnum: "0万",
	}
	album, err := AddAlbum(album)
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("%v", album)
}

// TestGetAlbum 查询专辑测试
func TestGetAlbum(t *testing.T) {
	album := &model.Album{
		//ID: 20000000,
		Name: "武器A",
	}
	albumList, err := GetAlbum(album)
	if err != nil {
		t.Logf("%v", err)
	}
	t.Log(albumList)
}

// TestMdfAlbum 修改专辑测试
func TestMdfAlbum(t *testing.T) {
	albumList, err := GetAlbum(&model.Album{ID: 20000000})
	if err != nil {
		t.Logf("%v", err)
	}
	if len(*albumList) == 0 {
		t.Error("no data")
		return
	}
	talbum:= &model.Album{
		ID: (*albumList)[0].ID,
		Name: "武器B",
	}
	err = MdfAlbum(talbum)
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("mdf album success")
}

// TestDelAlbum 删除专辑测试
func TestDelAlbum(t *testing.T) {
	err := DelAlbum(&model.Album{ID: 20000000})
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("del album success")
}

// TestAddMusic 插入音乐测试
func TestAddMusic(t *testing.T) {
	music := &model.Music{
		Name: "国际歌",
		Poster: "https://www.suvvm.work/images/ortrait.jpg",
		Path: "http://m8.music.126.net/20201119220648/17233129086daaf596237f43b218beb5/ymusic/1a32/22d0/301e/3964f63dc993257f280cb214cefc403a.mp3 ",
		Author: "suvvm",
	}
	music, err := AddMusic(music)
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("%v", music)
}

// TestGetMusic查询音乐测试
func TestGetMusic(t *testing.T) {
	music := &model.Music{
		//ID: 30000000,
		Name: "国际歌",
	}
	musicList, err := GetMusic(music)
	if err != nil {
		t.Logf("%v", err)
	}
	t.Log(musicList)
}

// TestMdfMusic 修改音乐测试
func TestMdfMusic(t *testing.T) {
	musicList, err := GetMusic(&model.Music{ID: 30000000})
	if err != nil {
		t.Logf("%v", err)
	}
	if len(*musicList) == 0 {
		t.Error("no data")
		return
	}
	tmusic:= &model.Music{
		ID: (*musicList)[0].ID,
		Author: "中央合唱团",
	}
	err = MdfMusic(tmusic)
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("mdf music success")
}

// TestDelMusic 删除音乐测试
func TestDelMusic(t *testing.T) {
	err := DelMusic(&model.Music{ID: 30000000})
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("del music success")
}

// TestAddUserAlbum 插入用户专辑关系
func TestAddUserAlbum(t *testing.T) {
	userAlbum := &model.UserAlbum{
		Uid: 10000000,
		Aid: 20000001,
	}
	userAlbum, err := AddUserAlbum(userAlbum)
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("%v", userAlbum)
}

// TestGetMusic查询用户专辑关系测试
func TestGetUserAlbum(t *testing.T) {
	userAlbum := &model.UserAlbum{
		//ID: 30000000,
		Uid: 10000000,
	}
	userAlbums, err := GetUserAlbum(userAlbum)
	if err != nil {
		t.Logf("%v", err)
	}
	t.Log(userAlbums)
}

// TestDelMusic 删除用户专辑关系测试
func TestDelUserAlbum(t *testing.T) {
	err := DelUserAlbum(&model.UserAlbum{Uid: 10000000})
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("del UserAlbum success")
}

// TestAddAlbumMusic 插入专辑音乐关系
func TestAddAlbumMusic(t *testing.T) {
	albumMusic := &model.AlbumMusic{
		Mid: 30000000,
		Aid: 20000001,
	}
	albumMusic, err := AddAlbumMusic(albumMusic)
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("%v", albumMusic)
}

// TestGetAlbumMusic查询专辑音乐关系测试
func TestGetAlbumMusic(t *testing.T) {
	albumMusic := &model.AlbumMusic{
		Aid: 20000001,
	}
	albumMusic, err := GetAlbumMusic(albumMusic)
	if err != nil {
		t.Logf("%v", err)
	}
	t.Log(albumMusic)
}

// TestDelMusic 删除专辑音乐关系测试
func TestDelAlbumMusic(t *testing.T) {
	err := DelAlbumMusic(&model.AlbumMusic{Aid: 20000001})
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("del AlbumMusic success")
}
