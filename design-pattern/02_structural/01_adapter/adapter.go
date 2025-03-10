package adapter

// MediaPlayer 媒体播放器接口（目标接口）
type MediaPlayer interface {
	Play(audioType, fileName string)
}

// AdvancedMediaPlayer 高级媒体播放器接口（适配者接口）
type AdvancedMediaPlayer interface {
	PlayVlc(fileName string)
	PlayMp4(fileName string)
}

// VlcPlayer 具体的VLC播放器
type VlcPlayer struct{}

func (v *VlcPlayer) PlayVlc(fileName string) {
	println("Playing vlc file. Name:", fileName)
}

func (v *VlcPlayer) PlayMp4(fileName string) {
	// 什么都不做
}

// Mp4Player 具体的MP4播放器
type Mp4Player struct{}

func (m *Mp4Player) PlayVlc(fileName string) {
	// 什么都不做
}

func (m *Mp4Player) PlayMp4(fileName string) {
	println("Playing mp4 file. Name:", fileName)
}

// MediaAdapter 媒体适配器
type MediaAdapter struct {
	advancedMusicPlayer AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) *MediaAdapter {
	var advancedMusicPlayer AdvancedMediaPlayer
	if audioType == "vlc" {
		advancedMusicPlayer = &VlcPlayer{}
	} else if audioType == "mp4" {
		advancedMusicPlayer = &Mp4Player{}
	}
	return &MediaAdapter{advancedMusicPlayer}
}

func (m *MediaAdapter) Play(audioType, fileName string) {
	if audioType == "vlc" {
		m.advancedMusicPlayer.PlayVlc(fileName)
	} else if audioType == "mp4" {
		m.advancedMusicPlayer.PlayMp4(fileName)
	}
}

// AudioPlayer 音频播放器
type AudioPlayer struct {
	mediaAdapter *MediaAdapter
}

func (a *AudioPlayer) Play(audioType, fileName string) {
	// 播放 mp3 音乐文件的内置支持
	if audioType == "mp3" {
		println("Playing mp3 file. Name:", fileName)
	} else if audioType == "vlc" || audioType == "mp4" {
		// mediaAdapter 提供了播放其他文件格式的支持
		a.mediaAdapter = NewMediaAdapter(audioType)
		a.mediaAdapter.Play(audioType, fileName)
	} else {
		println("Invalid media. ", audioType, " format not supported")
	}
}
